// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ProductInfo.proto

// pacakage names

package proto

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

//Definition of the Product Message
type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_0353433997bc450f, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProductID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductID) Reset()         { *m = ProductID{} }
func (m *ProductID) String() string { return proto.CompactTextString(m) }
func (*ProductID) ProtoMessage()    {}
func (*ProductID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0353433997bc450f, []int{1}
}

func (m *ProductID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductID.Unmarshal(m, b)
}
func (m *ProductID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductID.Marshal(b, m, deterministic)
}
func (m *ProductID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductID.Merge(m, src)
}
func (m *ProductID) XXX_Size() int {
	return xxx_messageInfo_ProductID.Size(m)
}
func (m *ProductID) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductID.DiscardUnknown(m)
}

var xxx_messageInfo_ProductID proto.InternalMessageInfo

func (m *ProductID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Product)(nil), "proto.Product")
	proto.RegisterType((*ProductID)(nil), "proto.ProductID")
}

func init() {
	proto.RegisterFile("proto/ProductInfo.proto", fileDescriptor_0353433997bc450f)
}

var fileDescriptor_0353433997bc450f = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x0f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x2e, 0xf1, 0xcc, 0x4b, 0xcb, 0xd7, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0x29, 0x25, 0x7f, 0x2e, 0x76, 0xa8, 0x9c, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62,
	0x6e, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x16, 0x52, 0xe0, 0xe2, 0x4e, 0x49, 0x2d, 0x4e, 0x2e,
	0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x93, 0x60, 0x06, 0x4b, 0x21, 0x0b, 0x29, 0x29, 0x72, 0x71,
	0xc2, 0x2c, 0x73, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x85, 0x9a, 0x0a, 0xe1,
	0x18, 0xe5, 0x72, 0x71, 0x23, 0xb9, 0x47, 0x48, 0x8f, 0x8b, 0x2b, 0x31, 0x25, 0x05, 0xee, 0x0a,
	0x88, 0xfb, 0xf4, 0xa0, 0x7c, 0x29, 0x01, 0x54, 0xbe, 0xa7, 0x0b, 0x48, 0x7d, 0x7a, 0x6a, 0x09,
	0x4c, 0x3d, 0x86, 0xbc, 0x14, 0x9a, 0x09, 0x49, 0x6c, 0x60, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x7f, 0x2c, 0x95, 0x0b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductInfoClient interface {
	AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error)
	GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error)
}

type productInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInfoClient(cc grpc.ClientConnInterface) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) AddProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProductID, error) {
	out := new(ProductID)
	err := c.cc.Invoke(ctx, "/proto.ProductInfo/addProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productInfoClient) GetProduct(ctx context.Context, in *ProductID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/proto.ProductInfo/getProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
type ProductInfoServer interface {
	AddProduct(context.Context, *Product) (*ProductID, error)
	GetProduct(context.Context, *ProductID) (*Product, error)
}

// UnimplementedProductInfoServer can be embedded to have forward compatible implementations.
type UnimplementedProductInfoServer struct {
}

func (*UnimplementedProductInfoServer) AddProduct(ctx context.Context, req *Product) (*ProductID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedProductInfoServer) GetProduct(ctx context.Context, req *ProductID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}

func RegisterProductInfoServer(s *grpc.Server, srv ProductInfoServer) {
	s.RegisterService(&_ProductInfo_serviceDesc, srv)
}

func _ProductInfo_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductInfo/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).AddProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductInfo_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductInfo/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).GetProduct(ctx, req.(*ProductID))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addProduct",
			Handler:    _ProductInfo_AddProduct_Handler,
		},
		{
			MethodName: "getProduct",
			Handler:    _ProductInfo_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ProductInfo.proto",
}