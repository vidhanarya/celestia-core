// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/types/evidence.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// DuplicateVoteEvidence contains evidence a validator signed two conflicting
// votes.
type DuplicateVoteEvidence struct {
	VoteA     *Vote     `protobuf:"bytes,1,opt,name=vote_a,json=voteA,proto3" json:"vote_a,omitempty"`
	VoteB     *Vote     `protobuf:"bytes,2,opt,name=vote_b,json=voteB,proto3" json:"vote_b,omitempty"`
	Timestamp time.Time `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *DuplicateVoteEvidence) Reset()         { *m = DuplicateVoteEvidence{} }
func (m *DuplicateVoteEvidence) String() string { return proto.CompactTextString(m) }
func (*DuplicateVoteEvidence) ProtoMessage()    {}
func (*DuplicateVoteEvidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_6825fabc78e0a168, []int{0}
}
func (m *DuplicateVoteEvidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DuplicateVoteEvidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DuplicateVoteEvidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DuplicateVoteEvidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DuplicateVoteEvidence.Merge(m, src)
}
func (m *DuplicateVoteEvidence) XXX_Size() int {
	return m.Size()
}
func (m *DuplicateVoteEvidence) XXX_DiscardUnknown() {
	xxx_messageInfo_DuplicateVoteEvidence.DiscardUnknown(m)
}

var xxx_messageInfo_DuplicateVoteEvidence proto.InternalMessageInfo

func (m *DuplicateVoteEvidence) GetVoteA() *Vote {
	if m != nil {
		return m.VoteA
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetVoteB() *Vote {
	if m != nil {
		return m.VoteB
	}
	return nil
}

func (m *DuplicateVoteEvidence) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type Evidence struct {
	// Types that are valid to be assigned to Sum:
	//	*Evidence_DuplicateVoteEvidence
	Sum isEvidence_Sum `protobuf_oneof:"sum"`
}

func (m *Evidence) Reset()         { *m = Evidence{} }
func (m *Evidence) String() string { return proto.CompactTextString(m) }
func (*Evidence) ProtoMessage()    {}
func (*Evidence) Descriptor() ([]byte, []int) {
	return fileDescriptor_6825fabc78e0a168, []int{1}
}
func (m *Evidence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evidence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evidence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evidence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evidence.Merge(m, src)
}
func (m *Evidence) XXX_Size() int {
	return m.Size()
}
func (m *Evidence) XXX_DiscardUnknown() {
	xxx_messageInfo_Evidence.DiscardUnknown(m)
}

var xxx_messageInfo_Evidence proto.InternalMessageInfo

type isEvidence_Sum interface {
	isEvidence_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Evidence_DuplicateVoteEvidence struct {
	DuplicateVoteEvidence *DuplicateVoteEvidence `protobuf:"bytes,1,opt,name=duplicate_vote_evidence,json=duplicateVoteEvidence,proto3,oneof" json:"duplicate_vote_evidence,omitempty"`
}

func (*Evidence_DuplicateVoteEvidence) isEvidence_Sum() {}

func (m *Evidence) GetSum() isEvidence_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Evidence) GetDuplicateVoteEvidence() *DuplicateVoteEvidence {
	if x, ok := m.GetSum().(*Evidence_DuplicateVoteEvidence); ok {
		return x.DuplicateVoteEvidence
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Evidence) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Evidence_DuplicateVoteEvidence)(nil),
	}
}

// EvidenceData contains any evidence of malicious wrong-doing by validators
type EvidenceData struct {
	Evidence []Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence"`
	Hash     []byte     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *EvidenceData) Reset()         { *m = EvidenceData{} }
func (m *EvidenceData) String() string { return proto.CompactTextString(m) }
func (*EvidenceData) ProtoMessage()    {}
func (*EvidenceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6825fabc78e0a168, []int{2}
}
func (m *EvidenceData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EvidenceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EvidenceData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EvidenceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvidenceData.Merge(m, src)
}
func (m *EvidenceData) XXX_Size() int {
	return m.Size()
}
func (m *EvidenceData) XXX_DiscardUnknown() {
	xxx_messageInfo_EvidenceData.DiscardUnknown(m)
}

var xxx_messageInfo_EvidenceData proto.InternalMessageInfo

func (m *EvidenceData) GetEvidence() []Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *EvidenceData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*DuplicateVoteEvidence)(nil), "tendermint.types.DuplicateVoteEvidence")
	proto.RegisterType((*Evidence)(nil), "tendermint.types.Evidence")
	proto.RegisterType((*EvidenceData)(nil), "tendermint.types.EvidenceData")
}

func init() { proto.RegisterFile("tendermint/types/evidence.proto", fileDescriptor_6825fabc78e0a168) }

var fileDescriptor_6825fabc78e0a168 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xbb, 0xf2, 0x27, 0xb8, 0x72, 0x30, 0x8d, 0x28, 0x21, 0xa6, 0x35, 0x5c, 0xf4, 0xc2,
	0x6e, 0xa2, 0x57, 0x2e, 0x36, 0x98, 0x78, 0x6e, 0x0c, 0x07, 0x2f, 0xb8, 0x6d, 0xc7, 0xa5, 0x49,
	0xcb, 0x36, 0xed, 0x96, 0x04, 0x9f, 0x82, 0xd7, 0xf1, 0x0d, 0x38, 0x72, 0xf4, 0xa4, 0x06, 0x5e,
	0xc4, 0x74, 0x69, 0x0b, 0x01, 0x12, 0x2f, 0xcd, 0x74, 0xe6, 0x37, 0x3b, 0xdf, 0x37, 0xbb, 0xd8,
	0x94, 0x30, 0xf1, 0x20, 0x0e, 0xfd, 0x89, 0xa4, 0x72, 0x16, 0x41, 0x42, 0x61, 0xea, 0x7b, 0x30,
	0x71, 0x81, 0x44, 0xb1, 0x90, 0x42, 0x3f, 0xdf, 0x02, 0x44, 0x01, 0x9d, 0x0b, 0x2e, 0xb8, 0x50,
	0x45, 0x9a, 0x45, 0x1b, 0xae, 0x63, 0x72, 0x21, 0x78, 0x00, 0x54, 0xfd, 0x39, 0xe9, 0x3b, 0x95,
	0x7e, 0x08, 0x89, 0x64, 0x61, 0x94, 0x03, 0xd7, 0x07, 0x93, 0xd4, 0x77, 0x53, 0xed, 0x7e, 0x22,
	0xdc, 0x1a, 0xa4, 0x51, 0xe0, 0xbb, 0x4c, 0xc2, 0x50, 0x48, 0x78, 0xca, 0x65, 0xe8, 0x3d, 0x5c,
	0x9f, 0x0a, 0x09, 0x23, 0xd6, 0x46, 0x37, 0xe8, 0xee, 0xec, 0xfe, 0x92, 0xec, 0x2b, 0x22, 0x19,
	0x6f, 0xd7, 0x32, 0xea, 0xb1, 0xc4, 0x9d, 0xf6, 0xc9, 0xff, 0xb8, 0xa5, 0x5b, 0xf8, 0xb4, 0x14,
	0xda, 0xae, 0xa8, 0x8e, 0x0e, 0xd9, 0x58, 0x21, 0x85, 0x15, 0xf2, 0x52, 0x10, 0x56, 0x63, 0xf1,
	0x6d, 0x6a, 0xf3, 0x1f, 0x13, 0xd9, 0xdb, 0xb6, 0xae, 0xc4, 0x8d, 0x52, 0x2d, 0xc3, 0x57, 0x5e,
	0x61, 0x63, 0xa4, 0x84, 0x14, 0xfb, 0xcc, 0xe5, 0xdf, 0x1e, 0xea, 0x39, 0xea, 0xfb, 0x59, 0xb3,
	0x5b, 0xde, 0xb1, 0x82, 0x55, 0xc3, 0x95, 0x24, 0x0d, 0xbb, 0x6f, 0xb8, 0x59, 0xa4, 0x06, 0x4c,
	0x32, 0xbd, 0x8f, 0x1b, 0x3b, 0xa3, 0x2a, 0xca, 0xc8, 0xc1, 0xa8, 0xf2, 0x90, 0x6a, 0x66, 0xc4,
	0x2e, 0x3b, 0x74, 0x1d, 0x57, 0xc7, 0x2c, 0x19, 0xab, 0xa5, 0x35, 0x6d, 0x15, 0x5b, 0xc3, 0xc5,
	0xca, 0x40, 0xcb, 0x95, 0x81, 0x7e, 0x57, 0x06, 0x9a, 0xaf, 0x0d, 0x6d, 0xb9, 0x36, 0xb4, 0xaf,
	0xb5, 0xa1, 0xbd, 0xf6, 0xb9, 0x2f, 0xc7, 0xa9, 0x43, 0x5c, 0x11, 0xd2, 0x80, 0x7d, 0xcc, 0x02,
	0xf0, 0x38, 0xc4, 0x3b, 0x61, 0xcf, 0x15, 0x71, 0xfe, 0x16, 0xe8, 0xfe, 0xbd, 0x3b, 0x75, 0x95,
	0x7f, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x51, 0x95, 0xe6, 0x00, 0x7c, 0x02, 0x00, 0x00,
}

func (m *DuplicateVoteEvidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DuplicateVoteEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DuplicateVoteEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvidence(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if m.VoteB != nil {
		{
			size, err := m.VoteB.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.VoteA != nil {
		{
			size, err := m.VoteA.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Evidence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Evidence_DuplicateVoteEvidence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evidence_DuplicateVoteEvidence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DuplicateVoteEvidence != nil {
		{
			size, err := m.DuplicateVoteEvidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvidence(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *EvidenceData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvidenceData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EvidenceData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvidence(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvidence(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvidence(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvidence(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DuplicateVoteEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VoteA != nil {
		l = m.VoteA.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	if m.VoteB != nil {
		l = m.VoteB.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovEvidence(uint64(l))
	return n
}

func (m *Evidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Evidence_DuplicateVoteEvidence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DuplicateVoteEvidence != nil {
		l = m.DuplicateVoteEvidence.Size()
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}
func (m *EvidenceData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
			l = e.Size()
			n += 1 + l + sovEvidence(uint64(l))
		}
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvidence(uint64(l))
	}
	return n
}

func sovEvidence(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvidence(x uint64) (n int) {
	return sovEvidence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DuplicateVoteEvidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: DuplicateVoteEvidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DuplicateVoteEvidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VoteA == nil {
				m.VoteA = &Vote{}
			}
			if err := m.VoteA.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteB", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VoteB == nil {
				m.VoteB = &Vote{}
			}
			if err := m.VoteB.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *Evidence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: Evidence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evidence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DuplicateVoteEvidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DuplicateVoteEvidence{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Evidence_DuplicateVoteEvidence{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func (m *EvidenceData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvidence
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
			return fmt.Errorf("proto: EvidenceData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvidenceData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
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
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidence = append(m.Evidence, Evidence{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvidence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvidence
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvidence
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvidence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvidence
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvidence
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
func skipEvidence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
					return 0, ErrIntOverflowEvidence
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
				return 0, ErrInvalidLengthEvidence
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvidence
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvidence
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvidence        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvidence          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvidence = fmt.Errorf("proto: unexpected end of group")
)
