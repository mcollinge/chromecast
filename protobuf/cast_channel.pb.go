// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cast_channel.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignatureAlgorithm int32

const (
	SignatureAlgorithm_UNSPECIFIED     SignatureAlgorithm = 0
	SignatureAlgorithm_RSASSA_PKCS1v15 SignatureAlgorithm = 1
	SignatureAlgorithm_RSASSA_PSS      SignatureAlgorithm = 2
)

var SignatureAlgorithm_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "RSASSA_PKCS1v15",
	2: "RSASSA_PSS",
}
var SignatureAlgorithm_value = map[string]int32{
	"UNSPECIFIED":     0,
	"RSASSA_PKCS1v15": 1,
	"RSASSA_PSS":      2,
}

func (x SignatureAlgorithm) Enum() *SignatureAlgorithm {
	p := new(SignatureAlgorithm)
	*p = x
	return p
}
func (x SignatureAlgorithm) String() string {
	return proto.EnumName(SignatureAlgorithm_name, int32(x))
}
func (x *SignatureAlgorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SignatureAlgorithm_value, data, "SignatureAlgorithm")
	if err != nil {
		return err
	}
	*x = SignatureAlgorithm(value)
	return nil
}
func (SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{0}
}

// Always pass a version of the protocol for future compatibility
// requirements.
type CastMessage_ProtocolVersion int32

const (
	CastMessage_CASTV2_1_0 CastMessage_ProtocolVersion = 0
)

var CastMessage_ProtocolVersion_name = map[int32]string{
	0: "CASTV2_1_0",
}
var CastMessage_ProtocolVersion_value = map[string]int32{
	"CASTV2_1_0": 0,
}

func (x CastMessage_ProtocolVersion) Enum() *CastMessage_ProtocolVersion {
	p := new(CastMessage_ProtocolVersion)
	*p = x
	return p
}
func (x CastMessage_ProtocolVersion) String() string {
	return proto.EnumName(CastMessage_ProtocolVersion_name, int32(x))
}
func (x *CastMessage_ProtocolVersion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CastMessage_ProtocolVersion_value, data, "CastMessage_ProtocolVersion")
	if err != nil {
		return err
	}
	*x = CastMessage_ProtocolVersion(value)
	return nil
}
func (CastMessage_ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{0, 0}
}

// Encoding and payload info follows.
// What type of data do we have in this message.
type CastMessage_PayloadType int32

const (
	CastMessage_STRING CastMessage_PayloadType = 0
	CastMessage_BINARY CastMessage_PayloadType = 1
)

var CastMessage_PayloadType_name = map[int32]string{
	0: "STRING",
	1: "BINARY",
}
var CastMessage_PayloadType_value = map[string]int32{
	"STRING": 0,
	"BINARY": 1,
}

func (x CastMessage_PayloadType) Enum() *CastMessage_PayloadType {
	p := new(CastMessage_PayloadType)
	*p = x
	return p
}
func (x CastMessage_PayloadType) String() string {
	return proto.EnumName(CastMessage_PayloadType_name, int32(x))
}
func (x *CastMessage_PayloadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CastMessage_PayloadType_value, data, "CastMessage_PayloadType")
	if err != nil {
		return err
	}
	*x = CastMessage_PayloadType(value)
	return nil
}
func (CastMessage_PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{0, 1}
}

type AuthError_ErrorType int32

const (
	AuthError_INTERNAL_ERROR                  AuthError_ErrorType = 0
	AuthError_NO_TLS                          AuthError_ErrorType = 1
	AuthError_SIGNATURE_ALGORITHM_UNAVAILABLE AuthError_ErrorType = 2
)

var AuthError_ErrorType_name = map[int32]string{
	0: "INTERNAL_ERROR",
	1: "NO_TLS",
	2: "SIGNATURE_ALGORITHM_UNAVAILABLE",
}
var AuthError_ErrorType_value = map[string]int32{
	"INTERNAL_ERROR": 0,
	"NO_TLS":         1,
	"SIGNATURE_ALGORITHM_UNAVAILABLE": 2,
}

func (x AuthError_ErrorType) Enum() *AuthError_ErrorType {
	p := new(AuthError_ErrorType)
	*p = x
	return p
}
func (x AuthError_ErrorType) String() string {
	return proto.EnumName(AuthError_ErrorType_name, int32(x))
}
func (x *AuthError_ErrorType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthError_ErrorType_value, data, "AuthError_ErrorType")
	if err != nil {
		return err
	}
	*x = AuthError_ErrorType(value)
	return nil
}
func (AuthError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{3, 0}
}

type CastMessage struct {
	ProtocolVersion *CastMessage_ProtocolVersion `protobuf:"varint,1,req,name=protocol_version,json=protocolVersion,enum=protobuf.CastMessage_ProtocolVersion" json:"protocol_version,omitempty"`
	// source and destination ids identify the origin and destination of the
	// message.  They are used to route messages between endpoints that share a
	// device-to-device channel.
	//
	// For messages between applications:
	//   - The sender application id is a unique identifier generated on behalf of
	//     the sender application.
	//   - The receiver id is always the the session id for the application.
	//
	// For messages to or from the sender or receiver platform, the special ids
	// 'sender-0' and 'receiver-0' can be used.
	//
	// For messages intended for all endpoints using a given channel, the
	// wildcard destination_id '*' can be used.
	SourceId      *string `protobuf:"bytes,2,req,name=source_id,json=sourceId" json:"source_id,omitempty"`
	DestinationId *string `protobuf:"bytes,3,req,name=destination_id,json=destinationId" json:"destination_id,omitempty"`
	// This is the core multiplexing key.  All messages are sent on a namespace
	// and endpoints sharing a channel listen on one or more namespaces.  The
	// namespace defines the protocol and semantics of the message.
	Namespace   *string                  `protobuf:"bytes,4,req,name=namespace" json:"namespace,omitempty"`
	PayloadType *CastMessage_PayloadType `protobuf:"varint,5,req,name=payload_type,json=payloadType,enum=protobuf.CastMessage_PayloadType" json:"payload_type,omitempty"`
	// Depending on payload_type, exactly one of the following optional fields
	// will always be set.
	PayloadUtf8          *string  `protobuf:"bytes,6,opt,name=payload_utf8,json=payloadUtf8" json:"payload_utf8,omitempty"`
	PayloadBinary        []byte   `protobuf:"bytes,7,opt,name=payload_binary,json=payloadBinary" json:"payload_binary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CastMessage) Reset()         { *m = CastMessage{} }
func (m *CastMessage) String() string { return proto.CompactTextString(m) }
func (*CastMessage) ProtoMessage()    {}
func (*CastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{0}
}
func (m *CastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CastMessage.Unmarshal(m, b)
}
func (m *CastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CastMessage.Marshal(b, m, deterministic)
}
func (dst *CastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CastMessage.Merge(dst, src)
}
func (m *CastMessage) XXX_Size() int {
	return xxx_messageInfo_CastMessage.Size(m)
}
func (m *CastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CastMessage proto.InternalMessageInfo

func (m *CastMessage) GetProtocolVersion() CastMessage_ProtocolVersion {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return CastMessage_CASTV2_1_0
}

func (m *CastMessage) GetSourceId() string {
	if m != nil && m.SourceId != nil {
		return *m.SourceId
	}
	return ""
}

func (m *CastMessage) GetDestinationId() string {
	if m != nil && m.DestinationId != nil {
		return *m.DestinationId
	}
	return ""
}

func (m *CastMessage) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *CastMessage) GetPayloadType() CastMessage_PayloadType {
	if m != nil && m.PayloadType != nil {
		return *m.PayloadType
	}
	return CastMessage_STRING
}

func (m *CastMessage) GetPayloadUtf8() string {
	if m != nil && m.PayloadUtf8 != nil {
		return *m.PayloadUtf8
	}
	return ""
}

func (m *CastMessage) GetPayloadBinary() []byte {
	if m != nil {
		return m.PayloadBinary
	}
	return nil
}

// Messages for authentication protocol between a sender and a receiver.
type AuthChallenge struct {
	SignatureAlgorithm   *SignatureAlgorithm `protobuf:"varint,1,opt,name=signature_algorithm,json=signatureAlgorithm,enum=protobuf.SignatureAlgorithm,def=1" json:"signature_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AuthChallenge) Reset()         { *m = AuthChallenge{} }
func (m *AuthChallenge) String() string { return proto.CompactTextString(m) }
func (*AuthChallenge) ProtoMessage()    {}
func (*AuthChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{1}
}
func (m *AuthChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthChallenge.Unmarshal(m, b)
}
func (m *AuthChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthChallenge.Marshal(b, m, deterministic)
}
func (dst *AuthChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthChallenge.Merge(dst, src)
}
func (m *AuthChallenge) XXX_Size() int {
	return xxx_messageInfo_AuthChallenge.Size(m)
}
func (m *AuthChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_AuthChallenge proto.InternalMessageInfo

const Default_AuthChallenge_SignatureAlgorithm SignatureAlgorithm = SignatureAlgorithm_RSASSA_PKCS1v15

func (m *AuthChallenge) GetSignatureAlgorithm() SignatureAlgorithm {
	if m != nil && m.SignatureAlgorithm != nil {
		return *m.SignatureAlgorithm
	}
	return Default_AuthChallenge_SignatureAlgorithm
}

type AuthResponse struct {
	Signature               []byte              `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"`
	ClientAuthCertificate   []byte              `protobuf:"bytes,2,req,name=client_auth_certificate,json=clientAuthCertificate" json:"client_auth_certificate,omitempty"`
	IntermediateCertificate [][]byte            `protobuf:"bytes,3,rep,name=intermediate_certificate,json=intermediateCertificate" json:"intermediate_certificate,omitempty"`
	SignatureAlgorithm      *SignatureAlgorithm `protobuf:"varint,4,opt,name=signature_algorithm,json=signatureAlgorithm,enum=protobuf.SignatureAlgorithm,def=1" json:"signature_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}            `json:"-"`
	XXX_unrecognized        []byte              `json:"-"`
	XXX_sizecache           int32               `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{2}
}
func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (dst *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(dst, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

const Default_AuthResponse_SignatureAlgorithm SignatureAlgorithm = SignatureAlgorithm_RSASSA_PKCS1v15

func (m *AuthResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AuthResponse) GetClientAuthCertificate() []byte {
	if m != nil {
		return m.ClientAuthCertificate
	}
	return nil
}

func (m *AuthResponse) GetIntermediateCertificate() [][]byte {
	if m != nil {
		return m.IntermediateCertificate
	}
	return nil
}

func (m *AuthResponse) GetSignatureAlgorithm() SignatureAlgorithm {
	if m != nil && m.SignatureAlgorithm != nil {
		return *m.SignatureAlgorithm
	}
	return Default_AuthResponse_SignatureAlgorithm
}

type AuthError struct {
	ErrorType            *AuthError_ErrorType `protobuf:"varint,1,req,name=error_type,json=errorType,enum=protobuf.AuthError_ErrorType" json:"error_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthError) Reset()         { *m = AuthError{} }
func (m *AuthError) String() string { return proto.CompactTextString(m) }
func (*AuthError) ProtoMessage()    {}
func (*AuthError) Descriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{3}
}
func (m *AuthError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthError.Unmarshal(m, b)
}
func (m *AuthError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthError.Marshal(b, m, deterministic)
}
func (dst *AuthError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthError.Merge(dst, src)
}
func (m *AuthError) XXX_Size() int {
	return xxx_messageInfo_AuthError.Size(m)
}
func (m *AuthError) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthError.DiscardUnknown(m)
}

var xxx_messageInfo_AuthError proto.InternalMessageInfo

func (m *AuthError) GetErrorType() AuthError_ErrorType {
	if m != nil && m.ErrorType != nil {
		return *m.ErrorType
	}
	return AuthError_INTERNAL_ERROR
}

type DeviceAuthMessage struct {
	// Request fields
	Challenge *AuthChallenge `protobuf:"bytes,1,opt,name=challenge" json:"challenge,omitempty"`
	// Response fields
	Response             *AuthResponse `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Error                *AuthError    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeviceAuthMessage) Reset()         { *m = DeviceAuthMessage{} }
func (m *DeviceAuthMessage) String() string { return proto.CompactTextString(m) }
func (*DeviceAuthMessage) ProtoMessage()    {}
func (*DeviceAuthMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cast_channel_8bed9ab37bcc9723, []int{4}
}
func (m *DeviceAuthMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAuthMessage.Unmarshal(m, b)
}
func (m *DeviceAuthMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAuthMessage.Marshal(b, m, deterministic)
}
func (dst *DeviceAuthMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAuthMessage.Merge(dst, src)
}
func (m *DeviceAuthMessage) XXX_Size() int {
	return xxx_messageInfo_DeviceAuthMessage.Size(m)
}
func (m *DeviceAuthMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAuthMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAuthMessage proto.InternalMessageInfo

func (m *DeviceAuthMessage) GetChallenge() *AuthChallenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *DeviceAuthMessage) GetResponse() *AuthResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DeviceAuthMessage) GetError() *AuthError {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*CastMessage)(nil), "protobuf.CastMessage")
	proto.RegisterType((*AuthChallenge)(nil), "protobuf.AuthChallenge")
	proto.RegisterType((*AuthResponse)(nil), "protobuf.AuthResponse")
	proto.RegisterType((*AuthError)(nil), "protobuf.AuthError")
	proto.RegisterType((*DeviceAuthMessage)(nil), "protobuf.DeviceAuthMessage")
	proto.RegisterEnum("protobuf.SignatureAlgorithm", SignatureAlgorithm_name, SignatureAlgorithm_value)
	proto.RegisterEnum("protobuf.CastMessage_ProtocolVersion", CastMessage_ProtocolVersion_name, CastMessage_ProtocolVersion_value)
	proto.RegisterEnum("protobuf.CastMessage_PayloadType", CastMessage_PayloadType_name, CastMessage_PayloadType_value)
	proto.RegisterEnum("protobuf.AuthError_ErrorType", AuthError_ErrorType_name, AuthError_ErrorType_value)
}

func init() { proto.RegisterFile("cast_channel.proto", fileDescriptor_cast_channel_8bed9ab37bcc9723) }

var fileDescriptor_cast_channel_8bed9ab37bcc9723 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x4e, 0xdb, 0x3c,
	0x14, 0xc7, 0x49, 0x0b, 0x7c, 0xe4, 0xb4, 0xb4, 0xf9, 0x8c, 0x36, 0x22, 0x8d, 0x69, 0x25, 0x13,
	0x52, 0xb7, 0x8b, 0x6a, 0x54, 0x62, 0x62, 0xd3, 0x6e, 0xd2, 0x92, 0x41, 0xb6, 0x12, 0x2a, 0x27,
	0x45, 0xda, 0xcd, 0x2c, 0x93, 0xba, 0xad, 0xa5, 0x90, 0x44, 0x89, 0x8b, 0xd4, 0x67, 0xe1, 0x09,
	0xf6, 0x94, 0x9b, 0xe2, 0x90, 0xa6, 0x85, 0x71, 0xb7, 0x9b, 0xc4, 0xf9, 0x9f, 0xdf, 0x89, 0x8f,
	0xcf, 0xf9, 0x1b, 0x90, 0x4f, 0x53, 0x41, 0xfc, 0x19, 0x0d, 0x43, 0x16, 0x74, 0xe2, 0x24, 0x12,
	0x11, 0xda, 0x91, 0xaf, 0x9b, 0xf9, 0xc4, 0xb8, 0xaf, 0x42, 0xad, 0x4f, 0x53, 0x71, 0xc9, 0xd2,
	0x94, 0x4e, 0x19, 0x1a, 0x82, 0x26, 0x63, 0x7e, 0x14, 0x90, 0x3b, 0x96, 0xa4, 0x3c, 0x0a, 0x75,
	0xa5, 0x55, 0x69, 0x37, 0xba, 0x47, 0x9d, 0x22, 0xa9, 0xb3, 0x92, 0xd0, 0x19, 0x3e, 0xd0, 0xd7,
	0x39, 0x8c, 0x9b, 0xf1, 0xba, 0x80, 0x5e, 0x81, 0x9a, 0x46, 0xf3, 0xc4, 0x67, 0x84, 0x8f, 0xf5,
	0x4a, 0xab, 0xd2, 0x56, 0xf1, 0x4e, 0x2e, 0xd8, 0x63, 0x74, 0x04, 0x8d, 0x31, 0x4b, 0x05, 0x0f,
	0xa9, 0xe0, 0x51, 0x98, 0x11, 0x55, 0x49, 0xec, 0xae, 0xa8, 0xf6, 0x18, 0x1d, 0x80, 0x1a, 0xd2,
	0x5b, 0x96, 0xc6, 0xd4, 0x67, 0xfa, 0xa6, 0x24, 0x4a, 0x01, 0x9d, 0x41, 0x3d, 0xa6, 0x8b, 0x20,
	0xa2, 0x63, 0x22, 0x16, 0x31, 0xd3, 0xb7, 0x64, 0xbd, 0x87, 0xcf, 0xd4, 0x9b, 0x93, 0xde, 0x22,
	0x66, 0xb8, 0x16, 0x97, 0x1f, 0xe8, 0xb0, 0xfc, 0xcb, 0x5c, 0x4c, 0x4e, 0xf5, 0xed, 0x96, 0xd2,
	0x56, 0x97, 0xc8, 0x48, 0x4c, 0x4e, 0xb3, 0x6a, 0x0b, 0xe4, 0x86, 0x87, 0x34, 0x59, 0xe8, 0xff,
	0xb5, 0x94, 0x76, 0x1d, 0xef, 0x3e, 0xa8, 0x3d, 0x29, 0x1a, 0x87, 0xd0, 0x7c, 0xd4, 0x15, 0xd4,
	0x00, 0xe8, 0x9b, 0xae, 0x77, 0xdd, 0x25, 0xc7, 0xe4, 0x83, 0xb6, 0x61, 0x1c, 0x41, 0x6d, 0xa5,
	0x10, 0x04, 0xb0, 0xed, 0x7a, 0xd8, 0x76, 0xce, 0xb5, 0x8d, 0x6c, 0xdd, 0xb3, 0x1d, 0x13, 0xff,
	0xd0, 0x14, 0x23, 0x82, 0x5d, 0x73, 0x2e, 0x66, 0xfd, 0x19, 0x0d, 0x02, 0x16, 0x4e, 0x19, 0xfa,
	0x09, 0x7b, 0x29, 0x9f, 0x86, 0x54, 0xcc, 0x13, 0x46, 0x68, 0x30, 0x8d, 0x12, 0x2e, 0x66, 0xb7,
	0xba, 0xd2, 0x52, 0xda, 0x8d, 0xee, 0x41, 0x79, 0x62, 0xb7, 0x80, 0xcc, 0x82, 0xf9, 0xdc, 0xc4,
	0xae, 0xe9, 0xba, 0x26, 0x19, 0x7e, 0xef, 0xbb, 0xc7, 0x77, 0xc7, 0x27, 0x18, 0xa5, 0x4f, 0x20,
	0xe3, 0xb7, 0x02, 0xf5, 0x6c, 0x47, 0xcc, 0xd2, 0x38, 0x0a, 0x53, 0x96, 0x75, 0x7e, 0x89, 0x49,
	0x23, 0xd4, 0x71, 0x29, 0xa0, 0x8f, 0xb0, 0xef, 0x07, 0x9c, 0x85, 0x82, 0xd0, 0xb9, 0x98, 0x11,
	0x9f, 0x25, 0x82, 0x4f, 0xb8, 0x4f, 0x05, 0x93, 0x93, 0xae, 0xe3, 0x17, 0x79, 0x58, 0x1e, 0xa2,
	0x0c, 0xa2, 0x4f, 0xa0, 0xf3, 0x50, 0xb0, 0xe4, 0x96, 0x8d, 0x39, 0x15, 0x6c, 0x2d, 0xb1, 0xda,
	0xaa, 0xb6, 0xeb, 0x78, 0x7f, 0x35, 0xbe, 0x9a, 0xfa, 0x4c, 0x07, 0x36, 0xff, 0x55, 0x07, 0xee,
	0x15, 0x50, 0xb3, 0x72, 0xad, 0x24, 0x89, 0x12, 0xf4, 0x05, 0x80, 0x65, 0x8b, 0xdc, 0x58, 0xf9,
	0x45, 0x78, 0x5d, 0x6e, 0xb2, 0x04, 0x3b, 0xf2, 0x29, 0x4d, 0xa5, 0xb2, 0x62, 0x69, 0x0c, 0x41,
	0x5d, 0xea, 0x08, 0x41, 0xc3, 0x76, 0x3c, 0x0b, 0x3b, 0xe6, 0x80, 0x58, 0x18, 0x5f, 0xe1, 0x7c,
	0xd6, 0xce, 0x15, 0xf1, 0x06, 0xae, 0xa6, 0xa0, 0xb7, 0xf0, 0xc6, 0xb5, 0xcf, 0x1d, 0xd3, 0x1b,
	0x61, 0x8b, 0x98, 0x83, 0xf3, 0x2b, 0x6c, 0x7b, 0x17, 0x97, 0x64, 0xe4, 0x98, 0xd7, 0xa6, 0x3d,
	0x30, 0x7b, 0x03, 0x4b, 0xab, 0x18, 0xbf, 0x14, 0xf8, 0xff, 0x8c, 0xdd, 0x71, 0x9f, 0x65, 0x5b,
	0x17, 0x97, 0xf6, 0x04, 0x54, 0xbf, 0xb0, 0x88, 0xf4, 0x42, 0xad, 0xbb, 0xbf, 0x5e, 0xe4, 0xd2,
	0x41, 0xb8, 0x24, 0x51, 0x17, 0x76, 0x92, 0x87, 0x39, 0xeb, 0x15, 0x99, 0xf5, 0x72, 0x3d, 0xab,
	0x70, 0x01, 0x5e, 0x72, 0xe8, 0x1d, 0x6c, 0xc9, 0xf3, 0xe9, 0x55, 0x99, 0xb0, 0xf7, 0x97, 0x5e,
	0xe0, 0x9c, 0x78, 0xff, 0x0d, 0xd0, 0xd3, 0x21, 0xa0, 0x26, 0xd4, 0x46, 0x8e, 0x3b, 0xb4, 0xfa,
	0xf6, 0x57, 0xdb, 0x3a, 0xd3, 0x36, 0xd0, 0x1e, 0x3c, 0x9e, 0x8b, 0xa6, 0x64, 0xf7, 0xa5, 0x10,
	0x5d, 0x57, 0xab, 0xf4, 0x2a, 0x17, 0xd5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0xde, 0x30,
	0x7c, 0xc9, 0x04, 0x00, 0x00,
}