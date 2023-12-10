// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: brood.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Spore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Topic     string                 `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Data      string                 `protobuf:"bytes,20,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Spore) Reset() {
	*x = Spore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spore) ProtoMessage() {}

func (x *Spore) ProtoReflect() protoreflect.Message {
	mi := &file_brood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spore.ProtoReflect.Descriptor instead.
func (*Spore) Descriptor() ([]byte, []int) {
	return file_brood_proto_rawDescGZIP(), []int{0}
}

func (x *Spore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Spore) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Spore) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Spore) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type JoinBrood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hive   string   `protobuf:"bytes,1,opt,name=hive,proto3" json:"hive,omitempty"`
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *JoinBrood) Reset() {
	*x = JoinBrood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinBrood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinBrood) ProtoMessage() {}

func (x *JoinBrood) ProtoReflect() protoreflect.Message {
	mi := &file_brood_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinBrood.ProtoReflect.Descriptor instead.
func (*JoinBrood) Descriptor() ([]byte, []int) {
	return file_brood_proto_rawDescGZIP(), []int{1}
}

func (x *JoinBrood) GetHive() string {
	if x != nil {
		return x.Hive
	}
	return ""
}

func (x *JoinBrood) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

var File_brood_proto protoreflect.FileDescriptor

var file_brood_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x72, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa6, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x55, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1b, 0x92, 0x41, 0x18, 0x4a, 0x16,
	0x22, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x31, 0x32, 0x2d, 0x31, 0x30, 0x54, 0x30, 0x30, 0x3a, 0x30,
	0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x0c, 0x92, 0x41, 0x09,
	0x0a, 0x07, 0x2a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x65, 0x22, 0x49, 0x0a, 0x09, 0x4a, 0x6f, 0x69,
	0x6e, 0x42, 0x72, 0x6f, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x3a, 0x10, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0x2a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x42,
	0x72, 0x6f, 0x6f, 0x64, 0x32, 0x54, 0x0a, 0x05, 0x62, 0x72, 0x6f, 0x6f, 0x64, 0x12, 0x37, 0x0a,
	0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x42,
	0x72, 0x6f, 0x6f, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f,
	0x6a, 0x6f, 0x69, 0x6e, 0x30, 0x01, 0x1a, 0x12, 0x92, 0x41, 0x0f, 0x12, 0x0d, 0x62, 0x72, 0x6f,
	0x6f, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0xb9, 0x01, 0x92, 0x41, 0x99,
	0x01, 0x12, 0x4a, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x6f, 0x64, 0x20, 0x41, 0x50, 0x49, 0x22, 0x38,
	0x0a, 0x06, 0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x12, 0x1f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73,
	0x73, 0x63, 0x62, 0x2f, 0x62, 0x72, 0x6f, 0x6f, 0x64, 0x1a, 0x0d, 0x64, 0x65, 0x76, 0x40, 0x69,
	0x74, 0x73, 0x73, 0x63, 0x62, 0x2e, 0x64, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x73, 0x63, 0x62, 0x2f, 0x62, 0x72,
	0x6f, 0x6f, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brood_proto_rawDescOnce sync.Once
	file_brood_proto_rawDescData = file_brood_proto_rawDesc
)

func file_brood_proto_rawDescGZIP() []byte {
	file_brood_proto_rawDescOnce.Do(func() {
		file_brood_proto_rawDescData = protoimpl.X.CompressGZIP(file_brood_proto_rawDescData)
	})
	return file_brood_proto_rawDescData
}

var file_brood_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_brood_proto_goTypes = []interface{}{
	(*Spore)(nil),                 // 0: pb.Spore
	(*JoinBrood)(nil),             // 1: pb.JoinBrood
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_brood_proto_depIdxs = []int32{
	2, // 0: pb.Spore.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: pb.brood.Join:input_type -> pb.JoinBrood
	0, // 2: pb.brood.Join:output_type -> pb.Spore
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_brood_proto_init() }
func file_brood_proto_init() {
	if File_brood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spore); i {
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
		file_brood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinBrood); i {
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
			RawDescriptor: file_brood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brood_proto_goTypes,
		DependencyIndexes: file_brood_proto_depIdxs,
		MessageInfos:      file_brood_proto_msgTypes,
	}.Build()
	File_brood_proto = out.File
	file_brood_proto_rawDesc = nil
	file_brood_proto_goTypes = nil
	file_brood_proto_depIdxs = nil
}