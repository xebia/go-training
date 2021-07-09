// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registration.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegistrationStatus int32

const (
	RegistrationStatus_REGISTRATION_UNKNOWN   RegistrationStatus = 0
	RegistrationStatus_REGISTRATION_PENDING   RegistrationStatus = 1
	RegistrationStatus_REGISTRATION_CONFIRMED RegistrationStatus = 2
)

var RegistrationStatus_name = map[int32]string{
	0: "REGISTRATION_UNKNOWN",
	1: "REGISTRATION_PENDING",
	2: "REGISTRATION_CONFIRMED",
}

var RegistrationStatus_value = map[string]int32{
	"REGISTRATION_UNKNOWN":   0,
	"REGISTRATION_PENDING":   1,
	"REGISTRATION_CONFIRMED": 2,
}

func (x RegistrationStatus) String() string {
	return proto.EnumName(RegistrationStatus_name, int32(x))
}

func (RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{0}
}

type RegisterPatientRequest struct {
	Patient              *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPatientRequest) Reset()         { *m = RegisterPatientRequest{} }
func (m *RegisterPatientRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterPatientRequest) ProtoMessage()    {}
func (*RegisterPatientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{0}
}

func (m *RegisterPatientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPatientRequest.Unmarshal(m, b)
}
func (m *RegisterPatientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPatientRequest.Marshal(b, m, deterministic)
}
func (m *RegisterPatientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPatientRequest.Merge(m, src)
}
func (m *RegisterPatientRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterPatientRequest.Size(m)
}
func (m *RegisterPatientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPatientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPatientRequest proto.InternalMessageInfo

func (m *RegisterPatientRequest) GetPatient() *Patient {
	if m != nil {
		return m.Patient
	}
	return nil
}

type RegisterPatientResponse struct {
	PatientUid           string   `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPatientResponse) Reset()         { *m = RegisterPatientResponse{} }
func (m *RegisterPatientResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterPatientResponse) ProtoMessage()    {}
func (*RegisterPatientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{1}
}

func (m *RegisterPatientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPatientResponse.Unmarshal(m, b)
}
func (m *RegisterPatientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPatientResponse.Marshal(b, m, deterministic)
}
func (m *RegisterPatientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPatientResponse.Merge(m, src)
}
func (m *RegisterPatientResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterPatientResponse.Size(m)
}
func (m *RegisterPatientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPatientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPatientResponse proto.InternalMessageInfo

func (m *RegisterPatientResponse) GetPatientUid() string {
	if m != nil {
		return m.PatientUid
	}
	return ""
}

type CompletePatientRegistrationRequest struct {
	Credentials          *RegistrationCredentials `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CompletePatientRegistrationRequest) Reset()         { *m = CompletePatientRegistrationRequest{} }
func (m *CompletePatientRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*CompletePatientRegistrationRequest) ProtoMessage()    {}
func (*CompletePatientRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{2}
}

func (m *CompletePatientRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompletePatientRegistrationRequest.Unmarshal(m, b)
}
func (m *CompletePatientRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompletePatientRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *CompletePatientRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletePatientRegistrationRequest.Merge(m, src)
}
func (m *CompletePatientRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_CompletePatientRegistrationRequest.Size(m)
}
func (m *CompletePatientRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletePatientRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompletePatientRegistrationRequest proto.InternalMessageInfo

func (m *CompletePatientRegistrationRequest) GetCredentials() *RegistrationCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type RegistrationCredentials struct {
	PatientUid           string   `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
	Pincode              int32    `protobuf:"varint,2,opt,name=pincode,proto3" json:"pincode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationCredentials) Reset()         { *m = RegistrationCredentials{} }
func (m *RegistrationCredentials) String() string { return proto.CompactTextString(m) }
func (*RegistrationCredentials) ProtoMessage()    {}
func (*RegistrationCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{3}
}

func (m *RegistrationCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationCredentials.Unmarshal(m, b)
}
func (m *RegistrationCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationCredentials.Marshal(b, m, deterministic)
}
func (m *RegistrationCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationCredentials.Merge(m, src)
}
func (m *RegistrationCredentials) XXX_Size() int {
	return xxx_messageInfo_RegistrationCredentials.Size(m)
}
func (m *RegistrationCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationCredentials proto.InternalMessageInfo

func (m *RegistrationCredentials) GetPatientUid() string {
	if m != nil {
		return m.PatientUid
	}
	return ""
}

func (m *RegistrationCredentials) GetPincode() int32 {
	if m != nil {
		return m.Pincode
	}
	return 0
}

type CompletePatientRegistrationResponse struct {
	Status               RegistrationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=service.RegistrationStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CompletePatientRegistrationResponse) Reset()         { *m = CompletePatientRegistrationResponse{} }
func (m *CompletePatientRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*CompletePatientRegistrationResponse) ProtoMessage()    {}
func (*CompletePatientRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{4}
}

func (m *CompletePatientRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompletePatientRegistrationResponse.Unmarshal(m, b)
}
func (m *CompletePatientRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompletePatientRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *CompletePatientRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletePatientRegistrationResponse.Merge(m, src)
}
func (m *CompletePatientRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_CompletePatientRegistrationResponse.Size(m)
}
func (m *CompletePatientRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletePatientRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompletePatientRegistrationResponse proto.InternalMessageInfo

func (m *CompletePatientRegistrationResponse) GetStatus() RegistrationStatus {
	if m != nil {
		return m.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

type Patient struct {
	BSN                  string             `protobuf:"bytes,1,opt,name=BSN,proto3" json:"BSN,omitempty"`
	FullName             string             `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Address              *Address           `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact              *Contact           `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
	Status               RegistrationStatus `protobuf:"varint,5,opt,name=status,proto3,enum=service.RegistrationStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Patient) Reset()         { *m = Patient{} }
func (m *Patient) String() string { return proto.CompactTextString(m) }
func (*Patient) ProtoMessage()    {}
func (*Patient) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{5}
}

func (m *Patient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Patient.Unmarshal(m, b)
}
func (m *Patient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Patient.Marshal(b, m, deterministic)
}
func (m *Patient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Patient.Merge(m, src)
}
func (m *Patient) XXX_Size() int {
	return xxx_messageInfo_Patient.Size(m)
}
func (m *Patient) XXX_DiscardUnknown() {
	xxx_messageInfo_Patient.DiscardUnknown(m)
}

var xxx_messageInfo_Patient proto.InternalMessageInfo

func (m *Patient) GetBSN() string {
	if m != nil {
		return m.BSN
	}
	return ""
}

func (m *Patient) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Patient) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Patient) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Patient) GetStatus() RegistrationStatus {
	if m != nil {
		return m.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

type Address struct {
	PostalCode           string   `protobuf:"bytes,1,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	HouseNumber          int64    `protobuf:"varint,2,opt,name=houseNumber,proto3" json:"houseNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{6}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Address) GetHouseNumber() int64 {
	if m != nil {
		return m.HouseNumber
	}
	return 0
}

type Contact struct {
	EmailAddress         string   `protobuf:"bytes,1,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_199f7aef77c18626, []int{7}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Contact) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func init() {
	proto.RegisterEnum("service.RegistrationStatus", RegistrationStatus_name, RegistrationStatus_value)
	proto.RegisterType((*RegisterPatientRequest)(nil), "service.RegisterPatientRequest")
	proto.RegisterType((*RegisterPatientResponse)(nil), "service.RegisterPatientResponse")
	proto.RegisterType((*CompletePatientRegistrationRequest)(nil), "service.CompletePatientRegistrationRequest")
	proto.RegisterType((*RegistrationCredentials)(nil), "service.RegistrationCredentials")
	proto.RegisterType((*CompletePatientRegistrationResponse)(nil), "service.CompletePatientRegistrationResponse")
	proto.RegisterType((*Patient)(nil), "service.Patient")
	proto.RegisterType((*Address)(nil), "service.Address")
	proto.RegisterType((*Contact)(nil), "service.Contact")
}

func init() { proto.RegisterFile("registration.proto", fileDescriptor_199f7aef77c18626) }

var fileDescriptor_199f7aef77c18626 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x95, 0x2d, 0xf4, 0x15, 0x41, 0x64, 0xd0, 0x88, 0x3a, 0x09, 0x22, 0x73, 0x99,
	0x06, 0xea, 0xa1, 0x3b, 0x71, 0xdc, 0xd2, 0x32, 0x55, 0x13, 0xee, 0xe4, 0x6e, 0x20, 0x71, 0x41,
	0x5e, 0xf2, 0xa0, 0x91, 0xd2, 0x38, 0xc4, 0x0e, 0x07, 0x3e, 0x1f, 0x5f, 0x86, 0x6f, 0x81, 0x92,
	0xd8, 0x5d, 0x9a, 0xc2, 0xb6, 0x5b, 0xfc, 0xf7, 0xff, 0xfd, 0xfd, 0xf3, 0xb3, 0x1d, 0x20, 0x05,
	0x7e, 0x4f, 0x94, 0x2e, 0x84, 0x4e, 0x64, 0x36, 0xca, 0x0b, 0xa9, 0x25, 0x71, 0x15, 0x16, 0x3f,
	0x93, 0x08, 0xe9, 0x04, 0x0e, 0x78, 0x3d, 0x8d, 0xc5, 0xa5, 0xd0, 0x09, 0x66, 0x9a, 0xe3, 0x8f,
	0x12, 0x95, 0x26, 0xc7, 0xe0, 0xe6, 0x8d, 0xe2, 0x3b, 0x81, 0x73, 0x34, 0x18, 0x7b, 0x23, 0x53,
	0x34, 0xb2, 0x4e, 0x6b, 0xa0, 0xef, 0xe1, 0xe5, 0x56, 0x8a, 0xca, 0x65, 0xa6, 0x90, 0xbc, 0x02,
	0x30, 0xae, 0xeb, 0x24, 0xae, 0x93, 0xfa, 0xbc, 0xa5, 0xd0, 0x25, 0xd0, 0x50, 0xae, 0xf2, 0x14,
	0x35, 0xae, 0x4b, 0x6f, 0x71, 0x2d, 0xcc, 0x19, 0x0c, 0xa2, 0x02, 0x63, 0xcc, 0x74, 0x22, 0x52,
	0x65, 0x80, 0x82, 0x35, 0x50, 0xbb, 0x24, 0xbc, 0xf5, 0xf1, 0x76, 0x11, 0x5d, 0x58, 0xc8, 0x2d,
	0xdf, 0x7d, 0x90, 0xc4, 0x07, 0x37, 0x4f, 0xb2, 0x48, 0xc6, 0xe8, 0xef, 0x06, 0xce, 0xd1, 0x1e,
	0xb7, 0x43, 0xfa, 0x05, 0xde, 0xdc, 0x89, 0x6f, 0xba, 0x70, 0x02, 0xfb, 0x4a, 0x0b, 0x5d, 0x36,
	0xe8, 0x4f, 0xc7, 0x87, 0xff, 0x44, 0x5f, 0xd4, 0x16, 0x6e, 0xac, 0xf4, 0xb7, 0x03, 0xae, 0x09,
	0x25, 0x1e, 0xf4, 0xce, 0x16, 0xcc, 0xa0, 0x55, 0x9f, 0x64, 0x08, 0x8f, 0xbf, 0x95, 0x69, 0xca,
	0xc4, 0xaa, 0x81, 0xea, 0xf3, 0xf5, 0xb8, 0x3a, 0x3b, 0x11, 0xc7, 0x05, 0x2a, 0xe5, 0xf7, 0x3a,
	0x67, 0x77, 0xda, 0xe8, 0xdc, 0x1a, 0x2a, 0x6f, 0x24, 0x33, 0x2d, 0x22, 0xed, 0x3f, 0xea, 0x78,
	0xc3, 0x46, 0xe7, 0xd6, 0xd0, 0xda, 0xc6, 0xde, 0xc3, 0xb7, 0x71, 0x01, 0xae, 0x59, 0xb4, 0xee,
	0xb3, 0x54, 0x5a, 0xa4, 0x61, 0xd5, 0x4a, 0xdb, 0xe7, 0xb5, 0x42, 0x02, 0x18, 0x2c, 0x65, 0xa9,
	0x90, 0x95, 0xab, 0x1b, 0x2c, 0xea, 0x6d, 0xf5, 0x78, 0x5b, 0xa2, 0x73, 0x70, 0x0d, 0x15, 0xa1,
	0xf0, 0x04, 0x57, 0x22, 0x49, 0x4d, 0xb8, 0x89, 0xdb, 0xd0, 0xaa, 0xc0, 0x7c, 0x29, 0xb3, 0x76,
	0x60, 0x9f, 0xb7, 0xa5, 0xe3, 0x18, 0xc8, 0x36, 0x3b, 0xf1, 0xe1, 0x05, 0x9f, 0x9e, 0xcf, 0x16,
	0x57, 0xfc, 0xf4, 0x6a, 0x36, 0x67, 0x5f, 0xaf, 0xd9, 0x05, 0x9b, 0x7f, 0x66, 0xde, 0xce, 0xd6,
	0xcc, 0xe5, 0x94, 0x4d, 0x66, 0xec, 0xdc, 0x73, 0xc8, 0x10, 0x0e, 0x36, 0x66, 0xc2, 0x39, 0xfb,
	0x30, 0xe3, 0x1f, 0xa7, 0x13, 0x6f, 0x77, 0xfc, 0xc7, 0x81, 0xe7, 0x1b, 0xcb, 0x34, 0x6d, 0x23,
	0x9f, 0xe0, 0x59, 0xe7, 0xe1, 0x90, 0xd7, 0x9d, 0x9e, 0x76, 0x1f, 0xe6, 0x30, 0xf8, 0xbf, 0xa1,
	0xb9, 0x6d, 0x74, 0x87, 0xfc, 0x82, 0xc3, 0x3b, 0xae, 0x25, 0x79, 0xdb, 0x3a, 0xe2, 0xfb, 0xde,
	0xde, 0xf0, 0xdd, 0xc3, 0xcc, 0x76, 0xed, 0x9b, 0xfd, 0xfa, 0x17, 0x73, 0xf2, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0x9f, 0x03, 0x05, 0x78, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	RegisterPatient(ctx context.Context, in *RegisterPatientRequest, opts ...grpc.CallOption) (*RegisterPatientResponse, error)
	CompletePatientRegistration(ctx context.Context, in *CompletePatientRegistrationRequest, opts ...grpc.CallOption) (*CompletePatientRegistrationResponse, error)
}

type registrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationServiceClient(cc *grpc.ClientConn) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) RegisterPatient(ctx context.Context, in *RegisterPatientRequest, opts ...grpc.CallOption) (*RegisterPatientResponse, error) {
	out := new(RegisterPatientResponse)
	err := c.cc.Invoke(ctx, "/service.RegistrationService/RegisterPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationServiceClient) CompletePatientRegistration(ctx context.Context, in *CompletePatientRegistrationRequest, opts ...grpc.CallOption) (*CompletePatientRegistrationResponse, error) {
	out := new(CompletePatientRegistrationResponse)
	err := c.cc.Invoke(ctx, "/service.RegistrationService/CompletePatientRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
type RegistrationServiceServer interface {
	RegisterPatient(context.Context, *RegisterPatientRequest) (*RegisterPatientResponse, error)
	CompletePatientRegistration(context.Context, *CompletePatientRegistrationRequest) (*CompletePatientRegistrationResponse, error)
}

// UnimplementedRegistrationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (*UnimplementedRegistrationServiceServer) RegisterPatient(ctx context.Context, req *RegisterPatientRequest) (*RegisterPatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPatient not implemented")
}
func (*UnimplementedRegistrationServiceServer) CompletePatientRegistration(ctx context.Context, req *CompletePatientRegistrationRequest) (*CompletePatientRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompletePatientRegistration not implemented")
}

func RegisterRegistrationServiceServer(s *grpc.Server, srv RegistrationServiceServer) {
	s.RegisterService(&_RegistrationService_serviceDesc, srv)
}

func _RegistrationService_RegisterPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).RegisterPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RegistrationService/RegisterPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).RegisterPatient(ctx, req.(*RegisterPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationService_CompletePatientRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletePatientRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).CompletePatientRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RegistrationService/CompletePatientRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).CompletePatientRegistration(ctx, req.(*CompletePatientRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPatient",
			Handler:    _RegistrationService_RegisterPatient_Handler,
		},
		{
			MethodName: "CompletePatientRegistration",
			Handler:    _RegistrationService_CompletePatientRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registration.proto",
}