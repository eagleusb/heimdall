/*
	Copyright 2019 Stellar Project

	Permission is hereby granted, free of charge, to any person obtaining a copy of
	this software and associated documentation files (the "Software"), to deal in the
	Software without restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
	and to permit persons to whom the Software is furnished to do so, subject to the
	following conditions:

	The above copyright notice and this permission notice shall be included in all copies
	or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
	PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE
	FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
	USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/gogo/protobuf/proto"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stellarproject/heimdall"
	v1 "github.com/stellarproject/heimdall/api/v1"
	"github.com/stellarproject/heimdall/client"
	"google.golang.org/grpc"
)

const (
	masterKey       = "heimdall:master"
	clusterKey      = "heimdall:key"
	keypairsKey     = "heimdall:keypairs"
	nodesKey        = "heimdall:nodes"
	nodeJoinKey     = "heimdall:join"
	peersKey        = "heimdall:peers"
	routesKey       = "heimdall:routes"
	peerIPsKey      = "heimdall:peerips"
	nodeIPsKey      = "heimdall:nodeips"
	nodeNetworksKey = "heimdall:nodenetworks"

	wireguardConfigPath = "/etc/wireguard/darknet.conf"
)

var (
	empty                    = &ptypes.Empty{}
	masterHeartbeatInterval  = time.Second * 5
	nodeHeartbeatInterval    = time.Second * 60
	nodeHeartbeatExpiry      = 86400
	peerConfigUpdateInterval = time.Second * 10

	// ErrRouteExists is returned when a requested route is already reserved
	ErrRouteExists = errors.New("route already reserved")
	// ErrNodeDoesNotExist is returned when an invalid node is requested
	ErrNodeDoesNotExist = errors.New("node does not exist")
)

// Server represents the Heimdall server
type Server struct {
	cfg       *heimdall.Config
	rpool     *redis.Pool
	wpool     *redis.Pool
	replicaCh chan struct{}
}

// NewServer returns a new Heimdall server
func NewServer(cfg *heimdall.Config) (*Server, error) {
	pool := getPool(cfg.RedisURL)
	return &Server{
		cfg:       cfg,
		rpool:     pool,
		wpool:     pool,
		replicaCh: make(chan struct{}, 1),
	}, nil
}

// Register enables callers to register this service with an existing GRPC server
func (s *Server) Register(server *grpc.Server) error {
	v1.RegisterHeimdallServer(server, s)
	return nil
}

// GenerateProfile generates a new Go profile
func (s *Server) GenerateProfile() (string, error) {
	tmpfile, err := ioutil.TempFile("", "heimdall-profile-")
	if err != nil {
		return "", err
	}
	runtime.GC()
	if err := pprof.WriteHeapProfile(tmpfile); err != nil {
		return "", err
	}
	tmpfile.Close()
	return tmpfile.Name(), nil
}

func (s *Server) Run() error {
	ctx := context.Background()
	// check peer address and make a grpc request for master info if present
	if s.cfg.GRPCPeerAddress != "" {
		logrus.Debugf("joining %s", s.cfg.GRPCPeerAddress)
		c, err := s.getClient(s.cfg.GRPCPeerAddress)
		if err != nil {
			return err
		}
		defer c.Close()

		master, err := c.Connect(s.cfg.ClusterKey)
		if err != nil {
			return err
		}

		logrus.Debugf("master info received: %+v", master)
		if err := s.joinMaster(master); err != nil {
			return err
		}

		go s.replicaMonitor()
	} else {
		// starting as master; remove existing key
		if err := s.configureNode(); err != nil {
			return err
		}
	}

	// ensure keypair
	if _, err := s.getOrCreateKeyPair(ctx, s.cfg.ID); err != nil {
		return err
	}

	// ensure node network subnet
	if err := s.ensureNetworkSubnet(ctx); err != nil {
		return err
	}

	// start node heartbeat to update in redis
	go s.nodeHeartbeat(ctx)

	// initial peer info update
	if err := s.updatePeerInfo(ctx); err != nil {
		return err
	}

	// initial config update
	if err := s.updatePeerConfig(ctx); err != nil {
		return err
	}

	// start peer config updater to configure wireguard as peers join
	go s.peerUpdater(ctx)

	// start listener for pub/sub
	errCh := make(chan error, 1)
	go func() {
		c := s.rpool.Get()
		defer c.Close()

		psc := redis.PubSubConn{Conn: c}
		psc.Subscribe(nodeJoinKey)
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				// TODO: handle join notify
				logrus.Debug("join notify")
			case redis.Subscription:
			default:
				logrus.Debugf("unknown message type %T", v)
			}
		}
	}()

	err := <-errCh
	return err
}

func (s *Server) Stop() error {
	s.rpool.Close()
	s.wpool.Close()
	return nil
}

func getPool(u string) *redis.Pool {
	pool := redis.NewPool(func() (redis.Conn, error) {
		conn, err := redis.DialURL(u)
		if err != nil {
			return nil, errors.Wrap(err, "unable to connect to redis")
		}
		return conn, nil
	}, 10)

	return pool
}

func (s *Server) ensureNetworkSubnet(ctx context.Context) error {
	network, err := redis.String(s.local(ctx, "GET", s.getNodeNetworkKey(s.cfg.ID)))
	if err != nil {
		if err != redis.ErrNil {
			return err
		}
		// allocate initial node subnet
		r, err := parseSubnetRange(s.cfg.NodeNetwork)
		if err != nil {
			return err
		}
		// iterate node networks to find first free
		nodeNetworkKeys, err := redis.Strings(s.local(ctx, "KEYS", s.getNodeNetworkKey("*")))
		if err != nil {
			return err
		}
		logrus.Debug(nodeNetworkKeys)
		lookup := map[string]struct{}{}
		for _, netKey := range nodeNetworkKeys {
			n, err := redis.String(s.local(ctx, "GET", netKey))
			if err != nil {
				return err
			}
			lookup[n] = struct{}{}
		}

		subnet := r.Subnet
		size, _ := subnet.Mask.Size()

		for {
			n, ok := nextSubnet(subnet, size)
			if !ok {
				return fmt.Errorf("error getting next subnet")
			}
			if _, exists := lookup[n.String()]; exists {
				subnet = n
				continue
			}
			logrus.Debugf("allocated network %s for %s", n.String(), s.cfg.ID)
			if err := s.updateNodeNetwork(ctx, n.String()); err != nil {
				return err
			}
			break
		}

		return nil
	}
	logrus.Debugf("node network: %s", network)
	return nil
}

func (s *Server) getOrCreateKeyPair(ctx context.Context, id string) (*v1.KeyPair, error) {
	key := s.getKeyPairKey(id)
	keyData, err := redis.Bytes(s.master(ctx, "GET", key))
	if err != nil {
		if err != redis.ErrNil {
			return nil, err
		}
		logrus.Debugf("generating new keypair for %s", s.cfg.ID)
		privateKey, publicKey, err := generateWireguardKeys(ctx)
		if err != nil {
			return nil, err
		}
		keyPair := &v1.KeyPair{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		}
		data, err := proto.Marshal(keyPair)
		if err != nil {
			return nil, err
		}
		if _, err := s.master(ctx, "SET", key, data); err != nil {
			return nil, err
		}
		return keyPair, nil
	}

	var keyPair v1.KeyPair
	if err := proto.Unmarshal(keyData, &keyPair); err != nil {
		return nil, err
	}
	return &keyPair, nil
}

func (s *Server) getNodeKey(id string) string {
	return fmt.Sprintf("%s:%s", nodesKey, id)
}

func (s *Server) getRouteKey(network string) string {
	return fmt.Sprintf("%s:%s", routesKey, network)
}

func (s *Server) getPeerKey(id string) string {
	return fmt.Sprintf("%s:%s", peersKey, id)
}

func (s *Server) getKeyPairKey(id string) string {
	return fmt.Sprintf("%s:%s", keypairsKey, id)
}

func (s *Server) getNodeNetworkKey(id string) string {
	return fmt.Sprintf("%s:%s", nodeNetworksKey, id)
}

func (s *Server) getClient(addr string) (*client.Client, error) {
	return client.NewClient(s.cfg.ID, addr)
}

func (s *Server) getClusterKey(ctx context.Context) (string, error) {
	return redis.String(s.local(ctx, "GET", clusterKey))
}

func (s *Server) local(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	return s.do(ctx, s.rpool, cmd, args...)
}

func (s *Server) master(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	return s.do(ctx, s.wpool, cmd, args...)
}

func (s *Server) do(ctx context.Context, pool *redis.Pool, cmd string, args ...interface{}) (interface{}, error) {
	conn, err := pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	r, err := conn.Do(cmd, args...)
	return r, err
}
