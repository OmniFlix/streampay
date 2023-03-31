// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: streampay/streampay.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type PaymentType int32

const (
	TypeDelayed    PaymentType = 0
	TypeContinuous PaymentType = 1
	TypePeriodic   PaymentType = 2
)

var PaymentType_name = map[int32]string{
	0: "PAYMENT_TYPE_DELAYED",
	1: "PAYMENT_TYPE_CONTINUOUS",
	2: "PAYMENT_TYPE_PERIODIC",
}

var PaymentType_value = map[string]int32{
	"PAYMENT_TYPE_DELAYED":    0,
	"PAYMENT_TYPE_CONTINUOUS": 1,
	"PAYMENT_TYPE_PERIODIC":   2,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}

func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fc790406387aa73, []int{0}
}

type PaymentStatus int32

const (
	StatusOpen      PaymentStatus = 0
	StatusCompleted PaymentStatus = 1
	StatusCancelled PaymentStatus = 2
)

var PaymentStatus_name = map[int32]string{
	0: "PAYMENT_STATUS_OPEN",
	1: "PAYMENT_STATUS_COMPLETED",
	2: "PAYMENT_STATUS_CANCELLED",
}

var PaymentStatus_value = map[string]int32{
	"PAYMENT_STATUS_OPEN":      0,
	"PAYMENT_STATUS_COMPLETED": 1,
	"PAYMENT_STATUS_CANCELLED": 2,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}

func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fc790406387aa73, []int{1}
}

type StreamPayment struct {
	Id                string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender            string      `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient         string      `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TotalAmount       types.Coin  `protobuf:"bytes,4,opt,name=total_amount,json=totalAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_amount" yaml:"total_amount"`
	StreamType        PaymentType `protobuf:"varint,5,opt,name=stream_type,json=streamType,proto3,enum=OmniFlix.streampay.v1.PaymentType" json:"stream_type,omitempty" yaml:"stream_type"`
	LockHeight        int64       `protobuf:"varint,6,opt,name=lock_height,json=lockHeight,proto3" json:"lock_height,omitempty" yaml:"lock_height"`
	StartTime         time.Time   `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime           time.Time   `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	TotalTransferred  types.Coin  `protobuf:"bytes,9,opt,name=total_transferred,json=totalTransferred,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"total_transferred" yaml:"total_transferred"`
	LastTransferredAt time.Time   `protobuf:"bytes,10,opt,name=last_transferred_at,json=lastTransferredAt,proto3,stdtime" json:"last_transferred_at" yaml:"last_transferred_at"`
}

func (m *StreamPayment) Reset()         { *m = StreamPayment{} }
func (m *StreamPayment) String() string { return proto.CompactTextString(m) }
func (*StreamPayment) ProtoMessage()    {}
func (*StreamPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fc790406387aa73, []int{0}
}
func (m *StreamPayment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamPayment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPayment.Merge(m, src)
}
func (m *StreamPayment) XXX_Size() int {
	return m.Size()
}
func (m *StreamPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPayment.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPayment proto.InternalMessageInfo

func (m *StreamPayment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamPayment) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *StreamPayment) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *StreamPayment) GetTotalAmount() types.Coin {
	if m != nil {
		return m.TotalAmount
	}
	return types.Coin{}
}

func (m *StreamPayment) GetStreamType() PaymentType {
	if m != nil {
		return m.StreamType
	}
	return TypeDelayed
}

func (m *StreamPayment) GetLockHeight() int64 {
	if m != nil {
		return m.LockHeight
	}
	return 0
}

func (m *StreamPayment) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *StreamPayment) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func (m *StreamPayment) GetTotalTransferred() types.Coin {
	if m != nil {
		return m.TotalTransferred
	}
	return types.Coin{}
}

func (m *StreamPayment) GetLastTransferredAt() time.Time {
	if m != nil {
		return m.LastTransferredAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("OmniFlix.streampay.v1.PaymentType", PaymentType_name, PaymentType_value)
	proto.RegisterEnum("OmniFlix.streampay.v1.PaymentStatus", PaymentStatus_name, PaymentStatus_value)
	proto.RegisterType((*StreamPayment)(nil), "OmniFlix.streampay.v1.StreamPayment")
}

func init() { proto.RegisterFile("streampay/streampay.proto", fileDescriptor_9fc790406387aa73) }

var fileDescriptor_9fc790406387aa73 = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0x8e, 0x03, 0x0d, 0x64, 0x02, 0xc1, 0x4c, 0x80, 0x1a, 0xb7, 0x75, 0xac, 0x1c, 0x4a, 0x4a,
	0x55, 0x5b, 0xa1, 0x87, 0x4a, 0xbd, 0xe5, 0x87, 0x51, 0xa3, 0x86, 0x24, 0x4a, 0x8c, 0x54, 0xda,
	0x83, 0x35, 0x89, 0x87, 0x60, 0x61, 0x7b, 0x2c, 0x7b, 0x82, 0x9a, 0xfb, 0x1e, 0x56, 0x39, 0x21,
	0xed, 0x75, 0x73, 0xda, 0x3d, 0xf1, 0x97, 0x70, 0xe4, 0xb8, 0x27, 0x58, 0xc1, 0x7f, 0xc0, 0x5f,
	0xb0, 0xb2, 0xc7, 0xd9, 0x98, 0x5d, 0x24, 0xb4, 0xa7, 0xcc, 0x7b, 0xef, 0xfb, 0xbe, 0x79, 0xef,
	0x7d, 0x13, 0x83, 0xdd, 0x80, 0xfa, 0x18, 0x39, 0x1e, 0x9a, 0xa8, 0x9f, 0x4f, 0x8a, 0xe7, 0x13,
	0x4a, 0xe0, 0x76, 0xc7, 0x71, 0xad, 0x43, 0xdb, 0xfa, 0x5f, 0x59, 0x54, 0x2e, 0x2a, 0xa2, 0x34,
	0x24, 0x81, 0x43, 0x02, 0x75, 0x80, 0x02, 0xac, 0x5e, 0x54, 0x06, 0x98, 0xa2, 0x8a, 0x3a, 0x24,
	0x96, 0xcb, 0x68, 0xe2, 0xd6, 0x88, 0x8c, 0x48, 0x74, 0x54, 0xc3, 0x53, 0x9c, 0x2d, 0x8e, 0x08,
	0x19, 0xd9, 0x58, 0x8d, 0xa2, 0xc1, 0xf8, 0x54, 0xa5, 0x96, 0x83, 0x03, 0x8a, 0x1c, 0x8f, 0x01,
	0x4a, 0x57, 0x19, 0xb0, 0xde, 0x8f, 0xee, 0xe9, 0xa2, 0x89, 0x83, 0x5d, 0x0a, 0xf3, 0x20, 0x6d,
	0x99, 0x02, 0x27, 0x73, 0xe5, 0x6c, 0x2f, 0x6d, 0x99, 0x70, 0x07, 0x64, 0x02, 0xec, 0x9a, 0xd8,
	0x17, 0xd2, 0x51, 0x2e, 0x8e, 0xe0, 0x8f, 0x20, 0xeb, 0xe3, 0xa1, 0xe5, 0x59, 0xd8, 0xa5, 0xc2,
	0x52, 0x54, 0x5a, 0x24, 0xe0, 0x2b, 0x0e, 0xac, 0x51, 0x42, 0x91, 0x6d, 0x20, 0x87, 0x8c, 0x5d,
	0x2a, 0x2c, 0xcb, 0x5c, 0x39, 0x77, 0xb0, 0xab, 0xb0, 0x31, 0x94, 0x70, 0x0c, 0x25, 0x1e, 0x43,
	0xa9, 0x13, 0xcb, 0xad, 0x1d, 0x5e, 0xdf, 0x16, 0x53, 0x8f, 0xb7, 0xc5, 0xc2, 0x04, 0x39, 0xf6,
	0x9f, 0xa5, 0x24, 0xb9, 0x74, 0x75, 0x57, 0xdc, 0x1b, 0x59, 0xf4, 0x6c, 0x3c, 0x50, 0x86, 0xc4,
	0x51, 0xe3, 0x55, 0xb0, 0x9f, 0xdf, 0x02, 0xf3, 0x5c, 0xa5, 0x13, 0x0f, 0x07, 0x91, 0x4e, 0x2f,
	0x17, 0x31, 0xab, 0x11, 0x11, 0xfe, 0x07, 0x72, 0x6c, 0x8b, 0x46, 0x88, 0x10, 0xbe, 0x93, 0xb9,
	0x72, 0xfe, 0xa0, 0xa4, 0x3c, 0xbb, 0x62, 0x25, 0xde, 0x80, 0x3e, 0xf1, 0x70, 0x6d, 0xe7, 0xf1,
	0xb6, 0x08, 0x59, 0x27, 0x09, 0x81, 0x52, 0x0f, 0xb0, 0x28, 0xc4, 0xc0, 0x3f, 0x40, 0xce, 0x26,
	0xc3, 0x73, 0xe3, 0x0c, 0x5b, 0xa3, 0x33, 0x2a, 0x64, 0x64, 0xae, 0xbc, 0x94, 0x24, 0x26, 0x8a,
	0xa5, 0x1e, 0x08, 0xa3, 0xbf, 0xa2, 0x00, 0xfe, 0x03, 0x40, 0x40, 0x91, 0x4f, 0x8d, 0xd0, 0x0d,
	0x61, 0x25, 0xda, 0x8c, 0xa8, 0x30, 0xab, 0x94, 0xb9, 0x55, 0x8a, 0x3e, 0xb7, 0xaa, 0xf6, 0x53,
	0xbc, 0x9a, 0xcd, 0x79, 0x43, 0x73, 0x6e, 0xe9, 0xf2, 0xae, 0xc8, 0xf5, 0xb2, 0x51, 0x22, 0x84,
	0xc3, 0x1e, 0x58, 0xc5, 0xae, 0xc9, 0x74, 0x57, 0x5f, 0xd4, 0xfd, 0x21, 0xd6, 0xdd, 0x60, 0xba,
	0x73, 0x26, 0x53, 0x5d, 0xc1, 0xae, 0x19, 0x69, 0xbe, 0xe1, 0xc0, 0x26, 0x73, 0x83, 0xfa, 0xc8,
	0x0d, 0x4e, 0xb1, 0xef, 0x63, 0x53, 0xc8, 0xbe, 0xe4, 0xe7, 0xdf, 0xb1, 0xb8, 0x90, 0xf4, 0x33,
	0xa1, 0xf0, 0x4d, 0xa6, 0xf2, 0x11, 0x5d, 0x5f, 0xb0, 0xa1, 0x0f, 0x0a, 0x36, 0x0a, 0x68, 0x52,
	0xd1, 0x40, 0x54, 0x00, 0x2f, 0x0e, 0xfd, 0x73, 0xdc, 0x97, 0x18, 0x9b, 0xf4, 0xb5, 0x08, 0x9b,
	0x7f, 0x33, 0xac, 0x24, 0x6e, 0xac, 0xd2, 0xfd, 0xb7, 0x1c, 0xc8, 0x25, 0x1e, 0x09, 0xfc, 0x05,
	0x6c, 0x75, 0xab, 0x27, 0x47, 0x5a, 0x5b, 0x37, 0xf4, 0x93, 0xae, 0x66, 0x34, 0xb4, 0x56, 0xf5,
	0x44, 0x6b, 0xf0, 0x29, 0x71, 0x63, 0x3a, 0x93, 0x73, 0x21, 0xa6, 0x81, 0x6d, 0x34, 0xc1, 0x26,
	0x54, 0xc1, 0xf7, 0x4f, 0xa0, 0xf5, 0x4e, 0x5b, 0x6f, 0xb6, 0x8f, 0x3b, 0xc7, 0x7d, 0x9e, 0x13,
	0xe1, 0x74, 0x26, 0xe7, 0x43, 0x74, 0x9d, 0xb8, 0xd4, 0x72, 0xc7, 0x64, 0x1c, 0xc0, 0x5f, 0xc1,
	0xf6, 0x13, 0x42, 0x57, 0xeb, 0x35, 0x3b, 0x8d, 0x66, 0x9d, 0x4f, 0x8b, 0xfc, 0x74, 0x26, 0xaf,
	0x85, 0xf0, 0x2e, 0xf6, 0x2d, 0x62, 0x5a, 0x43, 0x71, 0xf9, 0xf5, 0x3b, 0x29, 0xb5, 0xff, 0x9e,
	0x03, 0xeb, 0x71, 0x7b, 0x7d, 0x8a, 0xe8, 0x38, 0x80, 0x7b, 0xa0, 0x30, 0x17, 0xe9, 0xeb, 0x55,
	0xfd, 0xb8, 0x6f, 0x74, 0xba, 0x5a, 0x9b, 0x4f, 0x89, 0xf9, 0xe9, 0x4c, 0x06, 0x0c, 0xd4, 0xf1,
	0xb0, 0x0b, 0x2b, 0x40, 0xf8, 0x02, 0x58, 0xef, 0x1c, 0x75, 0x5b, 0x9a, 0xae, 0x35, 0x78, 0x4e,
	0x2c, 0x4c, 0x67, 0xf2, 0x06, 0x43, 0xd7, 0x89, 0xe3, 0xd9, 0x98, 0x62, 0xf3, 0x39, 0x4a, 0xb5,
	0x5d, 0xd7, 0x5a, 0x2d, 0xad, 0xc1, 0xa7, 0x9f, 0x50, 0x90, 0x3b, 0xc4, 0xb6, 0x8d, 0x4d, 0xd6,
	0x66, 0xad, 0x79, 0x7d, 0x2f, 0x71, 0x37, 0xf7, 0x12, 0xf7, 0xf1, 0x5e, 0xe2, 0x2e, 0x1f, 0xa4,
	0xd4, 0xcd, 0x83, 0x94, 0xfa, 0xf0, 0x20, 0xa5, 0xfe, 0x55, 0x13, 0xcf, 0x61, 0xfe, 0x17, 0x5d,
	0x7c, 0x1f, 0xd5, 0xe4, 0x39, 0x7a, 0x1b, 0x83, 0x4c, 0xe4, 0xef, 0xef, 0x9f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x8c, 0xe3, 0x0b, 0x46, 0x4f, 0x05, 0x00, 0x00,
}

func (m *StreamPayment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamPayment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamPayment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastTransferredAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastTransferredAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStreampay(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size, err := m.TotalTransferred.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStreampay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintStreampay(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x42
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintStreampay(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x3a
	if m.LockHeight != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.LockHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.StreamType != 0 {
		i = encodeVarintStreampay(dAtA, i, uint64(m.StreamType))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TotalAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStreampay(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintStreampay(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreampay(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreampay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamPayment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovStreampay(uint64(l))
	}
	l = m.TotalAmount.Size()
	n += 1 + l + sovStreampay(uint64(l))
	if m.StreamType != 0 {
		n += 1 + sovStreampay(uint64(m.StreamType))
	}
	if m.LockHeight != 0 {
		n += 1 + sovStreampay(uint64(m.LockHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovStreampay(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovStreampay(uint64(l))
	l = m.TotalTransferred.Size()
	n += 1 + l + sovStreampay(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastTransferredAt)
	n += 1 + l + sovStreampay(uint64(l))
	return n
}

func sovStreampay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreampay(x uint64) (n int) {
	return sovStreampay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamPayment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreampay
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
			return fmt.Errorf("proto: StreamPayment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamPayment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowStreampay
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockHeight", wireType)
			}
			m.LockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTransferred", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTransferred.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransferredAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreampay
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
				return ErrInvalidLengthStreampay
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreampay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastTransferredAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStreampay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreampay
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
func skipStreampay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreampay
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
					return 0, ErrIntOverflowStreampay
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
					return 0, ErrIntOverflowStreampay
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
				return 0, ErrInvalidLengthStreampay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreampay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreampay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreampay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreampay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreampay = fmt.Errorf("proto: unexpected end of group")
)
