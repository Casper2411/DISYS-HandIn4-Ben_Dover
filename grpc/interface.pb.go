// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: grpc/interface.proto

package DISYS_HandIn4_Ben_Dover

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LamportTimestamp int32 `protobuf:"varint,2,opt,name=lamportTimestamp,proto3" json:"lamportTimestamp,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetLamportTimestamp() int32 {
	if x != nil {
		return x.LamportTimestamp
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_grpc_interface_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_interface_proto protoreflect.FileDescriptor

var file_grpc_interface_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x52, 0x69, 0x63, 0x61, 0x72, 0x74, 0x41, 0x67,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x61, 0x22, 0x45, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x21, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x5b, 0x0a, 0x15, 0x52, 0x69, 0x63, 0x61, 0x72, 0x74, 0x41, 0x67, 0x72, 0x61, 0x77, 0x61,
	0x6c, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x52, 0x69, 0x63,
	0x61, 0x72, 0x74, 0x41, 0x67, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x61, 0x12, 0x17, 0x2e, 0x52, 0x69,
	0x63, 0x61, 0x72, 0x74, 0x41, 0x67, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x61, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x69, 0x63, 0x61, 0x72, 0x74, 0x41, 0x67, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x73, 0x70, 0x65, 0x72, 0x32, 0x34, 0x31, 0x31, 0x2f, 0x44,
	0x49, 0x53, 0x59, 0x53, 0x2d, 0x48, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x34, 0x2d, 0x42, 0x65, 0x6e,
	0x5f, 0x44, 0x6f, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_interface_proto_rawDescOnce sync.Once
	file_grpc_interface_proto_rawDescData = file_grpc_interface_proto_rawDesc
)

func file_grpc_interface_proto_rawDescGZIP() []byte {
	file_grpc_interface_proto_rawDescOnce.Do(func() {
		file_grpc_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_interface_proto_rawDescData)
	})
	return file_grpc_interface_proto_rawDescData
}

var file_grpc_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_interface_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: RicartAgrawala.Request
	(*Reply)(nil),   // 1: RicartAgrawala.Reply
}
var file_grpc_interface_proto_depIdxs = []int32{
	0, // 0: RicartAgrawala.RicartAgrawalaService.RicartAgrawala:input_type -> RicartAgrawala.Request
	1, // 1: RicartAgrawala.RicartAgrawalaService.RicartAgrawala:output_type -> RicartAgrawala.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_interface_proto_init() }
func file_grpc_interface_proto_init() {
	if File_grpc_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_grpc_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_interface_proto_goTypes,
		DependencyIndexes: file_grpc_interface_proto_depIdxs,
		MessageInfos:      file_grpc_interface_proto_msgTypes,
	}.Build()
	File_grpc_interface_proto = out.File
	file_grpc_interface_proto_rawDesc = nil
	file_grpc_interface_proto_goTypes = nil
	file_grpc_interface_proto_depIdxs = nil
}
