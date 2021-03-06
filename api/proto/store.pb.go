// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: api/proto/store.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetLastEventIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrailingLag int64   `protobuf:"varint,1,opt,name=trailing_lag,json=trailingLag,proto3" json:"trailing_lag,omitempty"`
	Filter      *Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetLastEventIDRequest) Reset() {
	*x = GetLastEventIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastEventIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastEventIDRequest) ProtoMessage() {}

func (x *GetLastEventIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastEventIDRequest.ProtoReflect.Descriptor instead.
func (*GetLastEventIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{0}
}

func (x *GetLastEventIDRequest) GetTrailingLag() int64 {
	if x != nil {
		return x.TrailingLag
	}
	return 0
}

func (x *GetLastEventIDRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetLastEventIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetLastEventIDReply) Reset() {
	*x = GetLastEventIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastEventIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastEventIDReply) ProtoMessage() {}

func (x *GetLastEventIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastEventIDReply.ProtoReflect.Descriptor instead.
func (*GetLastEventIDReply) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{1}
}

func (x *GetLastEventIDReply) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AfterEventId string  `protobuf:"bytes,1,opt,name=after_event_id,json=afterEventId,proto3" json:"after_event_id,omitempty"`
	Limit        int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	TrailingLag  int64   `protobuf:"varint,3,opt,name=trailing_lag,json=trailingLag,proto3" json:"trailing_lag,omitempty"`
	Filter       *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{2}
}

func (x *GetEventsRequest) GetAfterEventId() string {
	if x != nil {
		return x.AfterEventId
	}
	return ""
}

func (x *GetEventsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetEventsRequest) GetTrailingLag() int64 {
	if x != nil {
		return x.TrailingLag
	}
	return 0
}

func (x *GetEventsRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateTypes []string `protobuf:"bytes,1,rep,name=aggregate_types,json=aggregateTypes,proto3" json:"aggregate_types,omitempty"`
	Labels         []*Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Partitions     uint32   `protobuf:"varint,3,opt,name=partitions,proto3" json:"partitions,omitempty"`
	PartitionLow   uint32   `protobuf:"varint,4,opt,name=partitionLow,proto3" json:"partitionLow,omitempty"`
	PartitionHi    uint32   `protobuf:"varint,5,opt,name=partitionHi,proto3" json:"partitionHi,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetAggregateTypes() []string {
	if x != nil {
		return x.AggregateTypes
	}
	return nil
}

func (x *Filter) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Filter) GetPartitions() uint32 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

func (x *Filter) GetPartitionLow() uint32 {
	if x != nil {
		return x.PartitionLow
	}
	return 0
}

func (x *Filter) GetPartitionHi() uint32 {
	if x != nil {
		return x.PartitionHi
	}
	return 0
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{4}
}

func (x *Label) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetEventsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsReply) Reset() {
	*x = GetEventsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsReply) ProtoMessage() {}

func (x *GetEventsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsReply.ProtoReflect.Descriptor instead.
func (*GetEventsReply) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{5}
}

func (x *GetEventsReply) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AggregateId      string               `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateVersion uint32               `protobuf:"varint,3,opt,name=aggregate_version,json=aggregateVersion,proto3" json:"aggregate_version,omitempty"`
	AggregateIdHash  uint32               `protobuf:"varint,4,opt,name=aggregate_id_hash,json=aggregateIdHash,proto3" json:"aggregate_id_hash,omitempty"`
	AggregateType    string               `protobuf:"bytes,5,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	Kind             string               `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
	Body             []byte               `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	IdempotencyKey   string               `protobuf:"bytes,8,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	Labels           string               `protobuf:"bytes,9,opt,name=labels,proto3" json:"labels,omitempty"`
	CreatedAt        *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_proto_store_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *Event) GetAggregateVersion() uint32 {
	if x != nil {
		return x.AggregateVersion
	}
	return 0
}

func (x *Event) GetAggregateIdHash() uint32 {
	if x != nil {
		return x.AggregateIdHash
	}
	return 0
}

func (x *Event) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Event) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (x *Event) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

func (x *Event) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_proto_store_proto protoreflect.FileDescriptor

var file_api_proto_store_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x61, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x66, 0x74, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0xbd, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x77, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x22,
	0x2f, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xde, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x94, 0x01, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_store_proto_rawDescOnce sync.Once
	file_api_proto_store_proto_rawDescData = file_api_proto_store_proto_rawDesc
)

func file_api_proto_store_proto_rawDescGZIP() []byte {
	file_api_proto_store_proto_rawDescOnce.Do(func() {
		file_api_proto_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_store_proto_rawDescData)
	})
	return file_api_proto_store_proto_rawDescData
}

var file_api_proto_store_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_store_proto_goTypes = []interface{}{
	(*GetLastEventIDRequest)(nil), // 0: proto.GetLastEventIDRequest
	(*GetLastEventIDReply)(nil),   // 1: proto.GetLastEventIDReply
	(*GetEventsRequest)(nil),      // 2: proto.GetEventsRequest
	(*Filter)(nil),                // 3: proto.Filter
	(*Label)(nil),                 // 4: proto.Label
	(*GetEventsReply)(nil),        // 5: proto.GetEventsReply
	(*Event)(nil),                 // 6: proto.Event
	(*timestamp.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_api_proto_store_proto_depIdxs = []int32{
	3, // 0: proto.GetLastEventIDRequest.filter:type_name -> proto.Filter
	3, // 1: proto.GetEventsRequest.filter:type_name -> proto.Filter
	4, // 2: proto.Filter.labels:type_name -> proto.Label
	6, // 3: proto.GetEventsReply.events:type_name -> proto.Event
	7, // 4: proto.Event.created_at:type_name -> google.protobuf.Timestamp
	0, // 5: proto.Store.GetLastEventID:input_type -> proto.GetLastEventIDRequest
	2, // 6: proto.Store.GetEvents:input_type -> proto.GetEventsRequest
	1, // 7: proto.Store.GetLastEventID:output_type -> proto.GetLastEventIDReply
	5, // 8: proto.Store.GetEvents:output_type -> proto.GetEventsReply
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_store_proto_init() }
func file_api_proto_store_proto_init() {
	if File_api_proto_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastEventIDRequest); i {
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
		file_api_proto_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastEventIDReply); i {
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
		file_api_proto_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_api_proto_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_api_proto_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_api_proto_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsReply); i {
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
		file_api_proto_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_api_proto_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_store_proto_goTypes,
		DependencyIndexes: file_api_proto_store_proto_depIdxs,
		MessageInfos:      file_api_proto_store_proto_msgTypes,
	}.Build()
	File_api_proto_store_proto = out.File
	file_api_proto_store_proto_rawDesc = nil
	file_api_proto_store_proto_goTypes = nil
	file_api_proto_store_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	GetLastEventID(ctx context.Context, in *GetLastEventIDRequest, opts ...grpc.CallOption) (*GetLastEventIDReply, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsReply, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) GetLastEventID(ctx context.Context, in *GetLastEventIDRequest, opts ...grpc.CallOption) (*GetLastEventIDReply, error) {
	out := new(GetLastEventIDReply)
	err := c.cc.Invoke(ctx, "/proto.Store/GetLastEventID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsReply, error) {
	out := new(GetEventsReply)
	err := c.cc.Invoke(ctx, "/proto.Store/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	GetLastEventID(context.Context, *GetLastEventIDRequest) (*GetLastEventIDReply, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsReply, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) GetLastEventID(context.Context, *GetLastEventIDRequest) (*GetLastEventIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastEventID not implemented")
}
func (*UnimplementedStoreServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_GetLastEventID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastEventIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetLastEventID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/GetLastEventID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetLastEventID(ctx, req.(*GetLastEventIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastEventID",
			Handler:    _Store_GetLastEventID_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _Store_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/store.proto",
}
