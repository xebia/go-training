// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.2
// source: registration.proto

package regprotobuf

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

type RegistrationStatus int32

const (
	RegistrationStatus_REGISTRATION_UNKNOWN   RegistrationStatus = 0
	RegistrationStatus_REGISTRATION_PENDING   RegistrationStatus = 1
	RegistrationStatus_REGISTRATION_CONFIRMED RegistrationStatus = 2
	RegistrationStatus_REGISTRATION_BLOCKED   RegistrationStatus = 3
)

// Enum value maps for RegistrationStatus.
var (
	RegistrationStatus_name = map[int32]string{
		0: "REGISTRATION_UNKNOWN",
		1: "REGISTRATION_PENDING",
		2: "REGISTRATION_CONFIRMED",
		3: "REGISTRATION_BLOCKED",
	}
	RegistrationStatus_value = map[string]int32{
		"REGISTRATION_UNKNOWN":   0,
		"REGISTRATION_PENDING":   1,
		"REGISTRATION_CONFIRMED": 2,
		"REGISTRATION_BLOCKED":   3,
	}
)

func (x RegistrationStatus) Enum() *RegistrationStatus {
	p := new(RegistrationStatus)
	*p = x
	return p
}

func (x RegistrationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegistrationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_registration_proto_enumTypes[0].Descriptor()
}

func (RegistrationStatus) Type() protoreflect.EnumType {
	return &file_registration_proto_enumTypes[0]
}

func (x RegistrationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegistrationStatus.Descriptor instead.
func (RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{0}
}

type RegisterPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *RegisterPatientRequest) Reset() {
	*x = RegisterPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPatientRequest) ProtoMessage() {}

func (x *RegisterPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPatientRequest.ProtoReflect.Descriptor instead.
func (*RegisterPatientRequest) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterPatientRequest) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type RegisterPatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientUid string `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
}

func (x *RegisterPatientResponse) Reset() {
	*x = RegisterPatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPatientResponse) ProtoMessage() {}

func (x *RegisterPatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPatientResponse.ProtoReflect.Descriptor instead.
func (*RegisterPatientResponse) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterPatientResponse) GetPatientUid() string {
	if x != nil {
		return x.PatientUid
	}
	return ""
}

type CompletePatientRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientUid  string                   `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
	Credentials *RegistrationCredentials `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *CompletePatientRegistrationRequest) Reset() {
	*x = CompletePatientRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePatientRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePatientRegistrationRequest) ProtoMessage() {}

func (x *CompletePatientRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePatientRegistrationRequest.ProtoReflect.Descriptor instead.
func (*CompletePatientRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{2}
}

func (x *CompletePatientRegistrationRequest) GetPatientUid() string {
	if x != nil {
		return x.PatientUid
	}
	return ""
}

func (x *CompletePatientRegistrationRequest) GetCredentials() *RegistrationCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type RegistrationCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pincode int32 `protobuf:"varint,1,opt,name=pincode,proto3" json:"pincode,omitempty"`
}

func (x *RegistrationCredentials) Reset() {
	*x = RegistrationCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationCredentials) ProtoMessage() {}

func (x *RegistrationCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationCredentials.ProtoReflect.Descriptor instead.
func (*RegistrationCredentials) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{3}
}

func (x *RegistrationCredentials) GetPincode() int32 {
	if x != nil {
		return x.Pincode
	}
	return 0
}

type CompletePatientRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status RegistrationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=regprotobuf.RegistrationStatus" json:"status,omitempty"`
}

func (x *CompletePatientRegistrationResponse) Reset() {
	*x = CompletePatientRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePatientRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePatientRegistrationResponse) ProtoMessage() {}

func (x *CompletePatientRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePatientRegistrationResponse.ProtoReflect.Descriptor instead.
func (*CompletePatientRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{4}
}

func (x *CompletePatientRegistrationResponse) GetStatus() RegistrationStatus {
	if x != nil {
		return x.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BSN      string             `protobuf:"bytes,1,opt,name=BSN,proto3" json:"BSN,omitempty"`
	FullName string             `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Address  *Address           `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact  *Contact           `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	Status   RegistrationStatus `protobuf:"varint,5,opt,name=status,proto3,enum=regprotobuf.RegistrationStatus" json:"status,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{5}
}

func (x *Patient) GetBSN() string {
	if x != nil {
		return x.BSN
	}
	return ""
}

func (x *Patient) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Patient) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Patient) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Patient) GetStatus() RegistrationStatus {
	if x != nil {
		return x.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostalCode  string `protobuf:"bytes,1,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	HouseNumber int64  `protobuf:"varint,2,opt,name=houseNumber,proto3" json:"houseNumber,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{6}
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetHouseNumber() int64 {
	if x != nil {
		return x.HouseNumber
	}
	return 0
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	PhoneNumber  string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_registration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_registration_proto_rawDescGZIP(), []int{7}
}

func (x *Contact) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Contact) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

var File_registration_proto protoreflect.FileDescriptor

var file_registration_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x48, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x17, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x46, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5e, 0x0a, 0x23, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x53, 0x4e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x42, 0x53, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2a, 0x7e, 0x0a, 0x12, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x32, 0xfa, 0x01, 0x0a, 0x13,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x65, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x63, 0x47, 0x72, 0x6f, 0x6c, 0x2f,
	0x67, 0x6f, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registration_proto_rawDescOnce sync.Once
	file_registration_proto_rawDescData = file_registration_proto_rawDesc
)

func file_registration_proto_rawDescGZIP() []byte {
	file_registration_proto_rawDescOnce.Do(func() {
		file_registration_proto_rawDescData = protoimpl.X.CompressGZIP(file_registration_proto_rawDescData)
	})
	return file_registration_proto_rawDescData
}

var file_registration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_registration_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_registration_proto_goTypes = []interface{}{
	(RegistrationStatus)(0),                     // 0: regprotobuf.RegistrationStatus
	(*RegisterPatientRequest)(nil),              // 1: regprotobuf.RegisterPatientRequest
	(*RegisterPatientResponse)(nil),             // 2: regprotobuf.RegisterPatientResponse
	(*CompletePatientRegistrationRequest)(nil),  // 3: regprotobuf.CompletePatientRegistrationRequest
	(*RegistrationCredentials)(nil),             // 4: regprotobuf.RegistrationCredentials
	(*CompletePatientRegistrationResponse)(nil), // 5: regprotobuf.CompletePatientRegistrationResponse
	(*Patient)(nil),                             // 6: regprotobuf.Patient
	(*Address)(nil),                             // 7: regprotobuf.Address
	(*Contact)(nil),                             // 8: regprotobuf.Contact
}
var file_registration_proto_depIdxs = []int32{
	6, // 0: regprotobuf.RegisterPatientRequest.patient:type_name -> regprotobuf.Patient
	4, // 1: regprotobuf.CompletePatientRegistrationRequest.credentials:type_name -> regprotobuf.RegistrationCredentials
	0, // 2: regprotobuf.CompletePatientRegistrationResponse.status:type_name -> regprotobuf.RegistrationStatus
	7, // 3: regprotobuf.Patient.address:type_name -> regprotobuf.Address
	8, // 4: regprotobuf.Patient.contact:type_name -> regprotobuf.Contact
	0, // 5: regprotobuf.Patient.status:type_name -> regprotobuf.RegistrationStatus
	1, // 6: regprotobuf.RegistrationService.RegisterPatient:input_type -> regprotobuf.RegisterPatientRequest
	3, // 7: regprotobuf.RegistrationService.CompletePatientRegistration:input_type -> regprotobuf.CompletePatientRegistrationRequest
	2, // 8: regprotobuf.RegistrationService.RegisterPatient:output_type -> regprotobuf.RegisterPatientResponse
	5, // 9: regprotobuf.RegistrationService.CompletePatientRegistration:output_type -> regprotobuf.CompletePatientRegistrationResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_registration_proto_init() }
func file_registration_proto_init() {
	if File_registration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPatientRequest); i {
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
		file_registration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPatientResponse); i {
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
		file_registration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletePatientRegistrationRequest); i {
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
		file_registration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationCredentials); i {
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
		file_registration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletePatientRegistrationResponse); i {
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
		file_registration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_registration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_registration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_registration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_registration_proto_goTypes,
		DependencyIndexes: file_registration_proto_depIdxs,
		EnumInfos:         file_registration_proto_enumTypes,
		MessageInfos:      file_registration_proto_msgTypes,
	}.Build()
	File_registration_proto = out.File
	file_registration_proto_rawDesc = nil
	file_registration_proto_goTypes = nil
	file_registration_proto_depIdxs = nil
}
