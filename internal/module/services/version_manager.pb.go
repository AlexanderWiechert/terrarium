// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pb/terrarium/module/services/version_manager.proto

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

type TerminateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *module.Module `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *TerminateVersionRequest) Reset() {
	*x = TerminateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateVersionRequest) ProtoMessage() {}

func (x *TerminateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateVersionRequest.ProtoReflect.Descriptor instead.
func (*TerminateVersionRequest) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_version_manager_proto_rawDescGZIP(), []int{0}
}

func (x *TerminateVersionRequest) GetModule() *module.Module {
	if x != nil {
		return x.Module
	}
	return nil
}

type ListModuleVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *ListModuleVersionsRequest) Reset() {
	*x = ListModuleVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModuleVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModuleVersionsRequest) ProtoMessage() {}

func (x *ListModuleVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModuleVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListModuleVersionsRequest) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_version_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ListModuleVersionsRequest) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

type ListModuleVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []string `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *ListModuleVersionsResponse) Reset() {
	*x = ListModuleVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModuleVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModuleVersionsResponse) ProtoMessage() {}

func (x *ListModuleVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_terrarium_module_services_version_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModuleVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListModuleVersionsResponse) Descriptor() ([]byte, []int) {
	return file_pb_terrarium_module_services_version_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ListModuleVersionsResponse) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_pb_terrarium_module_services_version_manager_proto protoreflect.FileDescriptor

var file_pb_terrarium_module_services_version_manager_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x20, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4b, 0x0a, 0x17, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x33,
	0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa9, 0x03,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_terrarium_module_services_version_manager_proto_rawDescOnce sync.Once
	file_pb_terrarium_module_services_version_manager_proto_rawDescData = file_pb_terrarium_module_services_version_manager_proto_rawDesc
)

func file_pb_terrarium_module_services_version_manager_proto_rawDescGZIP() []byte {
	file_pb_terrarium_module_services_version_manager_proto_rawDescOnce.Do(func() {
		file_pb_terrarium_module_services_version_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_terrarium_module_services_version_manager_proto_rawDescData)
	})
	return file_pb_terrarium_module_services_version_manager_proto_rawDescData
}

var file_pb_terrarium_module_services_version_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_terrarium_module_services_version_manager_proto_goTypes = []interface{}{
	(*TerminateVersionRequest)(nil),    // 0: terrarium.module.services.TerminateVersionRequest
	(*ListModuleVersionsRequest)(nil),  // 1: terrarium.module.services.ListModuleVersionsRequest
	(*ListModuleVersionsResponse)(nil), // 2: terrarium.module.services.ListModuleVersionsResponse
	(*module.Module)(nil),              // 3: terrarium.module.Module
	(*module.BeginVersionRequest)(nil), // 4: terrarium.module.BeginVersionRequest
	(*module.Response)(nil),            // 5: terrarium.module.Response
}
var file_pb_terrarium_module_services_version_manager_proto_depIdxs = []int32{
	3, // 0: terrarium.module.services.TerminateVersionRequest.module:type_name -> terrarium.module.Module
	4, // 1: terrarium.module.services.VersionManager.BeginVersion:input_type -> terrarium.module.BeginVersionRequest
	0, // 2: terrarium.module.services.VersionManager.AbortVersion:input_type -> terrarium.module.services.TerminateVersionRequest
	0, // 3: terrarium.module.services.VersionManager.PublishVersion:input_type -> terrarium.module.services.TerminateVersionRequest
	1, // 4: terrarium.module.services.VersionManager.ListModuleVersions:input_type -> terrarium.module.services.ListModuleVersionsRequest
	5, // 5: terrarium.module.services.VersionManager.BeginVersion:output_type -> terrarium.module.Response
	5, // 6: terrarium.module.services.VersionManager.AbortVersion:output_type -> terrarium.module.Response
	5, // 7: terrarium.module.services.VersionManager.PublishVersion:output_type -> terrarium.module.Response
	2, // 8: terrarium.module.services.VersionManager.ListModuleVersions:output_type -> terrarium.module.services.ListModuleVersionsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_terrarium_module_services_version_manager_proto_init() }
func file_pb_terrarium_module_services_version_manager_proto_init() {
	if File_pb_terrarium_module_services_version_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_terrarium_module_services_version_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateVersionRequest); i {
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
		file_pb_terrarium_module_services_version_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModuleVersionsRequest); i {
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
		file_pb_terrarium_module_services_version_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModuleVersionsResponse); i {
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
			RawDescriptor: file_pb_terrarium_module_services_version_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_terrarium_module_services_version_manager_proto_goTypes,
		DependencyIndexes: file_pb_terrarium_module_services_version_manager_proto_depIdxs,
		MessageInfos:      file_pb_terrarium_module_services_version_manager_proto_msgTypes,
	}.Build()
	File_pb_terrarium_module_services_version_manager_proto = out.File
	file_pb_terrarium_module_services_version_manager_proto_rawDesc = nil
	file_pb_terrarium_module_services_version_manager_proto_goTypes = nil
	file_pb_terrarium_module_services_version_manager_proto_depIdxs = nil
}
