//*********************************************************************************************************************
//  Description: interface between router agents and SDK service manager
//
//  Copyright (c) 2018 Nokia
//*********************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: mpls_service.proto

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

/// Represents MPLS operation.
type MplsRoutePb_Operation int32

const (
	MplsRoutePb_INVALID_OP MplsRoutePb_Operation = 0 // Invalid operation
	MplsRoutePb_POP        MplsRoutePb_Operation = 1 // Pop operation
	MplsRoutePb_SWAP       MplsRoutePb_Operation = 2 // Swap operation
)

// Enum value maps for MplsRoutePb_Operation.
var (
	MplsRoutePb_Operation_name = map[int32]string{
		0: "INVALID_OP",
		1: "POP",
		2: "SWAP",
	}
	MplsRoutePb_Operation_value = map[string]int32{
		"INVALID_OP": 0,
		"POP":        1,
		"SWAP":       2,
	}
)

func (x MplsRoutePb_Operation) Enum() *MplsRoutePb_Operation {
	p := new(MplsRoutePb_Operation)
	*p = x
	return p
}

func (x MplsRoutePb_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MplsRoutePb_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_mpls_service_proto_enumTypes[0].Descriptor()
}

func (MplsRoutePb_Operation) Type() protoreflect.EnumType {
	return &file_mpls_service_proto_enumTypes[0]
}

func (x MplsRoutePb_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MplsRoutePb_Operation.Descriptor instead.
func (MplsRoutePb_Operation) EnumDescriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{1, 0}
}

//*
// Represents MPLS route key.
type MplsRouteKeyPb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopLabel *MplsLabel `protobuf:"bytes,1,opt,name=top_label,json=topLabel,proto3" json:"top_label,omitempty"` // Top label
}

func (x *MplsRouteKeyPb) Reset() {
	*x = MplsRouteKeyPb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteKeyPb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteKeyPb) ProtoMessage() {}

func (x *MplsRouteKeyPb) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteKeyPb.ProtoReflect.Descriptor instead.
func (*MplsRouteKeyPb) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{0}
}

func (x *MplsRouteKeyPb) GetTopLabel() *MplsLabel {
	if x != nil {
		return x.TopLabel
	}
	return nil
}

//*
// Represents MPLS route data.
type MplsRoutePb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NexthopGroupName string                `protobuf:"bytes,1,opt,name=nexthop_group_name,json=nexthopGroupName,proto3" json:"nexthop_group_name,omitempty"` // Next hop group name
	Operation        MplsRoutePb_Operation `protobuf:"varint,2,opt,name=operation,proto3,enum=srlinux.sdk.MplsRoutePb_Operation" json:"operation,omitempty"` // Operation such as POP or SWAP
	Preference       uint32                `protobuf:"varint,3,opt,name=preference,proto3" json:"preference,omitempty"`                                      // Route preference
}

func (x *MplsRoutePb) Reset() {
	*x = MplsRoutePb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRoutePb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRoutePb) ProtoMessage() {}

func (x *MplsRoutePb) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRoutePb.ProtoReflect.Descriptor instead.
func (*MplsRoutePb) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{1}
}

func (x *MplsRoutePb) GetNexthopGroupName() string {
	if x != nil {
		return x.NexthopGroupName
	}
	return ""
}

func (x *MplsRoutePb) GetOperation() MplsRoutePb_Operation {
	if x != nil {
		return x.Operation
	}
	return MplsRoutePb_INVALID_OP
}

func (x *MplsRoutePb) GetPreference() uint32 {
	if x != nil {
		return x.Preference
	}
	return 0
}

//*
// Represents MPLS route information; contains key and data.
type MplsRouteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  *MplsRouteKeyPb `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`   // MPLS route key
	Data *MplsRoutePb    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` // MPLS route data
}

func (x *MplsRouteInfo) Reset() {
	*x = MplsRouteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteInfo) ProtoMessage() {}

func (x *MplsRouteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteInfo.ProtoReflect.Descriptor instead.
func (*MplsRouteInfo) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{2}
}

func (x *MplsRouteInfo) GetKey() *MplsRouteKeyPb {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *MplsRouteInfo) GetData() *MplsRoutePb {
	if x != nil {
		return x.Data
	}
	return nil
}

//*
// Represents MPLS route add request, which can include one or more MPLS routes.
type MplsRouteAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*MplsRouteInfo `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"` // MPLS routes
}

func (x *MplsRouteAddRequest) Reset() {
	*x = MplsRouteAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteAddRequest) ProtoMessage() {}

func (x *MplsRouteAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteAddRequest.ProtoReflect.Descriptor instead.
func (*MplsRouteAddRequest) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{3}
}

func (x *MplsRouteAddRequest) GetRoutes() []*MplsRouteInfo {
	if x != nil {
		return x.Routes
	}
	return nil
}

//*
// Represents MPLS route add response.
type MplsRouteAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of MPLS route add request
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error string
}

func (x *MplsRouteAddResponse) Reset() {
	*x = MplsRouteAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteAddResponse) ProtoMessage() {}

func (x *MplsRouteAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteAddResponse.ProtoReflect.Descriptor instead.
func (*MplsRouteAddResponse) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{4}
}

func (x *MplsRouteAddResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_kSdkMgrSuccess
}

func (x *MplsRouteAddResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

//*
// Represents MPLS route delete request, which can include one or more MPLS routes.
type MplsRouteDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routes []*MplsRouteKeyPb `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"` // MPLS routes
}

func (x *MplsRouteDeleteRequest) Reset() {
	*x = MplsRouteDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteDeleteRequest) ProtoMessage() {}

func (x *MplsRouteDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteDeleteRequest.ProtoReflect.Descriptor instead.
func (*MplsRouteDeleteRequest) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{5}
}

func (x *MplsRouteDeleteRequest) GetRoutes() []*MplsRouteKeyPb {
	if x != nil {
		return x.Routes
	}
	return nil
}

//*
// Represents MPLS route delete response.
type MplsRouteDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   SdkMgrStatus `protobuf:"varint,1,opt,name=status,proto3,enum=srlinux.sdk.SdkMgrStatus" json:"status,omitempty"` // Status of MPLS route delete request
	ErrorStr string       `protobuf:"bytes,2,opt,name=error_str,json=errorStr,proto3" json:"error_str,omitempty"`            // Detailed error string
}

func (x *MplsRouteDeleteResponse) Reset() {
	*x = MplsRouteDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpls_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MplsRouteDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MplsRouteDeleteResponse) ProtoMessage() {}

func (x *MplsRouteDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mpls_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MplsRouteDeleteResponse.ProtoReflect.Descriptor instead.
func (*MplsRouteDeleteResponse) Descriptor() ([]byte, []int) {
	return file_mpls_service_proto_rawDescGZIP(), []int{6}
}

func (x *MplsRouteDeleteResponse) GetStatus() SdkMgrStatus {
	if x != nil {
		return x.Status
	}
	return SdkMgrStatus_kSdkMgrSuccess
}

func (x *MplsRouteDeleteResponse) GetErrorStr() string {
	if x != nil {
		return x.ErrorStr
	}
	return ""
}

var File_mpls_service_proto protoreflect.FileDescriptor

var file_mpls_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64,
	0x6b, 0x1a, 0x10, 0x73, 0x64, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x50, 0x62, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x08, 0x74, 0x6f, 0x70, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x4d,
	0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x62, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65,
	0x78, 0x74, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x50, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x09, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x50, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x57, 0x41, 0x50, 0x10, 0x02, 0x22, 0x6c, 0x0a, 0x0d, 0x4d, 0x70,
	0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x50, 0x62, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x50, 0x62, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x13, 0x4d, 0x70, 0x6c, 0x73,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70,
	0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x22, 0x66, 0x0a, 0x14, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x22, 0x4d, 0x0a, 0x16, 0x4d,
	0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x50, 0x62, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x17, 0x4d, 0x70,
	0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x74, 0x72, 0x32, 0xdd, 0x02, 0x0a, 0x16, 0x53, 0x64, 0x6b, 0x4d, 0x67, 0x72,
	0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x14, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x72, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5e, 0x0a, 0x0f, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75,
	0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x4d, 0x70, 0x6c, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x73,
	0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x6e, 0x64, 0x12, 0x18,
	0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x72, 0x6c, 0x69, 0x6e,
	0x75, 0x78, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpls_service_proto_rawDescOnce sync.Once
	file_mpls_service_proto_rawDescData = file_mpls_service_proto_rawDesc
)

func file_mpls_service_proto_rawDescGZIP() []byte {
	file_mpls_service_proto_rawDescOnce.Do(func() {
		file_mpls_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpls_service_proto_rawDescData)
	})
	return file_mpls_service_proto_rawDescData
}

var file_mpls_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mpls_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mpls_service_proto_goTypes = []interface{}{
	(MplsRoutePb_Operation)(0),      // 0: srlinux.sdk.MplsRoutePb.Operation
	(*MplsRouteKeyPb)(nil),          // 1: srlinux.sdk.MplsRouteKeyPb
	(*MplsRoutePb)(nil),             // 2: srlinux.sdk.MplsRoutePb
	(*MplsRouteInfo)(nil),           // 3: srlinux.sdk.MplsRouteInfo
	(*MplsRouteAddRequest)(nil),     // 4: srlinux.sdk.MplsRouteAddRequest
	(*MplsRouteAddResponse)(nil),    // 5: srlinux.sdk.MplsRouteAddResponse
	(*MplsRouteDeleteRequest)(nil),  // 6: srlinux.sdk.MplsRouteDeleteRequest
	(*MplsRouteDeleteResponse)(nil), // 7: srlinux.sdk.MplsRouteDeleteResponse
	(*MplsLabel)(nil),               // 8: srlinux.sdk.MplsLabel
	(SdkMgrStatus)(0),               // 9: srlinux.sdk.SdkMgrStatus
	(*SyncRequest)(nil),             // 10: srlinux.sdk.SyncRequest
	(*SyncResponse)(nil),            // 11: srlinux.sdk.SyncResponse
}
var file_mpls_service_proto_depIdxs = []int32{
	8,  // 0: srlinux.sdk.MplsRouteKeyPb.top_label:type_name -> srlinux.sdk.MplsLabel
	0,  // 1: srlinux.sdk.MplsRoutePb.operation:type_name -> srlinux.sdk.MplsRoutePb.Operation
	1,  // 2: srlinux.sdk.MplsRouteInfo.key:type_name -> srlinux.sdk.MplsRouteKeyPb
	2,  // 3: srlinux.sdk.MplsRouteInfo.data:type_name -> srlinux.sdk.MplsRoutePb
	3,  // 4: srlinux.sdk.MplsRouteAddRequest.routes:type_name -> srlinux.sdk.MplsRouteInfo
	9,  // 5: srlinux.sdk.MplsRouteAddResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	1,  // 6: srlinux.sdk.MplsRouteDeleteRequest.routes:type_name -> srlinux.sdk.MplsRouteKeyPb
	9,  // 7: srlinux.sdk.MplsRouteDeleteResponse.status:type_name -> srlinux.sdk.SdkMgrStatus
	4,  // 8: srlinux.sdk.SdkMgrMplsRouteService.MplsRouteAddOrUpdate:input_type -> srlinux.sdk.MplsRouteAddRequest
	6,  // 9: srlinux.sdk.SdkMgrMplsRouteService.MplsRouteDelete:input_type -> srlinux.sdk.MplsRouteDeleteRequest
	10, // 10: srlinux.sdk.SdkMgrMplsRouteService.SyncStart:input_type -> srlinux.sdk.SyncRequest
	10, // 11: srlinux.sdk.SdkMgrMplsRouteService.SyncEnd:input_type -> srlinux.sdk.SyncRequest
	5,  // 12: srlinux.sdk.SdkMgrMplsRouteService.MplsRouteAddOrUpdate:output_type -> srlinux.sdk.MplsRouteAddResponse
	7,  // 13: srlinux.sdk.SdkMgrMplsRouteService.MplsRouteDelete:output_type -> srlinux.sdk.MplsRouteDeleteResponse
	11, // 14: srlinux.sdk.SdkMgrMplsRouteService.SyncStart:output_type -> srlinux.sdk.SyncResponse
	11, // 15: srlinux.sdk.SdkMgrMplsRouteService.SyncEnd:output_type -> srlinux.sdk.SyncResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_mpls_service_proto_init() }
func file_mpls_service_proto_init() {
	if File_mpls_service_proto != nil {
		return
	}
	file_sdk_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpls_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteKeyPb); i {
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
		file_mpls_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRoutePb); i {
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
		file_mpls_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteInfo); i {
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
		file_mpls_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteAddRequest); i {
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
		file_mpls_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteAddResponse); i {
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
		file_mpls_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteDeleteRequest); i {
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
		file_mpls_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MplsRouteDeleteResponse); i {
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
			RawDescriptor: file_mpls_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpls_service_proto_goTypes,
		DependencyIndexes: file_mpls_service_proto_depIdxs,
		EnumInfos:         file_mpls_service_proto_enumTypes,
		MessageInfos:      file_mpls_service_proto_msgTypes,
	}.Build()
	File_mpls_service_proto = out.File
	file_mpls_service_proto_rawDesc = nil
	file_mpls_service_proto_goTypes = nil
	file_mpls_service_proto_depIdxs = nil
}
