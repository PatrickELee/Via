// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: api/proto/via.proto

package viapb

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

type Incident_DAY_OF_THE_WEEK int32

const (
	Incident_SUNDAY    Incident_DAY_OF_THE_WEEK = 0
	Incident_MONDAY    Incident_DAY_OF_THE_WEEK = 1
	Incident_TUESDAY   Incident_DAY_OF_THE_WEEK = 2
	Incident_WEDNESDAY Incident_DAY_OF_THE_WEEK = 3
	Incident_THURSDAY  Incident_DAY_OF_THE_WEEK = 4
	Incident_FRIDAY    Incident_DAY_OF_THE_WEEK = 5
	Incident_SATURDAY  Incident_DAY_OF_THE_WEEK = 6
)

// Enum value maps for Incident_DAY_OF_THE_WEEK.
var (
	Incident_DAY_OF_THE_WEEK_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	Incident_DAY_OF_THE_WEEK_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x Incident_DAY_OF_THE_WEEK) Enum() *Incident_DAY_OF_THE_WEEK {
	p := new(Incident_DAY_OF_THE_WEEK)
	*p = x
	return p
}

func (x Incident_DAY_OF_THE_WEEK) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Incident_DAY_OF_THE_WEEK) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_via_proto_enumTypes[0].Descriptor()
}

func (Incident_DAY_OF_THE_WEEK) Type() protoreflect.EnumType {
	return &file_api_proto_via_proto_enumTypes[0]
}

func (x Incident_DAY_OF_THE_WEEK) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Incident_DAY_OF_THE_WEEK.Descriptor instead.
func (Incident_DAY_OF_THE_WEEK) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{0, 0}
}

type Incident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lat  string                   `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long string                   `protobuf:"bytes,3,opt,name=long,proto3" json:"long,omitempty"`
	Time string                   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Day  Incident_DAY_OF_THE_WEEK `protobuf:"varint,5,opt,name=day,proto3,enum=viapb.Incident_DAY_OF_THE_WEEK" json:"day,omitempty"`
}

func (x *Incident) Reset() {
	*x = Incident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_via_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident) ProtoMessage() {}

func (x *Incident) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_via_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident.ProtoReflect.Descriptor instead.
func (*Incident) Descriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{0}
}

func (x *Incident) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Incident) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Incident) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

func (x *Incident) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Incident) GetDay() Incident_DAY_OF_THE_WEEK {
	if x != nil {
		return x.Day
	}
	return Incident_SUNDAY
}

type GetDangerProbabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Time     string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetDangerProbabilityRequest) Reset() {
	*x = GetDangerProbabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_via_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDangerProbabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDangerProbabilityRequest) ProtoMessage() {}

func (x *GetDangerProbabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_via_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDangerProbabilityRequest.ProtoReflect.Descriptor instead.
func (*GetDangerProbabilityRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{1}
}

func (x *GetDangerProbabilityRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetDangerProbabilityRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type GetDangerProbabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Danger string `protobuf:"bytes,1,opt,name=danger,proto3" json:"danger,omitempty"`
}

func (x *GetDangerProbabilityResponse) Reset() {
	*x = GetDangerProbabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_via_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDangerProbabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDangerProbabilityResponse) ProtoMessage() {}

func (x *GetDangerProbabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_via_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDangerProbabilityResponse.ProtoReflect.Descriptor instead.
func (*GetDangerProbabilityResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{2}
}

func (x *GetDangerProbabilityResponse) GetDanger() string {
	if x != nil {
		return x.Danger
	}
	return ""
}

type GetIncidentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetIncidentRequest) Reset() {
	*x = GetIncidentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_via_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIncidentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIncidentRequest) ProtoMessage() {}

func (x *GetIncidentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_via_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIncidentRequest.ProtoReflect.Descriptor instead.
func (*GetIncidentRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{3}
}

type GetIncidentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incident *Incident `protobuf:"bytes,1,opt,name=incident,proto3" json:"incident,omitempty"`
}

func (x *GetIncidentResponse) Reset() {
	*x = GetIncidentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_via_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIncidentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIncidentResponse) ProtoMessage() {}

func (x *GetIncidentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_via_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIncidentResponse.ProtoReflect.Descriptor instead.
func (*GetIncidentResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_via_proto_rawDescGZIP(), []int{4}
}

func (x *GetIncidentResponse) GetIncident() *Incident {
	if x != nil {
		return x.Incident
	}
	return nil
}

var File_api_proto_via_proto protoreflect.FileDescriptor

var file_api_proto_via_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x61, 0x70, 0x62, 0x22, 0xf6, 0x01, 0x0a,
	0x08, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1f, 0x2e, 0x76, 0x69, 0x61, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46, 0x5f, 0x54, 0x48, 0x45, 0x5f, 0x57, 0x45, 0x45,
	0x4b, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x6d, 0x0a, 0x0f, 0x44, 0x41, 0x59, 0x5f, 0x4f, 0x46,
	0x5f, 0x54, 0x48, 0x45, 0x5f, 0x57, 0x45, 0x45, 0x4b, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e,
	0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52,
	0x44, 0x41, 0x59, 0x10, 0x06, 0x22, 0x4d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x6e, 0x67,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x69,
	0x61, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x69, 0x6e,
	0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x32, 0xb0, 0x01, 0x0a, 0x03, 0x56, 0x69, 0x61, 0x12, 0x61,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x2e, 0x76, 0x69, 0x61, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x69, 0x61,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x19, 0x2e, 0x76, 0x69, 0x61, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x69,
	0x61, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x76, 0x69, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_via_proto_rawDescOnce sync.Once
	file_api_proto_via_proto_rawDescData = file_api_proto_via_proto_rawDesc
)

func file_api_proto_via_proto_rawDescGZIP() []byte {
	file_api_proto_via_proto_rawDescOnce.Do(func() {
		file_api_proto_via_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_via_proto_rawDescData)
	})
	return file_api_proto_via_proto_rawDescData
}

var file_api_proto_via_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_via_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_via_proto_goTypes = []interface{}{
	(Incident_DAY_OF_THE_WEEK)(0),        // 0: viapb.Incident.DAY_OF_THE_WEEK
	(*Incident)(nil),                     // 1: viapb.Incident
	(*GetDangerProbabilityRequest)(nil),  // 2: viapb.GetDangerProbabilityRequest
	(*GetDangerProbabilityResponse)(nil), // 3: viapb.GetDangerProbabilityResponse
	(*GetIncidentRequest)(nil),           // 4: viapb.GetIncidentRequest
	(*GetIncidentResponse)(nil),          // 5: viapb.GetIncidentResponse
}
var file_api_proto_via_proto_depIdxs = []int32{
	0, // 0: viapb.Incident.day:type_name -> viapb.Incident.DAY_OF_THE_WEEK
	1, // 1: viapb.GetIncidentResponse.incident:type_name -> viapb.Incident
	2, // 2: viapb.Via.GetDangerProbability:input_type -> viapb.GetDangerProbabilityRequest
	4, // 3: viapb.Via.GetIncident:input_type -> viapb.GetIncidentRequest
	3, // 4: viapb.Via.GetDangerProbability:output_type -> viapb.GetDangerProbabilityResponse
	5, // 5: viapb.Via.GetIncident:output_type -> viapb.GetIncidentResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_via_proto_init() }
func file_api_proto_via_proto_init() {
	if File_api_proto_via_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_via_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Incident); i {
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
		file_api_proto_via_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDangerProbabilityRequest); i {
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
		file_api_proto_via_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDangerProbabilityResponse); i {
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
		file_api_proto_via_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIncidentRequest); i {
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
		file_api_proto_via_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIncidentResponse); i {
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
			RawDescriptor: file_api_proto_via_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_via_proto_goTypes,
		DependencyIndexes: file_api_proto_via_proto_depIdxs,
		EnumInfos:         file_api_proto_via_proto_enumTypes,
		MessageInfos:      file_api_proto_via_proto_msgTypes,
	}.Build()
	File_api_proto_via_proto = out.File
	file_api_proto_via_proto_rawDesc = nil
	file_api_proto_via_proto_goTypes = nil
	file_api_proto_via_proto_depIdxs = nil
}
