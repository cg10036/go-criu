// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: sk-packet.proto

package sk_packet

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

type ScmEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   *uint32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Rights []uint32 `protobuf:"varint,2,rep,name=rights" json:"rights,omitempty"`
}

func (x *ScmEntry) Reset() {
	*x = ScmEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sk_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScmEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScmEntry) ProtoMessage() {}

func (x *ScmEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScmEntry.ProtoReflect.Descriptor instead.
func (*ScmEntry) Descriptor() ([]byte, []int) {
	return file_sk_packet_proto_rawDescGZIP(), []int{0}
}

func (x *ScmEntry) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ScmEntry) GetRights() []uint32 {
	if x != nil {
		return x.Rights
	}
	return nil
}

type SkPacketEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdFor  *uint32 `protobuf:"varint,1,req,name=id_for,json=idFor" json:"id_for,omitempty"`
	Length *uint32 `protobuf:"varint,2,req,name=length" json:"length,omitempty"`
	// Reserved for message address
	// optional bytes		addr	= 3;
	Scm []*ScmEntry `protobuf:"bytes,4,rep,name=scm" json:"scm,omitempty"`
}

func (x *SkPacketEntry) Reset() {
	*x = SkPacketEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sk_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkPacketEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkPacketEntry) ProtoMessage() {}

func (x *SkPacketEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkPacketEntry.ProtoReflect.Descriptor instead.
func (*SkPacketEntry) Descriptor() ([]byte, []int) {
	return file_sk_packet_proto_rawDescGZIP(), []int{1}
}

func (x *SkPacketEntry) GetIdFor() uint32 {
	if x != nil && x.IdFor != nil {
		return *x.IdFor
	}
	return 0
}

func (x *SkPacketEntry) GetLength() uint32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *SkPacketEntry) GetScm() []*ScmEntry {
	if x != nil {
		return x.Scm
	}
	return nil
}

var File_sk_packet_proto protoreflect.FileDescriptor

var file_sk_packet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x6b, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x37, 0x0a, 0x09, 0x73, 0x63, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x5e, 0x0a, 0x0f, 0x73, 0x6b,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x69,
	0x64, 0x46, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x03,
	0x73, 0x63, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x73, 0x63, 0x6d, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x73, 0x63, 0x6d,
}

var (
	file_sk_packet_proto_rawDescOnce sync.Once
	file_sk_packet_proto_rawDescData = file_sk_packet_proto_rawDesc
)

func file_sk_packet_proto_rawDescGZIP() []byte {
	file_sk_packet_proto_rawDescOnce.Do(func() {
		file_sk_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_sk_packet_proto_rawDescData)
	})
	return file_sk_packet_proto_rawDescData
}

var file_sk_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sk_packet_proto_goTypes = []interface{}{
	(*ScmEntry)(nil),      // 0: scm_entry
	(*SkPacketEntry)(nil), // 1: sk_packet_entry
}
var file_sk_packet_proto_depIdxs = []int32{
	0, // 0: sk_packet_entry.scm:type_name -> scm_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sk_packet_proto_init() }
func file_sk_packet_proto_init() {
	if File_sk_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sk_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScmEntry); i {
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
		file_sk_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkPacketEntry); i {
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
			RawDescriptor: file_sk_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sk_packet_proto_goTypes,
		DependencyIndexes: file_sk_packet_proto_depIdxs,
		MessageInfos:      file_sk_packet_proto_msgTypes,
	}.Build()
	File_sk_packet_proto = out.File
	file_sk_packet_proto_rawDesc = nil
	file_sk_packet_proto_goTypes = nil
	file_sk_packet_proto_depIdxs = nil
}
