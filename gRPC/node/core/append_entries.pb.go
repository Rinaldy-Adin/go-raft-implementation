// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: node/core/append_entries.proto

package gRPC

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppendEntriesArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term          int32                            `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderAddress *AppendEntriesArgs_LeaderAddress `protobuf:"bytes,2,opt,name=leaderAddress,proto3" json:"leaderAddress,omitempty"`
	PrevLogIndex  int32                            `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm   int32                            `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries       []*AppendEntriesArgs_LogEntry    `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit  int32                            `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (x *AppendEntriesArgs) Reset() {
	*x = AppendEntriesArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_append_entries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesArgs) ProtoMessage() {}

func (x *AppendEntriesArgs) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_append_entries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesArgs.ProtoReflect.Descriptor instead.
func (*AppendEntriesArgs) Descriptor() ([]byte, []int) {
	return file_node_core_append_entries_proto_rawDescGZIP(), []int{0}
}

func (x *AppendEntriesArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesArgs) GetLeaderAddress() *AppendEntriesArgs_LeaderAddress {
	if x != nil {
		return x.LeaderAddress
	}
	return nil
}

func (x *AppendEntriesArgs) GetPrevLogIndex() int32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesArgs) GetPrevLogTerm() int32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesArgs) GetEntries() []*AppendEntriesArgs_LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesArgs) GetLeaderCommit() int32 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type AppendEntriesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Term    int32 `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *AppendEntriesReply) Reset() {
	*x = AppendEntriesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_append_entries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesReply) ProtoMessage() {}

func (x *AppendEntriesReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_append_entries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesReply.ProtoReflect.Descriptor instead.
func (*AppendEntriesReply) Descriptor() ([]byte, []int) {
	return file_node_core_append_entries_proto_rawDescGZIP(), []int{1}
}

func (x *AppendEntriesReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AppendEntriesReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type AppendEntriesArgs_LeaderAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AppendEntriesArgs_LeaderAddress) Reset() {
	*x = AppendEntriesArgs_LeaderAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_append_entries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesArgs_LeaderAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesArgs_LeaderAddress) ProtoMessage() {}

func (x *AppendEntriesArgs_LeaderAddress) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_append_entries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesArgs_LeaderAddress.ProtoReflect.Descriptor instead.
func (*AppendEntriesArgs_LeaderAddress) Descriptor() ([]byte, []int) {
	return file_node_core_append_entries_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AppendEntriesArgs_LeaderAddress) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AppendEntriesArgs_LeaderAddress) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type AppendEntriesArgs_LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32  `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Value   string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AppendEntriesArgs_LogEntry) Reset() {
	*x = AppendEntriesArgs_LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_append_entries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesArgs_LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesArgs_LogEntry) ProtoMessage() {}

func (x *AppendEntriesArgs_LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_append_entries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesArgs_LogEntry.ProtoReflect.Descriptor instead.
func (*AppendEntriesArgs_LogEntry) Descriptor() ([]byte, []int) {
	return file_node_core_append_entries_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AppendEntriesArgs_LogEntry) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesArgs_LogEntry) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *AppendEntriesArgs_LogEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_node_core_append_entries_proto protoreflect.FileDescriptor

var file_node_core_append_entries_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x9f, 0x03, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65, 0x6e,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x4b, 0x0a, 0x0d, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73,
	0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0d,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54,
	0x65, 0x72, 0x6d, 0x12, 0x3a, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x1a, 0x33, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x4e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x32, 0x5c, 0x0a, 0x14,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70,
	0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x74, 0x75,
	0x62, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f,
	0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_core_append_entries_proto_rawDescOnce sync.Once
	file_node_core_append_entries_proto_rawDescData = file_node_core_append_entries_proto_rawDesc
)

func file_node_core_append_entries_proto_rawDescGZIP() []byte {
	file_node_core_append_entries_proto_rawDescOnce.Do(func() {
		file_node_core_append_entries_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_core_append_entries_proto_rawDescData)
	})
	return file_node_core_append_entries_proto_rawDescData
}

var file_node_core_append_entries_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_node_core_append_entries_proto_goTypes = []interface{}{
	(*AppendEntriesArgs)(nil),               // 0: core.AppendEntriesArgs
	(*AppendEntriesReply)(nil),              // 1: core.AppendEntriesReply
	(*AppendEntriesArgs_LeaderAddress)(nil), // 2: core.AppendEntriesArgs.LeaderAddress
	(*AppendEntriesArgs_LogEntry)(nil),      // 3: core.AppendEntriesArgs.LogEntry
}
var file_node_core_append_entries_proto_depIdxs = []int32{
	2, // 0: core.AppendEntriesArgs.leaderAddress:type_name -> core.AppendEntriesArgs.LeaderAddress
	3, // 1: core.AppendEntriesArgs.entries:type_name -> core.AppendEntriesArgs.LogEntry
	0, // 2: core.AppendEntriesService.AppendEntries:input_type -> core.AppendEntriesArgs
	1, // 3: core.AppendEntriesService.AppendEntries:output_type -> core.AppendEntriesReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_core_append_entries_proto_init() }
func file_node_core_append_entries_proto_init() {
	if File_node_core_append_entries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_core_append_entries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_core_append_entries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_core_append_entries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesArgs_LeaderAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_core_append_entries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesArgs_LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_core_append_entries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_core_append_entries_proto_goTypes,
		DependencyIndexes: file_node_core_append_entries_proto_depIdxs,
		MessageInfos:      file_node_core_append_entries_proto_msgTypes,
	}.Build()
	File_node_core_append_entries_proto = out.File
	file_node_core_append_entries_proto_rawDesc = nil
	file_node_core_append_entries_proto_goTypes = nil
	file_node_core_append_entries_proto_depIdxs = nil
}
