// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: riskscore.proto

package riskscoreProto

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

type GetScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Cr string `protobuf:"bytes,2,opt,name=cr,proto3" json:"cr,omitempty"`
}

func (x *GetScoreReq) Reset() {
	*x = GetScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riskscore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreReq) ProtoMessage() {}

func (x *GetScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_riskscore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoreReq.ProtoReflect.Descriptor instead.
func (*GetScoreReq) Descriptor() ([]byte, []int) {
	return file_riskscore_proto_rawDescGZIP(), []int{0}
}

func (x *GetScoreReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetScoreReq) GetCr() string {
	if x != nil {
		return x.Cr
	}
	return ""
}

type GetScoreResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score int32 `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *GetScoreResp) Reset() {
	*x = GetScoreResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_riskscore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoreResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoreResp) ProtoMessage() {}

func (x *GetScoreResp) ProtoReflect() protoreflect.Message {
	mi := &file_riskscore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoreResp.ProtoReflect.Descriptor instead.
func (*GetScoreResp) Descriptor() ([]byte, []int) {
	return file_riskscore_proto_rawDescGZIP(), []int{1}
}

func (x *GetScoreResp) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_riskscore_proto protoreflect.FileDescriptor

var file_riskscore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x2d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x63,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x72, 0x22, 0x24, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x32, 0x4a, 0x0a, 0x09, 0x52, 0x69, 0x73, 0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x69, 0x73,
	0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x24, 0x5a,
	0x22, 0x61, 0x72, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x2f, 0x72, 0x69, 0x73, 0x6b, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_riskscore_proto_rawDescOnce sync.Once
	file_riskscore_proto_rawDescData = file_riskscore_proto_rawDesc
)

func file_riskscore_proto_rawDescGZIP() []byte {
	file_riskscore_proto_rawDescOnce.Do(func() {
		file_riskscore_proto_rawDescData = protoimpl.X.CompressGZIP(file_riskscore_proto_rawDescData)
	})
	return file_riskscore_proto_rawDescData
}

var file_riskscore_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_riskscore_proto_goTypes = []interface{}{
	(*GetScoreReq)(nil),  // 0: riskscore.GetScoreReq
	(*GetScoreResp)(nil), // 1: riskscore.GetScoreResp
}
var file_riskscore_proto_depIdxs = []int32{
	0, // 0: riskscore.Riskscore.GetScore:input_type -> riskscore.GetScoreReq
	1, // 1: riskscore.Riskscore.GetScore:output_type -> riskscore.GetScoreResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_riskscore_proto_init() }
func file_riskscore_proto_init() {
	if File_riskscore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_riskscore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoreReq); i {
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
		file_riskscore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoreResp); i {
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
			RawDescriptor: file_riskscore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_riskscore_proto_goTypes,
		DependencyIndexes: file_riskscore_proto_depIdxs,
		MessageInfos:      file_riskscore_proto_msgTypes,
	}.Build()
	File_riskscore_proto = out.File
	file_riskscore_proto_rawDesc = nil
	file_riskscore_proto_goTypes = nil
	file_riskscore_proto_depIdxs = nil
}