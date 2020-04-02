// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow_template.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateWorkflowTemplateRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateWorkflowTemplateRequest) Reset()         { *m = CreateWorkflowTemplateRequest{} }
func (m *CreateWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowTemplateRequest) ProtoMessage()    {}
func (*CreateWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{0}
}

func (m *CreateWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowTemplateRequest.Merge(m, src)
}
func (m *CreateWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Size(m)
}
func (m *CreateWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowTemplateRequest proto.InternalMessageInfo

func (m *CreateWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateWorkflowTemplateRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type UpdateWorkflowTemplateVersionRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateWorkflowTemplateVersionRequest) Reset()         { *m = UpdateWorkflowTemplateVersionRequest{} }
func (m *UpdateWorkflowTemplateVersionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateWorkflowTemplateVersionRequest) ProtoMessage()    {}
func (*UpdateWorkflowTemplateVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{1}
}

func (m *UpdateWorkflowTemplateVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateWorkflowTemplateVersionRequest.Unmarshal(m, b)
}
func (m *UpdateWorkflowTemplateVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateWorkflowTemplateVersionRequest.Marshal(b, m, deterministic)
}
func (m *UpdateWorkflowTemplateVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateWorkflowTemplateVersionRequest.Merge(m, src)
}
func (m *UpdateWorkflowTemplateVersionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateWorkflowTemplateVersionRequest.Size(m)
}
func (m *UpdateWorkflowTemplateVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateWorkflowTemplateVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateWorkflowTemplateVersionRequest proto.InternalMessageInfo

func (m *UpdateWorkflowTemplateVersionRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *UpdateWorkflowTemplateVersionRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type GetWorkflowTemplateRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Version              int32    `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowTemplateRequest) Reset()         { *m = GetWorkflowTemplateRequest{} }
func (m *GetWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowTemplateRequest) ProtoMessage()    {}
func (*GetWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{2}
}

func (m *GetWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *GetWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowTemplateRequest.Merge(m, src)
}
func (m *GetWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Size(m)
}
func (m *GetWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowTemplateRequest proto.InternalMessageInfo

func (m *GetWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetWorkflowTemplateRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetWorkflowTemplateRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ListWorkflowTemplateVersionsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowTemplateVersionsRequest) Reset()         { *m = ListWorkflowTemplateVersionsRequest{} }
func (m *ListWorkflowTemplateVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplateVersionsRequest) ProtoMessage()    {}
func (*ListWorkflowTemplateVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{3}
}

func (m *ListWorkflowTemplateVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Unmarshal(m, b)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Merge(m, src)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplateVersionsRequest.Size(m)
}
func (m *ListWorkflowTemplateVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplateVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplateVersionsRequest proto.InternalMessageInfo

func (m *ListWorkflowTemplateVersionsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListWorkflowTemplateVersionsRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ListWorkflowTemplateVersionsResponse struct {
	Count                int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates    []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListWorkflowTemplateVersionsResponse) Reset()         { *m = ListWorkflowTemplateVersionsResponse{} }
func (m *ListWorkflowTemplateVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplateVersionsResponse) ProtoMessage()    {}
func (*ListWorkflowTemplateVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{4}
}

func (m *ListWorkflowTemplateVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Unmarshal(m, b)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Merge(m, src)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplateVersionsResponse.Size(m)
}
func (m *ListWorkflowTemplateVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplateVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplateVersionsResponse proto.InternalMessageInfo

func (m *ListWorkflowTemplateVersionsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkflowTemplateVersionsResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplates
	}
	return nil
}

type ListWorkflowTemplatesRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowTemplatesRequest) Reset()         { *m = ListWorkflowTemplatesRequest{} }
func (m *ListWorkflowTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplatesRequest) ProtoMessage()    {}
func (*ListWorkflowTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{5}
}

func (m *ListWorkflowTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Unmarshal(m, b)
}
func (m *ListWorkflowTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplatesRequest.Merge(m, src)
}
func (m *ListWorkflowTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplatesRequest.Size(m)
}
func (m *ListWorkflowTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplatesRequest proto.InternalMessageInfo

func (m *ListWorkflowTemplatesRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListWorkflowTemplatesResponse struct {
	Count                int32               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	WorkflowTemplates    []*WorkflowTemplate `protobuf:"bytes,2,rep,name=workflowTemplates,proto3" json:"workflowTemplates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListWorkflowTemplatesResponse) Reset()         { *m = ListWorkflowTemplatesResponse{} }
func (m *ListWorkflowTemplatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowTemplatesResponse) ProtoMessage()    {}
func (*ListWorkflowTemplatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{6}
}

func (m *ListWorkflowTemplatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Unmarshal(m, b)
}
func (m *ListWorkflowTemplatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowTemplatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowTemplatesResponse.Merge(m, src)
}
func (m *ListWorkflowTemplatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowTemplatesResponse.Size(m)
}
func (m *ListWorkflowTemplatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowTemplatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowTemplatesResponse proto.InternalMessageInfo

func (m *ListWorkflowTemplatesResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkflowTemplatesResponse) GetWorkflowTemplates() []*WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplates
	}
	return nil
}

type ArchiveWorkflowTemplateRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveWorkflowTemplateRequest) Reset()         { *m = ArchiveWorkflowTemplateRequest{} }
func (m *ArchiveWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*ArchiveWorkflowTemplateRequest) ProtoMessage()    {}
func (*ArchiveWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{7}
}

func (m *ArchiveWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *ArchiveWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *ArchiveWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveWorkflowTemplateRequest.Merge(m, src)
}
func (m *ArchiveWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_ArchiveWorkflowTemplateRequest.Size(m)
}
func (m *ArchiveWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveWorkflowTemplateRequest proto.InternalMessageInfo

func (m *ArchiveWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ArchiveWorkflowTemplateRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type ArchiveWorkflowTemplateResponse struct {
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,1,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ArchiveWorkflowTemplateResponse) Reset()         { *m = ArchiveWorkflowTemplateResponse{} }
func (m *ArchiveWorkflowTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*ArchiveWorkflowTemplateResponse) ProtoMessage()    {}
func (*ArchiveWorkflowTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{8}
}

func (m *ArchiveWorkflowTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveWorkflowTemplateResponse.Unmarshal(m, b)
}
func (m *ArchiveWorkflowTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveWorkflowTemplateResponse.Marshal(b, m, deterministic)
}
func (m *ArchiveWorkflowTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveWorkflowTemplateResponse.Merge(m, src)
}
func (m *ArchiveWorkflowTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_ArchiveWorkflowTemplateResponse.Size(m)
}
func (m *ArchiveWorkflowTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveWorkflowTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveWorkflowTemplateResponse proto.InternalMessageInfo

func (m *ArchiveWorkflowTemplateResponse) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type WorkflowTemplate struct {
	CreatedAt            string   `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Manifest             string   `protobuf:"bytes,5,opt,name=manifest,proto3" json:"manifest,omitempty"`
	IsLatest             bool     `protobuf:"varint,6,opt,name=isLatest,proto3" json:"isLatest,omitempty"`
	IsArchived           bool     `protobuf:"varint,7,opt,name=isArchived,proto3" json:"isArchived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{9}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WorkflowTemplate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WorkflowTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTemplate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowTemplate) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *WorkflowTemplate) GetIsLatest() bool {
	if m != nil {
		return m.IsLatest
	}
	return false
}

func (m *WorkflowTemplate) GetIsArchived() bool {
	if m != nil {
		return m.IsArchived
	}
	return false
}

func init() {
	proto.RegisterType((*CreateWorkflowTemplateRequest)(nil), "api.CreateWorkflowTemplateRequest")
	proto.RegisterType((*UpdateWorkflowTemplateVersionRequest)(nil), "api.UpdateWorkflowTemplateVersionRequest")
	proto.RegisterType((*GetWorkflowTemplateRequest)(nil), "api.GetWorkflowTemplateRequest")
	proto.RegisterType((*ListWorkflowTemplateVersionsRequest)(nil), "api.ListWorkflowTemplateVersionsRequest")
	proto.RegisterType((*ListWorkflowTemplateVersionsResponse)(nil), "api.ListWorkflowTemplateVersionsResponse")
	proto.RegisterType((*ListWorkflowTemplatesRequest)(nil), "api.ListWorkflowTemplatesRequest")
	proto.RegisterType((*ListWorkflowTemplatesResponse)(nil), "api.ListWorkflowTemplatesResponse")
	proto.RegisterType((*ArchiveWorkflowTemplateRequest)(nil), "api.ArchiveWorkflowTemplateRequest")
	proto.RegisterType((*ArchiveWorkflowTemplateResponse)(nil), "api.ArchiveWorkflowTemplateResponse")
	proto.RegisterType((*WorkflowTemplate)(nil), "api.WorkflowTemplate")
}

func init() {
	proto.RegisterFile("workflow_template.proto", fileDescriptor_b9a07547748a96e8)
}

var fileDescriptor_b9a07547748a96e8 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0x4e, 0x19, 0x03, 0xf6, 0xbc, 0x60, 0xa3, 0xb1, 0x21, 0x80, 0xcb, 0xe4, 0xb0, 0x13, 0x07,
	0xbd, 0x7a, 0x21, 0x1c, 0xbc, 0x70, 0x30, 0x8b, 0xe8, 0xd1, 0xd4, 0xad, 0xc4, 0x46, 0x58, 0xeb,
	0x5a, 0x20, 0xf1, 0xa4, 0x27, 0xff, 0x9a, 0x3f, 0xcb, 0xac, 0x0c, 0x44, 0xd8, 0x8c, 0x12, 0xf5,
	0xd6, 0xf7, 0xfa, 0xbd, 0xd7, 0xef, 0xfb, 0xf2, 0x6d, 0x70, 0x34, 0x17, 0xc9, 0xc3, 0x68, 0x2c,
	0xe6, 0xb7, 0x9a, 0x4d, 0xe4, 0x98, 0x6a, 0xd6, 0x95, 0x89, 0xd0, 0x02, 0x5b, 0x54, 0x72, 0xef,
	0x19, 0x41, 0xab, 0x9f, 0x30, 0xaa, 0xd9, 0x4d, 0x06, 0xbb, 0xca, 0x50, 0x01, 0x7b, 0x9c, 0x32,
	0xa5, 0x71, 0x13, 0x9c, 0x98, 0x4e, 0x98, 0x92, 0x34, 0x64, 0x04, 0xb9, 0xc8, 0x77, 0x82, 0x8f,
	0x06, 0xee, 0x41, 0x7d, 0xbe, 0x31, 0x48, 0x4a, 0x2e, 0xf2, 0xf7, 0x4e, 0x0f, 0xbb, 0x54, 0xf2,
	0xee, 0xd6, 0xd6, 0x2d, 0xb8, 0xf7, 0x8a, 0xa0, 0x33, 0x94, 0x51, 0x0e, 0x85, 0x6b, 0x96, 0x28,
	0x2e, 0xe2, 0x7f, 0x63, 0x32, 0x82, 0xc6, 0x05, 0xd3, 0xbb, 0x19, 0x51, 0x07, 0x6b, 0xca, 0x23,
	0xf3, 0xa2, 0x13, 0xa4, 0x47, 0x4c, 0xa0, 0x3a, 0x5b, 0x08, 0x20, 0x96, 0x8b, 0x7c, 0x3b, 0x58,
	0x96, 0xde, 0x10, 0x4e, 0x06, 0x5c, 0xe9, 0x02, 0xb9, 0x6a, 0xc7, 0x07, 0xbd, 0x17, 0x04, 0x9d,
	0xaf, 0xf7, 0x2a, 0x29, 0x62, 0xc5, 0xf0, 0x01, 0xd8, 0xa1, 0x98, 0xc6, 0xda, 0x2c, 0xb5, 0x83,
	0x45, 0x81, 0xfb, 0xb0, 0xbf, 0xe9, 0x88, 0x22, 0x25, 0xd7, 0x2a, 0x76, 0x70, 0x1b, 0xef, 0x9d,
	0x43, 0x33, 0x8f, 0xc2, 0xf7, 0x34, 0x79, 0x4f, 0xd0, 0x2a, 0x98, 0xfe, 0x7b, 0xe6, 0x97, 0xd0,
	0xee, 0x25, 0xe1, 0x3d, 0x9f, 0xb1, 0x5f, 0x0a, 0x80, 0x17, 0xc1, 0x71, 0xe1, 0xc6, 0x4c, 0x4f,
	0x5e, 0x68, 0xd1, 0xcf, 0x42, 0xfb, 0x86, 0xa0, 0xbe, 0x09, 0x4b, 0xa9, 0x86, 0xe6, 0xab, 0x8e,
	0x7a, 0x7a, 0x49, 0x75, 0xd5, 0xc8, 0xc9, 0x2a, 0x86, 0x72, 0xaa, 0xc4, 0x04, 0xd5, 0x09, 0xcc,
	0x79, 0x3d, 0xbf, 0xe5, 0x4f, 0xf9, 0xc5, 0x0d, 0xa8, 0x4d, 0x68, 0xcc, 0x47, 0x4c, 0x69, 0x62,
	0x9b, 0x89, 0x55, 0x9d, 0xde, 0x71, 0x35, 0x48, 0x1d, 0xd5, 0xa4, 0xe2, 0x22, 0xbf, 0x16, 0xac,
	0x6a, 0xdc, 0x06, 0xe0, 0x2a, 0xb3, 0x24, 0x22, 0x55, 0x73, 0xbb, 0xd6, 0xb9, 0xab, 0x98, 0x1f,
	0xd3, 0xd9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xb1, 0x9a, 0x5e, 0xb3, 0x04, 0x00, 0x00,
}
