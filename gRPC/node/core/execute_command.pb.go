// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: execute_command.proto

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

type ExecuteMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Vals []string `protobuf:"bytes,2,rep,name=vals,proto3" json:"vals,omitempty"`
}

func (x *ExecuteMsg) Reset() {
	*x = ExecuteMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteMsg) ProtoMessage() {}

func (x *ExecuteMsg) ProtoReflect() protoreflect.Message {
	mi := &file_execute_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteMsg.ProtoReflect.Descriptor instead.
func (*ExecuteMsg) Descriptor() ([]byte, []int) {
	return file_execute_command_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteMsg) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *ExecuteMsg) GetVals() []string {
	if x != nil {
		return x.Vals
	}
	return nil
}

type ExecuteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ExecuteRes) Reset() {
	*x = ExecuteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRes) ProtoMessage() {}

func (x *ExecuteRes) ProtoReflect() protoreflect.Message {
	mi := &file_execute_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRes.ProtoReflect.Descriptor instead.
func (*ExecuteRes) Descriptor() ([]byte, []int) {
	return file_execute_command_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ExecuteRes) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_execute_command_proto protoreflect.FileDescriptor

var file_execute_command_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x32, 0x0a,
	0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x76, 0x61, 0x6c,
	0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x41, 0x0a, 0x0b, 0x43, 0x6d, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x32,
	0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6d, 0x64, 0x12, 0x10, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x74, 0x75, 0x62, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_execute_command_proto_rawDescOnce sync.Once
	file_execute_command_proto_rawDescData = file_execute_command_proto_rawDesc
)

func file_execute_command_proto_rawDescGZIP() []byte {
	file_execute_command_proto_rawDescOnce.Do(func() {
		file_execute_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_execute_command_proto_rawDescData)
	})
	return file_execute_command_proto_rawDescData
}

var file_execute_command_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_execute_command_proto_goTypes = []interface{}{
	(*ExecuteMsg)(nil), // 0: core.ExecuteMsg
	(*ExecuteRes)(nil), // 1: core.ExecuteRes
}
var file_execute_command_proto_depIdxs = []int32{
	0, // 0: core.CmdExecutor.ExecuteCmd:input_type -> core.ExecuteMsg
	1, // 1: core.CmdExecutor.ExecuteCmd:output_type -> core.ExecuteRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_execute_command_proto_init() }
func file_execute_command_proto_init() {
	if File_execute_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execute_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteMsg); i {
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
		file_execute_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRes); i {
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
			RawDescriptor: file_execute_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_execute_command_proto_goTypes,
		DependencyIndexes: file_execute_command_proto_depIdxs,
		MessageInfos:      file_execute_command_proto_msgTypes,
	}.Build()
	File_execute_command_proto = out.File
	file_execute_command_proto_rawDesc = nil
	file_execute_command_proto_goTypes = nil
	file_execute_command_proto_depIdxs = nil
}
