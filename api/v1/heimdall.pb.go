// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stellarproject/heimdall/api/v1/heimdall.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Master struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GRPCAddress          string   `protobuf:"bytes,2,opt,name=grpc_address,json=grpcAddress,proto3" json:"grpc_address,omitempty"`
	RedisURL             string   `protobuf:"bytes,3,opt,name=redis_url,json=redisUrl,proto3" json:"redis_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Master) Reset()         { *m = Master{} }
func (m *Master) String() string { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()    {}
func (*Master) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{0}
}
func (m *Master) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Master.Unmarshal(m, b)
}
func (m *Master) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Master.Marshal(b, m, deterministic)
}
func (m *Master) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Master.Merge(m, src)
}
func (m *Master) XXX_Size() int {
	return xxx_messageInfo_Master.Size(m)
}
func (m *Master) XXX_DiscardUnknown() {
	xxx_messageInfo_Master.DiscardUnknown(m)
}

var xxx_messageInfo_Master proto.InternalMessageInfo

func (m *Master) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Master) GetGRPCAddress() string {
	if m != nil {
		return m.GRPCAddress
	}
	return ""
}

func (m *Master) GetRedisURL() string {
	if m != nil {
		return m.RedisURL
	}
	return ""
}

type ConnectRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterKey           string   `protobuf:"bytes,2,opt,name=cluster_key,json=clusterKey,proto3" json:"cluster_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{1}
}
func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ConnectRequest) GetClusterKey() string {
	if m != nil {
		return m.ClusterKey
	}
	return ""
}

type ConnectResponse struct {
	Master               *Master  `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{2}
}
func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

func (m *ConnectResponse) GetMaster() *Master {
	if m != nil {
		return m.Master
	}
	return nil
}

type KeyPair struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PublicKey            string   `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPair) Reset()         { *m = KeyPair{} }
func (m *KeyPair) String() string { return proto.CompactTextString(m) }
func (*KeyPair) ProtoMessage()    {}
func (*KeyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{3}
}
func (m *KeyPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPair.Unmarshal(m, b)
}
func (m *KeyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPair.Marshal(b, m, deterministic)
}
func (m *KeyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPair.Merge(m, src)
}
func (m *KeyPair) XXX_Size() int {
	return xxx_messageInfo_KeyPair.Size(m)
}
func (m *KeyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPair proto.InternalMessageInfo

func (m *KeyPair) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *KeyPair) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type Node struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	KeyPair              *KeyPair `protobuf:"bytes,3,opt,name=keypair,proto3" json:"keypair,omitempty"`
	EndpointIP           string   `protobuf:"bytes,4,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	EndpointPort         uint64   `protobuf:"varint,5,opt,name=endpoint_port,json=endpointPort,proto3" json:"endpoint_port,omitempty"`
	GatewayIP            string   `protobuf:"bytes,6,opt,name=gateway_ip,json=gatewayIp,proto3" json:"gateway_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{4}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Node) GetKeyPair() *KeyPair {
	if m != nil {
		return m.KeyPair
	}
	return nil
}

func (m *Node) GetEndpointIP() string {
	if m != nil {
		return m.EndpointIP
	}
	return ""
}

func (m *Node) GetEndpointPort() uint64 {
	if m != nil {
		return m.EndpointPort
	}
	return 0
}

func (m *Node) GetGatewayIP() string {
	if m != nil {
		return m.GatewayIP
	}
	return ""
}

type NodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{5}
}
func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

type NodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{6}
}
func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Peer struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KeyPair              *KeyPair `protobuf:"bytes,2,opt,name=keypair,proto3" json:"keypair,omitempty"`
	AllowedIPs           []string `protobuf:"bytes,3,rep,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	Endpoint             string   `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{7}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Peer) GetKeyPair() *KeyPair {
	if m != nil {
		return m.KeyPair
	}
	return nil
}

func (m *Peer) GetAllowedIPs() []string {
	if m != nil {
		return m.AllowedIPs
	}
	return nil
}

func (m *Peer) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type PeersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersRequest) Reset()         { *m = PeersRequest{} }
func (m *PeersRequest) String() string { return proto.CompactTextString(m) }
func (*PeersRequest) ProtoMessage()    {}
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{8}
}
func (m *PeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersRequest.Unmarshal(m, b)
}
func (m *PeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersRequest.Marshal(b, m, deterministic)
}
func (m *PeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersRequest.Merge(m, src)
}
func (m *PeersRequest) XXX_Size() int {
	return xxx_messageInfo_PeersRequest.Size(m)
}
func (m *PeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeersRequest proto.InternalMessageInfo

type PeersResponse struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersResponse) Reset()         { *m = PeersResponse{} }
func (m *PeersResponse) String() string { return proto.CompactTextString(m) }
func (*PeersResponse) ProtoMessage()    {}
func (*PeersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{9}
}
func (m *PeersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersResponse.Unmarshal(m, b)
}
func (m *PeersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersResponse.Marshal(b, m, deterministic)
}
func (m *PeersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersResponse.Merge(m, src)
}
func (m *PeersResponse) XXX_Size() int {
	return xxx_messageInfo_PeersResponse.Size(m)
}
func (m *PeersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeersResponse proto.InternalMessageInfo

func (m *PeersResponse) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Route struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Network              string   `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{10}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *Route) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type CreateRouteRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Network              string   `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRouteRequest) Reset()         { *m = CreateRouteRequest{} }
func (m *CreateRouteRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRouteRequest) ProtoMessage()    {}
func (*CreateRouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{11}
}
func (m *CreateRouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRouteRequest.Unmarshal(m, b)
}
func (m *CreateRouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRouteRequest.Marshal(b, m, deterministic)
}
func (m *CreateRouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRouteRequest.Merge(m, src)
}
func (m *CreateRouteRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRouteRequest.Size(m)
}
func (m *CreateRouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRouteRequest proto.InternalMessageInfo

func (m *CreateRouteRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *CreateRouteRequest) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type DeleteRouteRequest struct {
	Network              string   `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRouteRequest) Reset()         { *m = DeleteRouteRequest{} }
func (m *DeleteRouteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRouteRequest) ProtoMessage()    {}
func (*DeleteRouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{12}
}
func (m *DeleteRouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRouteRequest.Unmarshal(m, b)
}
func (m *DeleteRouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRouteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRouteRequest.Merge(m, src)
}
func (m *DeleteRouteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRouteRequest.Size(m)
}
func (m *DeleteRouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRouteRequest proto.InternalMessageInfo

func (m *DeleteRouteRequest) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type RoutesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesRequest) Reset()         { *m = RoutesRequest{} }
func (m *RoutesRequest) String() string { return proto.CompactTextString(m) }
func (*RoutesRequest) ProtoMessage()    {}
func (*RoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{13}
}
func (m *RoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesRequest.Unmarshal(m, b)
}
func (m *RoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesRequest.Marshal(b, m, deterministic)
}
func (m *RoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesRequest.Merge(m, src)
}
func (m *RoutesRequest) XXX_Size() int {
	return xxx_messageInfo_RoutesRequest.Size(m)
}
func (m *RoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesRequest proto.InternalMessageInfo

type RoutesResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesResponse) Reset()         { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()    {}
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6184fc395da86b1, []int{14}
}
func (m *RoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesResponse.Unmarshal(m, b)
}
func (m *RoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesResponse.Marshal(b, m, deterministic)
}
func (m *RoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesResponse.Merge(m, src)
}
func (m *RoutesResponse) XXX_Size() int {
	return xxx_messageInfo_RoutesResponse.Size(m)
}
func (m *RoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesResponse proto.InternalMessageInfo

func (m *RoutesResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*Master)(nil), "io.stellarproject.heimdall.api.v1.Master")
	proto.RegisterType((*ConnectRequest)(nil), "io.stellarproject.heimdall.api.v1.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "io.stellarproject.heimdall.api.v1.ConnectResponse")
	proto.RegisterType((*KeyPair)(nil), "io.stellarproject.heimdall.api.v1.KeyPair")
	proto.RegisterType((*Node)(nil), "io.stellarproject.heimdall.api.v1.Node")
	proto.RegisterType((*NodesRequest)(nil), "io.stellarproject.heimdall.api.v1.NodesRequest")
	proto.RegisterType((*NodesResponse)(nil), "io.stellarproject.heimdall.api.v1.NodesResponse")
	proto.RegisterType((*Peer)(nil), "io.stellarproject.heimdall.api.v1.Peer")
	proto.RegisterType((*PeersRequest)(nil), "io.stellarproject.heimdall.api.v1.PeersRequest")
	proto.RegisterType((*PeersResponse)(nil), "io.stellarproject.heimdall.api.v1.PeersResponse")
	proto.RegisterType((*Route)(nil), "io.stellarproject.heimdall.api.v1.Route")
	proto.RegisterType((*CreateRouteRequest)(nil), "io.stellarproject.heimdall.api.v1.CreateRouteRequest")
	proto.RegisterType((*DeleteRouteRequest)(nil), "io.stellarproject.heimdall.api.v1.DeleteRouteRequest")
	proto.RegisterType((*RoutesRequest)(nil), "io.stellarproject.heimdall.api.v1.RoutesRequest")
	proto.RegisterType((*RoutesResponse)(nil), "io.stellarproject.heimdall.api.v1.RoutesResponse")
}

func init() {
	proto.RegisterFile("github.com/stellarproject/heimdall/api/v1/heimdall.proto", fileDescriptor_b6184fc395da86b1)
}

var fileDescriptor_b6184fc395da86b1 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x6a, 0xeb, 0x46,
	0x18, 0xc5, 0x7f, 0xb2, 0xfd, 0xc9, 0x76, 0x60, 0x28, 0x41, 0xb8, 0x14, 0xb9, 0xca, 0xa2, 0x4e,
	0x09, 0x52, 0xec, 0x52, 0x28, 0x94, 0x42, 0xf3, 0xd7, 0x54, 0xa4, 0x0d, 0xee, 0xb4, 0xd9, 0x94,
	0x82, 0x91, 0xad, 0xa9, 0xa3, 0x44, 0xf6, 0x4c, 0x47, 0x63, 0x07, 0xaf, 0xba, 0xee, 0x4b, 0xf5,
	0x51, 0xbc, 0xf0, 0x8b, 0xb4, 0xcc, 0x68, 0xa4, 0x6b, 0xdf, 0x90, 0x7b, 0x15, 0xee, 0xdd, 0xf9,
	0x3b, 0xfa, 0xce, 0x99, 0x33, 0x67, 0xfe, 0x0c, 0xdf, 0xcc, 0x22, 0x71, 0xbf, 0x9c, 0xb8, 0x53,
	0x3a, 0xf7, 0x12, 0x41, 0xe2, 0x38, 0xe0, 0x8c, 0xd3, 0x07, 0x32, 0x15, 0xde, 0x3d, 0x89, 0xe6,
	0x61, 0x10, 0xc7, 0x5e, 0xc0, 0x22, 0x6f, 0x35, 0xc8, 0x6b, 0x97, 0x71, 0x2a, 0x28, 0xfa, 0x3c,
	0xa2, 0xee, 0x3e, 0xc3, 0xcd, 0x3b, 0x02, 0x16, 0xb9, 0xab, 0x41, 0xf7, 0x93, 0x19, 0x9d, 0x51,
	0xd5, 0xed, 0xc9, 0x5f, 0x29, 0xb1, 0xfb, 0xe9, 0x8c, 0xd2, 0x59, 0x4c, 0x3c, 0x55, 0x4d, 0x96,
	0x7f, 0x7a, 0x64, 0xce, 0xc4, 0x3a, 0xfd, 0xe8, 0xfc, 0x0d, 0xc6, 0xcf, 0x41, 0x22, 0x08, 0x47,
	0x87, 0x50, 0x8e, 0x42, 0xab, 0xd4, 0x2b, 0xf5, 0x9b, 0xe7, 0xc6, 0x76, 0x63, 0x97, 0xfd, 0x4b,
	0x5c, 0x8e, 0x42, 0x34, 0x84, 0xd6, 0x8c, 0xb3, 0xe9, 0x38, 0x08, 0x43, 0x4e, 0x92, 0xc4, 0x2a,
	0xab, 0x8e, 0x83, 0xed, 0xc6, 0x36, 0xaf, 0xf1, 0xe8, 0xe2, 0x2c, 0x85, 0xb1, 0x29, 0x9b, 0x74,
	0x81, 0x8e, 0xa1, 0xc9, 0x49, 0x18, 0x25, 0xe3, 0x25, 0x8f, 0xad, 0x8a, 0x22, 0xb4, 0xb6, 0x1b,
	0xbb, 0x81, 0x25, 0x78, 0x87, 0x7f, 0xc2, 0x0d, 0xf5, 0xf9, 0x8e, 0xc7, 0x8e, 0x0f, 0x9d, 0x0b,
	0xba, 0x58, 0x90, 0xa9, 0xc0, 0xe4, 0xaf, 0x25, 0x49, 0xc4, 0x8b, 0x46, 0x6c, 0x30, 0xa7, 0xf1,
	0x52, 0x7a, 0x1d, 0x3f, 0x92, 0x75, 0xea, 0x03, 0x83, 0x86, 0x6e, 0xc8, 0xda, 0xf9, 0x0d, 0x0e,
	0x72, 0xa9, 0x84, 0xd1, 0x45, 0x42, 0xd0, 0x19, 0x18, 0x73, 0x35, 0x3d, 0xa5, 0x67, 0x0e, 0x8f,
	0xdd, 0xf7, 0xa6, 0xe8, 0xa6, 0x79, 0x60, 0x4d, 0x74, 0x7c, 0xa8, 0xdf, 0x90, 0xf5, 0x28, 0x88,
	0xb8, 0x74, 0xc0, 0x78, 0xb4, 0x0a, 0x04, 0x51, 0x0e, 0x4a, 0xa9, 0x03, 0x0d, 0xdd, 0x90, 0x35,
	0xfa, 0x0c, 0x80, 0x2d, 0x27, 0x71, 0x34, 0xdd, 0x71, 0xd8, 0x4c, 0x11, 0x69, 0xf0, 0x9f, 0x32,
	0x54, 0x6f, 0x69, 0x48, 0x5e, 0x9c, 0x22, 0x82, 0xaa, 0x8c, 0x59, 0x33, 0xd5, 0x6f, 0xf4, 0x0b,
	0xd4, 0x1f, 0xc9, 0x9a, 0x05, 0x11, 0x57, 0x49, 0x9a, 0xc3, 0x2f, 0x0b, 0xcc, 0x41, 0x3b, 0x3e,
	0x37, 0xb7, 0x1b, 0x3b, 0xb3, 0x8f, 0x33, 0x1d, 0xe4, 0x81, 0x49, 0x16, 0x21, 0xa3, 0xd1, 0x42,
	0x8c, 0x23, 0x66, 0x55, 0x95, 0x8f, 0xce, 0x76, 0x63, 0xc3, 0x95, 0x86, 0xfd, 0x11, 0x86, 0xac,
	0xc5, 0x67, 0xe8, 0x08, 0xda, 0x39, 0x81, 0x51, 0x2e, 0xac, 0x5a, 0xaf, 0xd4, 0xaf, 0xe2, 0x56,
	0x06, 0x8e, 0x28, 0x17, 0xe8, 0x04, 0x60, 0x16, 0x08, 0xf2, 0x14, 0xac, 0xa5, 0xa8, 0xa1, 0x44,
	0xdb, 0xdb, 0x8d, 0xdd, 0xbc, 0x4e, 0x51, 0x7f, 0x84, 0x9b, 0xba, 0xc1, 0x67, 0x4e, 0x07, 0x5a,
	0x32, 0x8a, 0x44, 0xaf, 0xba, 0x73, 0x0b, 0x6d, 0x5d, 0xeb, 0xa5, 0xfb, 0x0e, 0x6a, 0x0b, 0x09,
	0x58, 0xa5, 0x5e, 0xa5, 0x6f, 0x0e, 0xbf, 0x28, 0x30, 0x6b, 0x29, 0x80, 0x53, 0x96, 0xf3, 0x6f,
	0x09, 0xaa, 0x23, 0xf2, 0x8e, 0x7d, 0xbd, 0x93, 0x6b, 0xf9, 0xe3, 0xe5, 0x1a, 0xc4, 0x31, 0x7d,
	0x22, 0xe1, 0x38, 0x62, 0x89, 0x55, 0xe9, 0x55, 0xb2, 0x5c, 0xcf, 0x52, 0xd8, 0x1f, 0x25, 0x18,
	0x74, 0x8b, 0xcf, 0x12, 0xd4, 0x85, 0x46, 0x16, 0x61, 0xba, 0x0a, 0x38, 0xaf, 0x65, 0x40, 0xd2,
	0xff, 0x6e, 0x40, 0xba, 0x7e, 0x13, 0x10, 0x93, 0xc0, 0x2b, 0x02, 0x92, 0x02, 0x38, 0x65, 0x39,
	0x3f, 0x40, 0x0d, 0xd3, 0xa5, 0x20, 0xe8, 0x08, 0xea, 0x32, 0xb2, 0x71, 0x9e, 0x12, 0x6c, 0x37,
	0xb6, 0x21, 0xb3, 0xf4, 0x2f, 0xb1, 0x21, 0x3f, 0xf9, 0x21, 0xb2, 0xa0, 0xbe, 0x20, 0xe2, 0x89,
	0xf2, 0x47, 0xbd, 0x39, 0xb3, 0xd2, 0xf9, 0x15, 0xd0, 0x05, 0x27, 0x81, 0x20, 0x4a, 0x2d, 0x3b,
	0xc4, 0x1f, 0x28, 0xea, 0x02, 0xba, 0x24, 0x31, 0x79, 0x4b, 0x74, 0xa7, 0xbf, 0xb4, 0xdf, 0x7f,
	0x00, 0x6d, 0xd5, 0x99, 0xa7, 0x85, 0xa1, 0x93, 0x01, 0x3a, 0xae, 0xef, 0xc1, 0xe0, 0x0a, 0xd1,
	0x79, 0xf5, 0x0b, 0xe4, 0x95, 0x8e, 0xae, 0x79, 0xc3, 0xff, 0xaa, 0xd0, 0xf8, 0x51, 0x77, 0x20,
	0x06, 0x75, 0x7d, 0xd9, 0xa0, 0x41, 0x01, 0xa5, 0xfd, 0x3b, 0xae, 0x3b, 0x7c, 0x0d, 0x45, 0x4f,
	0x60, 0x0e, 0x46, 0x3a, 0x25, 0x74, 0x5a, 0xd4, 0x7a, 0x16, 0x47, 0x77, 0xf0, 0x0a, 0x86, 0x1e,
	0xee, 0x0f, 0x30, 0x77, 0xd6, 0x15, 0x7d, 0x5d, 0xc4, 0xf1, 0xb3, 0x7d, 0xd0, 0x3d, 0x74, 0xd3,
	0xd7, 0xc7, 0xcd, 0x5e, 0x1f, 0xf7, 0x4a, 0xbe, 0x3e, 0x52, 0x7d, 0x67, 0x81, 0x0b, 0xa9, 0x3f,
	0xdf, 0x10, 0x2f, 0xaa, 0x3f, 0x40, 0x4d, 0x5d, 0x26, 0xc8, 0x2b, 0x78, 0x6b, 0xe4, 0x41, 0x9d,
	0x16, 0x27, 0xe8, 0x9c, 0x1e, 0xa0, 0xa6, 0xce, 0x65, 0xa1, 0xb1, 0x76, 0x4f, 0x74, 0xa1, 0xb1,
	0xf6, 0x8e, 0xfc, 0xb9, 0xfb, 0xfb, 0x49, 0xe1, 0xff, 0x0f, 0xdf, 0xae, 0x06, 0x13, 0x43, 0xe5,
	0xf2, 0xd5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xf6, 0x3d, 0x3c, 0x76, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeimdallClient is the client API for Heimdall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeimdallClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Routes(ctx context.Context, in *RoutesRequest, opts ...grpc.CallOption) (*RoutesResponse, error)
	CreateRoute(ctx context.Context, in *CreateRouteRequest, opts ...grpc.CallOption) (*types.Empty, error)
	DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error)
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error)
}

type heimdallClient struct {
	cc *grpc.ClientConn
}

func NewHeimdallClient(cc *grpc.ClientConn) HeimdallClient {
	return &heimdallClient{cc}
}

func (c *heimdallClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) Routes(ctx context.Context, in *RoutesRequest, opts ...grpc.CallOption) (*RoutesResponse, error) {
	out := new(RoutesResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/Routes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) CreateRoute(ctx context.Context, in *CreateRouteRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/CreateRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/DeleteRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error) {
	out := new(NodesResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heimdallClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error) {
	out := new(PeersResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.heimdall.api.v1.Heimdall/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeimdallServer is the server API for Heimdall service.
type HeimdallServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Routes(context.Context, *RoutesRequest) (*RoutesResponse, error)
	CreateRoute(context.Context, *CreateRouteRequest) (*types.Empty, error)
	DeleteRoute(context.Context, *DeleteRouteRequest) (*types.Empty, error)
	Nodes(context.Context, *NodesRequest) (*NodesResponse, error)
	Peers(context.Context, *PeersRequest) (*PeersResponse, error)
}

// UnimplementedHeimdallServer can be embedded to have forward compatible implementations.
type UnimplementedHeimdallServer struct {
}

func (*UnimplementedHeimdallServer) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedHeimdallServer) Routes(ctx context.Context, req *RoutesRequest) (*RoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routes not implemented")
}
func (*UnimplementedHeimdallServer) CreateRoute(ctx context.Context, req *CreateRouteRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoute not implemented")
}
func (*UnimplementedHeimdallServer) DeleteRoute(ctx context.Context, req *DeleteRouteRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoute not implemented")
}
func (*UnimplementedHeimdallServer) Nodes(ctx context.Context, req *NodesRequest) (*NodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nodes not implemented")
}
func (*UnimplementedHeimdallServer) Peers(ctx context.Context, req *PeersRequest) (*PeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}

func RegisterHeimdallServer(s *grpc.Server, srv HeimdallServer) {
	s.RegisterService(&_Heimdall_serviceDesc, srv)
}

func _Heimdall_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_Routes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).Routes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/Routes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).Routes(ctx, req.(*RoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_CreateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).CreateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/CreateRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).CreateRoute(ctx, req.(*CreateRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/DeleteRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).DeleteRoute(ctx, req.(*DeleteRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).Nodes(ctx, req.(*NodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heimdall_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeimdallServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.heimdall.api.v1.Heimdall/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeimdallServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heimdall_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.stellarproject.heimdall.api.v1.Heimdall",
	HandlerType: (*HeimdallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Heimdall_Connect_Handler,
		},
		{
			MethodName: "Routes",
			Handler:    _Heimdall_Routes_Handler,
		},
		{
			MethodName: "CreateRoute",
			Handler:    _Heimdall_CreateRoute_Handler,
		},
		{
			MethodName: "DeleteRoute",
			Handler:    _Heimdall_DeleteRoute_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Heimdall_Nodes_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _Heimdall_Peers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stellarproject/heimdall/api/v1/heimdall.proto",
}
