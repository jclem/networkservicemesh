// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registry.proto

package registry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkServiceEndpoint struct {
	NetworkServiceName        string            `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	Payload                   string            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	NetworkServiceManagerName string            `protobuf:"bytes,3,opt,name=network_service_manager_name,json=networkServiceManagerName,proto3" json:"network_service_manager_name,omitempty"`
	EndpointName              string            `protobuf:"bytes,4,opt,name=endpoint_name,json=endpointName,proto3" json:"endpoint_name,omitempty"`
	Labels                    map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                     string            `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

func (m *NetworkServiceEndpoint) Reset()         { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()    {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{0}
}
func (m *NetworkServiceEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpoint.Unmarshal(m, b)
}
func (m *NetworkServiceEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpoint.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpoint.Merge(dst, src)
}
func (m *NetworkServiceEndpoint) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpoint.Size(m)
}
func (m *NetworkServiceEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpoint proto.InternalMessageInfo

func (m *NetworkServiceEndpoint) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetNetworkServiceManagerName() string {
	if m != nil {
		return m.NetworkServiceManagerName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

func (m *NetworkServiceEndpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NetworkServiceEndpoint) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type NetworkService struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Matches              []*Match `protobuf:"bytes,3,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}
func (*NetworkService) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{1}
}
func (m *NetworkService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService.Unmarshal(m, b)
}
func (m *NetworkService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService.Marshal(b, m, deterministic)
}
func (dst *NetworkService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService.Merge(dst, src)
}
func (m *NetworkService) XXX_Size() int {
	return xxx_messageInfo_NetworkService.Size(m)
}
func (m *NetworkService) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService proto.InternalMessageInfo

func (m *NetworkService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkService) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *NetworkService) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type Match struct {
	SourceSelector       map[string]string `protobuf:"bytes,1,rep,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Routes               []*Destination    `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{2}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (dst *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(dst, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetSourceSelector() map[string]string {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *Match) GetRoutes() []*Destination {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Destination struct {
	DestinationSelector  map[string]string `protobuf:"bytes,1,rep,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Weight               uint32            `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Destination) Reset()         { *m = Destination{} }
func (m *Destination) String() string { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()    {}
func (*Destination) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{3}
}
func (m *Destination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destination.Unmarshal(m, b)
}
func (m *Destination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destination.Marshal(b, m, deterministic)
}
func (dst *Destination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destination.Merge(dst, src)
}
func (m *Destination) XXX_Size() int {
	return xxx_messageInfo_Destination.Size(m)
}
func (m *Destination) XXX_DiscardUnknown() {
	xxx_messageInfo_Destination.DiscardUnknown(m)
}

var xxx_messageInfo_Destination proto.InternalMessageInfo

func (m *Destination) GetDestinationSelector() map[string]string {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *Destination) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type NetworkServiceManager struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	LastSeen             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_seen,json=lastSeen,proto3" json:"last_seen,omitempty"`
	State                string               `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetworkServiceManager) Reset()         { *m = NetworkServiceManager{} }
func (m *NetworkServiceManager) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceManager) ProtoMessage()    {}
func (*NetworkServiceManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{4}
}
func (m *NetworkServiceManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceManager.Unmarshal(m, b)
}
func (m *NetworkServiceManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceManager.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceManager.Merge(dst, src)
}
func (m *NetworkServiceManager) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceManager.Size(m)
}
func (m *NetworkServiceManager) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceManager.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceManager proto.InternalMessageInfo

func (m *NetworkServiceManager) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServiceManager) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *NetworkServiceManager) GetLastSeen() *timestamp.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *NetworkServiceManager) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type RemoveNSERequest struct {
	EndpointName         string   `protobuf:"bytes,1,opt,name=endpoint_name,json=endpointName,proto3" json:"endpoint_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveNSERequest) Reset()         { *m = RemoveNSERequest{} }
func (m *RemoveNSERequest) String() string { return proto.CompactTextString(m) }
func (*RemoveNSERequest) ProtoMessage()    {}
func (*RemoveNSERequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{5}
}
func (m *RemoveNSERequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNSERequest.Unmarshal(m, b)
}
func (m *RemoveNSERequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNSERequest.Marshal(b, m, deterministic)
}
func (dst *RemoveNSERequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNSERequest.Merge(dst, src)
}
func (m *RemoveNSERequest) XXX_Size() int {
	return xxx_messageInfo_RemoveNSERequest.Size(m)
}
func (m *RemoveNSERequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNSERequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNSERequest proto.InternalMessageInfo

func (m *RemoveNSERequest) GetEndpointName() string {
	if m != nil {
		return m.EndpointName
	}
	return ""
}

type FindNetworkServiceRequest struct {
	NetworkServiceName   string   `protobuf:"bytes,1,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNetworkServiceRequest) Reset()         { *m = FindNetworkServiceRequest{} }
func (m *FindNetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*FindNetworkServiceRequest) ProtoMessage()    {}
func (*FindNetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{6}
}
func (m *FindNetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNetworkServiceRequest.Unmarshal(m, b)
}
func (m *FindNetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNetworkServiceRequest.Marshal(b, m, deterministic)
}
func (dst *FindNetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNetworkServiceRequest.Merge(dst, src)
}
func (m *FindNetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_FindNetworkServiceRequest.Size(m)
}
func (m *FindNetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindNetworkServiceRequest proto.InternalMessageInfo

func (m *FindNetworkServiceRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

type FindNetworkServiceResponse struct {
	Payload                 string                            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	NetworkService          *NetworkService                   `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	NetworkServiceManagers  map[string]*NetworkServiceManager `protobuf:"bytes,3,rep,name=network_service_managers,json=networkServiceManagers,proto3" json:"network_service_managers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkServiceEndpoints []*NetworkServiceEndpoint         `protobuf:"bytes,4,rep,name=network_service_endpoints,json=networkServiceEndpoints,proto3" json:"network_service_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                          `json:"-"`
	XXX_unrecognized        []byte                            `json:"-"`
	XXX_sizecache           int32                             `json:"-"`
}

func (m *FindNetworkServiceResponse) Reset()         { *m = FindNetworkServiceResponse{} }
func (m *FindNetworkServiceResponse) String() string { return proto.CompactTextString(m) }
func (*FindNetworkServiceResponse) ProtoMessage()    {}
func (*FindNetworkServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{7}
}
func (m *FindNetworkServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNetworkServiceResponse.Unmarshal(m, b)
}
func (m *FindNetworkServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNetworkServiceResponse.Marshal(b, m, deterministic)
}
func (dst *FindNetworkServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNetworkServiceResponse.Merge(dst, src)
}
func (m *FindNetworkServiceResponse) XXX_Size() int {
	return xxx_messageInfo_FindNetworkServiceResponse.Size(m)
}
func (m *FindNetworkServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNetworkServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindNetworkServiceResponse proto.InternalMessageInfo

func (m *FindNetworkServiceResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *FindNetworkServiceResponse) GetNetworkService() *NetworkService {
	if m != nil {
		return m.NetworkService
	}
	return nil
}

func (m *FindNetworkServiceResponse) GetNetworkServiceManagers() map[string]*NetworkServiceManager {
	if m != nil {
		return m.NetworkServiceManagers
	}
	return nil
}

func (m *FindNetworkServiceResponse) GetNetworkServiceEndpoints() []*NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkServiceEndpoints
	}
	return nil
}

type NSERegistration struct {
	NetworkService         *NetworkService         `protobuf:"bytes,1,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	NetworkServiceManager  *NetworkServiceManager  `protobuf:"bytes,2,opt,name=network_service_manager,json=networkServiceManager,proto3" json:"network_service_manager,omitempty"`
	NetworkserviceEndpoint *NetworkServiceEndpoint `protobuf:"bytes,3,opt,name=networkservice_endpoint,json=networkserviceEndpoint,proto3" json:"networkservice_endpoint,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *NSERegistration) Reset()         { *m = NSERegistration{} }
func (m *NSERegistration) String() string { return proto.CompactTextString(m) }
func (*NSERegistration) ProtoMessage()    {}
func (*NSERegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{8}
}
func (m *NSERegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSERegistration.Unmarshal(m, b)
}
func (m *NSERegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSERegistration.Marshal(b, m, deterministic)
}
func (dst *NSERegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSERegistration.Merge(dst, src)
}
func (m *NSERegistration) XXX_Size() int {
	return xxx_messageInfo_NSERegistration.Size(m)
}
func (m *NSERegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_NSERegistration.DiscardUnknown(m)
}

var xxx_messageInfo_NSERegistration proto.InternalMessageInfo

func (m *NSERegistration) GetNetworkService() *NetworkService {
	if m != nil {
		return m.NetworkService
	}
	return nil
}

func (m *NSERegistration) GetNetworkServiceManager() *NetworkServiceManager {
	if m != nil {
		return m.NetworkServiceManager
	}
	return nil
}

func (m *NSERegistration) GetNetworkserviceEndpoint() *NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkserviceEndpoint
	}
	return nil
}

type NetworkServiceEndpointList struct {
	NetworkServiceEndpoints []*NetworkServiceEndpoint `protobuf:"bytes,1,rep,name=network_service_endpoints,json=networkServiceEndpoints,proto3" json:"network_service_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *NetworkServiceEndpointList) Reset()         { *m = NetworkServiceEndpointList{} }
func (m *NetworkServiceEndpointList) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpointList) ProtoMessage()    {}
func (*NetworkServiceEndpointList) Descriptor() ([]byte, []int) {
	return fileDescriptor_registry_f962b46a720ff8bc, []int{9}
}
func (m *NetworkServiceEndpointList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpointList.Unmarshal(m, b)
}
func (m *NetworkServiceEndpointList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpointList.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpointList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpointList.Merge(dst, src)
}
func (m *NetworkServiceEndpointList) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpointList.Size(m)
}
func (m *NetworkServiceEndpointList) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpointList.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpointList proto.InternalMessageInfo

func (m *NetworkServiceEndpointList) GetNetworkServiceEndpoints() []*NetworkServiceEndpoint {
	if m != nil {
		return m.NetworkServiceEndpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "registry.NetworkServiceEndpoint")
	proto.RegisterMapType((map[string]string)(nil), "registry.NetworkServiceEndpoint.LabelsEntry")
	proto.RegisterType((*NetworkService)(nil), "registry.NetworkService")
	proto.RegisterType((*Match)(nil), "registry.Match")
	proto.RegisterMapType((map[string]string)(nil), "registry.Match.SourceSelectorEntry")
	proto.RegisterType((*Destination)(nil), "registry.Destination")
	proto.RegisterMapType((map[string]string)(nil), "registry.Destination.DestinationSelectorEntry")
	proto.RegisterType((*NetworkServiceManager)(nil), "registry.NetworkServiceManager")
	proto.RegisterType((*RemoveNSERequest)(nil), "registry.RemoveNSERequest")
	proto.RegisterType((*FindNetworkServiceRequest)(nil), "registry.FindNetworkServiceRequest")
	proto.RegisterType((*FindNetworkServiceResponse)(nil), "registry.FindNetworkServiceResponse")
	proto.RegisterMapType((map[string]*NetworkServiceManager)(nil), "registry.FindNetworkServiceResponse.NetworkServiceManagersEntry")
	proto.RegisterType((*NSERegistration)(nil), "registry.NSERegistration")
	proto.RegisterType((*NetworkServiceEndpointList)(nil), "registry.NetworkServiceEndpointList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceRegistryClient is the client API for NetworkServiceRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceRegistryClient interface {
	RegisterNSE(ctx context.Context, in *NSERegistration, opts ...grpc.CallOption) (*NSERegistration, error)
	RemoveNSE(ctx context.Context, in *RemoveNSERequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceRegistryClient(cc *grpc.ClientConn) NetworkServiceRegistryClient {
	return &networkServiceRegistryClient{cc}
}

func (c *networkServiceRegistryClient) RegisterNSE(ctx context.Context, in *NSERegistration, opts ...grpc.CallOption) (*NSERegistration, error) {
	out := new(NSERegistration)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceRegistry/RegisterNSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceRegistryClient) RemoveNSE(ctx context.Context, in *RemoveNSERequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceRegistry/RemoveNSE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceRegistryServer is the server API for NetworkServiceRegistry service.
type NetworkServiceRegistryServer interface {
	RegisterNSE(context.Context, *NSERegistration) (*NSERegistration, error)
	RemoveNSE(context.Context, *RemoveNSERequest) (*empty.Empty, error)
}

func RegisterNetworkServiceRegistryServer(s *grpc.Server, srv NetworkServiceRegistryServer) {
	s.RegisterService(&_NetworkServiceRegistry_serviceDesc, srv)
}

func _NetworkServiceRegistry_RegisterNSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NSERegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceRegistryServer).RegisterNSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceRegistry/RegisterNSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceRegistryServer).RegisterNSE(ctx, req.(*NSERegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServiceRegistry_RemoveNSE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNSERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceRegistryServer).RemoveNSE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceRegistry/RemoveNSE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceRegistryServer).RemoveNSE(ctx, req.(*RemoveNSERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServiceRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.NetworkServiceRegistry",
	HandlerType: (*NetworkServiceRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNSE",
			Handler:    _NetworkServiceRegistry_RegisterNSE_Handler,
		},
		{
			MethodName: "RemoveNSE",
			Handler:    _NetworkServiceRegistry_RemoveNSE_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

// NetworkServiceDiscoveryClient is the client API for NetworkServiceDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceDiscoveryClient interface {
	FindNetworkService(ctx context.Context, in *FindNetworkServiceRequest, opts ...grpc.CallOption) (*FindNetworkServiceResponse, error)
}

type networkServiceDiscoveryClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceDiscoveryClient(cc *grpc.ClientConn) NetworkServiceDiscoveryClient {
	return &networkServiceDiscoveryClient{cc}
}

func (c *networkServiceDiscoveryClient) FindNetworkService(ctx context.Context, in *FindNetworkServiceRequest, opts ...grpc.CallOption) (*FindNetworkServiceResponse, error) {
	out := new(FindNetworkServiceResponse)
	err := c.cc.Invoke(ctx, "/registry.NetworkServiceDiscovery/FindNetworkService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceDiscoveryServer is the server API for NetworkServiceDiscovery service.
type NetworkServiceDiscoveryServer interface {
	FindNetworkService(context.Context, *FindNetworkServiceRequest) (*FindNetworkServiceResponse, error)
}

func RegisterNetworkServiceDiscoveryServer(s *grpc.Server, srv NetworkServiceDiscoveryServer) {
	s.RegisterService(&_NetworkServiceDiscovery_serviceDesc, srv)
}

func _NetworkServiceDiscovery_FindNetworkService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceDiscoveryServer).FindNetworkService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NetworkServiceDiscovery/FindNetworkService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceDiscoveryServer).FindNetworkService(ctx, req.(*FindNetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServiceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.NetworkServiceDiscovery",
	HandlerType: (*NetworkServiceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindNetworkService",
			Handler:    _NetworkServiceDiscovery_FindNetworkService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

// NsmRegistryClient is the client API for NsmRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsmRegistryClient interface {
	RegisterNSM(ctx context.Context, in *NetworkServiceManager, opts ...grpc.CallOption) (*NetworkServiceManager, error)
	GetEndpoints(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkServiceEndpointList, error)
}

type nsmRegistryClient struct {
	cc *grpc.ClientConn
}

func NewNsmRegistryClient(cc *grpc.ClientConn) NsmRegistryClient {
	return &nsmRegistryClient{cc}
}

func (c *nsmRegistryClient) RegisterNSM(ctx context.Context, in *NetworkServiceManager, opts ...grpc.CallOption) (*NetworkServiceManager, error) {
	out := new(NetworkServiceManager)
	err := c.cc.Invoke(ctx, "/registry.NsmRegistry/RegisterNSM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nsmRegistryClient) GetEndpoints(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NetworkServiceEndpointList, error) {
	out := new(NetworkServiceEndpointList)
	err := c.cc.Invoke(ctx, "/registry.NsmRegistry/GetEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsmRegistryServer is the server API for NsmRegistry service.
type NsmRegistryServer interface {
	RegisterNSM(context.Context, *NetworkServiceManager) (*NetworkServiceManager, error)
	GetEndpoints(context.Context, *empty.Empty) (*NetworkServiceEndpointList, error)
}

func RegisterNsmRegistryServer(s *grpc.Server, srv NsmRegistryServer) {
	s.RegisterService(&_NsmRegistry_serviceDesc, srv)
}

func _NsmRegistry_RegisterNSM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceManager)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsmRegistryServer).RegisterNSM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NsmRegistry/RegisterNSM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsmRegistryServer).RegisterNSM(ctx, req.(*NetworkServiceManager))
	}
	return interceptor(ctx, in, info, handler)
}

func _NsmRegistry_GetEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsmRegistryServer).GetEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registry.NsmRegistry/GetEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsmRegistryServer).GetEndpoints(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsmRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registry.NsmRegistry",
	HandlerType: (*NsmRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNSM",
			Handler:    _NsmRegistry_RegisterNSM_Handler,
		},
		{
			MethodName: "GetEndpoints",
			Handler:    _NsmRegistry_GetEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry.proto",
}

func init() { proto.RegisterFile("registry.proto", fileDescriptor_registry_f962b46a720ff8bc) }

var fileDescriptor_registry_f962b46a720ff8bc = []byte{
	// 824 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6e, 0xfa, 0x46,
	0x10, 0x96, 0x81, 0x90, 0x64, 0x9c, 0x40, 0xb4, 0x49, 0x88, 0x71, 0x2b, 0x35, 0x22, 0x3d, 0xa4,
	0x52, 0x6b, 0x2a, 0xa2, 0x2a, 0xfd, 0x73, 0x48, 0xa3, 0x86, 0xf4, 0x02, 0x54, 0x32, 0x95, 0xaa,
	0x4a, 0x95, 0x90, 0x81, 0xa9, 0xe3, 0xc6, 0xf6, 0xba, 0xde, 0x85, 0xc8, 0x79, 0x82, 0x1e, 0xfa,
	0x0c, 0x7d, 0x82, 0xbe, 0x43, 0x8f, 0xbd, 0xf6, 0x15, 0xfa, 0x26, 0x3f, 0xb1, 0xb6, 0xc1, 0x36,
	0x36, 0x04, 0xfd, 0x2e, 0xd1, 0xae, 0x77, 0xe6, 0x9b, 0x99, 0x6f, 0xe6, 0x9b, 0x00, 0x35, 0x1f,
	0x4d, 0x8b, 0x71, 0x3f, 0xd0, 0x3c, 0x9f, 0x72, 0x4a, 0x0e, 0xe2, 0xbb, 0x7a, 0x63, 0x5a, 0xfc,
	0x69, 0x36, 0xd6, 0x26, 0xd4, 0x69, 0x9b, 0xd4, 0x36, 0x5c, 0xb3, 0x2d, 0x4c, 0xc6, 0xb3, 0x5f,
	0xdb, 0x1e, 0x0f, 0x3c, 0x64, 0x6d, 0x74, 0x3c, 0x1e, 0x84, 0x7f, 0x43, 0x77, 0xf5, 0x9b, 0xed,
	0x4e, 0xdc, 0x72, 0x90, 0x71, 0xc3, 0xf1, 0x56, 0xa7, 0xd0, 0xb9, 0xf5, 0x7f, 0x09, 0x1a, 0x03,
	0xe4, 0x2f, 0xd4, 0x7f, 0x1e, 0xa2, 0x3f, 0xb7, 0x26, 0xd8, 0x75, 0xa7, 0x1e, 0xb5, 0x5c, 0x4e,
	0x3e, 0x87, 0x33, 0x37, 0x7c, 0x19, 0xb1, 0xf0, 0x69, 0xe4, 0x1a, 0x0e, 0x2a, 0xd2, 0xa5, 0x74,
	0x7d, 0xa8, 0x13, 0x37, 0xe5, 0x35, 0x30, 0x1c, 0x24, 0x0a, 0xec, 0x7b, 0x46, 0x60, 0x53, 0x63,
	0xaa, 0x94, 0x84, 0x51, 0x7c, 0x25, 0x77, 0xf0, 0x61, 0x16, 0xcb, 0x31, 0x5c, 0xc3, 0x44, 0x3f,
	0xc4, 0x2c, 0x0b, 0xf3, 0x66, 0x1a, 0xb3, 0x1f, 0x5a, 0x08, 0xe8, 0x2b, 0x38, 0xc6, 0x28, 0xb1,
	0xd0, 0xa3, 0x22, 0x3c, 0x8e, 0xe2, 0x8f, 0xc2, 0xe8, 0x01, 0xaa, 0xb6, 0x31, 0x46, 0x9b, 0x29,
	0x7b, 0x97, 0xe5, 0x6b, 0xb9, 0xf3, 0xa9, 0xb6, 0x64, 0x3a, 0xbf, 0x46, 0xad, 0x27, 0xcc, 0xbb,
	0x2e, 0xf7, 0x03, 0x3d, 0xf2, 0x25, 0x67, 0xb0, 0xc7, 0xb8, 0xc1, 0x51, 0xa9, 0x8a, 0x10, 0xe1,
	0x45, 0xfd, 0x0a, 0xe4, 0x84, 0x31, 0x39, 0x81, 0xf2, 0x33, 0x06, 0x11, 0x17, 0x8b, 0xe3, 0xc2,
	0x6d, 0x6e, 0xd8, 0x33, 0x8c, 0x4a, 0x0f, 0x2f, 0x5f, 0x97, 0xbe, 0x94, 0x5a, 0x16, 0xd4, 0xd2,
	0xe1, 0x09, 0x81, 0x4a, 0x82, 0x4a, 0x71, 0xde, 0x40, 0xde, 0x27, 0xb0, 0xef, 0x18, 0x7c, 0xf2,
	0x84, 0x4c, 0x29, 0x8b, 0xba, 0xea, 0xab, 0xba, 0xfa, 0x8b, 0x07, 0x3d, 0x7e, 0x6f, 0xfd, 0x2b,
	0xc1, 0x9e, 0xf8, 0x44, 0x7a, 0x50, 0x67, 0x74, 0xe6, 0x4f, 0x70, 0xc4, 0xd0, 0xc6, 0x09, 0xa7,
	0xbe, 0x22, 0x09, 0xe7, 0xab, 0x8c, 0xb3, 0x36, 0x14, 0x66, 0xc3, 0xc8, 0x2a, 0xe4, 0xa2, 0xc6,
	0x52, 0x1f, 0xc9, 0x67, 0x50, 0xf5, 0xe9, 0x8c, 0x23, 0x53, 0x4a, 0x02, 0xe4, 0x7c, 0x05, 0xf2,
	0x80, 0x8c, 0x5b, 0xae, 0xc1, 0x2d, 0xea, 0xea, 0x91, 0x91, 0x7a, 0x0f, 0xa7, 0x39, 0xa8, 0x3b,
	0x91, 0xf6, 0x9f, 0x04, 0x72, 0x02, 0x9a, 0x18, 0x70, 0x36, 0x5d, 0x5d, 0xb3, 0x45, 0x69, 0xb9,
	0xf9, 0x24, 0xcf, 0xe9, 0xfa, 0x4e, 0xa7, 0xeb, 0x2f, 0xa4, 0x01, 0xd5, 0x17, 0xb4, 0xcc, 0x27,
	0x2e, 0xb2, 0x39, 0xd6, 0xa3, 0x9b, 0xfa, 0x08, 0x4a, 0x11, 0xd0, 0x4e, 0x25, 0xfd, 0x29, 0xc1,
	0xf9, 0x20, 0x6f, 0xc2, 0x73, 0xe7, 0xe1, 0x04, 0xca, 0x33, 0xdf, 0x8e, 0x50, 0x16, 0x47, 0x72,
	0x0b, 0x87, 0xb6, 0xc1, 0xf8, 0x88, 0x21, 0xba, 0x42, 0x31, 0x72, 0x47, 0xd5, 0x4c, 0x4a, 0x4d,
	0x1b, 0xb5, 0x58, 0xf1, 0xda, 0x8f, 0xb1, 0xc0, 0xf5, 0x83, 0x85, 0xf1, 0x10, 0xd1, 0x5d, 0x4d,
	0x74, 0x25, 0x31, 0xd1, 0xad, 0x5b, 0x38, 0xd1, 0xd1, 0xa1, 0x73, 0x1c, 0x0c, 0xbb, 0x3a, 0xfe,
	0x3e, 0x43, 0xc6, 0xd7, 0x65, 0x26, 0xad, 0xcb, 0xac, 0xd5, 0x87, 0xe6, 0xa3, 0xe5, 0x4e, 0xd3,
	0xa5, 0xc4, 0x08, 0x3b, 0x6f, 0x8d, 0xd6, 0x3f, 0x65, 0x50, 0xf3, 0xf0, 0x98, 0x47, 0x5d, 0x96,
	0xd2, 0x85, 0x94, 0xd6, 0xc5, 0x3d, 0xd4, 0x33, 0xa1, 0x04, 0x5b, 0x72, 0x47, 0x29, 0xd2, 0xbd,
	0x5e, 0x4b, 0xc7, 0x27, 0xaf, 0xa0, 0x14, 0xec, 0xa5, 0x58, 0x6b, 0xdf, 0xae, 0xb0, 0x8a, 0x93,
	0xd4, 0x72, 0xdb, 0x1a, 0xed, 0x95, 0x46, 0xee, 0x56, 0x63, 0xe4, 0x17, 0x68, 0x66, 0x63, 0xc7,
	0x34, 0x33, 0xa5, 0x22, 0x82, 0x5f, 0x6e, 0x5b, 0x60, 0xfa, 0x85, 0x9b, 0xfb, 0x9d, 0xa9, 0xbf,
	0xc1, 0x07, 0x1b, 0x92, 0xca, 0x99, 0xdb, 0x2f, 0x92, 0x73, 0x2b, 0x77, 0x3e, 0x2a, 0x0a, 0x1d,
	0xe1, 0x24, 0x07, 0xfb, 0x8f, 0x12, 0xd4, 0xc5, 0x10, 0x09, 0x87, 0x50, 0xaf, 0x39, 0xcd, 0x91,
	0x76, 0x6c, 0xce, 0x4f, 0x70, 0x51, 0xd0, 0x9c, 0xb7, 0xe6, 0x78, 0x9e, 0x4b, 0x3d, 0xf9, 0x79,
	0x09, 0x9c, 0x25, 0x3e, 0x92, 0xd5, 0x76, 0xde, 0x1b, 0x69, 0x80, 0xf8, 0x7b, 0xeb, 0x15, 0xd4,
	0x7c, 0x8f, 0x9e, 0xc5, 0xf8, 0xe6, 0x96, 0x4b, 0xef, 0xd9, 0xf2, 0xce, 0x5f, 0x52, 0xf6, 0x7f,
	0x79, 0xd4, 0x91, 0x80, 0x7c, 0x07, 0x72, 0x78, 0x46, 0x7f, 0x30, 0xec, 0x92, 0x66, 0x22, 0x48,
	0xba, 0x6f, 0x6a, 0xf1, 0x13, 0xb9, 0x83, 0xc3, 0xe5, 0xc2, 0x20, 0xea, 0xca, 0x2e, 0xbb, 0x45,
	0xd4, 0xc6, 0xda, 0x56, 0xea, 0x2e, 0x7e, 0xaf, 0x74, 0x5e, 0xe1, 0x22, 0x9d, 0xdf, 0x83, 0xc5,
	0x26, 0x74, 0x8e, 0x7e, 0x40, 0x46, 0x40, 0xd6, 0xe5, 0x45, 0xae, 0x36, 0x8b, 0x2f, 0x8c, 0xf6,
	0xf1, 0x5b, 0x14, 0xda, 0xf9, 0x5b, 0x02, 0x79, 0xc0, 0x9c, 0x25, 0x23, 0x3f, 0x24, 0x19, 0xe9,
	0x93, 0x6d, 0xa3, 0xa4, 0x6e, 0x33, 0x20, 0x3d, 0x38, 0xfa, 0x1e, 0xf9, 0xb2, 0x1b, 0xa4, 0x80,
	0x84, 0x64, 0xba, 0xc5, 0x93, 0x32, 0xae, 0x0a, 0xaf, 0x9b, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0xfb, 0x60, 0x1f, 0x2c, 0x0a, 0x00, 0x00,
}
