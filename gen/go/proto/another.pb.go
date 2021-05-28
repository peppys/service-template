// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: proto/another.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAllAnotherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Anothers []*Another `protobuf:"bytes,1,rep,name=anothers,proto3" json:"anothers,omitempty"`
}

func (x *ListAllAnotherResponse) Reset() {
	*x = ListAllAnotherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_another_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllAnotherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllAnotherResponse) ProtoMessage() {}

func (x *ListAllAnotherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_another_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllAnotherResponse.ProtoReflect.Descriptor instead.
func (*ListAllAnotherResponse) Descriptor() ([]byte, []int) {
	return file_proto_another_proto_rawDescGZIP(), []int{0}
}

func (x *ListAllAnotherResponse) GetAnothers() []*Another {
	if x != nil {
		return x.Anothers
	}
	return nil
}

type CreateAnotherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text   string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateAnotherRequest) Reset() {
	*x = CreateAnotherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_another_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnotherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnotherRequest) ProtoMessage() {}

func (x *CreateAnotherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_another_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnotherRequest.ProtoReflect.Descriptor instead.
func (*CreateAnotherRequest) Descriptor() ([]byte, []int) {
	return file_proto_another_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnotherRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateAnotherRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type GetAnotherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAnotherRequest) Reset() {
	*x = GetAnotherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_another_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnotherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnotherRequest) ProtoMessage() {}

func (x *GetAnotherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_another_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnotherRequest.ProtoReflect.Descriptor instead.
func (*GetAnotherRequest) Descriptor() ([]byte, []int) {
	return file_proto_another_proto_rawDescGZIP(), []int{2}
}

func (x *GetAnotherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Another struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Timestamp string `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Another) Reset() {
	*x = Another{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_another_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Another) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Another) ProtoMessage() {}

func (x *Another) ProtoReflect() protoreflect.Message {
	mi := &file_proto_another_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Another.ProtoReflect.Descriptor instead.
func (*Another) Descriptor() ([]byte, []int) {
	return file_proto_another_proto_rawDescGZIP(), []int{3}
}

func (x *Another) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Another) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Another) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Another) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_proto_another_proto protoreflect.FileDescriptor

var file_proto_another_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x16, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x08, 0x61, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a,
	0x07, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x32, 0x8a, 0x02, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x22, 0x09, 0x2f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_another_proto_rawDescOnce sync.Once
	file_proto_another_proto_rawDescData = file_proto_another_proto_rawDesc
)

func file_proto_another_proto_rawDescGZIP() []byte {
	file_proto_another_proto_rawDescOnce.Do(func() {
		file_proto_another_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_another_proto_rawDescData)
	})
	return file_proto_another_proto_rawDescData
}

var file_proto_another_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_another_proto_goTypes = []interface{}{
	(*ListAllAnotherResponse)(nil), // 0: template.ListAllAnotherResponse
	(*CreateAnotherRequest)(nil),   // 1: template.CreateAnotherRequest
	(*GetAnotherRequest)(nil),      // 2: template.GetAnotherRequest
	(*Another)(nil),                // 3: template.Another
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_proto_another_proto_depIdxs = []int32{
	3, // 0: template.ListAllAnotherResponse.anothers:type_name -> template.Another
	4, // 1: template.AnotherService.ListAll:input_type -> google.protobuf.Empty
	1, // 2: template.AnotherService.Create:input_type -> template.CreateAnotherRequest
	2, // 3: template.AnotherService.Get:input_type -> template.GetAnotherRequest
	0, // 4: template.AnotherService.ListAll:output_type -> template.ListAllAnotherResponse
	3, // 5: template.AnotherService.Create:output_type -> template.Another
	3, // 6: template.AnotherService.Get:output_type -> template.Another
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_another_proto_init() }
func file_proto_another_proto_init() {
	if File_proto_another_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_another_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllAnotherResponse); i {
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
		file_proto_another_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnotherRequest); i {
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
		file_proto_another_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnotherRequest); i {
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
		file_proto_another_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Another); i {
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
			RawDescriptor: file_proto_another_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_another_proto_goTypes,
		DependencyIndexes: file_proto_another_proto_depIdxs,
		MessageInfos:      file_proto_another_proto_msgTypes,
	}.Build()
	File_proto_another_proto = out.File
	file_proto_another_proto_rawDesc = nil
	file_proto_another_proto_goTypes = nil
	file_proto_another_proto_depIdxs = nil
}
