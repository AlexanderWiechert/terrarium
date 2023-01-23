// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/terrarium/module/services/tag_manager.proto

package services

import (
	module "github.com/terrariumcloud/terrarium/pkg/terrarium/module"
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

type PublishTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string   `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tags   []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *PublishTagRequest) Reset() {
	*x = PublishTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_tag_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishTagRequest) ProtoMessage() {}

func (x *PublishTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_tag_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishTagRequest.ProtoReflect.Descriptor instead.
func (*PublishTagRequest) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_tag_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PublishTagRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *PublishTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublishTagRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_tag_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_tag_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_tag_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_terrarium_module_services_tag_manager_proto protoreflect.FileDescriptor

var file_pb_terrarium_module_services_tag_manager_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74,
	0x61, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x20, 0x70, 0x62, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5b, 0x0a, 0x0a, 0x54, 0x61, 0x67,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x54, 0x61, 0x67, 0x12, 0x23, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_terrarium_module_services_tag_manager_proto_rawDescOnce sync.Once
	file_pb_terrarium_module_services_tag_manager_proto_rawDescData = file_pb_terrarium_module_services_tag_manager_proto_rawDesc
)

func file_pb_terrarium_module_services_tag_manager_proto_rawDescGZIP() []byte {
	file_pb_terrarium_module_services_tag_manager_proto_rawDescOnce.Do(func() {
		file_pb_terrarium_module_services_tag_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_terrarium_module_services_tag_manager_proto_rawDescData)
	})
	return file_pb_terrarium_module_services_tag_manager_proto_rawDescData
}

var file_pb_terrarium_module_services_tag_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_terrarium_module_services_tag_manager_proto_goTypes = []interface{}{
	(*PublishTagRequest)(nil),        // 0: terrarium.module.services.PublishTagRequest
	(*Response)(nil),                 // 1: terrarium.module.services.Response
	(*module.PublishTagRequest)(nil), // 2: terrarium.module.PublishTagRequest
	(*module.Response)(nil),          // 3: terrarium.module.Response
}
var file_pb_terrarium_module_services_tag_manager_proto_depIdxs = []int32{
	2, // 0: terrarium.module.services.TagManager.PublishTag:input_type -> terrarium.module.PublishTagRequest
	3, // 1: terrarium.module.services.TagManager.PublishTag:output_type -> terrarium.module.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_terrarium_module_services_tag_manager_proto_init() }
func file_pb_terrarium_module_services_tag_manager_proto_init() {
	if File_pb_terrarium_module_services_tag_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_terrarium_module_services_tag_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishTagRequest); i {
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
		file_pb_terrarium_module_services_tag_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_pb_terrarium_module_services_tag_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_terrarium_module_services_tag_manager_proto_goTypes,
		DependencyIndexes: file_pb_terrarium_module_services_tag_manager_proto_depIdxs,
		MessageInfos:      file_pb_terrarium_module_services_tag_manager_proto_msgTypes,
	}.Build()
	File_pb_terrarium_module_services_tag_manager_proto = out.File
	file_pb_terrarium_module_services_tag_manager_proto_rawDesc = nil
	file_pb_terrarium_module_services_tag_manager_proto_goTypes = nil
	file_pb_terrarium_module_services_tag_manager_proto_depIdxs = nil
}