// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: protos/stock.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StockUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceCode string  `protobuf:"bytes,1,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	Sku        string  `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Qty        float32 `protobuf:"fixed32,3,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *StockUpdateRequest) Reset() {
	*x = StockUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockUpdateRequest) ProtoMessage() {}

func (x *StockUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockUpdateRequest.ProtoReflect.Descriptor instead.
func (*StockUpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_stock_proto_rawDescGZIP(), []int{0}
}

func (x *StockUpdateRequest) GetSourceCode() string {
	if x != nil {
		return x.SourceCode
	}
	return ""
}

func (x *StockUpdateRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *StockUpdateRequest) GetQty() float32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type StockUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack bool `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *StockUpdateResponse) Reset() {
	*x = StockUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockUpdateResponse) ProtoMessage() {}

func (x *StockUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockUpdateResponse.ProtoReflect.Descriptor instead.
func (*StockUpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_stock_proto_rawDescGZIP(), []int{1}
}

func (x *StockUpdateResponse) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

var File_protos_stock_proto protoreflect.FileDescriptor

var file_protos_stock_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x59, 0x0a, 0x12, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x6b, 0x75, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32,
	0x4a, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6d, 0x75, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0xca, 0x02, 0x12, 0x4d, 0x75, 0x65, 0x6e, 0x63, 0x68, 0x44, 0x65, 0x76, 0x5c, 0x47,
	0x72, 0x70, 0x63, 0x44, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_stock_proto_rawDescOnce sync.Once
	file_protos_stock_proto_rawDescData = file_protos_stock_proto_rawDesc
)

func file_protos_stock_proto_rawDescGZIP() []byte {
	file_protos_stock_proto_rawDescOnce.Do(func() {
		file_protos_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_stock_proto_rawDescData)
	})
	return file_protos_stock_proto_rawDescData
}

var file_protos_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_stock_proto_goTypes = []interface{}{
	(*StockUpdateRequest)(nil),  // 0: stock.StockUpdateRequest
	(*StockUpdateResponse)(nil), // 1: stock.StockUpdateResponse
}
var file_protos_stock_proto_depIdxs = []int32{
	0, // 0: stock.Stock.Update:input_type -> stock.StockUpdateRequest
	1, // 1: stock.Stock.Update:output_type -> stock.StockUpdateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_stock_proto_init() }
func file_protos_stock_proto_init() {
	if File_protos_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockUpdateRequest); i {
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
		file_protos_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockUpdateResponse); i {
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
			RawDescriptor: file_protos_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_stock_proto_goTypes,
		DependencyIndexes: file_protos_stock_proto_depIdxs,
		MessageInfos:      file_protos_stock_proto_msgTypes,
	}.Build()
	File_protos_stock_proto = out.File
	file_protos_stock_proto_rawDesc = nil
	file_protos_stock_proto_goTypes = nil
	file_protos_stock_proto_depIdxs = nil
}
