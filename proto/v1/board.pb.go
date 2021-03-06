// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: v1/board.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId  string                       `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	OwnerId  string                       `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name     string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Metadata string                       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Members  []*CreateBoardRequest_Member `protobuf:"bytes,5,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *CreateBoardRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBoardRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *CreateBoardRequest) GetMembers() []*CreateBoardRequest_Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type UpdateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId  string                       `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name     string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Metadata string                       `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Members  []*UpdateBoardRequest_Member `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *UpdateBoardRequest) Reset() {
	*x = UpdateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardRequest) ProtoMessage() {}

func (x *UpdateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoardRequest) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *UpdateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBoardRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *UpdateBoardRequest) GetMembers() []*UpdateBoardRequest_Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type DeleteBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *DeleteBoardRequest) Reset() {
	*x = DeleteBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBoardRequest) ProtoMessage() {}

func (x *DeleteBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBoardRequest.ProtoReflect.Descriptor instead.
func (*DeleteBoardRequest) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

type BoardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex uint64   `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageSize  uint32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	BoardIds  []string `protobuf:"bytes,3,rep,name=board_ids,json=boardIds,proto3" json:"board_ids,omitempty"`
	OwnerIds  []string `protobuf:"bytes,4,rep,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	MemberIds []string `protobuf:"bytes,5,rep,name=member_ids,json=memberIds,proto3" json:"member_ids,omitempty"`
}

func (x *BoardsRequest) Reset() {
	*x = BoardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardsRequest) ProtoMessage() {}

func (x *BoardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardsRequest.ProtoReflect.Descriptor instead.
func (*BoardsRequest) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{3}
}

func (x *BoardsRequest) GetPageIndex() uint64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *BoardsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *BoardsRequest) GetBoardIds() []string {
	if x != nil {
		return x.BoardIds
	}
	return nil
}

func (x *BoardsRequest) GetOwnerIds() []string {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *BoardsRequest) GetMemberIds() []string {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

type BoardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  uint64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Boards []*BoardsResponse_Board `protobuf:"bytes,2,rep,name=boards,proto3" json:"boards,omitempty"`
}

func (x *BoardsResponse) Reset() {
	*x = BoardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardsResponse) ProtoMessage() {}

func (x *BoardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardsResponse.ProtoReflect.Descriptor instead.
func (*BoardsResponse) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{4}
}

func (x *BoardsResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BoardsResponse) GetBoards() []*BoardsResponse_Board {
	if x != nil {
		return x.Boards
	}
	return nil
}

type CreateBoardRequest_Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string   `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Roles    []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *CreateBoardRequest_Member) Reset() {
	*x = CreateBoardRequest_Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest_Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest_Member) ProtoMessage() {}

func (x *CreateBoardRequest_Member) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest_Member.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest_Member) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateBoardRequest_Member) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *CreateBoardRequest_Member) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UpdateBoardRequest_Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string   `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Roles    []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Delete   bool     `protobuf:"varint,3,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (x *UpdateBoardRequest_Member) Reset() {
	*x = UpdateBoardRequest_Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardRequest_Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardRequest_Member) ProtoMessage() {}

func (x *UpdateBoardRequest_Member) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardRequest_Member.ProtoReflect.Descriptor instead.
func (*UpdateBoardRequest_Member) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UpdateBoardRequest_Member) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *UpdateBoardRequest_Member) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UpdateBoardRequest_Member) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

type BoardsResponse_Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId   string                         `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	OwnerId   string                         `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Name      string                         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Metadata  string                         `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt *timestamppb.Timestamp         `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Members   []*BoardsResponse_Board_Member `protobuf:"bytes,6,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *BoardsResponse_Board) Reset() {
	*x = BoardsResponse_Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardsResponse_Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardsResponse_Board) ProtoMessage() {}

func (x *BoardsResponse_Board) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardsResponse_Board.ProtoReflect.Descriptor instead.
func (*BoardsResponse_Board) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{4, 0}
}

func (x *BoardsResponse_Board) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *BoardsResponse_Board) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *BoardsResponse_Board) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoardsResponse_Board) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *BoardsResponse_Board) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BoardsResponse_Board) GetMembers() []*BoardsResponse_Board_Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type BoardsResponse_Board_Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId string   `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Roles    []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *BoardsResponse_Board_Member) Reset() {
	*x = BoardsResponse_Board_Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_board_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardsResponse_Board_Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardsResponse_Board_Member) ProtoMessage() {}

func (x *BoardsResponse_Board_Member) ProtoReflect() protoreflect.Message {
	mi := &file_v1_board_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardsResponse_Board_Member.ProtoReflect.Descriptor instead.
func (*BoardsResponse_Board_Member) Descriptor() ([]byte, []int) {
	return file_v1_board_proto_rawDescGZIP(), []int{4, 0, 0}
}

func (x *BoardsResponse_Board_Member) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *BoardsResponse_Board_Member) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_v1_board_proto protoreflect.FileDescriptor

var file_v1_board_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x1a, 0x3b, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x22, 0xf3, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x1a, 0x53, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x0d, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22,
	0x87, 0x03, 0x0a, 0x0e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x1a, 0xa6, 0x02, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x3b, 0x0a, 0x06,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0x96, 0x02, 0x0a, 0x05, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x48, 0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x46, 0x75, 0x6e, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x4f, 0x72, 0x67, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_board_proto_rawDescOnce sync.Once
	file_v1_board_proto_rawDescData = file_v1_board_proto_rawDesc
)

func file_v1_board_proto_rawDescGZIP() []byte {
	file_v1_board_proto_rawDescOnce.Do(func() {
		file_v1_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_board_proto_rawDescData)
	})
	return file_v1_board_proto_rawDescData
}

var file_v1_board_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_board_proto_goTypes = []interface{}{
	(*CreateBoardRequest)(nil),          // 0: proto.v1.CreateBoardRequest
	(*UpdateBoardRequest)(nil),          // 1: proto.v1.UpdateBoardRequest
	(*DeleteBoardRequest)(nil),          // 2: proto.v1.DeleteBoardRequest
	(*BoardsRequest)(nil),               // 3: proto.v1.BoardsRequest
	(*BoardsResponse)(nil),              // 4: proto.v1.BoardsResponse
	(*CreateBoardRequest_Member)(nil),   // 5: proto.v1.CreateBoardRequest.Member
	(*UpdateBoardRequest_Member)(nil),   // 6: proto.v1.UpdateBoardRequest.Member
	(*BoardsResponse_Board)(nil),        // 7: proto.v1.BoardsResponse.Board
	(*BoardsResponse_Board_Member)(nil), // 8: proto.v1.BoardsResponse.Board.Member
	(*timestamppb.Timestamp)(nil),       // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 10: google.protobuf.Empty
}
var file_v1_board_proto_depIdxs = []int32{
	5,  // 0: proto.v1.CreateBoardRequest.members:type_name -> proto.v1.CreateBoardRequest.Member
	6,  // 1: proto.v1.UpdateBoardRequest.members:type_name -> proto.v1.UpdateBoardRequest.Member
	7,  // 2: proto.v1.BoardsResponse.boards:type_name -> proto.v1.BoardsResponse.Board
	9,  // 3: proto.v1.BoardsResponse.Board.created_at:type_name -> google.protobuf.Timestamp
	8,  // 4: proto.v1.BoardsResponse.Board.members:type_name -> proto.v1.BoardsResponse.Board.Member
	0,  // 5: proto.v1.Board.CreateBoard:input_type -> proto.v1.CreateBoardRequest
	1,  // 6: proto.v1.Board.UpdateBoard:input_type -> proto.v1.UpdateBoardRequest
	2,  // 7: proto.v1.Board.DeleteBoard:input_type -> proto.v1.DeleteBoardRequest
	3,  // 8: proto.v1.Board.GetBoards:input_type -> proto.v1.BoardsRequest
	10, // 9: proto.v1.Board.CreateBoard:output_type -> google.protobuf.Empty
	10, // 10: proto.v1.Board.UpdateBoard:output_type -> google.protobuf.Empty
	10, // 11: proto.v1.Board.DeleteBoard:output_type -> google.protobuf.Empty
	4,  // 12: proto.v1.Board.GetBoards:output_type -> proto.v1.BoardsResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_board_proto_init() }
func file_v1_board_proto_init() {
	if File_v1_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest); i {
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
		file_v1_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardRequest); i {
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
		file_v1_board_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBoardRequest); i {
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
		file_v1_board_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardsRequest); i {
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
		file_v1_board_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardsResponse); i {
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
		file_v1_board_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest_Member); i {
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
		file_v1_board_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardRequest_Member); i {
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
		file_v1_board_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardsResponse_Board); i {
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
		file_v1_board_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardsResponse_Board_Member); i {
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
			RawDescriptor: file_v1_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_board_proto_goTypes,
		DependencyIndexes: file_v1_board_proto_depIdxs,
		MessageInfos:      file_v1_board_proto_msgTypes,
	}.Build()
	File_v1_board_proto = out.File
	file_v1_board_proto_rawDesc = nil
	file_v1_board_proto_goTypes = nil
	file_v1_board_proto_depIdxs = nil
}
