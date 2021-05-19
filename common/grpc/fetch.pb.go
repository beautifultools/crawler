// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: common/grpc/fetch.proto

package grpc

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

type FetchURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchURL) Reset() {
	*x = FetchURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_grpc_fetch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchURL) ProtoMessage() {}

func (x *FetchURL) ProtoReflect() protoreflect.Message {
	mi := &file_common_grpc_fetch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchURL.ProtoReflect.Descriptor instead.
func (*FetchURL) Descriptor() ([]byte, []int) {
	return file_common_grpc_fetch_proto_rawDescGZIP(), []int{0}
}

func (x *FetchURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Header string `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body   string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_grpc_fetch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_common_grpc_fetch_proto_msgTypes[1]
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
	return file_common_grpc_fetch_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Response) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_grpc_fetch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_common_grpc_fetch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_common_grpc_fetch_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_common_grpc_fetch_proto protoreflect.FileDescriptor

var file_common_grpc_fetch_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x08, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x4a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x1b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x32, 0x2a, 0x0a, 0x07, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x12, 0x09, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x6c, 0x61, 0x6e,
	0x64, 0x2d, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_grpc_fetch_proto_rawDescOnce sync.Once
	file_common_grpc_fetch_proto_rawDescData = file_common_grpc_fetch_proto_rawDesc
)

func file_common_grpc_fetch_proto_rawDescGZIP() []byte {
	file_common_grpc_fetch_proto_rawDescOnce.Do(func() {
		file_common_grpc_fetch_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_grpc_fetch_proto_rawDescData)
	})
	return file_common_grpc_fetch_proto_rawDescData
}

var file_common_grpc_fetch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_grpc_fetch_proto_goTypes = []interface{}{
	(*FetchURL)(nil), // 0: FetchURL
	(*Response)(nil), // 1: Response
	(*Error)(nil),    // 2: Error
}
var file_common_grpc_fetch_proto_depIdxs = []int32{
	0, // 0: Fetcher.fetch:input_type -> FetchURL
	1, // 1: Fetcher.fetch:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_grpc_fetch_proto_init() }
func file_common_grpc_fetch_proto_init() {
	if File_common_grpc_fetch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_grpc_fetch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchURL); i {
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
		file_common_grpc_fetch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_grpc_fetch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_common_grpc_fetch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_grpc_fetch_proto_goTypes,
		DependencyIndexes: file_common_grpc_fetch_proto_depIdxs,
		MessageInfos:      file_common_grpc_fetch_proto_msgTypes,
	}.Build()
	File_common_grpc_fetch_proto = out.File
	file_common_grpc_fetch_proto_rawDesc = nil
	file_common_grpc_fetch_proto_goTypes = nil
	file_common_grpc_fetch_proto_depIdxs = nil
}
