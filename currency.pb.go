// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currency.proto

package restaurant

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Currency struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string               `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc60ed002193ea, []int{0}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Currency) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Currency) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Currency)(nil), "restaurant.Currency")
}

func init() {
	proto.RegisterFile("currency.proto", fileDescriptor_d3dc60ed002193ea)
}

var fileDescriptor_d3dc60ed002193ea = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0xfc, 0xfb, 0x2f, 0x76, 0x82, 0x11, 0x17, 0x91, 0x90, 0x8b, 0xc5, 0x53, 0x4f,
	0xdb, 0xd2, 0x82, 0xa0, 0xb7, 0x20, 0x22, 0x5e, 0xa3, 0x9e, 0x65, 0x93, 0x1d, 0xcb, 0x42, 0x37,
	0x1b, 0x36, 0x93, 0x42, 0xde, 0xcd, 0x47, 0xf2, 0x21, 0xa4, 0x9b, 0xc4, 0x16, 0x15, 0xc4, 0xdb,
	0xce, 0x37, 0xdf, 0x6f, 0xf6, 0x9b, 0x81, 0xa8, 0x68, 0x9c, 0xc3, 0xb2, 0x68, 0x45, 0xe5, 0x2c,
	0x59, 0x0e, 0x0e, 0x6b, 0x92, 0x8d, 0x93, 0x25, 0x25, 0x17, 0x6b, 0x6b, 0xd7, 0x1b, 0x9c, 0xfb,
	0x4e, 0xde, 0xbc, 0xce, 0x49, 0x9b, 0x5d, 0xd7, 0x54, 0x9d, 0x39, 0x09, 0x73, 0x59, 0xeb, 0xa2,
	0x2b, 0x2e, 0xdf, 0x18, 0x1c, 0xdd, 0xf6, 0xc3, 0x78, 0x04, 0x81, 0x56, 0x31, 0x9b, 0xb2, 0xd9,
	0x24, 0x0b, 0xb4, 0xe2, 0x1c, 0x46, 0xa5, 0x34, 0x18, 0x07, 0x5e, 0xf1, 0x6f, 0x7e, 0x0e, 0xe3,
	0xba, 0x35, 0xb9, 0xdd, 0xc4, 0xff, 0xbc, 0xda, 0x57, 0xfc, 0x1a, 0xa0, 0x70, 0x28, 0x09, 0xd5,
	0x8b, 0xa4, 0x78, 0x34, 0x65, 0xb3, 0x70, 0x99, 0x88, 0x2e, 0x8b, 0x18, 0xb2, 0x88, 0xa7, 0x21,
	0x4b, 0x36, 0xe9, 0xdd, 0x29, 0xed, 0xd0, 0xa6, 0x52, 0x03, 0xfa, 0xff, 0x77, 0xb4, 0x77, 0xa7,
	0xb4, 0x7c, 0x67, 0x70, 0x32, 0xc4, 0x7f, 0x44, 0xb7, 0xd5, 0x05, 0xf2, 0x15, 0x84, 0xa9, 0x52,
	0x9f, 0x4b, 0x9d, 0x89, 0xfd, 0x71, 0xc4, 0xa0, 0x26, 0xd1, 0xa1, 0xfa, 0xa0, 0xf8, 0x15, 0x44,
	0xcf, 0x7e, 0xea, 0x1f, 0xb9, 0x05, 0x44, 0x19, 0x1a, 0xbb, 0xdd, 0x73, 0x5f, 0x1c, 0xdf, 0x88,
	0x1b, 0x38, 0xbe, 0x47, 0xea, 0xed, 0x1a, 0x6b, 0x7e, 0x7a, 0x68, 0xb8, 0x33, 0x15, 0xb5, 0xc9,
	0x8f, 0x7f, 0x2f, 0x58, 0x3e, 0xf6, 0xd7, 0x58, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x35, 0x32,
	0x16, 0x06, 0x00, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	// Adds a new Curreny
	AddCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Id, error)
	// Updates current Currency
	UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Id, error)
	// Remove Curreny if not used by any restaurant
	RemoveCurrency(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error)
	// Get All the Currecny
	GetCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CurrencyService_GetCurrenciesClient, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) AddCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/restaurant.CurrencyService/AddCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/restaurant.CurrencyService/UpdateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) RemoveCurrency(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/restaurant.CurrencyService/RemoveCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) GetCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CurrencyService_GetCurrenciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CurrencyService_serviceDesc.Streams[0], "/restaurant.CurrencyService/GetCurrencies", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyServiceGetCurrenciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyService_GetCurrenciesClient interface {
	Recv() (*Currency, error)
	grpc.ClientStream
}

type currencyServiceGetCurrenciesClient struct {
	grpc.ClientStream
}

func (x *currencyServiceGetCurrenciesClient) Recv() (*Currency, error) {
	m := new(Currency)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
type CurrencyServiceServer interface {
	// Adds a new Curreny
	AddCurrency(context.Context, *Currency) (*Id, error)
	// Updates current Currency
	UpdateCurrency(context.Context, *Currency) (*Id, error)
	// Remove Curreny if not used by any restaurant
	RemoveCurrency(context.Context, *Id) (*Id, error)
	// Get All the Currecny
	GetCurrencies(*Empty, CurrencyService_GetCurrenciesServer) error
}

// UnimplementedCurrencyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (*UnimplementedCurrencyServiceServer) AddCurrency(ctx context.Context, req *Currency) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) UpdateCurrency(ctx context.Context, req *Currency) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) RemoveCurrency(ctx context.Context, req *Id) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCurrency not implemented")
}
func (*UnimplementedCurrencyServiceServer) GetCurrencies(req *Empty, srv CurrencyService_GetCurrenciesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}

func RegisterCurrencyServiceServer(s *grpc.Server, srv CurrencyServiceServer) {
	s.RegisterService(&_CurrencyService_serviceDesc, srv)
}

func _CurrencyService_AddCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).AddCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.CurrencyService/AddCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).AddCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_UpdateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.CurrencyService/UpdateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_RemoveCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).RemoveCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.CurrencyService/RemoveCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).RemoveCurrency(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_GetCurrencies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyServiceServer).GetCurrencies(m, &currencyServiceGetCurrenciesServer{stream})
}

type CurrencyService_GetCurrenciesServer interface {
	Send(*Currency) error
	grpc.ServerStream
}

type currencyServiceGetCurrenciesServer struct {
	grpc.ServerStream
}

func (x *currencyServiceGetCurrenciesServer) Send(m *Currency) error {
	return x.ServerStream.SendMsg(m)
}

var _CurrencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCurrency",
			Handler:    _CurrencyService_AddCurrency_Handler,
		},
		{
			MethodName: "UpdateCurrency",
			Handler:    _CurrencyService_UpdateCurrency_Handler,
		},
		{
			MethodName: "RemoveCurrency",
			Handler:    _CurrencyService_RemoveCurrency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCurrencies",
			Handler:       _CurrencyService_GetCurrencies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "currency.proto",
}
