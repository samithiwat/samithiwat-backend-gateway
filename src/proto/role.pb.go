// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: role.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       *Role    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RoleResponse) Reset() {
	*x = RoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResponse) ProtoMessage() {}

func (x *RoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResponse.ProtoReflect.Descriptor instead.
func (*RoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RoleResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RoleResponse) GetData() *Role {
	if x != nil {
		return x.Data
	}
	return nil
}

type RoleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       []*Role  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RoleListResponse) Reset() {
	*x = RoleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleListResponse) ProtoMessage() {}

func (x *RoleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleListResponse.ProtoReflect.Descriptor instead.
func (*RoleListResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RoleListResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RoleListResponse) GetData() []*Role {
	if x != nil {
		return x.Data
	}
	return nil
}

type RolePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Role             `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Meta  *PaginationMetadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *RolePagination) Reset() {
	*x = RolePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolePagination) ProtoMessage() {}

func (x *RolePagination) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolePagination.ProtoReflect.Descriptor instead.
func (*RolePagination) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *RolePagination) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RolePagination) GetMeta() *PaginationMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type RolePaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Errors     []string        `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	Data       *RolePagination `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RolePaginationResponse) Reset() {
	*x = RolePaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolePaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolePaginationResponse) ProtoMessage() {}

func (x *RolePaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolePaginationResponse.ProtoReflect.Descriptor instead.
func (*RolePaginationResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3}
}

func (x *RolePaginationResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RolePaginationResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *RolePaginationResponse) GetData() *RolePagination {
	if x != nil {
		return x.Data
	}
	return nil
}

type FindAllRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FindAllRoleRequest) Reset() {
	*x = FindAllRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRoleRequest) ProtoMessage() {}

func (x *FindAllRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRoleRequest.ProtoReflect.Descriptor instead.
func (*FindAllRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllRoleRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindAllRoleRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type FindOneRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRoleRequest) Reset() {
	*x = FindOneRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRoleRequest) ProtoMessage() {}

func (x *FindOneRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRoleRequest.ProtoReflect.Descriptor instead.
func (*FindOneRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5}
}

func (x *FindOneRoleRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindMultiRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *FindMultiRoleRequest) Reset() {
	*x = FindMultiRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMultiRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMultiRoleRequest) ProtoMessage() {}

func (x *FindMultiRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMultiRoleRequest.ProtoReflect.Descriptor instead.
func (*FindMultiRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{6}
}

func (x *FindMultiRoleRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRoleRequest) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRoleRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x1a, 0x09, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0c, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x69, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64,
	0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a,
	0x0e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x7a, 0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x12,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x32, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x22, 0x32, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xfb, 0x02, 0x0a, 0x0b, 0x52, 0x6f,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x1a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_role_proto_goTypes = []interface{}{
	(*RoleResponse)(nil),           // 0: role.RoleResponse
	(*RoleListResponse)(nil),       // 1: role.RoleListResponse
	(*RolePagination)(nil),         // 2: role.RolePagination
	(*RolePaginationResponse)(nil), // 3: role.RolePaginationResponse
	(*FindAllRoleRequest)(nil),     // 4: role.FindAllRoleRequest
	(*FindOneRoleRequest)(nil),     // 5: role.FindOneRoleRequest
	(*FindMultiRoleRequest)(nil),   // 6: role.FindMultiRoleRequest
	(*CreateRoleRequest)(nil),      // 7: role.CreateRoleRequest
	(*UpdateRoleRequest)(nil),      // 8: role.UpdateRoleRequest
	(*DeleteRoleRequest)(nil),      // 9: role.DeleteRoleRequest
	(*Role)(nil),                   // 10: dto.Role
	(*PaginationMetadata)(nil),     // 11: common.PaginationMetadata
}
var file_role_proto_depIdxs = []int32{
	10, // 0: role.RoleResponse.data:type_name -> dto.Role
	10, // 1: role.RoleListResponse.data:type_name -> dto.Role
	10, // 2: role.RolePagination.items:type_name -> dto.Role
	11, // 3: role.RolePagination.meta:type_name -> common.PaginationMetadata
	2,  // 4: role.RolePaginationResponse.data:type_name -> role.RolePagination
	10, // 5: role.CreateRoleRequest.role:type_name -> dto.Role
	10, // 6: role.UpdateRoleRequest.role:type_name -> dto.Role
	4,  // 7: role.RoleService.FindAll:input_type -> role.FindAllRoleRequest
	5,  // 8: role.RoleService.FindOne:input_type -> role.FindOneRoleRequest
	6,  // 9: role.RoleService.FindMulti:input_type -> role.FindMultiRoleRequest
	7,  // 10: role.RoleService.Create:input_type -> role.CreateRoleRequest
	8,  // 11: role.RoleService.Update:input_type -> role.UpdateRoleRequest
	9,  // 12: role.RoleService.Delete:input_type -> role.DeleteRoleRequest
	3,  // 13: role.RoleService.FindAll:output_type -> role.RolePaginationResponse
	0,  // 14: role.RoleService.FindOne:output_type -> role.RoleResponse
	1,  // 15: role.RoleService.FindMulti:output_type -> role.RoleListResponse
	0,  // 16: role.RoleService.Create:output_type -> role.RoleResponse
	0,  // 17: role.RoleService.Update:output_type -> role.RoleResponse
	0,  // 18: role.RoleService.Delete:output_type -> role.RoleResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	file_dto_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleResponse); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleListResponse); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolePagination); i {
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
		file_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolePaginationResponse); i {
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
		file_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRoleRequest); i {
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
		file_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRoleRequest); i {
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
		file_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMultiRoleRequest); i {
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
		file_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_role_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_role_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleServiceClient interface {
	FindAll(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*RolePaginationResponse, error)
	FindOne(ctx context.Context, in *FindOneRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	FindMulti(ctx context.Context, in *FindMultiRoleRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
	Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
	Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) FindAll(ctx context.Context, in *FindAllRoleRequest, opts ...grpc.CallOption) (*RolePaginationResponse, error) {
	out := new(RolePaginationResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindOne(ctx context.Context, in *FindOneRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindMulti(ctx context.Context, in *FindMultiRoleRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/FindMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Create(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Update(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) Delete(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*RoleResponse, error) {
	out := new(RoleResponse)
	err := c.cc.Invoke(ctx, "/role.RoleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
type RoleServiceServer interface {
	FindAll(context.Context, *FindAllRoleRequest) (*RolePaginationResponse, error)
	FindOne(context.Context, *FindOneRoleRequest) (*RoleResponse, error)
	FindMulti(context.Context, *FindMultiRoleRequest) (*RoleListResponse, error)
	Create(context.Context, *CreateRoleRequest) (*RoleResponse, error)
	Update(context.Context, *UpdateRoleRequest) (*RoleResponse, error)
	Delete(context.Context, *DeleteRoleRequest) (*RoleResponse, error)
}

// UnimplementedRoleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (*UnimplementedRoleServiceServer) FindAll(context.Context, *FindAllRoleRequest) (*RolePaginationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedRoleServiceServer) FindOne(context.Context, *FindOneRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedRoleServiceServer) FindMulti(context.Context, *FindMultiRoleRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMulti not implemented")
}
func (*UnimplementedRoleServiceServer) Create(context.Context, *CreateRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRoleServiceServer) Update(context.Context, *UpdateRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRoleServiceServer) Delete(context.Context, *DeleteRoleRequest) (*RoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRoleServiceServer(s *grpc.Server, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindAll(ctx, req.(*FindAllRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindOne(ctx, req.(*FindOneRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMultiRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/FindMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindMulti(ctx, req.(*FindMultiRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Create(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Update(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.RoleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).Delete(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "role.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _RoleService_FindAll_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _RoleService_FindOne_Handler,
		},
		{
			MethodName: "FindMulti",
			Handler:    _RoleService_FindMulti_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RoleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RoleService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
