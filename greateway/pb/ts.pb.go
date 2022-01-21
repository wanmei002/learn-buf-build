// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: ts.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Dial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Dial) Reset() {
	*x = Dial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dial) ProtoMessage() {}

func (x *Dial) ProtoReflect() protoreflect.Message {
	mi := &file_ts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dial.ProtoReflect.Descriptor instead.
func (*Dial) Descriptor() ([]byte, []int) {
	return file_ts_proto_rawDescGZIP(), []int{0}
}

func (x *Dial) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Dial) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Dial) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Ask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *AskContext `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Ask) Reset() {
	*x = Ask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ask) ProtoMessage() {}

func (x *Ask) ProtoReflect() protoreflect.Message {
	mi := &file_ts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ask.ProtoReflect.Descriptor instead.
func (*Ask) Descriptor() ([]byte, []int) {
	return file_ts_proto_rawDescGZIP(), []int{1}
}

func (x *Ask) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Ask) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Ask) GetData() *AskContext {
	if x != nil {
		return x.Data
	}
	return nil
}

type AskContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port     string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CanCall  bool   `protobuf:"varint,4,opt,name=can_call,json=canCall,proto3" json:"can_call,omitempty"`
}

func (x *AskContext) Reset() {
	*x = AskContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskContext) ProtoMessage() {}

func (x *AskContext) ProtoReflect() protoreflect.Message {
	mi := &file_ts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskContext.ProtoReflect.Descriptor instead.
func (*AskContext) Descriptor() ([]byte, []int) {
	return file_ts_proto_rawDescGZIP(), []int{2}
}

func (x *AskContext) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AskContext) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *AskContext) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AskContext) GetCanCall() bool {
	if x != nil {
		return x.CanCall
	}
	return false
}

var File_ts_proto protoreflect.FileDescriptor

var file_ts_proto_rawDesc = []byte{
	0x0a, 0x08, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x77, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4a, 0x0a, 0x04, 0x44, 0x69, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x03, 0x41,
	0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x6b, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x32, 0x53, 0x0a, 0x0c,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0b,
	0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x67, 0x77,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x1a, 0x0a, 0x2e, 0x67, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x73, 0x6b, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x67, 0x74, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ts_proto_rawDescOnce sync.Once
	file_ts_proto_rawDescData = file_ts_proto_rawDesc
)

func file_ts_proto_rawDescGZIP() []byte {
	file_ts_proto_rawDescOnce.Do(func() {
		file_ts_proto_rawDescData = protoimpl.X.CompressGZIP(file_ts_proto_rawDescData)
	})
	return file_ts_proto_rawDescData
}

var file_ts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ts_proto_goTypes = []interface{}{
	(*Dial)(nil),       // 0: gw.v1.Dial
	(*Ask)(nil),        // 1: gw.v1.Ask
	(*AskContext)(nil), // 2: gw.v1.AskContext
}
var file_ts_proto_depIdxs = []int32{
	2, // 0: gw.v1.Ask.data:type_name -> gw.v1.AskContext
	0, // 1: gw.v1.PhoneService.DialRequest:input_type -> gw.v1.Dial
	1, // 2: gw.v1.PhoneService.DialRequest:output_type -> gw.v1.Ask
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ts_proto_init() }
func file_ts_proto_init() {
	if File_ts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dial); i {
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
		file_ts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ask); i {
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
		file_ts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskContext); i {
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
			RawDescriptor: file_ts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ts_proto_goTypes,
		DependencyIndexes: file_ts_proto_depIdxs,
		MessageInfos:      file_ts_proto_msgTypes,
	}.Build()
	File_ts_proto = out.File
	file_ts_proto_rawDesc = nil
	file_ts_proto_goTypes = nil
	file_ts_proto_depIdxs = nil
}
