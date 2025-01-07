// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: seccomp.proto

package seccomp

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

type SeccompFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter []byte  `protobuf:"bytes,1,req,name=filter" json:"filter,omitempty"`
	Prev   *uint32 `protobuf:"varint,2,opt,name=prev" json:"prev,omitempty"`
	Flags  *uint32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
}

func (x *SeccompFilter) Reset() {
	*x = SeccompFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seccomp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeccompFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeccompFilter) ProtoMessage() {}

func (x *SeccompFilter) ProtoReflect() protoreflect.Message {
	mi := &file_seccomp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeccompFilter.ProtoReflect.Descriptor instead.
func (*SeccompFilter) Descriptor() ([]byte, []int) {
	return file_seccomp_proto_rawDescGZIP(), []int{0}
}

func (x *SeccompFilter) GetFilter() []byte {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *SeccompFilter) GetPrev() uint32 {
	if x != nil && x.Prev != nil {
		return *x.Prev
	}
	return 0
}

func (x *SeccompFilter) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

type SeccompEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeccompFilters []*SeccompFilter `protobuf:"bytes,1,rep,name=seccomp_filters,json=seccompFilters" json:"seccomp_filters,omitempty"`
}

func (x *SeccompEntry) Reset() {
	*x = SeccompEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seccomp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeccompEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeccompEntry) ProtoMessage() {}

func (x *SeccompEntry) ProtoReflect() protoreflect.Message {
	mi := &file_seccomp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeccompEntry.ProtoReflect.Descriptor instead.
func (*SeccompEntry) Descriptor() ([]byte, []int) {
	return file_seccomp_proto_rawDescGZIP(), []int{1}
}

func (x *SeccompEntry) GetSeccompFilters() []*SeccompFilter {
	if x != nil {
		return x.SeccompFilters
	}
	return nil
}

var File_seccomp_proto protoreflect.FileDescriptor

var file_seccomp_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x52, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0c, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x22, 0x49, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0e,
	0x73, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
}

var (
	file_seccomp_proto_rawDescOnce sync.Once
	file_seccomp_proto_rawDescData = file_seccomp_proto_rawDesc
)

func file_seccomp_proto_rawDescGZIP() []byte {
	file_seccomp_proto_rawDescOnce.Do(func() {
		file_seccomp_proto_rawDescData = protoimpl.X.CompressGZIP(file_seccomp_proto_rawDescData)
	})
	return file_seccomp_proto_rawDescData
}

var file_seccomp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_seccomp_proto_goTypes = []interface{}{
	(*SeccompFilter)(nil), // 0: seccomp_filter
	(*SeccompEntry)(nil),  // 1: seccomp_entry
}
var file_seccomp_proto_depIdxs = []int32{
	0, // 0: seccomp_entry.seccomp_filters:type_name -> seccomp_filter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_seccomp_proto_init() }
func file_seccomp_proto_init() {
	if File_seccomp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seccomp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeccompFilter); i {
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
		file_seccomp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeccompEntry); i {
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
			RawDescriptor: file_seccomp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_seccomp_proto_goTypes,
		DependencyIndexes: file_seccomp_proto_depIdxs,
		MessageInfos:      file_seccomp_proto_msgTypes,
	}.Build()
	File_seccomp_proto = out.File
	file_seccomp_proto_rawDesc = nil
	file_seccomp_proto_goTypes = nil
	file_seccomp_proto_depIdxs = nil
}
