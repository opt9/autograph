// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/containeranalysis/v1alpha1/source_context.proto

package containeranalysis // import "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1alpha1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of an alias.
type AliasContext_Kind int32

const (
	// Unknown.
	AliasContext_KIND_UNSPECIFIED AliasContext_Kind = 0
	// Git tag.
	AliasContext_FIXED AliasContext_Kind = 1
	// Git branch.
	AliasContext_MOVABLE AliasContext_Kind = 2
	// Used to specify non-standard aliases. For example, if a Git repo has a
	// ref named "refs/foo/bar".
	AliasContext_OTHER AliasContext_Kind = 4
)

var AliasContext_Kind_name = map[int32]string{
	0: "KIND_UNSPECIFIED",
	1: "FIXED",
	2: "MOVABLE",
	4: "OTHER",
}

var AliasContext_Kind_value = map[string]int32{
	"KIND_UNSPECIFIED": 0,
	"FIXED":            1,
	"MOVABLE":          2,
	"OTHER":            4,
}

func (x AliasContext_Kind) String() string {
	return proto.EnumName(AliasContext_Kind_name, int32(x))
}

func (AliasContext_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{1, 0}
}

// A SourceContext is a reference to a tree of files. A SourceContext together
// with a path point to a unique revision of a single file or directory.
type SourceContext struct {
	// A SourceContext can refer any one of the following types of repositories.
	//
	// Types that are valid to be assigned to Context:
	//	*SourceContext_CloudRepo
	//	*SourceContext_Gerrit
	//	*SourceContext_Git
	Context isSourceContext_Context `protobuf_oneof:"context"`
	// Labels with user defined metadata.
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SourceContext) Reset()         { *m = SourceContext{} }
func (m *SourceContext) String() string { return proto.CompactTextString(m) }
func (*SourceContext) ProtoMessage()    {}
func (*SourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{0}
}
func (m *SourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceContext.Unmarshal(m, b)
}
func (m *SourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceContext.Marshal(b, m, deterministic)
}
func (m *SourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceContext.Merge(m, src)
}
func (m *SourceContext) XXX_Size() int {
	return xxx_messageInfo_SourceContext.Size(m)
}
func (m *SourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_SourceContext proto.InternalMessageInfo

type isSourceContext_Context interface {
	isSourceContext_Context()
}

type SourceContext_CloudRepo struct {
	CloudRepo *CloudRepoSourceContext `protobuf:"bytes,1,opt,name=cloud_repo,json=cloudRepo,proto3,oneof"`
}

type SourceContext_Gerrit struct {
	Gerrit *GerritSourceContext `protobuf:"bytes,2,opt,name=gerrit,proto3,oneof"`
}

type SourceContext_Git struct {
	Git *GitSourceContext `protobuf:"bytes,3,opt,name=git,proto3,oneof"`
}

func (*SourceContext_CloudRepo) isSourceContext_Context() {}

func (*SourceContext_Gerrit) isSourceContext_Context() {}

func (*SourceContext_Git) isSourceContext_Context() {}

func (m *SourceContext) GetContext() isSourceContext_Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SourceContext) GetCloudRepo() *CloudRepoSourceContext {
	if x, ok := m.GetContext().(*SourceContext_CloudRepo); ok {
		return x.CloudRepo
	}
	return nil
}

func (m *SourceContext) GetGerrit() *GerritSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Gerrit); ok {
		return x.Gerrit
	}
	return nil
}

func (m *SourceContext) GetGit() *GitSourceContext {
	if x, ok := m.GetContext().(*SourceContext_Git); ok {
		return x.Git
	}
	return nil
}

func (m *SourceContext) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SourceContext) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SourceContext_OneofMarshaler, _SourceContext_OneofUnmarshaler, _SourceContext_OneofSizer, []interface{}{
		(*SourceContext_CloudRepo)(nil),
		(*SourceContext_Gerrit)(nil),
		(*SourceContext_Git)(nil),
	}
}

func _SourceContext_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SourceContext)
	// context
	switch x := m.Context.(type) {
	case *SourceContext_CloudRepo:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudRepo); err != nil {
			return err
		}
	case *SourceContext_Gerrit:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gerrit); err != nil {
			return err
		}
	case *SourceContext_Git:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Git); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SourceContext.Context has unexpected type %T", x)
	}
	return nil
}

func _SourceContext_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SourceContext)
	switch tag {
	case 1: // context.cloud_repo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CloudRepoSourceContext)
		err := b.DecodeMessage(msg)
		m.Context = &SourceContext_CloudRepo{msg}
		return true, err
	case 2: // context.gerrit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GerritSourceContext)
		err := b.DecodeMessage(msg)
		m.Context = &SourceContext_Gerrit{msg}
		return true, err
	case 3: // context.git
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GitSourceContext)
		err := b.DecodeMessage(msg)
		m.Context = &SourceContext_Git{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SourceContext_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SourceContext)
	// context
	switch x := m.Context.(type) {
	case *SourceContext_CloudRepo:
		s := proto.Size(x.CloudRepo)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SourceContext_Gerrit:
		s := proto.Size(x.Gerrit)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SourceContext_Git:
		s := proto.Size(x.Git)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An alias to a repo revision.
type AliasContext struct {
	// The alias kind.
	Kind AliasContext_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=google.devtools.containeranalysis.v1alpha1.AliasContext_Kind" json:"kind,omitempty"`
	// The alias name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliasContext) Reset()         { *m = AliasContext{} }
func (m *AliasContext) String() string { return proto.CompactTextString(m) }
func (*AliasContext) ProtoMessage()    {}
func (*AliasContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{1}
}
func (m *AliasContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliasContext.Unmarshal(m, b)
}
func (m *AliasContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliasContext.Marshal(b, m, deterministic)
}
func (m *AliasContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliasContext.Merge(m, src)
}
func (m *AliasContext) XXX_Size() int {
	return xxx_messageInfo_AliasContext.Size(m)
}
func (m *AliasContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AliasContext.DiscardUnknown(m)
}

var xxx_messageInfo_AliasContext proto.InternalMessageInfo

func (m *AliasContext) GetKind() AliasContext_Kind {
	if m != nil {
		return m.Kind
	}
	return AliasContext_KIND_UNSPECIFIED
}

func (m *AliasContext) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A CloudRepoSourceContext denotes a particular revision in a Google Cloud
// Source Repo.
type CloudRepoSourceContext struct {
	// The ID of the repo.
	RepoId *RepoId `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	// A revision in a Cloud Repo can be identified by either its revision ID or
	// its alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*CloudRepoSourceContext_RevisionId
	//	*CloudRepoSourceContext_AliasContext
	Revision             isCloudRepoSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CloudRepoSourceContext) Reset()         { *m = CloudRepoSourceContext{} }
func (m *CloudRepoSourceContext) String() string { return proto.CompactTextString(m) }
func (*CloudRepoSourceContext) ProtoMessage()    {}
func (*CloudRepoSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{2}
}
func (m *CloudRepoSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudRepoSourceContext.Unmarshal(m, b)
}
func (m *CloudRepoSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudRepoSourceContext.Marshal(b, m, deterministic)
}
func (m *CloudRepoSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudRepoSourceContext.Merge(m, src)
}
func (m *CloudRepoSourceContext) XXX_Size() int {
	return xxx_messageInfo_CloudRepoSourceContext.Size(m)
}
func (m *CloudRepoSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudRepoSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_CloudRepoSourceContext proto.InternalMessageInfo

func (m *CloudRepoSourceContext) GetRepoId() *RepoId {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type isCloudRepoSourceContext_Revision interface {
	isCloudRepoSourceContext_Revision()
}

type CloudRepoSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type CloudRepoSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,3,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*CloudRepoSourceContext_RevisionId) isCloudRepoSourceContext_Revision() {}

func (*CloudRepoSourceContext_AliasContext) isCloudRepoSourceContext_Revision() {}

func (m *CloudRepoSourceContext) GetRevision() isCloudRepoSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *CloudRepoSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

func (m *CloudRepoSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*CloudRepoSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CloudRepoSourceContext) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CloudRepoSourceContext_OneofMarshaler, _CloudRepoSourceContext_OneofUnmarshaler, _CloudRepoSourceContext_OneofSizer, []interface{}{
		(*CloudRepoSourceContext_RevisionId)(nil),
		(*CloudRepoSourceContext_AliasContext)(nil),
	}
}

func _CloudRepoSourceContext_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CloudRepoSourceContext)
	// revision
	switch x := m.Revision.(type) {
	case *CloudRepoSourceContext_RevisionId:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RevisionId)
	case *CloudRepoSourceContext_AliasContext:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AliasContext); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CloudRepoSourceContext.Revision has unexpected type %T", x)
	}
	return nil
}

func _CloudRepoSourceContext_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CloudRepoSourceContext)
	switch tag {
	case 2: // revision.revision_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Revision = &CloudRepoSourceContext_RevisionId{x}
		return true, err
	case 3: // revision.alias_context
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AliasContext)
		err := b.DecodeMessage(msg)
		m.Revision = &CloudRepoSourceContext_AliasContext{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CloudRepoSourceContext_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CloudRepoSourceContext)
	// revision
	switch x := m.Revision.(type) {
	case *CloudRepoSourceContext_RevisionId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.RevisionId)))
		n += len(x.RevisionId)
	case *CloudRepoSourceContext_AliasContext:
		s := proto.Size(x.AliasContext)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A SourceContext referring to a Gerrit project.
type GerritSourceContext struct {
	// The URI of a running Gerrit instance.
	HostUri string `protobuf:"bytes,1,opt,name=host_uri,json=hostUri,proto3" json:"host_uri,omitempty"`
	// The full project name within the host. Projects may be nested, so
	// "project/subproject" is a valid project name. The "repo name" is
	// the hostURI/project.
	GerritProject string `protobuf:"bytes,2,opt,name=gerrit_project,json=gerritProject,proto3" json:"gerrit_project,omitempty"`
	// A revision in a Gerrit project can be identified by either its revision ID
	// or its alias.
	//
	// Types that are valid to be assigned to Revision:
	//	*GerritSourceContext_RevisionId
	//	*GerritSourceContext_AliasContext
	Revision             isGerritSourceContext_Revision `protobuf_oneof:"revision"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GerritSourceContext) Reset()         { *m = GerritSourceContext{} }
func (m *GerritSourceContext) String() string { return proto.CompactTextString(m) }
func (*GerritSourceContext) ProtoMessage()    {}
func (*GerritSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{3}
}
func (m *GerritSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritSourceContext.Unmarshal(m, b)
}
func (m *GerritSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritSourceContext.Marshal(b, m, deterministic)
}
func (m *GerritSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritSourceContext.Merge(m, src)
}
func (m *GerritSourceContext) XXX_Size() int {
	return xxx_messageInfo_GerritSourceContext.Size(m)
}
func (m *GerritSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GerritSourceContext proto.InternalMessageInfo

func (m *GerritSourceContext) GetHostUri() string {
	if m != nil {
		return m.HostUri
	}
	return ""
}

func (m *GerritSourceContext) GetGerritProject() string {
	if m != nil {
		return m.GerritProject
	}
	return ""
}

type isGerritSourceContext_Revision interface {
	isGerritSourceContext_Revision()
}

type GerritSourceContext_RevisionId struct {
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3,oneof"`
}

type GerritSourceContext_AliasContext struct {
	AliasContext *AliasContext `protobuf:"bytes,4,opt,name=alias_context,json=aliasContext,proto3,oneof"`
}

func (*GerritSourceContext_RevisionId) isGerritSourceContext_Revision() {}

func (*GerritSourceContext_AliasContext) isGerritSourceContext_Revision() {}

func (m *GerritSourceContext) GetRevision() isGerritSourceContext_Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *GerritSourceContext) GetRevisionId() string {
	if x, ok := m.GetRevision().(*GerritSourceContext_RevisionId); ok {
		return x.RevisionId
	}
	return ""
}

func (m *GerritSourceContext) GetAliasContext() *AliasContext {
	if x, ok := m.GetRevision().(*GerritSourceContext_AliasContext); ok {
		return x.AliasContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GerritSourceContext) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GerritSourceContext_OneofMarshaler, _GerritSourceContext_OneofUnmarshaler, _GerritSourceContext_OneofSizer, []interface{}{
		(*GerritSourceContext_RevisionId)(nil),
		(*GerritSourceContext_AliasContext)(nil),
	}
}

func _GerritSourceContext_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GerritSourceContext)
	// revision
	switch x := m.Revision.(type) {
	case *GerritSourceContext_RevisionId:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RevisionId)
	case *GerritSourceContext_AliasContext:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AliasContext); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GerritSourceContext.Revision has unexpected type %T", x)
	}
	return nil
}

func _GerritSourceContext_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GerritSourceContext)
	switch tag {
	case 3: // revision.revision_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Revision = &GerritSourceContext_RevisionId{x}
		return true, err
	case 4: // revision.alias_context
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AliasContext)
		err := b.DecodeMessage(msg)
		m.Revision = &GerritSourceContext_AliasContext{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GerritSourceContext_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GerritSourceContext)
	// revision
	switch x := m.Revision.(type) {
	case *GerritSourceContext_RevisionId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.RevisionId)))
		n += len(x.RevisionId)
	case *GerritSourceContext_AliasContext:
		s := proto.Size(x.AliasContext)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GitSourceContext denotes a particular revision in a third party Git
// repository (e.g., GitHub).
type GitSourceContext struct {
	// Git repository URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Required.
	// Git commit hash.
	RevisionId           string   `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitSourceContext) Reset()         { *m = GitSourceContext{} }
func (m *GitSourceContext) String() string { return proto.CompactTextString(m) }
func (*GitSourceContext) ProtoMessage()    {}
func (*GitSourceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{4}
}
func (m *GitSourceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitSourceContext.Unmarshal(m, b)
}
func (m *GitSourceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitSourceContext.Marshal(b, m, deterministic)
}
func (m *GitSourceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitSourceContext.Merge(m, src)
}
func (m *GitSourceContext) XXX_Size() int {
	return xxx_messageInfo_GitSourceContext.Size(m)
}
func (m *GitSourceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_GitSourceContext.DiscardUnknown(m)
}

var xxx_messageInfo_GitSourceContext proto.InternalMessageInfo

func (m *GitSourceContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GitSourceContext) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

// A unique identifier for a Cloud Repo.
type RepoId struct {
	// A cloud repo can be identified by either its project ID and repository name
	// combination, or its globally unique identifier.
	//
	// Types that are valid to be assigned to Id:
	//	*RepoId_ProjectRepoId
	//	*RepoId_Uid
	Id                   isRepoId_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RepoId) Reset()         { *m = RepoId{} }
func (m *RepoId) String() string { return proto.CompactTextString(m) }
func (*RepoId) ProtoMessage()    {}
func (*RepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{5}
}
func (m *RepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoId.Unmarshal(m, b)
}
func (m *RepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoId.Marshal(b, m, deterministic)
}
func (m *RepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoId.Merge(m, src)
}
func (m *RepoId) XXX_Size() int {
	return xxx_messageInfo_RepoId.Size(m)
}
func (m *RepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoId.DiscardUnknown(m)
}

var xxx_messageInfo_RepoId proto.InternalMessageInfo

type isRepoId_Id interface {
	isRepoId_Id()
}

type RepoId_ProjectRepoId struct {
	ProjectRepoId *ProjectRepoId `protobuf:"bytes,1,opt,name=project_repo_id,json=projectRepoId,proto3,oneof"`
}

type RepoId_Uid struct {
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3,oneof"`
}

func (*RepoId_ProjectRepoId) isRepoId_Id() {}

func (*RepoId_Uid) isRepoId_Id() {}

func (m *RepoId) GetId() isRepoId_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RepoId) GetProjectRepoId() *ProjectRepoId {
	if x, ok := m.GetId().(*RepoId_ProjectRepoId); ok {
		return x.ProjectRepoId
	}
	return nil
}

func (m *RepoId) GetUid() string {
	if x, ok := m.GetId().(*RepoId_Uid); ok {
		return x.Uid
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RepoId) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RepoId_OneofMarshaler, _RepoId_OneofUnmarshaler, _RepoId_OneofSizer, []interface{}{
		(*RepoId_ProjectRepoId)(nil),
		(*RepoId_Uid)(nil),
	}
}

func _RepoId_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RepoId)
	// id
	switch x := m.Id.(type) {
	case *RepoId_ProjectRepoId:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProjectRepoId); err != nil {
			return err
		}
	case *RepoId_Uid:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Uid)
	case nil:
	default:
		return fmt.Errorf("RepoId.Id has unexpected type %T", x)
	}
	return nil
}

func _RepoId_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RepoId)
	switch tag {
	case 1: // id.project_repo_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProjectRepoId)
		err := b.DecodeMessage(msg)
		m.Id = &RepoId_ProjectRepoId{msg}
		return true, err
	case 2: // id.uid
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Id = &RepoId_Uid{x}
		return true, err
	default:
		return false, nil
	}
}

func _RepoId_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RepoId)
	// id
	switch x := m.Id.(type) {
	case *RepoId_ProjectRepoId:
		s := proto.Size(x.ProjectRepoId)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RepoId_Uid:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Uid)))
		n += len(x.Uid)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Selects a repo using a Google Cloud Platform project ID (e.g.,
// winged-cargo-31) and a repo name within that project.
type ProjectRepoId struct {
	// The ID of the project.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The name of the repo. Leave empty for the default repo.
	RepoName             string   `protobuf:"bytes,2,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRepoId) Reset()         { *m = ProjectRepoId{} }
func (m *ProjectRepoId) String() string { return proto.CompactTextString(m) }
func (*ProjectRepoId) ProtoMessage()    {}
func (*ProjectRepoId) Descriptor() ([]byte, []int) {
	return fileDescriptor_49411dfe5c54cc3e, []int{6}
}
func (m *ProjectRepoId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRepoId.Unmarshal(m, b)
}
func (m *ProjectRepoId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRepoId.Marshal(b, m, deterministic)
}
func (m *ProjectRepoId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRepoId.Merge(m, src)
}
func (m *ProjectRepoId) XXX_Size() int {
	return xxx_messageInfo_ProjectRepoId.Size(m)
}
func (m *ProjectRepoId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRepoId.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRepoId proto.InternalMessageInfo

func (m *ProjectRepoId) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ProjectRepoId) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func init() {
	proto.RegisterType((*SourceContext)(nil), "google.devtools.containeranalysis.v1alpha1.SourceContext")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.containeranalysis.v1alpha1.SourceContext.LabelsEntry")
	proto.RegisterType((*AliasContext)(nil), "google.devtools.containeranalysis.v1alpha1.AliasContext")
	proto.RegisterType((*CloudRepoSourceContext)(nil), "google.devtools.containeranalysis.v1alpha1.CloudRepoSourceContext")
	proto.RegisterType((*GerritSourceContext)(nil), "google.devtools.containeranalysis.v1alpha1.GerritSourceContext")
	proto.RegisterType((*GitSourceContext)(nil), "google.devtools.containeranalysis.v1alpha1.GitSourceContext")
	proto.RegisterType((*RepoId)(nil), "google.devtools.containeranalysis.v1alpha1.RepoId")
	proto.RegisterType((*ProjectRepoId)(nil), "google.devtools.containeranalysis.v1alpha1.ProjectRepoId")
	proto.RegisterEnum("google.devtools.containeranalysis.v1alpha1.AliasContext_Kind", AliasContext_Kind_name, AliasContext_Kind_value)
}

func init() {
	proto.RegisterFile("google/devtools/containeranalysis/v1alpha1/source_context.proto", fileDescriptor_49411dfe5c54cc3e)
}

var fileDescriptor_49411dfe5c54cc3e = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x5d, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0xe3, 0x38, 0x24, 0xf8, 0x84, 0x70, 0xa3, 0xb9, 0xe8, 0x2a, 0x97, 0x7b, 0xab, 0x52,
	0x4b, 0x48, 0xa8, 0x0f, 0xb6, 0x48, 0x5f, 0xa0, 0x1f, 0x42, 0x24, 0x18, 0x62, 0x85, 0x42, 0x6a,
	0x4a, 0xd5, 0x0f, 0x55, 0xd6, 0x60, 0x8f, 0xcc, 0x14, 0xe3, 0xb1, 0xc6, 0x76, 0x54, 0x56, 0xd0,
	0x97, 0xae, 0xa2, 0x8b, 0xe8, 0x12, 0xba, 0x95, 0xae, 0xa0, 0xef, 0xd5, 0x8c, 0x6d, 0xc9, 0x01,
	0x2a, 0x91, 0x4a, 0x7d, 0xca, 0xcc, 0x99, 0x99, 0xdf, 0xf9, 0x9f, 0x33, 0xff, 0x89, 0x61, 0x27,
	0x60, 0x2c, 0x08, 0x89, 0xe9, 0x93, 0x69, 0xca, 0x58, 0x98, 0x98, 0x1e, 0x8b, 0x52, 0x4c, 0x23,
	0xc2, 0x71, 0x84, 0xc3, 0xab, 0x84, 0x26, 0xe6, 0x74, 0x13, 0x87, 0xf1, 0x39, 0xde, 0x34, 0x13,
	0x96, 0x71, 0x8f, 0xb8, 0x62, 0x07, 0xf9, 0x98, 0x1a, 0x31, 0x67, 0x29, 0x43, 0x0f, 0x73, 0x80,
	0x51, 0x02, 0x8c, 0x1b, 0x00, 0xa3, 0x04, 0xac, 0xfe, 0x5f, 0x24, 0xc3, 0x31, 0x35, 0x71, 0x14,
	0xb1, 0x14, 0xa7, 0x94, 0x45, 0x49, 0x4e, 0xd2, 0xbf, 0xa9, 0xd0, 0x39, 0x91, 0x29, 0x86, 0x79,
	0x06, 0xe4, 0x01, 0x78, 0x21, 0xcb, 0x7c, 0x97, 0x93, 0x98, 0xf5, 0x94, 0x35, 0x65, 0xa3, 0xdd,
	0x1f, 0x18, 0x77, 0x4f, 0x68, 0x0c, 0xc5, 0x69, 0x87, 0xc4, 0x6c, 0x86, 0x3b, 0xaa, 0x39, 0x9a,
	0x57, 0xae, 0xa0, 0x37, 0xd0, 0x0c, 0x08, 0xe7, 0x34, 0xed, 0xd5, 0x65, 0x82, 0x9d, 0x79, 0x12,
	0x1c, 0xc8, 0x93, 0xd7, 0xe9, 0x05, 0x10, 0x4d, 0x40, 0x0d, 0x68, 0xda, 0x53, 0x25, 0xf7, 0xe9,
	0x5c, 0xdc, 0x9b, 0x50, 0x81, 0x42, 0xef, 0xa1, 0x19, 0xe2, 0x33, 0x12, 0x26, 0xbd, 0xc6, 0x9a,
	0xba, 0xd1, 0xee, 0x5b, 0xf3, 0x40, 0x67, 0x88, 0xc6, 0xa1, 0xe4, 0x58, 0x51, 0xca, 0xaf, 0x9c,
	0x02, 0xba, 0xba, 0x0d, 0xed, 0x4a, 0x18, 0x75, 0x41, 0xbd, 0x20, 0x57, 0xb2, 0xf1, 0x9a, 0x23,
	0x86, 0x68, 0x05, 0x16, 0xa6, 0x38, 0xcc, 0x88, 0xec, 0x95, 0xe6, 0xe4, 0x93, 0xc7, 0xf5, 0x2d,
	0x65, 0xa0, 0x41, 0xab, 0x30, 0x86, 0xfe, 0x55, 0x81, 0xa5, 0xdd, 0x90, 0xe2, 0xa4, 0xbc, 0xc7,
	0x17, 0xd0, 0xb8, 0xa0, 0x91, 0x2f, 0x41, 0xcb, 0xfd, 0x67, 0xf3, 0x68, 0xae, 0x72, 0x8c, 0x31,
	0x8d, 0x7c, 0x47, 0xa2, 0x10, 0x82, 0x46, 0x84, 0x2f, 0x4b, 0x1d, 0x72, 0xac, 0xef, 0x40, 0x43,
	0xec, 0x40, 0x2b, 0xd0, 0x1d, 0xdb, 0x47, 0x7b, 0xee, 0xe9, 0xd1, 0xc9, 0xc4, 0x1a, 0xda, 0xfb,
	0xb6, 0xb5, 0xd7, 0xad, 0x21, 0x0d, 0x16, 0xf6, 0xed, 0xd7, 0xd6, 0x5e, 0x57, 0x41, 0x6d, 0x68,
	0x3d, 0x3f, 0x7e, 0xb5, 0x3b, 0x38, 0xb4, 0xba, 0x75, 0x11, 0x3f, 0x7e, 0x39, 0xb2, 0x9c, 0x6e,
	0x43, 0xff, 0xa1, 0xc0, 0x3f, 0xb7, 0x5b, 0x06, 0x8d, 0xa1, 0x25, 0x4c, 0xe8, 0x52, 0xbf, 0xf0,
	0x61, 0x7f, 0x9e, 0x2a, 0x04, 0xcf, 0xf6, 0x9d, 0x26, 0x97, 0xbf, 0xe8, 0x01, 0xb4, 0x39, 0x99,
	0xd2, 0x84, 0xb2, 0x48, 0x00, 0x65, 0x0d, 0xa3, 0x9a, 0x03, 0x65, 0xd0, 0xf6, 0x91, 0x0b, 0x1d,
	0x2c, 0x4a, 0x2f, 0x5f, 0x5b, 0x61, 0xa2, 0xad, 0xdf, 0xed, 0xdd, 0xa8, 0xe6, 0x2c, 0xe1, 0xca,
	0x7c, 0x00, 0xb0, 0x58, 0xa6, 0xd3, 0xbf, 0x2b, 0xf0, 0xf7, 0x2d, 0x4e, 0x46, 0xff, 0xc2, 0xe2,
	0x39, 0x4b, 0x52, 0x37, 0xe3, 0xb4, 0x30, 0x41, 0x4b, 0xcc, 0x4f, 0x39, 0x45, 0xeb, 0xb0, 0x9c,
	0x9b, 0xdc, 0x8d, 0x39, 0xfb, 0x40, 0xbc, 0xb4, 0xb8, 0x89, 0x4e, 0x1e, 0x9d, 0xe4, 0xc1, 0xeb,
	0x95, 0xaa, 0x77, 0xa9, 0xb4, 0xf1, 0x07, 0x2b, 0xb5, 0xa0, 0x7b, 0xfd, 0x69, 0x09, 0x97, 0x67,
	0x3c, 0x2c, 0x5d, 0x9e, 0xf1, 0x10, 0xdd, 0xbf, 0xe5, 0x7e, 0xaa, 0x9a, 0xf5, 0x4f, 0x0a, 0x34,
	0xf3, 0x3b, 0x45, 0x1e, 0xfc, 0x55, 0x74, 0xc0, 0x9d, 0x35, 0xc8, 0xf6, 0x3c, 0x05, 0x14, 0xfd,
	0xca, 0x99, 0xa3, 0x9a, 0xd3, 0x89, 0xab, 0x01, 0x84, 0x40, 0xcd, 0x2a, 0x46, 0x11, 0x93, 0x41,
	0x03, 0xea, 0xd4, 0xd7, 0xc7, 0xd0, 0x99, 0x39, 0x8b, 0xee, 0x01, 0x94, 0x7a, 0x0a, 0x29, 0x9a,
	0xa3, 0x15, 0x11, 0xdb, 0x47, 0xff, 0x81, 0x26, 0x65, 0x56, 0x1e, 0xcf, 0xa2, 0x08, 0x1c, 0xe1,
	0x4b, 0x32, 0xf8, 0xac, 0xc0, 0xba, 0xc7, 0x2e, 0x4b, 0xe1, 0xbf, 0xd6, 0x3b, 0x51, 0xde, 0xbe,
	0x2b, 0x36, 0x05, 0x2c, 0xc4, 0x51, 0x60, 0x30, 0x1e, 0x98, 0x01, 0x89, 0xe4, 0x3f, 0xb9, 0x99,
	0x2f, 0xe1, 0x98, 0x26, 0x77, 0xf9, 0xae, 0x3c, 0xb9, 0xb1, 0xf4, 0xa5, 0xae, 0x1e, 0x0c, 0x77,
	0xcf, 0x9a, 0x92, 0xf6, 0xe8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xcf, 0x59, 0x43, 0xa4,
	0x06, 0x00, 0x00,
}
