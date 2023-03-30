// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streampay/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgStreamSend struct {
	Sender     string      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient  string      `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount     types.Coin  `protobuf:"bytes,3,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	EndTime    time.Time   `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	StreamType PaymentType `protobuf:"varint,5,opt,name=stream_type,json=streamType,proto3,enum=OmniFlix.paymentstream.streampay.PaymentType" json:"stream_type,omitempty" yaml:"stream_type"`
}

func (m *MsgStreamSend) Reset()         { *m = MsgStreamSend{} }
func (m *MsgStreamSend) String() string { return proto.CompactTextString(m) }
func (*MsgStreamSend) ProtoMessage()    {}
func (*MsgStreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc5bf14bcdfaf29, []int{0}
}
func (m *MsgStreamSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStreamSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStreamSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStreamSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStreamSend.Merge(m, src)
}
func (m *MsgStreamSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgStreamSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStreamSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStreamSend proto.InternalMessageInfo

func (m *MsgStreamSend) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStreamSend) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgStreamSend) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *MsgStreamSend) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *MsgStreamSend) GetStreamType() PaymentType {
	if m != nil {
		return m.StreamType
	}
	return TypeDelayed
}

type MsgStreamSendResponse struct {
}

func (m *MsgStreamSendResponse) Reset()         { *m = MsgStreamSendResponse{} }
func (m *MsgStreamSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStreamSendResponse) ProtoMessage()    {}
func (*MsgStreamSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfc5bf14bcdfaf29, []int{1}
}
func (m *MsgStreamSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStreamSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStreamSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStreamSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStreamSendResponse.Merge(m, src)
}
func (m *MsgStreamSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStreamSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStreamSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStreamSendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStreamSend)(nil), "OmniFlix.paymentstream.streampay.MsgStreamSend")
	proto.RegisterType((*MsgStreamSendResponse)(nil), "OmniFlix.paymentstream.streampay.MsgStreamSendResponse")
}

func init() { proto.RegisterFile("streampay/tx.proto", fileDescriptor_bfc5bf14bcdfaf29) }

var fileDescriptor_bfc5bf14bcdfaf29 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x8d, 0x5b, 0x08, 0xd4, 0x15, 0x20, 0x59, 0x50, 0xa6, 0x01, 0xcd, 0x44, 0xb3, 0x21, 0x9b,
	0xd8, 0x6a, 0x40, 0x42, 0x62, 0x19, 0x24, 0x76, 0x05, 0x34, 0xed, 0x8a, 0x4d, 0xe5, 0xc9, 0xdc,
	0x0e, 0x16, 0xf1, 0x43, 0xb1, 0x53, 0x65, 0x16, 0xfc, 0x43, 0xbf, 0x83, 0x2f, 0xe9, 0xb2, 0x4b,
	0x56, 0x29, 0x4a, 0xc4, 0x0f, 0xf4, 0x0b, 0x90, 0xc7, 0x33, 0x6d, 0xba, 0xaa, 0x58, 0xcd, 0xbd,
	0xd7, 0xe7, 0x9e, 0xe3, 0xe3, 0x33, 0x98, 0x58, 0x37, 0x03, 0x2e, 0x0d, 0xaf, 0x98, 0x5b, 0x50,
	0x33, 0xd3, 0x4e, 0x93, 0xfe, 0x17, 0xa9, 0xc4, 0xa7, 0xa9, 0x58, 0x50, 0xc3, 0x2b, 0x09, 0xca,
	0x05, 0x0c, 0xbd, 0x81, 0xf6, 0xe2, 0x89, 0xb6, 0x52, 0x5b, 0x96, 0x73, 0x0b, 0xec, 0xec, 0x20,
	0x07, 0xc7, 0x0f, 0xd8, 0x44, 0x0b, 0x15, 0x18, 0x7a, 0xcf, 0x4b, 0x5d, 0xea, 0xba, 0x64, 0xbe,
	0x6a, 0xa6, 0x49, 0xa9, 0x75, 0x39, 0x05, 0x56, 0x77, 0xf9, 0xfc, 0x94, 0x39, 0x21, 0xc1, 0x3a,
	0x2e, 0x4d, 0x03, 0xd8, 0xbf, 0xbd, 0xcc, 0x4d, 0x15, 0x8e, 0xd2, 0xbf, 0x5b, 0xf8, 0xc9, 0xa1,
	0x2d, 0x8f, 0xea, 0xf1, 0x11, 0xa8, 0x82, 0xec, 0xe1, 0xae, 0x05, 0x55, 0xc0, 0x2c, 0x42, 0x7d,
	0x34, 0xd8, 0xc9, 0x9a, 0x8e, 0xbc, 0xc6, 0x3b, 0x33, 0x98, 0x08, 0x23, 0x40, 0xb9, 0x68, 0xab,
	0x3e, 0xba, 0x1d, 0x90, 0x1c, 0x77, 0xb9, 0xd4, 0x73, 0xe5, 0xa2, 0xed, 0x3e, 0x1a, 0xec, 0x8e,
	0xf6, 0x69, 0xb0, 0x42, 0xbd, 0x15, 0xda, 0x58, 0xa1, 0x1f, 0xb5, 0x50, 0x63, 0x76, 0xb1, 0x4c,
	0x3a, 0xbf, 0xae, 0x92, 0x37, 0xa5, 0x70, 0xdf, 0xe7, 0x39, 0x9d, 0x68, 0xc9, 0x1a, 0xdf, 0xe1,
	0x33, 0xb4, 0xc5, 0x0f, 0xe6, 0x2a, 0x03, 0xb6, 0x5e, 0xc8, 0x1a, 0x66, 0x92, 0xe1, 0xc7, 0xa0,
	0x8a, 0x13, 0xef, 0x2e, 0x7a, 0x50, 0xab, 0xf4, 0x68, 0xb0, 0x4e, 0x5b, 0xeb, 0xf4, 0xb8, 0xb5,
	0x3e, 0x7e, 0xe5, 0x65, 0xae, 0x97, 0xc9, 0xb3, 0x8a, 0xcb, 0xe9, 0x87, 0xb4, 0xdd, 0x4c, 0xcf,
	0xaf, 0x12, 0x94, 0x3d, 0x02, 0x55, 0x78, 0x28, 0x39, 0xc5, 0xbb, 0xe1, 0x49, 0x4e, 0xbc, 0x60,
	0xf4, 0xb0, 0x8f, 0x06, 0x4f, 0x47, 0x43, 0x7a, 0x5f, 0x52, 0xf4, 0x6b, 0x98, 0x1f, 0x57, 0x06,
	0xc6, 0x7b, 0xd7, 0xcb, 0x84, 0x04, 0x95, 0x0d, 0xae, 0x34, 0xc3, 0xa1, 0xf3, 0x98, 0xf4, 0x25,
	0x7e, 0x71, 0xe7, 0x99, 0x33, 0xb0, 0x46, 0x2b, 0x0b, 0xa3, 0x9f, 0x78, 0xfb, 0xd0, 0x96, 0xe4,
	0x0c, 0xe3, 0x8d, 0x0c, 0xd8, 0xfd, 0x17, 0xb8, 0xc3, 0xd6, 0x7b, 0xff, 0x9f, 0x0b, 0xad, 0xfc,
	0xf8, 0xf3, 0xc5, 0x2a, 0x46, 0x97, 0xab, 0x18, 0xfd, 0x59, 0xc5, 0xe8, 0x7c, 0x1d, 0x77, 0x2e,
	0xd7, 0x71, 0xe7, 0xf7, 0x3a, 0xee, 0x7c, 0x7b, 0xb7, 0x11, 0x4f, 0x4b, 0xce, 0x1a, 0xf2, 0x61,
	0xa0, 0x65, 0x0b, 0xb6, 0xf1, 0x9b, 0xfb, 0xc0, 0xf2, 0x6e, 0x9d, 0xc4, 0xdb, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0xf5, 0x6b, 0x04, 0x00, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	StreamSend(ctx context.Context, in *MsgStreamSend, opts ...grpc.CallOption) (*MsgStreamSendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StreamSend(ctx context.Context, in *MsgStreamSend, opts ...grpc.CallOption) (*MsgStreamSendResponse, error) {
	out := new(MsgStreamSendResponse)
	err := c.cc.Invoke(ctx, "/OmniFlix.paymentstream.streampay.Msg/StreamSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	StreamSend(context.Context, *MsgStreamSend) (*MsgStreamSendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StreamSend(ctx context.Context, req *MsgStreamSend) (*MsgStreamSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StreamSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStreamSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StreamSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OmniFlix.paymentstream.streampay.Msg/StreamSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StreamSend(ctx, req.(*MsgStreamSend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OmniFlix.paymentstream.streampay.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamSend",
			Handler:    _Msg_StreamSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streampay/tx.proto",
}

func (m *MsgStreamSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStreamSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStreamSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StreamType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StreamType))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStreamSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStreamSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStreamSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStreamSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovTx(uint64(l))
	if m.StreamType != 0 {
		n += 1 + sovTx(uint64(m.StreamType))
	}
	return n
}

func (m *MsgStreamSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStreamSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStreamSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStreamSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamType", wireType)
			}
			m.StreamType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamType |= PaymentType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStreamSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStreamSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStreamSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)