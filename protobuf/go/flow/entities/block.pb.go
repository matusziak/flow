// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type BlockStatus int32

const (
	BlockStatus_BLOCK_UNKNOWN   BlockStatus = 0
	BlockStatus_BLOCK_FINALIZED BlockStatus = 1
	BlockStatus_BLOCK_SEALED    BlockStatus = 2
)

var BlockStatus_name = map[int32]string{
	0: "BLOCK_UNKNOWN",
	1: "BLOCK_FINALIZED",
	2: "BLOCK_SEALED",
}

var BlockStatus_value = map[string]int32{
	"BLOCK_UNKNOWN":   0,
	"BLOCK_FINALIZED": 1,
	"BLOCK_SEALED":    2,
}

func (x BlockStatus) String() string {
	return proto.EnumName(BlockStatus_name, int32(x))
}

func (BlockStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1afc3335f2172fc, []int{0}
}

type Block struct {
	Id                       []byte                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId                 []byte                  `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Height                   uint64                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp                *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CollectionGuarantees     []*CollectionGuarantee  `protobuf:"bytes,5,rep,name=collection_guarantees,json=collectionGuarantees,proto3" json:"collection_guarantees,omitempty"`
	BlockSeals               []*BlockSeal            `protobuf:"bytes,6,rep,name=block_seals,json=blockSeals,proto3" json:"block_seals,omitempty"`
	Signatures               [][]byte                `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty"`
	ExecutionReceiptMetaList []*ExecutionReceiptMeta `protobuf:"bytes,8,rep,name=execution_receipt_metaList,json=executionReceiptMetaList,proto3" json:"execution_receipt_metaList,omitempty"`
	ExecutionResultList      []*ExecutionResult      `protobuf:"bytes,9,rep,name=execution_result_list,json=executionResultList,proto3" json:"execution_result_list,omitempty"`
	BlockHeader              *BlockHeader            `protobuf:"bytes,10,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1afc3335f2172fc, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Block) GetParentId() []byte {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetCollectionGuarantees() []*CollectionGuarantee {
	if m != nil {
		return m.CollectionGuarantees
	}
	return nil
}

func (m *Block) GetBlockSeals() []*BlockSeal {
	if m != nil {
		return m.BlockSeals
	}
	return nil
}

func (m *Block) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Block) GetExecutionReceiptMetaList() []*ExecutionReceiptMeta {
	if m != nil {
		return m.ExecutionReceiptMetaList
	}
	return nil
}

func (m *Block) GetExecutionResultList() []*ExecutionResult {
	if m != nil {
		return m.ExecutionResultList
	}
	return nil
}

func (m *Block) GetBlockHeader() *BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func init() {
	proto.RegisterEnum("flow.entities.BlockStatus", BlockStatus_name, BlockStatus_value)
	proto.RegisterType((*Block)(nil), "flow.entities.Block")
}

func init() { proto.RegisterFile("flow/entities/block.proto", fileDescriptor_c1afc3335f2172fc) }

var fileDescriptor_c1afc3335f2172fc = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x18, 0xc5, 0x69, 0xd7, 0x95, 0xf5, 0x6b, 0x07, 0xc5, 0x63, 0xc8, 0x14, 0x54, 0xa2, 0xc1, 0xa1,
	0xe2, 0x90, 0xa0, 0x71, 0x81, 0x03, 0x87, 0x75, 0x2b, 0xa3, 0x5a, 0xe9, 0x90, 0x07, 0x9a, 0xb4,
	0x4b, 0xe4, 0x24, 0xdf, 0x52, 0x8b, 0x34, 0xae, 0x62, 0x47, 0xf0, 0x67, 0xf1, 0x27, 0xa2, 0xd8,
	0x4d, 0xdb, 0x94, 0x72, 0xa9, 0xea, 0xf7, 0x7e, 0x7e, 0x5f, 0xf2, 0xec, 0xc0, 0xf3, 0xfb, 0x44,
	0xfe, 0xf2, 0x30, 0xd5, 0x42, 0x0b, 0x54, 0x5e, 0x90, 0xc8, 0xf0, 0xa7, 0xbb, 0xc8, 0xa4, 0x96,
	0xe4, 0xb0, 0xb0, 0xdc, 0xd2, 0xea, 0xbd, 0x8a, 0xa5, 0x8c, 0x13, 0xf4, 0x8c, 0x19, 0xe4, 0xf7,
	0x9e, 0x16, 0x73, 0x54, 0x9a, 0xcf, 0x17, 0x96, 0xef, 0xf5, 0xab, 0x51, 0xa1, 0x4c, 0x12, 0x0c,
	0xb5, 0x90, 0xe9, 0x6e, 0xdf, 0x8c, 0xf2, 0x15, 0xf2, 0x64, 0xe9, 0xbf, 0xa9, 0xfa, 0xf8, 0x1b,
	0xc3, 0xbc, 0xd8, 0xee, 0x67, 0xa8, 0xf2, 0x44, 0x2f, 0x29, 0x67, 0x57, 0xca, 0x0c, 0x79, 0x84,
	0x99, 0x25, 0x4e, 0xfe, 0x34, 0x60, 0x7f, 0x58, 0xc8, 0xe4, 0x11, 0xd4, 0x45, 0x44, 0x6b, 0x4e,
	0x6d, 0xd0, 0x61, 0x75, 0x11, 0x91, 0x17, 0xd0, 0x5a, 0xf0, 0x0c, 0x53, 0xed, 0x8b, 0x88, 0xd6,
	0x8d, 0x7c, 0x60, 0x85, 0x71, 0x44, 0x9e, 0x41, 0x73, 0x86, 0x22, 0x9e, 0x69, 0xba, 0xe7, 0xd4,
	0x06, 0x0d, 0xb6, 0x5c, 0x91, 0x0f, 0xd0, 0x5a, 0xbd, 0x29, 0x6d, 0x38, 0xb5, 0x41, 0xfb, 0xb4,
	0xe7, 0xda, 0x2e, 0xdc, 0xb2, 0x0b, 0xf7, 0x7b, 0x49, 0xb0, 0x35, 0x4c, 0x6e, 0xe1, 0x78, 0x5d,
	0x82, 0x1f, 0xe7, 0x3c, 0xe3, 0xa9, 0x46, 0x54, 0x74, 0xdf, 0xd9, 0x1b, 0xb4, 0x4f, 0x4f, 0xdc,
	0x4a, 0xc1, 0xee, 0xf9, 0x8a, 0xbd, 0x2c, 0x51, 0xf6, 0x34, 0xfc, 0x57, 0x54, 0xe4, 0x23, 0xb4,
	0xd7, 0xed, 0x29, 0xda, 0x34, 0x71, 0x74, 0x2b, 0xce, 0x54, 0x70, 0x83, 0x3c, 0x61, 0x10, 0x94,
	0x7f, 0x15, 0xe9, 0x03, 0x28, 0x11, 0xa7, 0x5c, 0xe7, 0x19, 0x2a, 0xfa, 0xd0, 0xd9, 0x1b, 0x74,
	0xd8, 0x86, 0x42, 0x38, 0xf4, 0x36, 0x8b, 0x0f, 0x51, 0x2c, 0xb4, 0x3f, 0x47, 0xcd, 0x27, 0x42,
	0x69, 0x7a, 0x60, 0x26, 0xbd, 0xde, 0x9a, 0x34, 0x2a, 0x37, 0x30, 0xcb, 0x7f, 0x45, 0xcd, 0x19,
	0xc5, 0x1d, 0x6a, 0x11, 0x42, 0x18, 0x1c, 0x6f, 0x9f, 0xad, 0x9f, 0x14, 0xe9, 0x2d, 0x93, 0xde,
	0xff, 0x7f, 0x7a, 0x81, 0xb2, 0x23, 0xac, 0x0a, 0x26, 0xf3, 0x13, 0x74, 0x36, 0x6f, 0x02, 0x85,
	0xe5, 0x39, 0xed, 0xa8, 0xe4, 0x8b, 0x21, 0x98, 0x6d, 0xd0, 0x2e, 0xde, 0x5e, 0x42, 0xdb, 0xd6,
	0xa5, 0xb9, 0xce, 0x15, 0x79, 0x02, 0x87, 0xc3, 0xc9, 0xf5, 0xf9, 0x95, 0xff, 0x63, 0x7a, 0x35,
	0xbd, 0xbe, 0x9d, 0x76, 0x1f, 0x90, 0x23, 0x78, 0x6c, 0xa5, 0xcf, 0xe3, 0xe9, 0xd9, 0x64, 0x7c,
	0x37, 0xba, 0xe8, 0xd6, 0x48, 0x17, 0x3a, 0x56, 0xbc, 0x19, 0x9d, 0x4d, 0x46, 0x17, 0xdd, 0xfa,
	0xf0, 0x1b, 0xbc, 0x94, 0x59, 0xec, 0xca, 0xd4, 0x0c, 0x5e, 0x5d, 0x8f, 0xf2, 0x09, 0xee, 0xde,
	0xc5, 0x42, 0xcf, 0xf2, 0xc0, 0x0d, 0xe5, 0xdc, 0xb3, 0x90, 0x67, 0x7e, 0x56, 0x1f, 0x55, 0x2c,
	0xbd, 0xca, 0x05, 0x0f, 0x9a, 0xc6, 0x7a, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x63, 0x03, 0xba,
	0xf0, 0xa9, 0x03, 0x00, 0x00,
}
