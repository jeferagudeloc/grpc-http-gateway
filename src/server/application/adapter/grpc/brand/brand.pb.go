// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: brand.proto

package brand

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

type GetBrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []*Brand `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *GetBrandResponse) Reset() {
	*x = GetBrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandResponse) ProtoMessage() {}

func (x *GetBrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandResponse.ProtoReflect.Descriptor instead.
func (*GetBrandResponse) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{0}
}

func (x *GetBrandResponse) GetBrands() []*Brand {
	if x != nil {
		return x.Brands
	}
	return nil
}

type GetBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandName string `protobuf:"bytes,1,opt,name=brandName,proto3" json:"brandName,omitempty"`
}

func (x *GetBrandRequest) Reset() {
	*x = GetBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandRequest) ProtoMessage() {}

func (x *GetBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandRequest.ProtoReflect.Descriptor instead.
func (*GetBrandRequest) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{1}
}

func (x *GetBrandRequest) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

type Brand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status       string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CreationTime string `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
}

func (x *Brand) Reset() {
	*x = Brand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brand) ProtoMessage() {}

func (x *Brand) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brand.ProtoReflect.Descriptor instead.
func (*Brand) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{2}
}

func (x *Brand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Brand) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Brand) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

var File_brand_proto protoreflect.FileDescriptor

var file_brand_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x2f,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x57, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x4c, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_brand_proto_rawDescOnce sync.Once
	file_brand_proto_rawDescData = file_brand_proto_rawDesc
)

func file_brand_proto_rawDescGZIP() []byte {
	file_brand_proto_rawDescOnce.Do(func() {
		file_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_brand_proto_rawDescData)
	})
	return file_brand_proto_rawDescData
}

var file_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_brand_proto_goTypes = []interface{}{
	(*GetBrandResponse)(nil), // 0: brand.GetBrandResponse
	(*GetBrandRequest)(nil),  // 1: brand.GetBrandRequest
	(*Brand)(nil),            // 2: brand.Brand
}
var file_brand_proto_depIdxs = []int32{
	2, // 0: brand.GetBrandResponse.brands:type_name -> brand.Brand
	1, // 1: brand.BrandHandler.GetBrands:input_type -> brand.GetBrandRequest
	0, // 2: brand.BrandHandler.GetBrands:output_type -> brand.GetBrandResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_brand_proto_init() }
func file_brand_proto_init() {
	if File_brand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandResponse); i {
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
		file_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandRequest); i {
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
		file_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brand); i {
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
			RawDescriptor: file_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brand_proto_goTypes,
		DependencyIndexes: file_brand_proto_depIdxs,
		MessageInfos:      file_brand_proto_msgTypes,
	}.Build()
	File_brand_proto = out.File
	file_brand_proto_rawDesc = nil
	file_brand_proto_goTypes = nil
	file_brand_proto_depIdxs = nil
}
