// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/proto/experiment.proto

package pubsub

import (
	segmenters "github.com/gojek/xp/common/segmenters"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Experiment_Type int32

const (
	Experiment_A_B        Experiment_Type = 0
	Experiment_Switchback Experiment_Type = 1
)

// Enum value maps for Experiment_Type.
var (
	Experiment_Type_name = map[int32]string{
		0: "A_B",
		1: "Switchback",
	}
	Experiment_Type_value = map[string]int32{
		"A_B":        0,
		"Switchback": 1,
	}
)

func (x Experiment_Type) Enum() *Experiment_Type {
	p := new(Experiment_Type)
	*p = x
	return p
}

func (x Experiment_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Experiment_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_experiment_proto_enumTypes[0].Descriptor()
}

func (Experiment_Type) Type() protoreflect.EnumType {
	return &file_api_proto_experiment_proto_enumTypes[0]
}

func (x Experiment_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Experiment_Type.Descriptor instead.
func (Experiment_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{2, 0}
}

type Experiment_Status int32

const (
	Experiment_Active   Experiment_Status = 0
	Experiment_Inactive Experiment_Status = 1
)

// Enum value maps for Experiment_Status.
var (
	Experiment_Status_name = map[int32]string{
		0: "Active",
		1: "Inactive",
	}
	Experiment_Status_value = map[string]int32{
		"Active":   0,
		"Inactive": 1,
	}
)

func (x Experiment_Status) Enum() *Experiment_Status {
	p := new(Experiment_Status)
	*p = x
	return p
}

func (x Experiment_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Experiment_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_experiment_proto_enumTypes[1].Descriptor()
}

func (Experiment_Status) Type() protoreflect.EnumType {
	return &file_api_proto_experiment_proto_enumTypes[1]
}

func (x Experiment_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Experiment_Status.Descriptor instead.
func (Experiment_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{2, 1}
}

type Experiment_Tier int32

const (
	Experiment_Default  Experiment_Tier = 0
	Experiment_Override Experiment_Tier = 1
)

// Enum value maps for Experiment_Tier.
var (
	Experiment_Tier_name = map[int32]string{
		0: "Default",
		1: "Override",
	}
	Experiment_Tier_value = map[string]int32{
		"Default":  0,
		"Override": 1,
	}
)

func (x Experiment_Tier) Enum() *Experiment_Tier {
	p := new(Experiment_Tier)
	*p = x
	return p
}

func (x Experiment_Tier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Experiment_Tier) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_experiment_proto_enumTypes[2].Descriptor()
}

func (Experiment_Tier) Type() protoreflect.EnumType {
	return &file_api_proto_experiment_proto_enumTypes[2]
}

func (x Experiment_Tier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Experiment_Tier.Descriptor instead.
func (Experiment_Tier) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{2, 2}
}

type ExperimentCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *ExperimentCreated) Reset() {
	*x = ExperimentCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentCreated) ProtoMessage() {}

func (x *ExperimentCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentCreated.ProtoReflect.Descriptor instead.
func (*ExperimentCreated) Descriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *ExperimentCreated) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

type ExperimentUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Experiment *Experiment `protobuf:"bytes,1,opt,name=experiment,proto3" json:"experiment,omitempty"`
}

func (x *ExperimentUpdated) Reset() {
	*x = ExperimentUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_experiment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentUpdated) ProtoMessage() {}

func (x *ExperimentUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_experiment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentUpdated.ProtoReflect.Descriptor instead.
func (*ExperimentUpdated) Descriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{1}
}

func (x *ExperimentUpdated) GetExperiment() *Experiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

type Experiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId  int64                                     `protobuf:"varint,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Status     Experiment_Status                         `protobuf:"varint,3,opt,name=status,proto3,enum=pubsub.Experiment_Status" json:"status,omitempty"`
	Name       string                                    `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Segments   map[string]*segmenters.ListSegmenterValue `protobuf:"bytes,5,rep,name=segments,proto3" json:"segments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type       Experiment_Type                           `protobuf:"varint,6,opt,name=type,proto3,enum=pubsub.Experiment_Type" json:"type,omitempty"`
	Interval   int32                                     `protobuf:"varint,7,opt,name=interval,proto3" json:"interval,omitempty"`
	Tier       Experiment_Tier                           `protobuf:"varint,8,opt,name=tier,proto3,enum=pubsub.Experiment_Tier" json:"tier,omitempty"`
	StartTime  *timestamppb.Timestamp                    `protobuf:"bytes,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamppb.Timestamp                    `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Treatments []*ExperimentTreatment                    `protobuf:"bytes,11,rep,name=treatments,proto3" json:"treatments,omitempty"`
	UpdatedAt  *timestamppb.Timestamp                    `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Experiment) Reset() {
	*x = Experiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_experiment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experiment) ProtoMessage() {}

func (x *Experiment) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_experiment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experiment.ProtoReflect.Descriptor instead.
func (*Experiment) Descriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{2}
}

func (x *Experiment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Experiment) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *Experiment) GetStatus() Experiment_Status {
	if x != nil {
		return x.Status
	}
	return Experiment_Active
}

func (x *Experiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Experiment) GetSegments() map[string]*segmenters.ListSegmenterValue {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *Experiment) GetType() Experiment_Type {
	if x != nil {
		return x.Type
	}
	return Experiment_A_B
}

func (x *Experiment) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Experiment) GetTier() Experiment_Tier {
	if x != nil {
		return x.Tier
	}
	return Experiment_Default
}

func (x *Experiment) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Experiment) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Experiment) GetTreatments() []*ExperimentTreatment {
	if x != nil {
		return x.Treatments
	}
	return nil
}

func (x *Experiment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ExperimentTreatment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Traffic uint32           `protobuf:"varint,2,opt,name=traffic,proto3" json:"traffic,omitempty"`
	Config  *structpb.Struct `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ExperimentTreatment) Reset() {
	*x = ExperimentTreatment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_experiment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentTreatment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentTreatment) ProtoMessage() {}

func (x *ExperimentTreatment) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_experiment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentTreatment.ProtoReflect.Descriptor instead.
func (*ExperimentTreatment) Descriptor() ([]byte, []int) {
	return file_api_proto_experiment_proto_rawDescGZIP(), []int{3}
}

func (x *ExperimentTreatment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExperimentTreatment) GetTraffic() uint32 {
	if x != nil {
		return x.Traffic
	}
	return 0
}

func (x *ExperimentTreatment) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_api_proto_experiment_proto protoreflect.FileDescriptor

var file_api_proto_experiment_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0xe5, 0x05, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2b, 0x0a,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x74,
	0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x1a, 0x5b, 0x0a, 0x0d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x1f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x5f, 0x42,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x62, 0x61, 0x63, 0x6b,
	0x10, 0x01, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x22, 0x21, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f,
	0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x10, 0x01, 0x22, 0x74, 0x0a, 0x13, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x2f,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_experiment_proto_rawDescOnce sync.Once
	file_api_proto_experiment_proto_rawDescData = file_api_proto_experiment_proto_rawDesc
)

func file_api_proto_experiment_proto_rawDescGZIP() []byte {
	file_api_proto_experiment_proto_rawDescOnce.Do(func() {
		file_api_proto_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_experiment_proto_rawDescData)
	})
	return file_api_proto_experiment_proto_rawDescData
}

var file_api_proto_experiment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_proto_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_experiment_proto_goTypes = []interface{}{
	(Experiment_Type)(0),                  // 0: pubsub.Experiment.Type
	(Experiment_Status)(0),                // 1: pubsub.Experiment.Status
	(Experiment_Tier)(0),                  // 2: pubsub.Experiment.Tier
	(*ExperimentCreated)(nil),             // 3: pubsub.ExperimentCreated
	(*ExperimentUpdated)(nil),             // 4: pubsub.ExperimentUpdated
	(*Experiment)(nil),                    // 5: pubsub.Experiment
	(*ExperimentTreatment)(nil),           // 6: pubsub.ExperimentTreatment
	nil,                                   // 7: pubsub.Experiment.SegmentsEntry
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
	(*structpb.Struct)(nil),               // 9: google.protobuf.Struct
	(*segmenters.ListSegmenterValue)(nil), // 10: segmenters.ListSegmenterValue
}
var file_api_proto_experiment_proto_depIdxs = []int32{
	5,  // 0: pubsub.ExperimentCreated.experiment:type_name -> pubsub.Experiment
	5,  // 1: pubsub.ExperimentUpdated.experiment:type_name -> pubsub.Experiment
	1,  // 2: pubsub.Experiment.status:type_name -> pubsub.Experiment.Status
	7,  // 3: pubsub.Experiment.segments:type_name -> pubsub.Experiment.SegmentsEntry
	0,  // 4: pubsub.Experiment.type:type_name -> pubsub.Experiment.Type
	2,  // 5: pubsub.Experiment.tier:type_name -> pubsub.Experiment.Tier
	8,  // 6: pubsub.Experiment.start_time:type_name -> google.protobuf.Timestamp
	8,  // 7: pubsub.Experiment.end_time:type_name -> google.protobuf.Timestamp
	6,  // 8: pubsub.Experiment.treatments:type_name -> pubsub.ExperimentTreatment
	8,  // 9: pubsub.Experiment.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 10: pubsub.ExperimentTreatment.config:type_name -> google.protobuf.Struct
	10, // 11: pubsub.Experiment.SegmentsEntry.value:type_name -> segmenters.ListSegmenterValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_proto_experiment_proto_init() }
func file_api_proto_experiment_proto_init() {
	if File_api_proto_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentCreated); i {
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
		file_api_proto_experiment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentUpdated); i {
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
		file_api_proto_experiment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experiment); i {
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
		file_api_proto_experiment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentTreatment); i {
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
			RawDescriptor: file_api_proto_experiment_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_experiment_proto_goTypes,
		DependencyIndexes: file_api_proto_experiment_proto_depIdxs,
		EnumInfos:         file_api_proto_experiment_proto_enumTypes,
		MessageInfos:      file_api_proto_experiment_proto_msgTypes,
	}.Build()
	File_api_proto_experiment_proto = out.File
	file_api_proto_experiment_proto_rawDesc = nil
	file_api_proto_experiment_proto_goTypes = nil
	file_api_proto_experiment_proto_depIdxs = nil
}
