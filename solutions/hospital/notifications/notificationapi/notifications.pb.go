// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.2
// source: notifications.proto

package notificationapi

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

type DeliveryStatus int32

const (
	DeliveryStatus_UNKNOWN   DeliveryStatus = 0
	DeliveryStatus_DELIVERED DeliveryStatus = 1
	DeliveryStatus_PENDING   DeliveryStatus = 2
	DeliveryStatus_FAILED    DeliveryStatus = 3
)

// Enum value maps for DeliveryStatus.
var (
	DeliveryStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "DELIVERED",
		2: "PENDING",
		3: "FAILED",
	}
	DeliveryStatus_value = map[string]int32{
		"UNKNOWN":   0,
		"DELIVERED": 1,
		"PENDING":   2,
		"FAILED":    3,
	}
)

func (x DeliveryStatus) Enum() *DeliveryStatus {
	p := new(DeliveryStatus)
	*p = x
	return p
}

func (x DeliveryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_notifications_proto_enumTypes[0].Descriptor()
}

func (DeliveryStatus) Type() protoreflect.EnumType {
	return &file_notifications_proto_enumTypes[0]
}

func (x DeliveryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryStatus.Descriptor instead.
func (DeliveryStatus) EnumDescriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{0}
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *EmailMessage `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetEmail() *EmailMessage {
	if x != nil {
		return x.Email
	}
	return nil
}

type EmailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientEmailAddress string `protobuf:"bytes,1,opt,name=recipientEmailAddress,proto3" json:"recipientEmailAddress,omitempty"`
	Subject               string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                  string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *EmailMessage) Reset() {
	*x = EmailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailMessage) ProtoMessage() {}

func (x *EmailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailMessage.ProtoReflect.Descriptor instead.
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *EmailMessage) GetRecipientEmailAddress() string {
	if x != nil {
		return x.RecipientEmailAddress
	}
	return ""
}

func (x *EmailMessage) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type SendSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sms *SmsMessage `protobuf:"bytes,1,opt,name=sms,proto3" json:"sms,omitempty"`
}

func (x *SendSmsRequest) Reset() {
	*x = SendSmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsRequest) ProtoMessage() {}

func (x *SendSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsRequest.ProtoReflect.Descriptor instead.
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *SendSmsRequest) GetSms() *SmsMessage {
	if x != nil {
		return x.Sms
	}
	return nil
}

type SendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status DeliveryStatus `protobuf:"varint,1,opt,name=status,proto3,enum=notificationapi.DeliveryStatus" json:"status,omitempty"`
}

func (x *SendReply) Reset() {
	*x = SendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReply) ProtoMessage() {}

func (x *SendReply) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReply.ProtoReflect.Descriptor instead.
func (*SendReply) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{3}
}

func (x *SendReply) GetStatus() DeliveryStatus {
	if x != nil {
		return x.Status
	}
	return DeliveryStatus_UNKNOWN
}

type SmsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientPhoneNumber string `protobuf:"bytes,1,opt,name=recipientPhoneNumber,proto3" json:"recipientPhoneNumber,omitempty"`
	Body                 string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SmsMessage) Reset() {
	*x = SmsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsMessage) ProtoMessage() {}

func (x *SmsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsMessage.ProtoReflect.Descriptor instead.
func (*SmsMessage) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{4}
}

func (x *SmsMessage) GetRecipientPhoneNumber() string {
	if x != nil {
		return x.RecipientPhoneNumber
	}
	return ""
}

func (x *SmsMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_notifications_proto protoreflect.FileDescriptor

var file_notifications_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x22, 0x47, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x72, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x34, 0x0a, 0x15, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x3f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x73, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6d, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x03, 0x73, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x6d,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x2a, 0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x32, 0xa6, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d,
	0x73, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x61, 0x72, 0x63, 0x47, 0x72, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_proto_rawDescOnce sync.Once
	file_notifications_proto_rawDescData = file_notifications_proto_rawDesc
)

func file_notifications_proto_rawDescGZIP() []byte {
	file_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_proto_rawDescData)
	})
	return file_notifications_proto_rawDescData
}

var file_notifications_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notifications_proto_goTypes = []interface{}{
	(DeliveryStatus)(0),      // 0: notificationapi.DeliveryStatus
	(*SendEmailRequest)(nil), // 1: notificationapi.SendEmailRequest
	(*EmailMessage)(nil),     // 2: notificationapi.EmailMessage
	(*SendSmsRequest)(nil),   // 3: notificationapi.SendSmsRequest
	(*SendReply)(nil),        // 4: notificationapi.SendReply
	(*SmsMessage)(nil),       // 5: notificationapi.SmsMessage
}
var file_notifications_proto_depIdxs = []int32{
	2, // 0: notificationapi.SendEmailRequest.email:type_name -> notificationapi.EmailMessage
	5, // 1: notificationapi.SendSmsRequest.sms:type_name -> notificationapi.SmsMessage
	0, // 2: notificationapi.SendReply.status:type_name -> notificationapi.DeliveryStatus
	1, // 3: notificationapi.Notification.SendEmail:input_type -> notificationapi.SendEmailRequest
	3, // 4: notificationapi.Notification.SendSms:input_type -> notificationapi.SendSmsRequest
	4, // 5: notificationapi.Notification.SendEmail:output_type -> notificationapi.SendReply
	4, // 6: notificationapi.Notification.SendSms:output_type -> notificationapi.SendReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_notifications_proto_init() }
func file_notifications_proto_init() {
	if File_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailMessage); i {
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
		file_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsRequest); i {
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
		file_notifications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReply); i {
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
		file_notifications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsMessage); i {
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
			RawDescriptor: file_notifications_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_proto_depIdxs,
		EnumInfos:         file_notifications_proto_enumTypes,
		MessageInfos:      file_notifications_proto_msgTypes,
	}.Build()
	File_notifications_proto = out.File
	file_notifications_proto_rawDesc = nil
	file_notifications_proto_goTypes = nil
	file_notifications_proto_depIdxs = nil
}
