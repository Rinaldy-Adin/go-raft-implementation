// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: node/core/request_vote_rpc.proto

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

type RequestVoteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  int32 `protobuf:"varint,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex int32 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  int32 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *RequestVoteArgs) Reset() {
	*x = RequestVoteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_request_vote_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteArgs) ProtoMessage() {}

func (x *RequestVoteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_request_vote_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteArgs.ProtoReflect.Descriptor instead.
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return file_node_core_request_vote_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteArgs) GetCandidateId() int32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *RequestVoteArgs) GetLastLogIndex() int32 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteArgs) GetLastLogTerm() int32 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_core_request_vote_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_core_request_vote_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_node_core_request_vote_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

var File_node_core_request_vote_rpc_proto protoreflect.FileDescriptor

var file_node_core_request_vote_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x48, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x32, 0x48, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x18, 0x5a, 0x16,
	0x74, 0x75, 0x62, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x66,
	0x74, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_core_request_vote_rpc_proto_rawDescOnce sync.Once
	file_node_core_request_vote_rpc_proto_rawDescData = file_node_core_request_vote_rpc_proto_rawDesc
)

func file_node_core_request_vote_rpc_proto_rawDescGZIP() []byte {
	file_node_core_request_vote_rpc_proto_rawDescOnce.Do(func() {
		file_node_core_request_vote_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_core_request_vote_rpc_proto_rawDescData)
	})
	return file_node_core_request_vote_rpc_proto_rawDescData
}

var file_node_core_request_vote_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_core_request_vote_rpc_proto_goTypes = []interface{}{
	(*RequestVoteArgs)(nil),  // 0: core.RequestVoteArgs
	(*RequestVoteReply)(nil), // 1: core.RequestVoteReply
}
var file_node_core_request_vote_rpc_proto_depIdxs = []int32{
	0, // 0: core.RaftNode.RequestVote:input_type -> core.RequestVoteArgs
	1, // 1: core.RaftNode.RequestVote:output_type -> core.RequestVoteReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_core_request_vote_rpc_proto_init() }
func file_node_core_request_vote_rpc_proto_init() {
	if File_node_core_request_vote_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_core_request_vote_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteArgs); i {
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
		file_node_core_request_vote_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteReply); i {
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
			RawDescriptor: file_node_core_request_vote_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_core_request_vote_rpc_proto_goTypes,
		DependencyIndexes: file_node_core_request_vote_rpc_proto_depIdxs,
		MessageInfos:      file_node_core_request_vote_rpc_proto_msgTypes,
	}.Build()
	File_node_core_request_vote_rpc_proto = out.File
	file_node_core_request_vote_rpc_proto_rawDesc = nil
	file_node_core_request_vote_rpc_proto_goTypes = nil
	file_node_core_request_vote_rpc_proto_depIdxs = nil
}
