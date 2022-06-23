// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/services/module_session_manager.proto

package services

import (
	terrarium "github.com/terrariumcloud/terrarium-grpc-gateway/pkg/terrarium"
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

type BeginVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *terrarium.VersionedModule `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *BeginVersionRequest) Reset() {
	*x = BeginVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_module_session_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginVersionRequest) ProtoMessage() {}

func (x *BeginVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_module_session_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginVersionRequest.ProtoReflect.Descriptor instead.
func (*BeginVersionRequest) Descriptor() ([]byte, []int) {
	return file_internal_services_module_session_manager_proto_rawDescGZIP(), []int{0}
}

func (x *BeginVersionRequest) GetModule() *terrarium.VersionedModule {
	if x != nil {
		return x.Module
	}
	return nil
}

type TerminateVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionKey string `protobuf:"bytes,1,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
}

func (x *TerminateVersionRequest) Reset() {
	*x = TerminateVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_services_module_session_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateVersionRequest) ProtoMessage() {}

func (x *TerminateVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_services_module_session_manager_proto_msgTypes[1]
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
	return file_internal_services_module_session_manager_proto_rawDescGZIP(), []int{1}
}

func (x *TerminateVersionRequest) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

var File_internal_services_module_session_manager_proto protoreflect.FileDescriptor

var file_internal_services_module_session_manager_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x1a, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x13, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x17, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x32, 0xdc, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6f, 0x0a, 0x0c, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x71, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_services_module_session_manager_proto_rawDescOnce sync.Once
	file_internal_services_module_session_manager_proto_rawDescData = file_internal_services_module_session_manager_proto_rawDesc
)

func file_internal_services_module_session_manager_proto_rawDescGZIP() []byte {
	file_internal_services_module_session_manager_proto_rawDescOnce.Do(func() {
		file_internal_services_module_session_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_services_module_session_manager_proto_rawDescData)
	})
	return file_internal_services_module_session_manager_proto_rawDescData
}

var file_internal_services_module_session_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_services_module_session_manager_proto_goTypes = []interface{}{
	(*BeginVersionRequest)(nil),                 // 0: terrarium.services.module.BeginVersionRequest
	(*TerminateVersionRequest)(nil),             // 1: terrarium.services.module.TerminateVersionRequest
	(*terrarium.VersionedModule)(nil),           // 2: terrarium.module.VersionedModule
	(*terrarium.BeginVersionResponse)(nil),      // 3: terrarium.module.BeginVersionResponse
	(*terrarium.TransactionStatusResponse)(nil), // 4: terrarium.module.TransactionStatusResponse
}
var file_internal_services_module_session_manager_proto_depIdxs = []int32{
	2, // 0: terrarium.services.module.BeginVersionRequest.module:type_name -> terrarium.module.VersionedModule
	0, // 1: terrarium.services.module.SessionManager.BeginVersion:input_type -> terrarium.services.module.BeginVersionRequest
	1, // 2: terrarium.services.module.SessionManager.AbortVersion:input_type -> terrarium.services.module.TerminateVersionRequest
	1, // 3: terrarium.services.module.SessionManager.PublishVersion:input_type -> terrarium.services.module.TerminateVersionRequest
	3, // 4: terrarium.services.module.SessionManager.BeginVersion:output_type -> terrarium.module.BeginVersionResponse
	4, // 5: terrarium.services.module.SessionManager.AbortVersion:output_type -> terrarium.module.TransactionStatusResponse
	4, // 6: terrarium.services.module.SessionManager.PublishVersion:output_type -> terrarium.module.TransactionStatusResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_services_module_session_manager_proto_init() }
func file_internal_services_module_session_manager_proto_init() {
	if File_internal_services_module_session_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_services_module_session_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginVersionRequest); i {
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
		file_internal_services_module_session_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_services_module_session_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_services_module_session_manager_proto_goTypes,
		DependencyIndexes: file_internal_services_module_session_manager_proto_depIdxs,
		MessageInfos:      file_internal_services_module_session_manager_proto_msgTypes,
	}.Build()
	File_internal_services_module_session_manager_proto = out.File
	file_internal_services_module_session_manager_proto_rawDesc = nil
	file_internal_services_module_session_manager_proto_goTypes = nil
	file_internal_services_module_session_manager_proto_depIdxs = nil
}