// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: group.proto

package proto

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

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LeaderID string  `protobuf:"bytes,2,opt,name=leaderID,proto3" json:"leaderID,omitempty"`
	Token    string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Members  []*User `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetLeaderID() string {
	if x != nil {
		return x.LeaderID
	}
	return ""
}

func (x *Group) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Group) GetMembers() []*User {
	if x != nil {
		return x.Members
	}
	return nil
}

//Find One
type FindOneGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindOneGroupRequest) Reset() {
	*x = FindOneGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneGroupRequest) ProtoMessage() {}

func (x *FindOneGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneGroupRequest.ProtoReflect.Descriptor instead.
func (*FindOneGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *FindOneGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindOneGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *FindOneGroupResponse) Reset() {
	*x = FindOneGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneGroupResponse) ProtoMessage() {}

func (x *FindOneGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneGroupResponse.ProtoReflect.Descriptor instead.
func (*FindOneGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type FindByTokenGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *FindByTokenGroupRequest) Reset() {
	*x = FindByTokenGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByTokenGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByTokenGroupRequest) ProtoMessage() {}

func (x *FindByTokenGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByTokenGroupRequest.ProtoReflect.Descriptor instead.
func (*FindByTokenGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *FindByTokenGroupRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FindByTokenGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *FindByTokenGroupResponse) Reset() {
	*x = FindByTokenGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByTokenGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByTokenGroupResponse) ProtoMessage() {}

func (x *FindByTokenGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByTokenGroupResponse.ProtoReflect.Descriptor instead.
func (*FindByTokenGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *FindByTokenGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	LeaderId string `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGroupRequest) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *UpdateGroupRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

type UpdateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupResponse) Reset() {
	*x = UpdateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupResponse) ProtoMessage() {}

func (x *UpdateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type JoinGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *JoinGroupRequest) Reset() {
	*x = JoinGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupRequest) ProtoMessage() {}

func (x *JoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{9}
}

func (x *JoinGroupRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JoinGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type JoinGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *JoinGroupResponse) Reset() {
	*x = JoinGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupResponse) ProtoMessage() {}

func (x *JoinGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{10}
}

func (x *JoinGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type DeleteMemberGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LeaderId string `protobuf:"bytes,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
}

func (x *DeleteMemberGroupRequest) Reset() {
	*x = DeleteMemberGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemberGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberGroupRequest) ProtoMessage() {}

func (x *DeleteMemberGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteMemberGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteMemberGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteMemberGroupRequest) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

type DeleteMemberGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *DeleteMemberGroupResponse) Reset() {
	*x = DeleteMemberGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMemberGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMemberGroupResponse) ProtoMessage() {}

func (x *DeleteMemberGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMemberGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteMemberGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteMemberGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type LeaveGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LeaveGroupRequest) Reset() {
	*x = LeaveGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGroupRequest) ProtoMessage() {}

func (x *LeaveGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGroupRequest.ProtoReflect.Descriptor instead.
func (*LeaveGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{13}
}

func (x *LeaveGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LeaveGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *LeaveGroupResponse) Reset() {
	*x = LeaveGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveGroupResponse) ProtoMessage() {}

func (x *LeaveGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveGroupResponse.ProtoReflect.Descriptor instead.
func (*LeaveGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{14}
}

func (x *LeaveGroupResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6f, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x22, 0x2d, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3a, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2f, 0x0a, 0x17,
	0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a,
	0x18, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2c, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x54, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x40, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x11, 0x4a, 0x6f, 0x69,
	0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x4e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3f, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x22, 0x2b, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x38, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x32, 0xfe, 0x03, 0x0a, 0x0c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_group_proto_goTypes = []interface{}{
	(*Group)(nil),                     // 0: group.Group
	(*FindOneGroupRequest)(nil),       // 1: group.FindOneGroupRequest
	(*FindOneGroupResponse)(nil),      // 2: group.FindOneGroupResponse
	(*FindByTokenGroupRequest)(nil),   // 3: group.FindByTokenGroupRequest
	(*FindByTokenGroupResponse)(nil),  // 4: group.FindByTokenGroupResponse
	(*CreateGroupRequest)(nil),        // 5: group.CreateGroupRequest
	(*CreateGroupResponse)(nil),       // 6: group.CreateGroupResponse
	(*UpdateGroupRequest)(nil),        // 7: group.UpdateGroupRequest
	(*UpdateGroupResponse)(nil),       // 8: group.UpdateGroupResponse
	(*JoinGroupRequest)(nil),          // 9: group.JoinGroupRequest
	(*JoinGroupResponse)(nil),         // 10: group.JoinGroupResponse
	(*DeleteMemberGroupRequest)(nil),  // 11: group.DeleteMemberGroupRequest
	(*DeleteMemberGroupResponse)(nil), // 12: group.DeleteMemberGroupResponse
	(*LeaveGroupRequest)(nil),         // 13: group.LeaveGroupRequest
	(*LeaveGroupResponse)(nil),        // 14: group.LeaveGroupResponse
	(*User)(nil),                      // 15: user.User
}
var file_group_proto_depIdxs = []int32{
	15, // 0: group.Group.members:type_name -> user.User
	0,  // 1: group.FindOneGroupResponse.group:type_name -> group.Group
	0,  // 2: group.FindByTokenGroupResponse.group:type_name -> group.Group
	0,  // 3: group.CreateGroupResponse.group:type_name -> group.Group
	0,  // 4: group.UpdateGroupRequest.group:type_name -> group.Group
	0,  // 5: group.UpdateGroupResponse.group:type_name -> group.Group
	0,  // 6: group.JoinGroupResponse.group:type_name -> group.Group
	0,  // 7: group.DeleteMemberGroupResponse.group:type_name -> group.Group
	0,  // 8: group.LeaveGroupResponse.group:type_name -> group.Group
	1,  // 9: group.GroupService.FindOne:input_type -> group.FindOneGroupRequest
	3,  // 10: group.GroupService.FindByToken:input_type -> group.FindByTokenGroupRequest
	5,  // 11: group.GroupService.Create:input_type -> group.CreateGroupRequest
	7,  // 12: group.GroupService.Update:input_type -> group.UpdateGroupRequest
	9,  // 13: group.GroupService.Join:input_type -> group.JoinGroupRequest
	11, // 14: group.GroupService.DeleteMember:input_type -> group.DeleteMemberGroupRequest
	13, // 15: group.GroupService.Leave:input_type -> group.LeaveGroupRequest
	2,  // 16: group.GroupService.FindOne:output_type -> group.FindOneGroupResponse
	4,  // 17: group.GroupService.FindByToken:output_type -> group.FindByTokenGroupResponse
	6,  // 18: group.GroupService.Create:output_type -> group.CreateGroupResponse
	8,  // 19: group.GroupService.Update:output_type -> group.UpdateGroupResponse
	10, // 20: group.GroupService.Join:output_type -> group.JoinGroupResponse
	12, // 21: group.GroupService.DeleteMember:output_type -> group.DeleteMemberGroupResponse
	14, // 22: group.GroupService.Leave:output_type -> group.LeaveGroupResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneGroupRequest); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneGroupResponse); i {
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
		file_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByTokenGroupRequest); i {
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
		file_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByTokenGroupResponse); i {
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
		file_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupResponse); i {
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
		file_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupRequest); i {
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
		file_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupResponse); i {
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
		file_group_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemberGroupRequest); i {
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
		file_group_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMemberGroupResponse); i {
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
		file_group_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGroupRequest); i {
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
		file_group_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveGroupResponse); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
