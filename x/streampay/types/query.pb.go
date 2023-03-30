// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streampay/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryStreamPaymentsRequest struct {
	Sender     string             `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient  string             `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryStreamPaymentsRequest) Reset()         { *m = QueryStreamPaymentsRequest{} }
func (m *QueryStreamPaymentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStreamPaymentsRequest) ProtoMessage()    {}
func (*QueryStreamPaymentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec493855485c26d, []int{0}
}
func (m *QueryStreamPaymentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStreamPaymentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStreamPaymentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStreamPaymentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStreamPaymentsRequest.Merge(m, src)
}
func (m *QueryStreamPaymentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStreamPaymentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStreamPaymentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStreamPaymentsRequest proto.InternalMessageInfo

func (m *QueryStreamPaymentsRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *QueryStreamPaymentsRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *QueryStreamPaymentsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryStreamPaymentsResponse struct {
	StreamPayments []StreamPayment     `protobuf:"bytes,1,rep,name=stream_payments,json=streamPayments,proto3" json:"stream_payments" yaml:"stream_payments"`
	Pagination     *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryStreamPaymentsResponse) Reset()         { *m = QueryStreamPaymentsResponse{} }
func (m *QueryStreamPaymentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStreamPaymentsResponse) ProtoMessage()    {}
func (*QueryStreamPaymentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec493855485c26d, []int{1}
}
func (m *QueryStreamPaymentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStreamPaymentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStreamPaymentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStreamPaymentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStreamPaymentsResponse.Merge(m, src)
}
func (m *QueryStreamPaymentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStreamPaymentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStreamPaymentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStreamPaymentsResponse proto.InternalMessageInfo

func (m *QueryStreamPaymentsResponse) GetStreamPayments() []StreamPayment {
	if m != nil {
		return m.StreamPayments
	}
	return nil
}

func (m *QueryStreamPaymentsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryStreamPaymentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryStreamPaymentRequest) Reset()         { *m = QueryStreamPaymentRequest{} }
func (m *QueryStreamPaymentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStreamPaymentRequest) ProtoMessage()    {}
func (*QueryStreamPaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec493855485c26d, []int{2}
}
func (m *QueryStreamPaymentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStreamPaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStreamPaymentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStreamPaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStreamPaymentRequest.Merge(m, src)
}
func (m *QueryStreamPaymentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryStreamPaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStreamPaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStreamPaymentRequest proto.InternalMessageInfo

func (m *QueryStreamPaymentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStreamPaymentResponse struct {
	StreamPayment *StreamPayment `protobuf:"bytes,1,opt,name=stream_payment,json=streamPayment,proto3" json:"stream_payment,omitempty"`
}

func (m *QueryStreamPaymentResponse) Reset()         { *m = QueryStreamPaymentResponse{} }
func (m *QueryStreamPaymentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStreamPaymentResponse) ProtoMessage()    {}
func (*QueryStreamPaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dec493855485c26d, []int{3}
}
func (m *QueryStreamPaymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryStreamPaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryStreamPaymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryStreamPaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStreamPaymentResponse.Merge(m, src)
}
func (m *QueryStreamPaymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryStreamPaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStreamPaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStreamPaymentResponse proto.InternalMessageInfo

func (m *QueryStreamPaymentResponse) GetStreamPayment() *StreamPayment {
	if m != nil {
		return m.StreamPayment
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryStreamPaymentsRequest)(nil), "OmniFlix.paymentstream.streampay.QueryStreamPaymentsRequest")
	proto.RegisterType((*QueryStreamPaymentsResponse)(nil), "OmniFlix.paymentstream.streampay.QueryStreamPaymentsResponse")
	proto.RegisterType((*QueryStreamPaymentRequest)(nil), "OmniFlix.paymentstream.streampay.QueryStreamPaymentRequest")
	proto.RegisterType((*QueryStreamPaymentResponse)(nil), "OmniFlix.paymentstream.streampay.QueryStreamPaymentResponse")
}

func init() { proto.RegisterFile("streampay/query.proto", fileDescriptor_dec493855485c26d) }

var fileDescriptor_dec493855485c26d = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x4c, 0x4c, 0x9a, 0x27, 0x0a, 0x18, 0x36, 0x65, 0xd9, 0x94, 0x45, 0x39, 0x40,
	0x05, 0x9a, 0xad, 0x15, 0x4e, 0x30, 0x2e, 0x3b, 0x8c, 0x1b, 0x8c, 0x20, 0x71, 0xe0, 0x82, 0xdc,
	0xd6, 0x0a, 0x96, 0x1a, 0x3b, 0x8b, 0x5d, 0xd4, 0x08, 0x71, 0xe1, 0x09, 0x90, 0xb8, 0xf1, 0x0e,
	0x88, 0xd7, 0xd8, 0x71, 0x12, 0x17, 0x4e, 0xd3, 0xd4, 0x72, 0xe1, 0xca, 0x13, 0xa0, 0xda, 0x4e,
	0x4b, 0x46, 0x51, 0x45, 0x6f, 0x8e, 0xfd, 0x7d, 0xff, 0xef, 0xf7, 0xff, 0xc7, 0x86, 0x1b, 0x4a,
	0x17, 0x8c, 0x66, 0x39, 0x2d, 0xc9, 0xc9, 0x80, 0x15, 0x25, 0xce, 0x0b, 0xa9, 0x25, 0x8a, 0x9e,
	0x65, 0x82, 0x1f, 0xf5, 0xf9, 0x10, 0xe7, 0xb4, 0xcc, 0x98, 0xd0, 0xb6, 0x0c, 0x4f, 0xab, 0x83,
	0x9d, 0x54, 0xca, 0xb4, 0xcf, 0x08, 0xcd, 0x39, 0xa1, 0x42, 0x48, 0x4d, 0x35, 0x97, 0x42, 0xd9,
	0xfe, 0xe0, 0x6e, 0x57, 0xaa, 0x4c, 0x2a, 0xd2, 0xa1, 0x8a, 0x59, 0x61, 0xf2, 0x76, 0xbf, 0xc3,
	0x34, 0xdd, 0x27, 0x39, 0x4d, 0xb9, 0x30, 0xc5, 0xae, 0x76, 0x6b, 0x86, 0x30, 0x5d, 0xb9, 0xa3,
	0x5b, 0xa9, 0x4c, 0xa5, 0x59, 0x92, 0xc9, 0xca, 0xee, 0xc6, 0x9f, 0x01, 0x0c, 0x9e, 0x4f, 0x34,
	0x5f, 0x98, 0xf2, 0x63, 0x47, 0x98, 0xb0, 0x93, 0x01, 0x53, 0x1a, 0x6d, 0xc2, 0x55, 0xc5, 0x44,
	0x8f, 0x15, 0x3e, 0x88, 0x40, 0x6b, 0x2d, 0x71, 0x5f, 0x68, 0x07, 0xae, 0x15, 0xac, 0xcb, 0x73,
	0xce, 0x84, 0xf6, 0x3d, 0x73, 0x34, 0xdb, 0x40, 0x47, 0x10, 0xce, 0xc8, 0xfc, 0x95, 0x08, 0xb4,
	0xd6, 0xdb, 0xb7, 0xb1, 0xb5, 0x81, 0x27, 0x36, 0xb0, 0xcd, 0xc7, 0xd9, 0xc0, 0xc7, 0x34, 0x65,
	0x6e, 0x62, 0xf2, 0x47, 0x67, 0x7c, 0x01, 0xe0, 0xf6, 0x5c, 0x38, 0x95, 0x4b, 0xa1, 0x18, 0x1a,
	0xc2, 0x6b, 0xd6, 0xe5, 0xeb, 0x2a, 0x59, 0x1f, 0x44, 0x2b, 0xad, 0xf5, 0x36, 0xc1, 0x8b, 0x32,
	0xc7, 0x35, 0xc9, 0xc3, 0xf0, 0xf4, 0x7c, 0xb7, 0xf1, 0xeb, 0x7c, 0x77, 0xb3, 0xa4, 0x59, 0xff,
	0x61, 0x7c, 0x49, 0x35, 0x4e, 0x9a, 0xaa, 0x46, 0x80, 0x9e, 0xd4, 0x1c, 0x7a, 0xc6, 0xe1, 0x9d,
	0x85, 0x0e, 0x2d, 0x76, 0xcd, 0xe2, 0x3d, 0xb8, 0xf5, 0xb7, 0xc3, 0x2a, 0xfd, 0x26, 0xf4, 0x78,
	0xcf, 0x25, 0xef, 0xf1, 0x5e, 0xac, 0xe7, 0xfd, 0xab, 0x69, 0x1a, 0x2f, 0x61, 0xb3, 0xce, 0x6d,
	0x3a, 0xff, 0x3f, 0x8c, 0xe4, 0x6a, 0xcd, 0x6c, 0xfb, 0xa7, 0x07, 0xaf, 0x98, 0xb1, 0xe8, 0x0b,
	0x80, 0x37, 0x6c, 0x29, 0x17, 0xe9, 0x34, 0x8b, 0x83, 0xc5, 0xfa, 0xff, 0xbe, 0x61, 0xc1, 0xe3,
	0x25, 0xbb, 0xad, 0xe9, 0x78, 0xfb, 0xc3, 0xb7, 0x1f, 0x9f, 0xbc, 0x0d, 0x74, 0xd3, 0xdd, 0x77,
	0x2e, 0xd2, 0xbd, 0x4a, 0x07, 0x7d, 0x05, 0xf0, 0xfa, 0x65, 0x5e, 0xf4, 0x68, 0x99, 0x81, 0x15,
	0xed, 0xc1, 0x72, 0xcd, 0x0e, 0x36, 0x32, 0xb0, 0x01, 0xf2, 0xe7, 0xc0, 0x92, 0x77, 0xbc, 0xf7,
	0xfe, 0xf0, 0xe9, 0xe9, 0x28, 0x04, 0x67, 0xa3, 0x10, 0x5c, 0x8c, 0x42, 0xf0, 0x71, 0x1c, 0x36,
	0xce, 0xc6, 0x61, 0xe3, 0xfb, 0x38, 0x6c, 0xbc, 0x7a, 0x90, 0x72, 0xfd, 0x66, 0xd0, 0xc1, 0x5d,
	0x99, 0x91, 0x8a, 0x81, 0xb8, 0xe6, 0x3d, 0x2b, 0x47, 0x86, 0xb3, 0x47, 0x4f, 0x74, 0x99, 0x33,
	0xd5, 0x59, 0x35, 0xaf, 0xfc, 0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xc1, 0xf0, 0x52,
	0x9b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	StreamingPayments(ctx context.Context, in *QueryStreamPaymentsRequest, opts ...grpc.CallOption) (*QueryStreamPaymentsResponse, error)
	StreamingPayment(ctx context.Context, in *QueryStreamPaymentRequest, opts ...grpc.CallOption) (*QueryStreamPaymentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) StreamingPayments(ctx context.Context, in *QueryStreamPaymentsRequest, opts ...grpc.CallOption) (*QueryStreamPaymentsResponse, error) {
	out := new(QueryStreamPaymentsResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.paymentstream.streampay.Query/StreamingPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamingPayment(ctx context.Context, in *QueryStreamPaymentRequest, opts ...grpc.CallOption) (*QueryStreamPaymentResponse, error) {
	out := new(QueryStreamPaymentResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.paymentstream.streampay.Query/StreamingPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	StreamingPayments(context.Context, *QueryStreamPaymentsRequest) (*QueryStreamPaymentsResponse, error)
	StreamingPayment(context.Context, *QueryStreamPaymentRequest) (*QueryStreamPaymentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) StreamingPayments(ctx context.Context, req *QueryStreamPaymentsRequest) (*QueryStreamPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamingPayments not implemented")
}
func (*UnimplementedQueryServer) StreamingPayment(ctx context.Context, req *QueryStreamPaymentRequest) (*QueryStreamPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamingPayment not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_StreamingPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStreamPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StreamingPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.paymentstream.streampay.Query/StreamingPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StreamingPayments(ctx, req.(*QueryStreamPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamingPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStreamPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StreamingPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.paymentstream.streampay.Query/StreamingPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StreamingPayment(ctx, req.(*QueryStreamPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OmniFlix.paymentstream.streampay.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamingPayments",
			Handler:    _Query_StreamingPayments_Handler,
		},
		{
			MethodName: "StreamingPayment",
			Handler:    _Query_StreamingPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streampay/query.proto",
}

func (m *QueryStreamPaymentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStreamPaymentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStreamPaymentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStreamPaymentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStreamPaymentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStreamPaymentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StreamPayments) > 0 {
		for iNdEx := len(m.StreamPayments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StreamPayments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryStreamPaymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStreamPaymentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStreamPaymentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryStreamPaymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryStreamPaymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryStreamPaymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamPayment != nil {
		{
			size, err := m.StreamPayment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryStreamPaymentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStreamPaymentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StreamPayments) > 0 {
		for _, e := range m.StreamPayments {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStreamPaymentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryStreamPaymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StreamPayment != nil {
		l = m.StreamPayment.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryStreamPaymentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStreamPaymentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStreamPaymentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStreamPaymentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStreamPaymentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStreamPaymentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamPayments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamPayments = append(m.StreamPayments, StreamPayment{})
			if err := m.StreamPayments[len(m.StreamPayments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStreamPaymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStreamPaymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStreamPaymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryStreamPaymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryStreamPaymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryStreamPaymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamPayment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StreamPayment == nil {
				m.StreamPayment = &StreamPayment{}
			}
			if err := m.StreamPayment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)