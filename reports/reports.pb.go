// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: reports/reports.proto

package reports

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmptyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_reports_reports_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{0}
}

type TimeBasedReport struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      int32                  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Distance      int32                  `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	Heartbeats    []int64                `protobuf:"varint,4,rep,packed,name=heartbeats,proto3" json:"heartbeats,omitempty"`
	Breaths       []int64                `protobuf:"varint,5,rep,packed,name=breaths,proto3" json:"breaths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeBasedReport) Reset() {
	*x = TimeBasedReport{}
	mi := &file_reports_reports_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeBasedReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeBasedReport) ProtoMessage() {}

func (x *TimeBasedReport) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeBasedReport.ProtoReflect.Descriptor instead.
func (*TimeBasedReport) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{1}
}

func (x *TimeBasedReport) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *TimeBasedReport) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TimeBasedReport) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *TimeBasedReport) GetHeartbeats() []int64 {
	if x != nil {
		return x.Heartbeats
	}
	return nil
}

func (x *TimeBasedReport) GetBreaths() []int64 {
	if x != nil {
		return x.Breaths
	}
	return nil
}

type DeviceControlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      int32                  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceControlRequest) Reset() {
	*x = DeviceControlRequest{}
	mi := &file_reports_reports_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceControlRequest) ProtoMessage() {}

func (x *DeviceControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceControlRequest.ProtoReflect.Descriptor instead.
func (*DeviceControlRequest) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceControlRequest) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type DeviceControlResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceControlResponse) Reset() {
	*x = DeviceControlResponse{}
	mi := &file_reports_reports_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceControlResponse) ProtoMessage() {}

func (x *DeviceControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceControlResponse.ProtoReflect.Descriptor instead.
func (*DeviceControlResponse) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceControlResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeviceStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      int32                  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Muted         bool                   `protobuf:"varint,2,opt,name=muted,proto3" json:"muted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceStatus) Reset() {
	*x = DeviceStatus{}
	mi := &file_reports_reports_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatus) ProtoMessage() {}

func (x *DeviceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatus.ProtoReflect.Descriptor instead.
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{4}
}

func (x *DeviceStatus) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *DeviceStatus) GetMuted() bool {
	if x != nil {
		return x.Muted
	}
	return false
}

type DeviceStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Statuses      []*DeviceStatus        `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceStatusResponse) Reset() {
	*x = DeviceStatusResponse{}
	mi := &file_reports_reports_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceStatusResponse) ProtoMessage() {}

func (x *DeviceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceStatusResponse.ProtoReflect.Descriptor instead.
func (*DeviceStatusResponse) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceStatusResponse) GetStatuses() []*DeviceStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type EventBasedReport struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      int32                  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Timestamp     int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EventId       int32                  `protobuf:"varint,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventData     []int32                `protobuf:"varint,4,rep,packed,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventBasedReport) Reset() {
	*x = EventBasedReport{}
	mi := &file_reports_reports_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventBasedReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBasedReport) ProtoMessage() {}

func (x *EventBasedReport) ProtoReflect() protoreflect.Message {
	mi := &file_reports_reports_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBasedReport.ProtoReflect.Descriptor instead.
func (*EventBasedReport) Descriptor() ([]byte, []int) {
	return file_reports_reports_proto_rawDescGZIP(), []int{6}
}

func (x *EventBasedReport) GetDeviceId() int32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *EventBasedReport) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *EventBasedReport) GetEventId() int32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *EventBasedReport) GetEventData() []int32 {
	if x != nil {
		return x.EventData
	}
	return nil
}

var File_reports_reports_proto protoreflect.FileDescriptor

const file_reports_reports_proto_rawDesc = "" +
	"\n" +
	"\x15reports/reports.proto\x12\areports\"\x0e\n" +
	"\fEmptyRequest\"\xa2\x01\n" +
	"\x0fTimeBasedReport\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\x05R\bdeviceId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\x12\x1a\n" +
	"\bdistance\x18\x03 \x01(\x05R\bdistance\x12\x1e\n" +
	"\n" +
	"heartbeats\x18\x04 \x03(\x03R\n" +
	"heartbeats\x12\x18\n" +
	"\abreaths\x18\x05 \x03(\x03R\abreaths\"3\n" +
	"\x14DeviceControlRequest\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\x05R\bdeviceId\"1\n" +
	"\x15DeviceControlResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"A\n" +
	"\fDeviceStatus\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\x05R\bdeviceId\x12\x14\n" +
	"\x05muted\x18\x02 \x01(\bR\x05muted\"I\n" +
	"\x14DeviceStatusResponse\x121\n" +
	"\bstatuses\x18\x01 \x03(\v2\x15.reports.DeviceStatusR\bstatuses\"\x87\x01\n" +
	"\x10EventBasedReport\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\x05R\bdeviceId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\x12\x19\n" +
	"\bevent_id\x18\x03 \x01(\x05R\aeventId\x12\x1d\n" +
	"\n" +
	"event_data\x18\x04 \x03(\x05R\teventData2\x92\x03\n" +
	"\rReportService\x12K\n" +
	"\x16StreamTimeBasedReports\x12\x15.reports.EmptyRequest\x1a\x18.reports.TimeBasedReport0\x01\x12M\n" +
	"\x17StreamEventBasedReports\x12\x15.reports.EmptyRequest\x1a\x19.reports.EventBasedReport0\x01\x12K\n" +
	"\n" +
	"MuteDevice\x12\x1d.reports.DeviceControlRequest\x1a\x1e.reports.DeviceControlResponse\x12M\n" +
	"\fUnmuteDevice\x12\x1d.reports.DeviceControlRequest\x1a\x1e.reports.DeviceControlResponse\x12I\n" +
	"\x11GetDeviceStatuses\x12\x15.reports.EmptyRequest\x1a\x1d.reports.DeviceStatusResponseB\x18Z\x16data-generator/reportsb\x06proto3"

var (
	file_reports_reports_proto_rawDescOnce sync.Once
	file_reports_reports_proto_rawDescData []byte
)

func file_reports_reports_proto_rawDescGZIP() []byte {
	file_reports_reports_proto_rawDescOnce.Do(func() {
		file_reports_reports_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_reports_reports_proto_rawDesc), len(file_reports_reports_proto_rawDesc)))
	})
	return file_reports_reports_proto_rawDescData
}

var file_reports_reports_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_reports_reports_proto_goTypes = []any{
	(*EmptyRequest)(nil),          // 0: reports.EmptyRequest
	(*TimeBasedReport)(nil),       // 1: reports.TimeBasedReport
	(*DeviceControlRequest)(nil),  // 2: reports.DeviceControlRequest
	(*DeviceControlResponse)(nil), // 3: reports.DeviceControlResponse
	(*DeviceStatus)(nil),          // 4: reports.DeviceStatus
	(*DeviceStatusResponse)(nil),  // 5: reports.DeviceStatusResponse
	(*EventBasedReport)(nil),      // 6: reports.EventBasedReport
}
var file_reports_reports_proto_depIdxs = []int32{
	4, // 0: reports.DeviceStatusResponse.statuses:type_name -> reports.DeviceStatus
	0, // 1: reports.ReportService.StreamTimeBasedReports:input_type -> reports.EmptyRequest
	0, // 2: reports.ReportService.StreamEventBasedReports:input_type -> reports.EmptyRequest
	2, // 3: reports.ReportService.MuteDevice:input_type -> reports.DeviceControlRequest
	2, // 4: reports.ReportService.UnmuteDevice:input_type -> reports.DeviceControlRequest
	0, // 5: reports.ReportService.GetDeviceStatuses:input_type -> reports.EmptyRequest
	1, // 6: reports.ReportService.StreamTimeBasedReports:output_type -> reports.TimeBasedReport
	6, // 7: reports.ReportService.StreamEventBasedReports:output_type -> reports.EventBasedReport
	3, // 8: reports.ReportService.MuteDevice:output_type -> reports.DeviceControlResponse
	3, // 9: reports.ReportService.UnmuteDevice:output_type -> reports.DeviceControlResponse
	5, // 10: reports.ReportService.GetDeviceStatuses:output_type -> reports.DeviceStatusResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reports_reports_proto_init() }
func file_reports_reports_proto_init() {
	if File_reports_reports_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_reports_reports_proto_rawDesc), len(file_reports_reports_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reports_reports_proto_goTypes,
		DependencyIndexes: file_reports_reports_proto_depIdxs,
		MessageInfos:      file_reports_reports_proto_msgTypes,
	}.Build()
	File_reports_reports_proto = out.File
	file_reports_reports_proto_goTypes = nil
	file_reports_reports_proto_depIdxs = nil
}
