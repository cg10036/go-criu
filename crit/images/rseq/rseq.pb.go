// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: rseq.proto

package rseq

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

type RseqEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RseqAbiPointer *uint64 `protobuf:"varint,1,req,name=rseq_abi_pointer,json=rseqAbiPointer" json:"rseq_abi_pointer,omitempty"`
	RseqAbiSize    *uint32 `protobuf:"varint,2,req,name=rseq_abi_size,json=rseqAbiSize" json:"rseq_abi_size,omitempty"`
	Signature      *uint32 `protobuf:"varint,3,req,name=signature" json:"signature,omitempty"`
	RseqCsPointer  *uint64 `protobuf:"varint,4,opt,name=rseq_cs_pointer,json=rseqCsPointer" json:"rseq_cs_pointer,omitempty"`
}

func (x *RseqEntry) Reset() {
	*x = RseqEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rseq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RseqEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RseqEntry) ProtoMessage() {}

func (x *RseqEntry) ProtoReflect() protoreflect.Message {
	mi := &file_rseq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RseqEntry.ProtoReflect.Descriptor instead.
func (*RseqEntry) Descriptor() ([]byte, []int) {
	return file_rseq_proto_rawDescGZIP(), []int{0}
}

func (x *RseqEntry) GetRseqAbiPointer() uint64 {
	if x != nil && x.RseqAbiPointer != nil {
		return *x.RseqAbiPointer
	}
	return 0
}

func (x *RseqEntry) GetRseqAbiSize() uint32 {
	if x != nil && x.RseqAbiSize != nil {
		return *x.RseqAbiSize
	}
	return 0
}

func (x *RseqEntry) GetSignature() uint32 {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return 0
}

func (x *RseqEntry) GetRseqCsPointer() uint64 {
	if x != nil && x.RseqCsPointer != nil {
		return *x.RseqCsPointer
	}
	return 0
}

var File_rseq_proto protoreflect.FileDescriptor

var file_rseq_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x73, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a,
	0x0a, 0x72, 0x73, 0x65, 0x71, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x73, 0x65, 0x71, 0x5f, 0x61, 0x62, 0x69, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x73, 0x65, 0x71, 0x41, 0x62, 0x69, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x73, 0x65, 0x71, 0x5f, 0x61, 0x62,
	0x69, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x73,
	0x65, 0x71, 0x41, 0x62, 0x69, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x73, 0x65, 0x71, 0x5f,
	0x63, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0d, 0x72, 0x73, 0x65, 0x71, 0x43, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
}

var (
	file_rseq_proto_rawDescOnce sync.Once
	file_rseq_proto_rawDescData = file_rseq_proto_rawDesc
)

func file_rseq_proto_rawDescGZIP() []byte {
	file_rseq_proto_rawDescOnce.Do(func() {
		file_rseq_proto_rawDescData = protoimpl.X.CompressGZIP(file_rseq_proto_rawDescData)
	})
	return file_rseq_proto_rawDescData
}

var file_rseq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rseq_proto_goTypes = []interface{}{
	(*RseqEntry)(nil), // 0: rseq_entry
}
var file_rseq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rseq_proto_init() }
func file_rseq_proto_init() {
	if File_rseq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rseq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RseqEntry); i {
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
			RawDescriptor: file_rseq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rseq_proto_goTypes,
		DependencyIndexes: file_rseq_proto_depIdxs,
		MessageInfos:      file_rseq_proto_msgTypes,
	}.Build()
	File_rseq_proto = out.File
	file_rseq_proto_rawDesc = nil
	file_rseq_proto_goTypes = nil
	file_rseq_proto_depIdxs = nil
}
