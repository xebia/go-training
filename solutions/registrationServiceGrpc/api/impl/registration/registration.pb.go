// Code generated by protoc-gen-go. DO NOT EDIT.
// source: registration.proto

package registration

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	Status               RegistrationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.RegistrationStatus" json:"status,omitempty"`
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
	FullName             string             `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Status               RegistrationStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.RegistrationStatus" json:"status,omitempty"`
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

func (m *Patient) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Patient) GetStatus() RegistrationStatus {
	if m != nil {
		return m.Status
	}
	return RegistrationStatus_REGISTRATION_UNKNOWN
}

func init() {
	proto.RegisterEnum("api.RegistrationStatus", RegistrationStatus_name, RegistrationStatus_value)
	proto.RegisterType((*RegisterPatientRequest)(nil), "api.RegisterPatientRequest")
	proto.RegisterType((*RegisterPatientResponse)(nil), "api.RegisterPatientResponse")
	proto.RegisterType((*CompletePatientRegistrationRequest)(nil), "api.CompletePatientRegistrationRequest")
	proto.RegisterType((*RegistrationCredentials)(nil), "api.RegistrationCredentials")
	proto.RegisterType((*CompletePatientRegistrationResponse)(nil), "api.CompletePatientRegistrationResponse")
	proto.RegisterType((*Patient)(nil), "api.Patient")
}

func init() { proto.RegisterFile("registration.proto", fileDescriptor_199f7aef77c18626) }

var fileDescriptor_199f7aef77c18626 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x49, 0x10, 0x2b, 0xfc, 0x37, 0x41, 0x65, 0xa6, 0xb6, 0xca, 0x0a, 0x9a, 0x3c, 0x09,
	0x2a, 0x0e, 0x89, 0x54, 0x2e, 0xb0, 0x03, 0x02, 0xba, 0x32, 0x55, 0x88, 0x74, 0x72, 0xb7, 0x21,
	0x71, 0x41, 0xa6, 0xf9, 0x53, 0x2c, 0x52, 0xdb, 0x24, 0x2e, 0x17, 0xc4, 0x85, 0x57, 0xe0, 0x2d,
	0xb8, 0xf0, 0x30, 0xbc, 0x02, 0x0f, 0x82, 0xe4, 0x3a, 0xab, 0xb7, 0xa0, 0x75, 0x47, 0xfb, 0xfb,
	0xfe, 0x5f, 0x7e, 0xb1, 0x3f, 0x03, 0x29, 0x70, 0x26, 0x4a, 0x53, 0x70, 0x23, 0x94, 0x8c, 0x75,
	0xa1, 0x8c, 0x22, 0xd7, 0xb9, 0x16, 0x51, 0x77, 0xa6, 0xd4, 0x2c, 0xc7, 0x84, 0x6b, 0x91, 0x70,
	0x29, 0x95, 0xb1, 0x8e, 0x72, 0x69, 0xa1, 0xcf, 0xa1, 0xc5, 0xec, 0x20, 0x16, 0x47, 0xdc, 0x08,
	0x94, 0x86, 0xe1, 0x97, 0x05, 0x96, 0x86, 0x3c, 0x80, 0x86, 0x5e, 0xee, 0x74, 0x82, 0xdd, 0xa0,
	0xb7, 0xd9, 0xdf, 0x8a, 0xb9, 0x16, 0x71, 0xe5, 0xaa, 0x44, 0xfa, 0x14, 0xda, 0xb5, 0x84, 0x52,
	0x2b, 0x59, 0x22, 0xb9, 0x0f, 0xe0, 0x5c, 0x27, 0x22, 0xb3, 0x29, 0xb7, 0x98, 0xb7, 0x43, 0x33,
	0xa0, 0x03, 0x35, 0xd7, 0x39, 0x1a, 0x3c, 0x1b, 0x5d, 0xfd, 0x44, 0x05, 0xf2, 0x0c, 0x36, 0xa7,
	0x05, 0x66, 0x28, 0x8d, 0xe0, 0x79, 0xe9, 0x60, 0xba, 0x16, 0xc6, 0xb7, 0x0f, 0x56, 0x1e, 0xe6,
	0x0f, 0xd0, 0x49, 0x05, 0x58, 0xf3, 0xad, 0x03, 0x24, 0x1d, 0x68, 0x68, 0x21, 0xa7, 0x2a, 0xc3,
	0x4e, 0xb8, 0x1b, 0xf4, 0x6e, 0xb0, 0x6a, 0x49, 0x4f, 0x61, 0xef, 0x52, 0x74, 0x77, 0x02, 0x09,
	0x6c, 0x94, 0x86, 0x9b, 0xc5, 0x12, 0xfb, 0x76, 0xbf, 0x5d, 0xc3, 0x9e, 0x58, 0x99, 0x39, 0x1b,
	0x3d, 0x85, 0x86, 0xcb, 0x23, 0x11, 0xdc, 0xfc, 0xb8, 0xc8, 0xf3, 0x94, 0xcf, 0xd1, 0xa1, 0x9d,
	0xad, 0xbd, 0xdc, 0xf0, 0x4a, 0xb9, 0x8f, 0x32, 0x20, 0x75, 0x95, 0x74, 0x60, 0x9b, 0x0d, 0x0f,
	0x47, 0x93, 0x63, 0xf6, 0xe2, 0x78, 0x34, 0x4e, 0xdf, 0x9f, 0xa4, 0xaf, 0xd3, 0xf1, 0xdb, 0xb4,
	0x79, 0xad, 0xa6, 0x1c, 0x0d, 0xd3, 0x83, 0x51, 0x7a, 0xd8, 0x0c, 0x48, 0x04, 0xad, 0x73, 0xca,
	0x60, 0x9c, 0xbe, 0x1a, 0xb1, 0x37, 0xc3, 0x83, 0x66, 0xd8, 0xff, 0x1d, 0xc2, 0xdd, 0x73, 0x9f,
	0xc1, 0xe2, 0xab, 0x98, 0x22, 0xf9, 0x0c, 0x77, 0x2e, 0x74, 0x84, 0xec, 0x78, 0xc4, 0x17, 0xbb,
	0x17, 0x75, 0xff, 0x2f, 0x2e, 0x0f, 0x95, 0xde, 0xfb, 0xf1, 0xe7, 0xef, 0xcf, 0xb0, 0x4d, 0xb7,
	0x6c, 0xa7, 0xdd, 0x75, 0xed, 0x57, 0x85, 0x24, 0xbf, 0x02, 0xd8, 0xb9, 0xe4, 0x6e, 0xc8, 0x43,
	0x1b, 0xbe, 0xbe, 0x78, 0x51, 0x6f, 0xbd, 0xd1, 0x11, 0x3d, 0xb1, 0x44, 0xfd, 0x68, 0xcf, 0x27,
	0x4a, 0xbe, 0xad, 0x9a, 0xf4, 0x3d, 0x99, 0x7e, 0xe2, 0x79, 0x8e, 0x72, 0x86, 0xfb, 0x7e, 0x39,
	0x5f, 0xb6, 0xde, 0x6d, 0xc7, 0x71, 0x22, 0xe6, 0x3a, 0x4f, 0xfc, 0x07, 0xfc, 0x61, 0xc3, 0x3e,
	0xcf, 0xc7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x06, 0x38, 0x06, 0xd7, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/api.RegistrationService/RegisterPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationServiceClient) CompletePatientRegistration(ctx context.Context, in *CompletePatientRegistrationRequest, opts ...grpc.CallOption) (*CompletePatientRegistrationResponse, error) {
	out := new(CompletePatientRegistrationResponse)
	err := c.cc.Invoke(ctx, "/api.RegistrationService/CompletePatientRegistration", in, out, opts...)
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
		FullMethod: "/api.RegistrationService/RegisterPatient",
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
		FullMethod: "/api.RegistrationService/CompletePatientRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).CompletePatientRegistration(ctx, req.(*CompletePatientRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RegistrationService",
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
