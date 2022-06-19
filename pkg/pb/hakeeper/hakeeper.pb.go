// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hakeeper.proto

package hakeeper

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ConfigChangeType indicates config change command type.
type ConfigChangeType int32

const (
	AddNode    ConfigChangeType = 0
	RemoveNode ConfigChangeType = 1
	StartNode  ConfigChangeType = 2
	StopNode   ConfigChangeType = 3
)

var ConfigChangeType_name = map[int32]string{
	0: "AddNode",
	1: "RemoveNode",
	2: "StartNode",
	3: "StopNode",
}

var ConfigChangeType_value = map[string]int32{
	"AddNode":    0,
	"RemoveNode": 1,
	"StartNode":  2,
	"StopNode":   3,
}

func (x ConfigChangeType) String() string {
	return proto.EnumName(ConfigChangeType_name, int32(x))
}

func (ConfigChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{0}
}

// ServiceType specifies type of service
type ServiceType int32

const (
	LogService ServiceType = 0
	DnService  ServiceType = 1
)

var ServiceType_name = map[int32]string{
	0: "LogService",
	1: "DnService",
}

var ServiceType_value = map[string]int32{
	"LogService": 0,
	"DnService":  1,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{1}
}

// Replica of the shard
type Replica struct {
	ShardID              uint64   `protobuf:"varint,1,opt,name=ShardID,proto3" json:"ShardID,omitempty"`
	ReplicaID            uint64   `protobuf:"varint,2,opt,name=ReplicaID,proto3" json:"ReplicaID,omitempty"`
	Epoch                uint64   `protobuf:"varint,3,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Replica) Reset()         { *m = Replica{} }
func (m *Replica) String() string { return proto.CompactTextString(m) }
func (*Replica) ProtoMessage()    {}
func (*Replica) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{0}
}
func (m *Replica) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Replica) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Replica.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Replica) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Replica.Merge(m, src)
}
func (m *Replica) XXX_Size() int {
	return m.Size()
}
func (m *Replica) XXX_DiscardUnknown() {
	xxx_messageInfo_Replica.DiscardUnknown(m)
}

var xxx_messageInfo_Replica proto.InternalMessageInfo

func (m *Replica) GetShardID() uint64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *Replica) GetReplicaID() uint64 {
	if m != nil {
		return m.ReplicaID
	}
	return 0
}

func (m *Replica) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// ConfigChange is the detail of a config change.
type ConfigChange struct {
	Replica              Replica          `protobuf:"bytes,1,opt,name=Replica,proto3" json:"Replica"`
	ChangeType           ConfigChangeType `protobuf:"varint,2,opt,name=ChangeType,proto3,enum=hakeeper.ConfigChangeType" json:"ChangeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConfigChange) Reset()         { *m = ConfigChange{} }
func (m *ConfigChange) String() string { return proto.CompactTextString(m) }
func (*ConfigChange) ProtoMessage()    {}
func (*ConfigChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{1}
}
func (m *ConfigChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigChange.Merge(m, src)
}
func (m *ConfigChange) XXX_Size() int {
	return m.Size()
}
func (m *ConfigChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigChange.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigChange proto.InternalMessageInfo

func (m *ConfigChange) GetReplica() Replica {
	if m != nil {
		return m.Replica
	}
	return Replica{}
}

func (m *ConfigChange) GetChangeType() ConfigChangeType {
	if m != nil {
		return m.ChangeType
	}
	return AddNode
}

// ScheduleCommand contains a shard schedule command.
type ScheduleCommand struct {
	UUID                 string       `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ConfigChange         ConfigChange `protobuf:"bytes,2,opt,name=ConfigChange,proto3" json:"ConfigChange"`
	ServiceType          ServiceType  `protobuf:"varint,3,opt,name=ServiceType,proto3,enum=hakeeper.ServiceType" json:"ServiceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScheduleCommand) Reset()         { *m = ScheduleCommand{} }
func (m *ScheduleCommand) String() string { return proto.CompactTextString(m) }
func (*ScheduleCommand) ProtoMessage()    {}
func (*ScheduleCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{2}
}
func (m *ScheduleCommand) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduleCommand.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduleCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleCommand.Merge(m, src)
}
func (m *ScheduleCommand) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleCommand proto.InternalMessageInfo

func (m *ScheduleCommand) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *ScheduleCommand) GetConfigChange() ConfigChange {
	if m != nil {
		return m.ConfigChange
	}
	return ConfigChange{}
}

func (m *ScheduleCommand) GetServiceType() ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return LogService
}

type CommandBatch struct {
	Term                 uint64            `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	Commands             []ScheduleCommand `protobuf:"bytes,2,rep,name=Commands,proto3" json:"Commands"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CommandBatch) Reset()         { *m = CommandBatch{} }
func (m *CommandBatch) String() string { return proto.CompactTextString(m) }
func (*CommandBatch) ProtoMessage()    {}
func (*CommandBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1506f3aa5330eb, []int{3}
}
func (m *CommandBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommandBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommandBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommandBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandBatch.Merge(m, src)
}
func (m *CommandBatch) XXX_Size() int {
	return m.Size()
}
func (m *CommandBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CommandBatch proto.InternalMessageInfo

func (m *CommandBatch) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *CommandBatch) GetCommands() []ScheduleCommand {
	if m != nil {
		return m.Commands
	}
	return nil
}

func init() {
	proto.RegisterEnum("hakeeper.ConfigChangeType", ConfigChangeType_name, ConfigChangeType_value)
	proto.RegisterEnum("hakeeper.ServiceType", ServiceType_name, ServiceType_value)
	proto.RegisterType((*Replica)(nil), "hakeeper.Replica")
	proto.RegisterType((*ConfigChange)(nil), "hakeeper.ConfigChange")
	proto.RegisterType((*ScheduleCommand)(nil), "hakeeper.ScheduleCommand")
	proto.RegisterType((*CommandBatch)(nil), "hakeeper.CommandBatch")
}

func init() { proto.RegisterFile("hakeeper.proto", fileDescriptor_5e1506f3aa5330eb) }

var fileDescriptor_5e1506f3aa5330eb = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0xae, 0xd2, 0x40,
	0x14, 0xed, 0x50, 0x14, 0xb8, 0x45, 0xac, 0x13, 0x35, 0x95, 0x98, 0x4a, 0xba, 0x22, 0x44, 0x21,
	0xd6, 0x85, 0x89, 0x6e, 0x14, 0x70, 0x41, 0x42, 0x5c, 0xb4, 0x10, 0x97, 0xa6, 0xb4, 0x43, 0xdb,
	0x48, 0x3b, 0x4d, 0x2d, 0x24, 0x2e, 0xfc, 0x17, 0x3f, 0x87, 0x25, 0x5f, 0x60, 0xde, 0xe3, 0x4b,
	0x5e, 0x7a, 0x3b, 0x94, 0xf6, 0xe5, 0xed, 0xee, 0xe9, 0x3d, 0xe7, 0x9e, 0x73, 0x26, 0x85, 0x5e,
	0xe0, 0xfc, 0x62, 0x2c, 0x61, 0xe9, 0x38, 0x49, 0x79, 0xc6, 0x69, 0xfb, 0x82, 0xfb, 0xef, 0xfc,
	0x30, 0x0b, 0xf6, 0x9b, 0xb1, 0xcb, 0xa3, 0x89, 0xcf, 0x7d, 0x3e, 0x41, 0xc2, 0x66, 0xbf, 0x45,
	0x84, 0x00, 0xa7, 0x42, 0x68, 0xfc, 0x80, 0x96, 0xc5, 0x92, 0x5d, 0xe8, 0x3a, 0x54, 0x83, 0x96,
	0x1d, 0x38, 0xa9, 0xb7, 0x98, 0x6b, 0x64, 0x40, 0x86, 0x4d, 0xeb, 0x02, 0xe9, 0x6b, 0xe8, 0x08,
	0xd2, 0x62, 0xae, 0x35, 0x70, 0x77, 0xfd, 0x40, 0x9f, 0xc3, 0xa3, 0x6f, 0x09, 0x77, 0x03, 0x4d,
	0xc6, 0x4d, 0x01, 0x8c, 0xbf, 0xd0, 0x9d, 0xf1, 0x78, 0x1b, 0xfa, 0xb3, 0xc0, 0x89, 0x7d, 0x46,
	0xdf, 0x97, 0x46, 0x78, 0x5d, 0x31, 0x9f, 0x8d, 0xcb, 0x0e, 0x62, 0x31, 0x6d, 0x1e, 0xff, 0xbf,
	0x91, 0xac, 0x32, 0xd0, 0x27, 0x80, 0x42, 0xbc, 0xfa, 0x93, 0x30, 0xf4, 0xed, 0x99, 0xfd, 0xab,
	0xaa, 0x7a, 0x3e, 0x67, 0x58, 0x15, 0xb6, 0xf1, 0x8f, 0xc0, 0x53, 0xdb, 0x0d, 0x98, 0xb7, 0xdf,
	0xb1, 0x19, 0x8f, 0x22, 0x27, 0xf6, 0x28, 0x85, 0xe6, 0x7a, 0x2d, 0xda, 0x75, 0x2c, 0x9c, 0xe9,
	0x97, 0x7a, 0x4c, 0x74, 0x51, 0xcc, 0x97, 0x0f, 0xbb, 0x88, 0x80, 0xf5, 0x62, 0x1f, 0x41, 0xb1,
	0x59, 0x7a, 0x08, 0xdd, 0x22, 0xa6, 0x8c, 0x31, 0x5f, 0x5c, 0x0f, 0x54, 0x96, 0x56, 0x95, 0x69,
	0xfc, 0xcc, 0xad, 0x31, 0xd9, 0xd4, 0xc9, 0xdc, 0x20, 0x8f, 0xb7, 0x62, 0x69, 0x24, 0x1e, 0x1f,
	0x67, 0xfa, 0x19, 0xda, 0x82, 0xf3, 0x5b, 0x6b, 0x0c, 0xe4, 0xa1, 0x62, 0xbe, 0xaa, 0x5c, 0xae,
	0xf7, 0x13, 0xe9, 0x4a, 0xc1, 0x68, 0x09, 0xea, 0xfd, 0x37, 0xa2, 0x0a, 0xb4, 0xbe, 0x7a, 0xde,
	0x77, 0xee, 0x31, 0x55, 0xa2, 0x3d, 0x00, 0x8b, 0x45, 0xfc, 0xc0, 0x10, 0x13, 0xfa, 0x04, 0x3a,
	0x76, 0xe6, 0xa4, 0x19, 0xc2, 0x06, 0xed, 0x42, 0xdb, 0xce, 0x78, 0x82, 0x48, 0x1e, 0xbd, 0xad,
	0xf5, 0xcc, 0xb5, 0x4b, 0xee, 0x8b, 0x2f, 0xaa, 0x94, 0x6b, 0xe7, 0xf1, 0x05, 0x92, 0xa9, 0x7a,
	0xba, 0xd5, 0xa5, 0xe3, 0x59, 0x27, 0xa7, 0xb3, 0x4e, 0x6e, 0xce, 0x3a, 0xd9, 0x3c, 0xc6, 0x1f,
	0xee, 0xc3, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x25, 0xdc, 0x25, 0xbb, 0x02, 0x00, 0x00,
}

func (m *Replica) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Replica) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Replica) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Epoch != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if m.ReplicaID != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.ReplicaID))
		i--
		dAtA[i] = 0x10
	}
	if m.ShardID != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.ShardID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConfigChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ChangeType != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.ChangeType))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Replica.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHakeeper(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduleCommand) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleCommand) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleCommand) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ServiceType != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.ServiceType))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.ConfigChange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHakeeper(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintHakeeper(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommandBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommandBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommandBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Commands) > 0 {
		for iNdEx := len(m.Commands) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commands[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHakeeper(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Term != 0 {
		i = encodeVarintHakeeper(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHakeeper(dAtA []byte, offset int, v uint64) int {
	offset -= sovHakeeper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Replica) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ShardID != 0 {
		n += 1 + sovHakeeper(uint64(m.ShardID))
	}
	if m.ReplicaID != 0 {
		n += 1 + sovHakeeper(uint64(m.ReplicaID))
	}
	if m.Epoch != 0 {
		n += 1 + sovHakeeper(uint64(m.Epoch))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Replica.Size()
	n += 1 + l + sovHakeeper(uint64(l))
	if m.ChangeType != 0 {
		n += 1 + sovHakeeper(uint64(m.ChangeType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ScheduleCommand) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovHakeeper(uint64(l))
	}
	l = m.ConfigChange.Size()
	n += 1 + l + sovHakeeper(uint64(l))
	if m.ServiceType != 0 {
		n += 1 + sovHakeeper(uint64(m.ServiceType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommandBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Term != 0 {
		n += 1 + sovHakeeper(uint64(m.Term))
	}
	if len(m.Commands) > 0 {
		for _, e := range m.Commands {
			l = e.Size()
			n += 1 + l + sovHakeeper(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHakeeper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHakeeper(x uint64) (n int) {
	return sovHakeeper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Replica) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHakeeper
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
			return fmt.Errorf("proto: Replica: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Replica: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardID", wireType)
			}
			m.ShardID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShardID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaID", wireType)
			}
			m.ReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicaID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHakeeper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHakeeper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHakeeper
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
			return fmt.Errorf("proto: ConfigChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
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
				return ErrInvalidLengthHakeeper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHakeeper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Replica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeType", wireType)
			}
			m.ChangeType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChangeType |= ConfigChangeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHakeeper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHakeeper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleCommand) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHakeeper
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
			return fmt.Errorf("proto: ScheduleCommand: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleCommand: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
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
				return ErrInvalidLengthHakeeper
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHakeeper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigChange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
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
				return ErrInvalidLengthHakeeper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHakeeper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConfigChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceType", wireType)
			}
			m.ServiceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceType |= ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHakeeper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHakeeper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommandBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHakeeper
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
			return fmt.Errorf("proto: CommandBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommandBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commands", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHakeeper
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
				return ErrInvalidLengthHakeeper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHakeeper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commands = append(m.Commands, ScheduleCommand{})
			if err := m.Commands[len(m.Commands)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHakeeper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHakeeper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHakeeper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHakeeper
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
					return 0, ErrIntOverflowHakeeper
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
					return 0, ErrIntOverflowHakeeper
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
				return 0, ErrInvalidLengthHakeeper
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHakeeper
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHakeeper
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHakeeper        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHakeeper          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHakeeper = fmt.Errorf("proto: unexpected end of group")
)
