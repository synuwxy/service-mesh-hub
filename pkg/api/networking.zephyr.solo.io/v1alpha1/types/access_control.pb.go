// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/networking/v1alpha1/access_control.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//access control policies apply ALLOW policies to communication in a mesh
//access control policies specify the following:
//ALLOW those requests:
//- originating from from **source pods**
//- sent to **destination pods**
//- matching the indicated request criteria (allowed_paths, allowed_methods, allowed_ports)
//if no access control policies are present, all traffic in the mesh will be set to ALLOW
type AccessControlPolicySpec struct {
	//
	//requests originating from these pods will have the rule applied
	//leave empty to have all pods in the mesh apply these policies
	//
	//note that access control policies are mapped to source pods by their
	//service account. if other pods share the same service account,
	//this access control rule will apply to those pods as well.
	//
	//for fine-grained access control policies, ensure that your
	//service accounts properly reflect the desired
	//boundary for your access control policies
	SourceSelector *types.IdentitySelector `protobuf:"bytes,2,opt,name=source_selector,json=sourceSelector,proto3" json:"source_selector,omitempty"`
	//
	//requests destined for these pods will have the rule applied
	//leave empty to apply to all destination pods in the mesh
	DestinationSelector *types.Selector `protobuf:"bytes,3,opt,name=destination_selector,json=destinationSelector,proto3" json:"destination_selector,omitempty"`
	//
	//Optional. A list of HTTP paths or gRPC methods to allow.
	//gRPC methods must be presented as fully-qualified name in the form of
	//"/packageName.serviceName/methodName" and are case sensitive.
	//Exact match, prefix match, and suffix match are supported for paths.
	//For example, the path "/books/review" matches
	//"/books/review" (exact match), "*books/" (suffix match), or "/books*" (prefix match),
	//
	//If not specified, it allows to any path.
	AllowedPaths []string `protobuf:"bytes,4,rep,name=allowed_paths,json=allowedPaths,proto3" json:"allowed_paths,omitempty"`
	//
	//Optional. A list of HTTP methods to allow (e.g., "GET", "POST").
	//It is ignored in gRPC case because the value is always "POST".
	//If not specified, allows any method.
	AllowedMethods []types.HttpMethodValue `protobuf:"varint,5,rep,packed,name=allowed_methods,json=allowedMethods,proto3,enum=core.zephyr.solo.io.HttpMethodValue" json:"allowed_methods,omitempty"`
	//
	//Optional. A list of ports which to allow
	//if not set any port is allowed
	AllowedPorts         []uint32 `protobuf:"varint,6,rep,packed,name=allowed_ports,json=allowedPorts,proto3" json:"allowed_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessControlPolicySpec) Reset()         { *m = AccessControlPolicySpec{} }
func (m *AccessControlPolicySpec) String() string { return proto.CompactTextString(m) }
func (*AccessControlPolicySpec) ProtoMessage()    {}
func (*AccessControlPolicySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace7a2fd75022f8b, []int{0}
}
func (m *AccessControlPolicySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlPolicySpec.Unmarshal(m, b)
}
func (m *AccessControlPolicySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlPolicySpec.Marshal(b, m, deterministic)
}
func (m *AccessControlPolicySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlPolicySpec.Merge(m, src)
}
func (m *AccessControlPolicySpec) XXX_Size() int {
	return xxx_messageInfo_AccessControlPolicySpec.Size(m)
}
func (m *AccessControlPolicySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlPolicySpec.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlPolicySpec proto.InternalMessageInfo

func (m *AccessControlPolicySpec) GetSourceSelector() *types.IdentitySelector {
	if m != nil {
		return m.SourceSelector
	}
	return nil
}

func (m *AccessControlPolicySpec) GetDestinationSelector() *types.Selector {
	if m != nil {
		return m.DestinationSelector
	}
	return nil
}

func (m *AccessControlPolicySpec) GetAllowedPaths() []string {
	if m != nil {
		return m.AllowedPaths
	}
	return nil
}

func (m *AccessControlPolicySpec) GetAllowedMethods() []types.HttpMethodValue {
	if m != nil {
		return m.AllowedMethods
	}
	return nil
}

func (m *AccessControlPolicySpec) GetAllowedPorts() []uint32 {
	if m != nil {
		return m.AllowedPorts
	}
	return nil
}

type AccessControlPolicyStatus struct {
	TranslationStatus    *types.ComputedStatus                        `protobuf:"bytes,1,opt,name=translation_status,json=translationStatus,proto3" json:"translation_status,omitempty"`
	TranslatorErrors     []*AccessControlPolicyStatus_TranslatorError `protobuf:"bytes,2,rep,name=translator_errors,json=translatorErrors,proto3" json:"translator_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *AccessControlPolicyStatus) Reset()         { *m = AccessControlPolicyStatus{} }
func (m *AccessControlPolicyStatus) String() string { return proto.CompactTextString(m) }
func (*AccessControlPolicyStatus) ProtoMessage()    {}
func (*AccessControlPolicyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace7a2fd75022f8b, []int{1}
}
func (m *AccessControlPolicyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlPolicyStatus.Unmarshal(m, b)
}
func (m *AccessControlPolicyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlPolicyStatus.Marshal(b, m, deterministic)
}
func (m *AccessControlPolicyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlPolicyStatus.Merge(m, src)
}
func (m *AccessControlPolicyStatus) XXX_Size() int {
	return xxx_messageInfo_AccessControlPolicyStatus.Size(m)
}
func (m *AccessControlPolicyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlPolicyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlPolicyStatus proto.InternalMessageInfo

func (m *AccessControlPolicyStatus) GetTranslationStatus() *types.ComputedStatus {
	if m != nil {
		return m.TranslationStatus
	}
	return nil
}

func (m *AccessControlPolicyStatus) GetTranslatorErrors() []*AccessControlPolicyStatus_TranslatorError {
	if m != nil {
		return m.TranslatorErrors
	}
	return nil
}

// TODO use a shared Status message with TrafficPolicy once autopilot allows for it
type AccessControlPolicyStatus_TranslatorError struct {
	// ID representing a translator that translates TrafficPolicy to Mesh-specific config
	TranslatorId         string   `protobuf:"bytes,1,opt,name=translator_id,json=translatorId,proto3" json:"translator_id,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessControlPolicyStatus_TranslatorError) Reset() {
	*m = AccessControlPolicyStatus_TranslatorError{}
}
func (m *AccessControlPolicyStatus_TranslatorError) String() string {
	return proto.CompactTextString(m)
}
func (*AccessControlPolicyStatus_TranslatorError) ProtoMessage() {}
func (*AccessControlPolicyStatus_TranslatorError) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace7a2fd75022f8b, []int{1, 0}
}
func (m *AccessControlPolicyStatus_TranslatorError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlPolicyStatus_TranslatorError.Unmarshal(m, b)
}
func (m *AccessControlPolicyStatus_TranslatorError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlPolicyStatus_TranslatorError.Marshal(b, m, deterministic)
}
func (m *AccessControlPolicyStatus_TranslatorError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlPolicyStatus_TranslatorError.Merge(m, src)
}
func (m *AccessControlPolicyStatus_TranslatorError) XXX_Size() int {
	return xxx_messageInfo_AccessControlPolicyStatus_TranslatorError.Size(m)
}
func (m *AccessControlPolicyStatus_TranslatorError) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlPolicyStatus_TranslatorError.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlPolicyStatus_TranslatorError proto.InternalMessageInfo

func (m *AccessControlPolicyStatus_TranslatorError) GetTranslatorId() string {
	if m != nil {
		return m.TranslatorId
	}
	return ""
}

func (m *AccessControlPolicyStatus_TranslatorError) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*AccessControlPolicySpec)(nil), "networking.zephyr.solo.io.AccessControlPolicySpec")
	proto.RegisterType((*AccessControlPolicyStatus)(nil), "networking.zephyr.solo.io.AccessControlPolicyStatus")
	proto.RegisterType((*AccessControlPolicyStatus_TranslatorError)(nil), "networking.zephyr.solo.io.AccessControlPolicyStatus.TranslatorError")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/networking/v1alpha1/access_control.proto", fileDescriptor_ace7a2fd75022f8b)
}

var fileDescriptor_ace7a2fd75022f8b = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd5, 0x04, 0x2a, 0x65, 0x9b, 0x26, 0x60, 0x90, 0x70, 0x23, 0x21, 0x45, 0x2d, 0x48,
	0x16, 0xa2, 0x6b, 0x35, 0x9c, 0x00, 0x0a, 0x12, 0x7d, 0x08, 0x0a, 0x0e, 0xe2, 0x01, 0x1e, 0xac,
	0xed, 0x7a, 0x14, 0x2f, 0xb5, 0x3d, 0xcb, 0xee, 0x98, 0x2a, 0x5c, 0x8e, 0x4b, 0x70, 0x20, 0xe4,
	0xb5, 0x13, 0xa7, 0x21, 0x55, 0x1e, 0xf7, 0xd7, 0x3f, 0xdf, 0xfc, 0x33, 0x9a, 0x65, 0xd3, 0x85,
	0xa2, 0xb4, 0xbc, 0xe6, 0x12, 0xf3, 0xd0, 0x62, 0x86, 0xe7, 0x0a, 0xc3, 0x1c, 0x6c, 0x7a, 0xae,
	0x0d, 0xfe, 0x00, 0x49, 0x36, 0x14, 0x5a, 0x85, 0x05, 0xd0, 0x2d, 0x9a, 0x1b, 0x55, 0x2c, 0xc2,
	0x5f, 0x17, 0x22, 0xd3, 0xa9, 0xb8, 0x08, 0x85, 0x94, 0x60, 0x6d, 0x2c, 0xb1, 0x20, 0x83, 0x19,
	0xd7, 0x06, 0x09, 0xbd, 0x93, 0xd6, 0xc9, 0x7f, 0x83, 0x4e, 0x97, 0x86, 0x57, 0x54, 0xae, 0x70,
	0xf4, 0xea, 0x7f, 0xac, 0x44, 0x03, 0x2d, 0x30, 0x25, 0xd2, 0x35, 0x66, 0xc4, 0xf7, 0x79, 0x2d,
	0x64, 0x20, 0x09, 0x4d, 0xe3, 0x7f, 0xbd, 0xd7, 0x4f, 0x82, 0x4a, 0x5b, 0xbb, 0x4f, 0xff, 0x76,
	0xd8, 0xb3, 0xb7, 0x2e, 0xfd, 0x65, 0x1d, 0x7e, 0x86, 0x99, 0x92, 0xcb, 0xb9, 0x06, 0xe9, 0x7d,
	0x62, 0x43, 0x8b, 0xa5, 0x91, 0x10, 0xaf, 0x5a, 0xf8, 0x9d, 0xf1, 0x41, 0x70, 0x34, 0x79, 0xc9,
	0x2b, 0xe2, 0xd6, 0x50, 0xfc, 0x2a, 0x81, 0x82, 0x14, 0x2d, 0xe7, 0x8d, 0x39, 0x1a, 0xd4, 0xd5,
	0xab, 0xb7, 0x37, 0x63, 0x4f, 0x13, 0xb0, 0xa4, 0x0a, 0x41, 0x0a, 0x8b, 0x16, 0xda, 0x75, 0xd0,
	0xe7, 0x3b, 0xa1, 0x6b, 0xd8, 0x93, 0x8d, 0xd2, 0x35, 0xf1, 0x8c, 0x1d, 0x8b, 0x2c, 0xc3, 0x5b,
	0x48, 0x62, 0x2d, 0x28, 0xb5, 0xfe, 0x83, 0x71, 0x37, 0xe8, 0x45, 0xfd, 0x46, 0x9c, 0x55, 0x9a,
	0x37, 0x65, 0xc3, 0x95, 0x29, 0x07, 0x4a, 0x31, 0xb1, 0xfe, 0xc3, 0x71, 0x37, 0x18, 0x4c, 0x5e,
	0xec, 0xec, 0xf8, 0x91, 0x48, 0x4f, 0x9d, 0xef, 0xab, 0xc8, 0x4a, 0x88, 0x06, 0x4d, 0x71, 0xad,
	0xd9, 0x3b, 0x3d, 0xd1, 0x90, 0xf5, 0x0f, 0xc7, 0xdd, 0xe0, 0xb8, 0xed, 0x59, 0x69, 0xa7, 0x7f,
	0x3a, 0xec, 0x64, 0xd7, 0x5a, 0xdd, 0xea, 0xbd, 0x88, 0x79, 0x64, 0x44, 0x61, 0xb3, 0x66, 0x11,
	0x4e, 0xf5, 0x0f, 0xdc, 0x1a, 0xce, 0x76, 0x86, 0xba, 0xc4, 0x5c, 0x97, 0x04, 0x49, 0x0d, 0x88,
	0x1e, 0x6f, 0x94, 0x37, 0xcc, 0x9f, 0x6c, 0x2d, 0xa2, 0x89, 0xc1, 0x18, 0x34, 0xd6, 0xef, 0x8c,
	0xbb, 0xc1, 0xd1, 0xe4, 0x3d, 0xbf, 0xf7, 0x12, 0xf9, 0xbd, 0x21, 0xf9, 0x97, 0x35, 0xed, 0x43,
	0x05, 0x8b, 0x1e, 0xd1, 0x5d, 0xc1, 0x8e, 0xbe, 0xb3, 0xe1, 0x96, 0xa9, 0x5a, 0xce, 0x46, 0x0a,
	0x95, 0xb8, 0xa1, 0x7a, 0x51, 0xbf, 0x15, 0xaf, 0x92, 0xca, 0xe4, 0xf2, 0xc5, 0x39, 0x58, 0x2b,
	0x16, 0xe0, 0xae, 0xaa, 0x17, 0xf5, 0x9d, 0x38, 0xad, 0xb5, 0x77, 0xf3, 0x6f, 0x9f, 0xf7, 0x7e,
	0x47, 0x7d, 0xb3, 0xd8, 0xfa, 0x92, 0x5b, 0xe3, 0xb5, 0x47, 0x4f, 0x4b, 0x0d, 0xf6, 0xfa, 0xd0,
	0x1d, 0xfd, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x05, 0x43, 0x56, 0xe2, 0xea, 0x03, 0x00,
	0x00,
}
