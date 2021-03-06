// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: relay.proto

package domain

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

type IncomingEvent_Type int32

const (
	IncomingEvent_NONE       IncomingEvent_Type = 0
	IncomingEvent_CONNECT    IncomingEvent_Type = 1
	IncomingEvent_DISCONNECT IncomingEvent_Type = 2
)

// Enum value maps for IncomingEvent_Type.
var (
	IncomingEvent_Type_name = map[int32]string{
		0: "NONE",
		1: "CONNECT",
		2: "DISCONNECT",
	}
	IncomingEvent_Type_value = map[string]int32{
		"NONE":       0,
		"CONNECT":    1,
		"DISCONNECT": 2,
	}
)

func (x IncomingEvent_Type) Enum() *IncomingEvent_Type {
	p := new(IncomingEvent_Type)
	*p = x
	return p
}

func (x IncomingEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IncomingEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_proto_enumTypes[0].Descriptor()
}

func (IncomingEvent_Type) Type() protoreflect.EnumType {
	return &file_relay_proto_enumTypes[0]
}

func (x IncomingEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IncomingEvent_Type.Descriptor instead.
func (IncomingEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{0, 0}
}

type IncomingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string             `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Data     []byte             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Type     IncomingEvent_Type `protobuf:"varint,3,opt,name=type,proto3,enum=domain.IncomingEvent_Type" json:"type,omitempty"`
}

func (x *IncomingEvent) Reset() {
	*x = IncomingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingEvent) ProtoMessage() {}

func (x *IncomingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingEvent.ProtoReflect.Descriptor instead.
func (*IncomingEvent) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{0}
}

func (x *IncomingEvent) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IncomingEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IncomingEvent) GetType() IncomingEvent_Type {
	if x != nil {
		return x.Type
	}
	return IncomingEvent_NONE
}

type OutgoingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientsList []string `protobuf:"bytes,1,rep,name=clients_list,json=clientsList,proto3" json:"clients_list,omitempty"`
	Data        []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OutgoingEvent) Reset() {
	*x = OutgoingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutgoingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingEvent) ProtoMessage() {}

func (x *OutgoingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingEvent.ProtoReflect.Descriptor instead.
func (*OutgoingEvent) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{1}
}

func (x *OutgoingEvent) GetClientsList() []string {
	if x != nil {
		return x.ClientsList
	}
	return nil
}

func (x *OutgoingEvent) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_relay_proto protoreflect.FileDescriptor

var file_relay_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x02, 0x22, 0x46, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x67, 0x6f,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x74, 0x61, 0x71, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_proto_rawDescOnce sync.Once
	file_relay_proto_rawDescData = file_relay_proto_rawDesc
)

func file_relay_proto_rawDescGZIP() []byte {
	file_relay_proto_rawDescOnce.Do(func() {
		file_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_proto_rawDescData)
	})
	return file_relay_proto_rawDescData
}

var file_relay_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relay_proto_goTypes = []interface{}{
	(IncomingEvent_Type)(0), // 0: domain.IncomingEvent.Type
	(*IncomingEvent)(nil),   // 1: domain.IncomingEvent
	(*OutgoingEvent)(nil),   // 2: domain.OutgoingEvent
}
var file_relay_proto_depIdxs = []int32{
	0, // 0: domain.IncomingEvent.type:type_name -> domain.IncomingEvent.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relay_proto_init() }
func file_relay_proto_init() {
	if File_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingEvent); i {
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
		file_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutgoingEvent); i {
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
			RawDescriptor: file_relay_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relay_proto_goTypes,
		DependencyIndexes: file_relay_proto_depIdxs,
		EnumInfos:         file_relay_proto_enumTypes,
		MessageInfos:      file_relay_proto_msgTypes,
	}.Build()
	File_relay_proto = out.File
	file_relay_proto_rawDesc = nil
	file_relay_proto_goTypes = nil
	file_relay_proto_depIdxs = nil
}
