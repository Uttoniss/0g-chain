// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: light/light.proto

package light

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

// SampleRequest contains the blob to sample (by batch and blob index) and required sample times
type SampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId        []byte `protobuf:"bytes,1,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	BatchHeaderHash []byte `protobuf:"bytes,2,opt,name=batch_header_hash,json=batchHeaderHash,proto3" json:"batch_header_hash,omitempty"`
	BlobIndex       uint32 `protobuf:"varint,3,opt,name=blob_index,json=blobIndex,proto3" json:"blob_index,omitempty"`
	Times           uint32 `protobuf:"varint,4,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *SampleRequest) Reset() {
	*x = SampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_light_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequest) ProtoMessage() {}

func (x *SampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_light_light_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequest.ProtoReflect.Descriptor instead.
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return file_light_light_proto_rawDescGZIP(), []int{0}
}

func (x *SampleRequest) GetStreamId() []byte {
	if x != nil {
		return x.StreamId
	}
	return nil
}

func (x *SampleRequest) GetBatchHeaderHash() []byte {
	if x != nil {
		return x.BatchHeaderHash
	}
	return nil
}

func (x *SampleRequest) GetBlobIndex() uint32 {
	if x != nil {
		return x.BlobIndex
	}
	return 0
}

func (x *SampleRequest) GetTimes() uint32 {
	if x != nil {
		return x.Times
	}
	return 0
}

// SampleReply contains the sample result
type SampleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SampleReply) Reset() {
	*x = SampleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_light_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleReply) ProtoMessage() {}

func (x *SampleReply) ProtoReflect() protoreflect.Message {
	mi := &file_light_light_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleReply.ProtoReflect.Descriptor instead.
func (*SampleReply) Descriptor() ([]byte, []int) {
	return file_light_light_proto_rawDescGZIP(), []int{1}
}

func (x *SampleReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchHeaderHash []byte `protobuf:"bytes,1,opt,name=batch_header_hash,json=batchHeaderHash,proto3" json:"batch_header_hash,omitempty"`
	BlobIndex       uint32 `protobuf:"varint,2,opt,name=blob_index,json=blobIndex,proto3" json:"blob_index,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_light_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_light_light_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_light_light_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveRequest) GetBatchHeaderHash() []byte {
	if x != nil {
		return x.BatchHeaderHash
	}
	return nil
}

func (x *RetrieveRequest) GetBlobIndex() uint32 {
	if x != nil {
		return x.BlobIndex
	}
	return 0
}

type RetrieveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RetrieveReply) Reset() {
	*x = RetrieveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_light_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReply) ProtoMessage() {}

func (x *RetrieveReply) ProtoReflect() protoreflect.Message {
	mi := &file_light_light_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveReply.ProtoReflect.Descriptor instead.
func (*RetrieveReply) Descriptor() ([]byte, []int) {
	return file_light_light_proto_rawDescGZIP(), []int{3}
}

func (x *RetrieveReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RetrieveReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_light_light_proto protoreflect.FileDescriptor

var file_light_light_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x62, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x79,
	0x0a, 0x05, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x30,
	0x67, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x2f, 0x72, 0x75, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_light_light_proto_rawDescOnce sync.Once
	file_light_light_proto_rawDescData = file_light_light_proto_rawDesc
)

func file_light_light_proto_rawDescGZIP() []byte {
	file_light_light_proto_rawDescOnce.Do(func() {
		file_light_light_proto_rawDescData = protoimpl.X.CompressGZIP(file_light_light_proto_rawDescData)
	})
	return file_light_light_proto_rawDescData
}

var file_light_light_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_light_light_proto_goTypes = []interface{}{
	(*SampleRequest)(nil),   // 0: light.SampleRequest
	(*SampleReply)(nil),     // 1: light.SampleReply
	(*RetrieveRequest)(nil), // 2: light.RetrieveRequest
	(*RetrieveReply)(nil),   // 3: light.RetrieveReply
}
var file_light_light_proto_depIdxs = []int32{
	0, // 0: light.Light.Sample:input_type -> light.SampleRequest
	2, // 1: light.Light.Retrieve:input_type -> light.RetrieveRequest
	1, // 2: light.Light.Sample:output_type -> light.SampleReply
	3, // 3: light.Light.Retrieve:output_type -> light.RetrieveReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_light_light_proto_init() }
func file_light_light_proto_init() {
	if File_light_light_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_light_light_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequest); i {
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
		file_light_light_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleReply); i {
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
		file_light_light_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_light_light_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveReply); i {
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
			RawDescriptor: file_light_light_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_light_light_proto_goTypes,
		DependencyIndexes: file_light_light_proto_depIdxs,
		MessageInfos:      file_light_light_proto_msgTypes,
	}.Build()
	File_light_light_proto = out.File
	file_light_light_proto_rawDesc = nil
	file_light_light_proto_goTypes = nil
	file_light_light_proto_depIdxs = nil
}