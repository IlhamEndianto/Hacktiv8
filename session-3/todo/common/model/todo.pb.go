// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: common/model/todo.proto

package model

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Todo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetData() []*Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Todo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllRequest) GetData() []*Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Todo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetByIDResponse) Reset() {
	*x = GetByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDResponse) ProtoMessage() {}

func (x *GetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDResponse.ProtoReflect.Descriptor instead.
func (*GetByIDResponse) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIDResponse) GetData() *Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

type MutationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MutationResponse) Reset() {
	*x = MutationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutationResponse) ProtoMessage() {}

func (x *MutationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutationResponse.ProtoReflect.Descriptor instead.
func (*MutationResponse) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{5}
}

func (x *MutationResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_model_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_model_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_common_model_todo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_common_model_todo_proto protoreflect.FileDescriptor

var file_common_model_todo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x10, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x85,
	0x04, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x23, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f,
	0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x69, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x23, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a, 0x22, 0x06, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x12,
	0x6b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x1a,
	0x0b, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x33, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6c, 0x68, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x69, 0x61, 0x6e,
	0x74, 0x6f, 0x2f, 0x48, 0x61, 0x63, 0x6b, 0x74, 0x69, 0x76, 0x38, 0x2f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2d, 0x33, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_model_todo_proto_rawDescOnce sync.Once
	file_common_model_todo_proto_rawDescData = file_common_model_todo_proto_rawDesc
)

func file_common_model_todo_proto_rawDescGZIP() []byte {
	file_common_model_todo_proto_rawDescOnce.Do(func() {
		file_common_model_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_model_todo_proto_rawDescData)
	})
	return file_common_model_todo_proto_rawDescData
}

var file_common_model_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_model_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),             // 0: session3.todo.model.Todo
	(*GetAllResponse)(nil),   // 1: session3.todo.model.GetAllResponse
	(*GetAllRequest)(nil),    // 2: session3.todo.model.GetAllRequest
	(*GetByIDRequest)(nil),   // 3: session3.todo.model.GetByIDRequest
	(*GetByIDResponse)(nil),  // 4: session3.todo.model.GetByIDResponse
	(*MutationResponse)(nil), // 5: session3.todo.model.MutationResponse
	(*UpdateRequest)(nil),    // 6: session3.todo.model.UpdateRequest
	(*DeleteRequest)(nil),    // 7: session3.todo.model.DeleteRequest
	(*emptypb.Empty)(nil),    // 8: google.protobuf.Empty
}
var file_common_model_todo_proto_depIdxs = []int32{
	0, // 0: session3.todo.model.GetAllResponse.data:type_name -> session3.todo.model.Todo
	0, // 1: session3.todo.model.GetAllRequest.data:type_name -> session3.todo.model.Todo
	0, // 2: session3.todo.model.GetByIDResponse.data:type_name -> session3.todo.model.Todo
	8, // 3: session3.todo.model.TodoService.GetAll:input_type -> google.protobuf.Empty
	3, // 4: session3.todo.model.TodoService.GetByID:input_type -> session3.todo.model.GetByIDRequest
	0, // 5: session3.todo.model.TodoService.Create:input_type -> session3.todo.model.Todo
	6, // 6: session3.todo.model.TodoService.Update:input_type -> session3.todo.model.UpdateRequest
	7, // 7: session3.todo.model.TodoService.Delete:input_type -> session3.todo.model.DeleteRequest
	1, // 8: session3.todo.model.TodoService.GetAll:output_type -> session3.todo.model.GetAllResponse
	4, // 9: session3.todo.model.TodoService.GetByID:output_type -> session3.todo.model.GetByIDResponse
	5, // 10: session3.todo.model.TodoService.Create:output_type -> session3.todo.model.MutationResponse
	5, // 11: session3.todo.model.TodoService.Update:output_type -> session3.todo.model.MutationResponse
	5, // 12: session3.todo.model.TodoService.Delete:output_type -> session3.todo.model.MutationResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_model_todo_proto_init() }
func file_common_model_todo_proto_init() {
	if File_common_model_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_model_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_common_model_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_common_model_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_common_model_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
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
		file_common_model_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDResponse); i {
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
		file_common_model_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutationResponse); i {
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
		file_common_model_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_common_model_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
			RawDescriptor: file_common_model_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_model_todo_proto_goTypes,
		DependencyIndexes: file_common_model_todo_proto_depIdxs,
		MessageInfos:      file_common_model_todo_proto_msgTypes,
	}.Build()
	File_common_model_todo_proto = out.File
	file_common_model_todo_proto_rawDesc = nil
	file_common_model_todo_proto_goTypes = nil
	file_common_model_todo_proto_depIdxs = nil
}
