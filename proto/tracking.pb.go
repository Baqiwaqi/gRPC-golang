// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: service/tracking.proto

package golang_grpc_servive

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

type NoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParams) Reset() {
	*x = NoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParams) ProtoMessage() {}

func (x *NoParams) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParams.ProtoReflect.Descriptor instead.
func (*NoParams) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{0}
}

type MyService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MyService) Reset() {
	*x = MyService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService) ProtoMessage() {}

func (x *MyService) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService.ProtoReflect.Descriptor instead.
func (*MyService) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1}
}

type MyService_Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number    string             `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Airline   *MyService_Airline `protobuf:"bytes,3,opt,name=airline,proto3" json:"airline,omitempty"`
	Departure *MyService_Airport `protobuf:"bytes,4,opt,name=departure,proto3" json:"departure,omitempty"`
	Arrival   *MyService_Airport `protobuf:"bytes,5,opt,name=arrival,proto3" json:"arrival,omitempty"`
}

func (x *MyService_Flight) Reset() {
	*x = MyService_Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Flight) ProtoMessage() {}

func (x *MyService_Flight) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Flight.ProtoReflect.Descriptor instead.
func (*MyService_Flight) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MyService_Flight) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MyService_Flight) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *MyService_Flight) GetAirline() *MyService_Airline {
	if x != nil {
		return x.Airline
	}
	return nil
}

func (x *MyService_Flight) GetDeparture() *MyService_Airport {
	if x != nil {
		return x.Departure
	}
	return nil
}

func (x *MyService_Flight) GetArrival() *MyService_Airport {
	if x != nil {
		return x.Arrival
	}
	return nil
}

type MyService_Airport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *MyService_Airport) Reset() {
	*x = MyService_Airport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Airport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Airport) ProtoMessage() {}

func (x *MyService_Airport) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Airport.ProtoReflect.Descriptor instead.
func (*MyService_Airport) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 1}
}

func (x *MyService_Airport) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MyService_Airport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MyService_Airport) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MyService_Airport) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type MyService_Airline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *MyService_Airline) Reset() {
	*x = MyService_Airline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Airline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Airline) ProtoMessage() {}

func (x *MyService_Airline) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Airline.ProtoReflect.Descriptor instead.
func (*MyService_Airline) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 2}
}

func (x *MyService_Airline) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MyService_Airline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MyService_Airline) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type MyService_Flight_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MyService_Flight_Request) Reset() {
	*x = MyService_Flight_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Flight_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Flight_Request) ProtoMessage() {}

func (x *MyService_Flight_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Flight_Request.ProtoReflect.Descriptor instead.
func (*MyService_Flight_Request) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *MyService_Flight_Request) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MyService_Flight_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flight *MyService_Flight `protobuf:"bytes,1,opt,name=flight,proto3" json:"flight,omitempty"`
}

func (x *MyService_Flight_Response) Reset() {
	*x = MyService_Flight_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Flight_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Flight_Response) ProtoMessage() {}

func (x *MyService_Flight_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Flight_Response.ProtoReflect.Descriptor instead.
func (*MyService_Flight_Response) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *MyService_Flight_Response) GetFlight() *MyService_Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

type MyService_Flight_ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*MyService_Flight `protobuf:"bytes,1,rep,name=flights,proto3" json:"flights,omitempty"`
}

func (x *MyService_Flight_ListResponse) Reset() {
	*x = MyService_Flight_ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Flight_ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Flight_ListResponse) ProtoMessage() {}

func (x *MyService_Flight_ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Flight_ListResponse.ProtoReflect.Descriptor instead.
func (*MyService_Flight_ListResponse) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 0, 2}
}

func (x *MyService_Flight_ListResponse) GetFlights() []*MyService_Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

type MyService_Flight_NoParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MyService_Flight_NoParams) Reset() {
	*x = MyService_Flight_NoParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_tracking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyService_Flight_NoParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyService_Flight_NoParams) ProtoMessage() {}

func (x *MyService_Flight_NoParams) ProtoReflect() protoreflect.Message {
	mi := &file_service_tracking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyService_Flight_NoParams.ProtoReflect.Descriptor instead.
func (*MyService_Flight_NoParams) Descriptor() ([]byte, []int) {
	return file_service_tracking_proto_rawDescGZIP(), []int{1, 0, 3}
}

var File_service_tracking_proto protoreflect.FileDescriptor

var file_service_tracking_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x0a, 0x0a, 0x08, 0x4e, 0x6f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0xe0, 0x04, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0xa4, 0x03, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x61, 0x69, 0x72, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c,
	0x1a, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x44, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x1a, 0x4a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x1a, 0x0a, 0x0a,
	0x08, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x5f, 0x0a, 0x07, 0x41, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x4b, 0x0a, 0x07, 0x41, 0x69,
	0x72, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x32, 0xc2, 0x02, 0x0a, 0x0e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x60, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x28, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x28, 0x2e,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x66, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4d,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x71, 0x69,
	0x77, 0x61, 0x71, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_tracking_proto_rawDescOnce sync.Once
	file_service_tracking_proto_rawDescData = file_service_tracking_proto_rawDesc
)

func file_service_tracking_proto_rawDescGZIP() []byte {
	file_service_tracking_proto_rawDescOnce.Do(func() {
		file_service_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_tracking_proto_rawDescData)
	})
	return file_service_tracking_proto_rawDescData
}

var file_service_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_tracking_proto_goTypes = []interface{}{
	(*NoParams)(nil),                      // 0: flightTracking.NoParams
	(*MyService)(nil),                     // 1: flightTracking.MyService
	(*MyService_Flight)(nil),              // 2: flightTracking.MyService.Flight
	(*MyService_Airport)(nil),             // 3: flightTracking.MyService.Airport
	(*MyService_Airline)(nil),             // 4: flightTracking.MyService.Airline
	(*MyService_Flight_Request)(nil),      // 5: flightTracking.MyService.Flight.Request
	(*MyService_Flight_Response)(nil),     // 6: flightTracking.MyService.Flight.Response
	(*MyService_Flight_ListResponse)(nil), // 7: flightTracking.MyService.Flight.ListResponse
	(*MyService_Flight_NoParams)(nil),     // 8: flightTracking.MyService.Flight.NoParams
}
var file_service_tracking_proto_depIdxs = []int32{
	4, // 0: flightTracking.MyService.Flight.airline:type_name -> flightTracking.MyService.Airline
	3, // 1: flightTracking.MyService.Flight.departure:type_name -> flightTracking.MyService.Airport
	3, // 2: flightTracking.MyService.Flight.arrival:type_name -> flightTracking.MyService.Airport
	2, // 3: flightTracking.MyService.Flight.Response.flight:type_name -> flightTracking.MyService.Flight
	2, // 4: flightTracking.MyService.Flight.ListResponse.flights:type_name -> flightTracking.MyService.Flight
	5, // 5: flightTracking.FlightTracking.GetFlight:input_type -> flightTracking.MyService.Flight.Request
	5, // 6: flightTracking.FlightTracking.StreamFlights:input_type -> flightTracking.MyService.Flight.Request
	5, // 7: flightTracking.FlightTracking.ListFlights:input_type -> flightTracking.MyService.Flight.Request
	6, // 8: flightTracking.FlightTracking.GetFlight:output_type -> flightTracking.MyService.Flight.Response
	6, // 9: flightTracking.FlightTracking.StreamFlights:output_type -> flightTracking.MyService.Flight.Response
	7, // 10: flightTracking.FlightTracking.ListFlights:output_type -> flightTracking.MyService.Flight.ListResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_tracking_proto_init() }
func file_service_tracking_proto_init() {
	if File_service_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParams); i {
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
		file_service_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService); i {
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
		file_service_tracking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Flight); i {
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
		file_service_tracking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Airport); i {
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
		file_service_tracking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Airline); i {
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
		file_service_tracking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Flight_Request); i {
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
		file_service_tracking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Flight_Response); i {
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
		file_service_tracking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Flight_ListResponse); i {
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
		file_service_tracking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyService_Flight_NoParams); i {
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
			RawDescriptor: file_service_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_tracking_proto_goTypes,
		DependencyIndexes: file_service_tracking_proto_depIdxs,
		MessageInfos:      file_service_tracking_proto_msgTypes,
	}.Build()
	File_service_tracking_proto = out.File
	file_service_tracking_proto_rawDesc = nil
	file_service_tracking_proto_goTypes = nil
	file_service_tracking_proto_depIdxs = nil
}
