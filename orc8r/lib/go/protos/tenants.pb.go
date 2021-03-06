// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/tenants.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetTenantRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTenantRequest) Reset()         { *m = GetTenantRequest{} }
func (m *GetTenantRequest) String() string { return proto.CompactTextString(m) }
func (*GetTenantRequest) ProtoMessage()    {}
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fa0ed566a19fde, []int{0}
}

func (m *GetTenantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTenantRequest.Unmarshal(m, b)
}
func (m *GetTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTenantRequest.Marshal(b, m, deterministic)
}
func (m *GetTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTenantRequest.Merge(m, src)
}
func (m *GetTenantRequest) XXX_Size() int {
	return xxx_messageInfo_GetTenantRequest.Size(m)
}
func (m *GetTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTenantRequest proto.InternalMessageInfo

func (m *GetTenantRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Tenant struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Networks             []string `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tenant) Reset()         { *m = Tenant{} }
func (m *Tenant) String() string { return proto.CompactTextString(m) }
func (*Tenant) ProtoMessage()    {}
func (*Tenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fa0ed566a19fde, []int{1}
}

func (m *Tenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tenant.Unmarshal(m, b)
}
func (m *Tenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tenant.Marshal(b, m, deterministic)
}
func (m *Tenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tenant.Merge(m, src)
}
func (m *Tenant) XXX_Size() int {
	return xxx_messageInfo_Tenant.Size(m)
}
func (m *Tenant) XXX_DiscardUnknown() {
	xxx_messageInfo_Tenant.DiscardUnknown(m)
}

var xxx_messageInfo_Tenant proto.InternalMessageInfo

func (m *Tenant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tenant) GetNetworks() []string {
	if m != nil {
		return m.Networks
	}
	return nil
}

type TenantList struct {
	Tenants              []*IDAndTenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TenantList) Reset()         { *m = TenantList{} }
func (m *TenantList) String() string { return proto.CompactTextString(m) }
func (*TenantList) ProtoMessage()    {}
func (*TenantList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fa0ed566a19fde, []int{2}
}

func (m *TenantList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TenantList.Unmarshal(m, b)
}
func (m *TenantList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TenantList.Marshal(b, m, deterministic)
}
func (m *TenantList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TenantList.Merge(m, src)
}
func (m *TenantList) XXX_Size() int {
	return xxx_messageInfo_TenantList.Size(m)
}
func (m *TenantList) XXX_DiscardUnknown() {
	xxx_messageInfo_TenantList.DiscardUnknown(m)
}

var xxx_messageInfo_TenantList proto.InternalMessageInfo

func (m *TenantList) GetTenants() []*IDAndTenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

type IDAndTenant struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tenant               *Tenant  `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDAndTenant) Reset()         { *m = IDAndTenant{} }
func (m *IDAndTenant) String() string { return proto.CompactTextString(m) }
func (*IDAndTenant) ProtoMessage()    {}
func (*IDAndTenant) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fa0ed566a19fde, []int{3}
}

func (m *IDAndTenant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDAndTenant.Unmarshal(m, b)
}
func (m *IDAndTenant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDAndTenant.Marshal(b, m, deterministic)
}
func (m *IDAndTenant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDAndTenant.Merge(m, src)
}
func (m *IDAndTenant) XXX_Size() int {
	return xxx_messageInfo_IDAndTenant.Size(m)
}
func (m *IDAndTenant) XXX_DiscardUnknown() {
	xxx_messageInfo_IDAndTenant.DiscardUnknown(m)
}

var xxx_messageInfo_IDAndTenant proto.InternalMessageInfo

func (m *IDAndTenant) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IDAndTenant) GetTenant() *Tenant {
	if m != nil {
		return m.Tenant
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTenantRequest)(nil), "magma.orc8r.GetTenantRequest")
	proto.RegisterType((*Tenant)(nil), "magma.orc8r.Tenant")
	proto.RegisterType((*TenantList)(nil), "magma.orc8r.TenantList")
	proto.RegisterType((*IDAndTenant)(nil), "magma.orc8r.IDAndTenant")
}

func init() { proto.RegisterFile("orc8r/protos/tenants.proto", fileDescriptor_e7fa0ed566a19fde) }

var fileDescriptor_e7fa0ed566a19fde = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x9b, 0xf4, 0x4f, 0xff, 0x66, 0x52, 0x8b, 0x8e, 0x07, 0xd3, 0x88, 0x10, 0xf6, 0x14,
	0x10, 0x12, 0x88, 0x97, 0x22, 0x14, 0x6c, 0x2d, 0x14, 0xc5, 0x53, 0x2a, 0x1e, 0xbc, 0xa5, 0xcd,
	0x50, 0x82, 0x49, 0x56, 0x93, 0x55, 0x3f, 0xa7, 0xdf, 0x48, 0xdc, 0x5d, 0x43, 0x53, 0x2b, 0xe2,
	0x29, 0x99, 0x97, 0xdf, 0xbc, 0xbc, 0x19, 0x06, 0x5c, 0x5e, 0xad, 0x46, 0x55, 0xf8, 0x54, 0x71,
	0xc1, 0xeb, 0x50, 0x50, 0x99, 0x94, 0xa2, 0x0e, 0x64, 0x89, 0x76, 0x91, 0xac, 0x8b, 0x24, 0x90,
	0x84, 0x3b, 0x6c, 0x81, 0x2b, 0x5e, 0x14, 0xbc, 0x54, 0x1c, 0x63, 0x70, 0x30, 0x27, 0x71, 0x27,
	0x7b, 0x63, 0x7a, 0x7e, 0xa1, 0x5a, 0xe0, 0x00, 0xcc, 0x2c, 0x75, 0x0c, 0xcf, 0xf0, 0xbb, 0xb1,
	0x99, 0xa5, 0x6c, 0x04, 0x3d, 0x05, 0x20, 0xc2, 0xbf, 0x32, 0x29, 0x48, 0x7e, 0xb3, 0x62, 0xf9,
	0x8e, 0x2e, 0xec, 0x95, 0x24, 0xde, 0x78, 0xf5, 0x58, 0x3b, 0xa6, 0xd7, 0xf5, 0xad, 0xb8, 0xa9,
	0xd9, 0x25, 0x80, 0xea, 0xbc, 0xcd, 0x6a, 0x81, 0x11, 0xfc, 0xd7, 0x21, 0x1d, 0xc3, 0xeb, 0xfa,
	0x76, 0xe4, 0x04, 0x1b, 0x29, 0x83, 0xeb, 0xd9, 0xa4, 0x4c, 0x75, 0x92, 0x2f, 0x90, 0xdd, 0x80,
	0xbd, 0xa1, 0x6f, 0x47, 0xc3, 0x33, 0xe8, 0x29, 0xd2, 0x31, 0x3d, 0xc3, 0xb7, 0xa3, 0xa3, 0x96,
	0xa3, 0x36, 0xd3, 0x48, 0xf4, 0x6e, 0xc2, 0x40, 0x49, 0xf5, 0x82, 0xaa, 0xd7, 0x6c, 0x45, 0x38,
	0x86, 0xfd, 0x39, 0x89, 0x49, 0x9e, 0x6b, 0x1d, 0x0f, 0x5b, 0x06, 0xf7, 0x3c, 0x4b, 0xdd, 0xe3,
	0x1d, 0x9e, 0x9f, 0xf3, 0xb0, 0x0e, 0x4e, 0xc0, 0x6a, 0xb6, 0x87, 0xa7, 0x2d, 0x6e, 0x7b, 0xab,
	0xee, 0xae, 0x68, 0xac, 0x83, 0x63, 0xe8, 0x5f, 0x55, 0x94, 0x08, 0xd2, 0x2e, 0x3f, 0xee, 0xc4,
	0xfd, 0x1e, 0x8d, 0x75, 0xf0, 0x02, 0xac, 0x45, 0x93, 0xe0, 0x8f, 0xbd, 0x53, 0xe8, 0xcf, 0x28,
	0xa7, 0xe6, 0xd7, 0xbf, 0x0c, 0xb0, 0xcb, 0x63, 0x7a, 0xf2, 0x30, 0x94, 0x6a, 0xa8, 0x4e, 0x2c,
	0xcf, 0x96, 0xe1, 0x9a, 0xeb, 0x4b, 0x5b, 0xf6, 0xe4, 0xf3, 0xfc, 0x23, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x90, 0xe8, 0xaa, 0xa9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TenantsServiceClient is the client API for TenantsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantsServiceClient interface {
	GetAllTenants(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TenantList, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*Void, error)
	SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*Void, error)
	DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Void, error)
}

type tenantsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantsServiceClient(cc grpc.ClientConnInterface) TenantsServiceClient {
	return &tenantsServiceClient{cc}
}

func (c *tenantsServiceClient) GetAllTenants(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TenantList, error) {
	out := new(TenantList)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetAllTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) CreateTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) SetTenant(ctx context.Context, in *IDAndTenant, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/SetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantsServiceClient) DeleteTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/magma.orc8r.TenantsService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantsServiceServer is the server API for TenantsService service.
type TenantsServiceServer interface {
	GetAllTenants(context.Context, *Void) (*TenantList, error)
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	CreateTenant(context.Context, *IDAndTenant) (*Void, error)
	SetTenant(context.Context, *IDAndTenant) (*Void, error)
	DeleteTenant(context.Context, *GetTenantRequest) (*Void, error)
}

// UnimplementedTenantsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantsServiceServer struct {
}

func (*UnimplementedTenantsServiceServer) GetAllTenants(ctx context.Context, req *Void) (*TenantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTenants not implemented")
}
func (*UnimplementedTenantsServiceServer) GetTenant(ctx context.Context, req *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) CreateTenant(ctx context.Context, req *IDAndTenant) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) SetTenant(ctx context.Context, req *IDAndTenant) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTenant not implemented")
}
func (*UnimplementedTenantsServiceServer) DeleteTenant(ctx context.Context, req *GetTenantRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}

func RegisterTenantsServiceServer(s *grpc.Server, srv TenantsServiceServer) {
	s.RegisterService(&_TenantsService_serviceDesc, srv)
}

func _TenantsService_GetAllTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetAllTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetAllTenants(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).CreateTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_SetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDAndTenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).SetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/SetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).SetTenant(ctx, req.(*IDAndTenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantsService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.TenantsService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantsServiceServer).DeleteTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.TenantsService",
	HandlerType: (*TenantsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTenants",
			Handler:    _TenantsService_GetAllTenants_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantsService_GetTenant_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _TenantsService_CreateTenant_Handler,
		},
		{
			MethodName: "SetTenant",
			Handler:    _TenantsService_SetTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantsService_DeleteTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/tenants.proto",
}
