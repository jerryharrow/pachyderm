// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: server/worker/pipeline/transform/transform.proto

package transform

import (
	pfs "github.com/pachyderm/pachyderm/v2/src/pfs"
	pps "github.com/pachyderm/pachyderm/v2/src/pps"
	datum "github.com/pachyderm/pachyderm/v2/src/server/worker/datum"
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

type CreateParallelDatumsTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job           *pps.Job       `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Salt          string         `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	FileSetId     string         `protobuf:"bytes,3,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	BaseFileSetId string         `protobuf:"bytes,4,opt,name=base_file_set_id,json=baseFileSetId,proto3" json:"base_file_set_id,omitempty"`
	PathRange     *pfs.PathRange `protobuf:"bytes,5,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *CreateParallelDatumsTask) Reset() {
	*x = CreateParallelDatumsTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParallelDatumsTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParallelDatumsTask) ProtoMessage() {}

func (x *CreateParallelDatumsTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParallelDatumsTask.ProtoReflect.Descriptor instead.
func (*CreateParallelDatumsTask) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{0}
}

func (x *CreateParallelDatumsTask) GetJob() *pps.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *CreateParallelDatumsTask) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *CreateParallelDatumsTask) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *CreateParallelDatumsTask) GetBaseFileSetId() string {
	if x != nil {
		return x.BaseFileSetId
	}
	return ""
}

func (x *CreateParallelDatumsTask) GetPathRange() *pfs.PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type CreateParallelDatumsTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSetId string       `protobuf:"bytes,1,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	Stats     *datum.Stats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *CreateParallelDatumsTaskResult) Reset() {
	*x = CreateParallelDatumsTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParallelDatumsTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParallelDatumsTaskResult) ProtoMessage() {}

func (x *CreateParallelDatumsTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParallelDatumsTaskResult.ProtoReflect.Descriptor instead.
func (*CreateParallelDatumsTaskResult) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{1}
}

func (x *CreateParallelDatumsTaskResult) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *CreateParallelDatumsTaskResult) GetStats() *datum.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type CreateSerialDatumsTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job            *pps.Job       `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Salt           string         `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	FileSetId      string         `protobuf:"bytes,3,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	BaseMetaCommit *pfs.Commit    `protobuf:"bytes,4,opt,name=base_meta_commit,json=baseMetaCommit,proto3" json:"base_meta_commit,omitempty"`
	NoSkip         bool           `protobuf:"varint,5,opt,name=no_skip,json=noSkip,proto3" json:"no_skip,omitempty"`
	PathRange      *pfs.PathRange `protobuf:"bytes,6,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
}

func (x *CreateSerialDatumsTask) Reset() {
	*x = CreateSerialDatumsTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSerialDatumsTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSerialDatumsTask) ProtoMessage() {}

func (x *CreateSerialDatumsTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSerialDatumsTask.ProtoReflect.Descriptor instead.
func (*CreateSerialDatumsTask) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSerialDatumsTask) GetJob() *pps.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *CreateSerialDatumsTask) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *CreateSerialDatumsTask) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *CreateSerialDatumsTask) GetBaseMetaCommit() *pfs.Commit {
	if x != nil {
		return x.BaseMetaCommit
	}
	return nil
}

func (x *CreateSerialDatumsTask) GetNoSkip() bool {
	if x != nil {
		return x.NoSkip
	}
	return false
}

func (x *CreateSerialDatumsTask) GetPathRange() *pfs.PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

type CreateSerialDatumsTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSetId             string       `protobuf:"bytes,1,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	OutputDeleteFileSetId string       `protobuf:"bytes,2,opt,name=output_delete_file_set_id,json=outputDeleteFileSetId,proto3" json:"output_delete_file_set_id,omitempty"`
	MetaDeleteFileSetId   string       `protobuf:"bytes,3,opt,name=meta_delete_file_set_id,json=metaDeleteFileSetId,proto3" json:"meta_delete_file_set_id,omitempty"`
	Stats                 *datum.Stats `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *CreateSerialDatumsTaskResult) Reset() {
	*x = CreateSerialDatumsTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSerialDatumsTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSerialDatumsTaskResult) ProtoMessage() {}

func (x *CreateSerialDatumsTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSerialDatumsTaskResult.ProtoReflect.Descriptor instead.
func (*CreateSerialDatumsTaskResult) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSerialDatumsTaskResult) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *CreateSerialDatumsTaskResult) GetOutputDeleteFileSetId() string {
	if x != nil {
		return x.OutputDeleteFileSetId
	}
	return ""
}

func (x *CreateSerialDatumsTaskResult) GetMetaDeleteFileSetId() string {
	if x != nil {
		return x.MetaDeleteFileSetId
	}
	return ""
}

func (x *CreateSerialDatumsTaskResult) GetStats() *datum.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type CreateDatumSetsTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSetId string         `protobuf:"bytes,1,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	PathRange *pfs.PathRange `protobuf:"bytes,2,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
	SetSpec   *datum.SetSpec `protobuf:"bytes,3,opt,name=set_spec,json=setSpec,proto3" json:"set_spec,omitempty"`
}

func (x *CreateDatumSetsTask) Reset() {
	*x = CreateDatumSetsTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatumSetsTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatumSetsTask) ProtoMessage() {}

func (x *CreateDatumSetsTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatumSetsTask.ProtoReflect.Descriptor instead.
func (*CreateDatumSetsTask) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDatumSetsTask) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *CreateDatumSetsTask) GetPathRange() *pfs.PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

func (x *CreateDatumSetsTask) GetSetSpec() *datum.SetSpec {
	if x != nil {
		return x.SetSpec
	}
	return nil
}

type CreateDatumSetsTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatumSets []*pfs.PathRange `protobuf:"bytes,1,rep,name=datum_sets,json=datumSets,proto3" json:"datum_sets,omitempty"`
}

func (x *CreateDatumSetsTaskResult) Reset() {
	*x = CreateDatumSetsTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatumSetsTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatumSetsTaskResult) ProtoMessage() {}

func (x *CreateDatumSetsTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatumSetsTaskResult.ProtoReflect.Descriptor instead.
func (*CreateDatumSetsTaskResult) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDatumSetsTaskResult) GetDatumSets() []*pfs.PathRange {
	if x != nil {
		return x.DatumSets
	}
	return nil
}

type DatumSetTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job          *pps.Job       `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	FileSetId    string         `protobuf:"bytes,2,opt,name=file_set_id,json=fileSetId,proto3" json:"file_set_id,omitempty"`
	PathRange    *pfs.PathRange `protobuf:"bytes,3,opt,name=path_range,json=pathRange,proto3" json:"path_range,omitempty"`
	OutputCommit *pfs.Commit    `protobuf:"bytes,4,opt,name=output_commit,json=outputCommit,proto3" json:"output_commit,omitempty"`
}

func (x *DatumSetTask) Reset() {
	*x = DatumSetTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumSetTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumSetTask) ProtoMessage() {}

func (x *DatumSetTask) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumSetTask.ProtoReflect.Descriptor instead.
func (*DatumSetTask) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{6}
}

func (x *DatumSetTask) GetJob() *pps.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *DatumSetTask) GetFileSetId() string {
	if x != nil {
		return x.FileSetId
	}
	return ""
}

func (x *DatumSetTask) GetPathRange() *pfs.PathRange {
	if x != nil {
		return x.PathRange
	}
	return nil
}

func (x *DatumSetTask) GetOutputCommit() *pfs.Commit {
	if x != nil {
		return x.OutputCommit
	}
	return nil
}

type DatumSetTaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputFileSetId string       `protobuf:"bytes,1,opt,name=output_file_set_id,json=outputFileSetId,proto3" json:"output_file_set_id,omitempty"`
	MetaFileSetId   string       `protobuf:"bytes,2,opt,name=meta_file_set_id,json=metaFileSetId,proto3" json:"meta_file_set_id,omitempty"`
	Stats           *datum.Stats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *DatumSetTaskResult) Reset() {
	*x = DatumSetTaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatumSetTaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatumSetTaskResult) ProtoMessage() {}

func (x *DatumSetTaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_worker_pipeline_transform_transform_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatumSetTaskResult.ProtoReflect.Descriptor instead.
func (*DatumSetTaskResult) Descriptor() ([]byte, []int) {
	return file_server_worker_pipeline_transform_transform_proto_rawDescGZIP(), []int{7}
}

func (x *DatumSetTaskResult) GetOutputFileSetId() string {
	if x != nil {
		return x.OutputFileSetId
	}
	return ""
}

func (x *DatumSetTaskResult) GetMetaFileSetId() string {
	if x != nil {
		return x.MetaFileSetId
	}
	return ""
}

func (x *DatumSetTaskResult) GetStats() *datum.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

var File_server_worker_pipeline_transform_transform_proto protoreflect.FileDescriptor

var file_server_worker_pipeline_transform_transform_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x0d, 0x70, 0x66, 0x73, 0x2f, 0x70, 0x66, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x70, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x73, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x70, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x22, 0x64, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6c,
	0x6c, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x73, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x70, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f,
	0x62, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x0e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x6e, 0x6f, 0x53, 0x6b, 0x69, 0x70, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x1c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x75, 0x6d,
	0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x19, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22,
	0x92, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x53,
	0x65, 0x74, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x66,
	0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09,
	0x70, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x61,
	0x74, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x73, 0x65, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x4d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x75, 0x6d, 0x53, 0x65, 0x74, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x30, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x53,
	0x65, 0x74, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x53, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x70, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x0c, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x44,
	0x61, 0x74, 0x75, 0x6d, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2b, 0x0a, 0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64,
	0x65, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_worker_pipeline_transform_transform_proto_rawDescOnce sync.Once
	file_server_worker_pipeline_transform_transform_proto_rawDescData = file_server_worker_pipeline_transform_transform_proto_rawDesc
)

func file_server_worker_pipeline_transform_transform_proto_rawDescGZIP() []byte {
	file_server_worker_pipeline_transform_transform_proto_rawDescOnce.Do(func() {
		file_server_worker_pipeline_transform_transform_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_worker_pipeline_transform_transform_proto_rawDescData)
	})
	return file_server_worker_pipeline_transform_transform_proto_rawDescData
}

var file_server_worker_pipeline_transform_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_worker_pipeline_transform_transform_proto_goTypes = []interface{}{
	(*CreateParallelDatumsTask)(nil),       // 0: pachyderm.worker.pipeline.transform.CreateParallelDatumsTask
	(*CreateParallelDatumsTaskResult)(nil), // 1: pachyderm.worker.pipeline.transform.CreateParallelDatumsTaskResult
	(*CreateSerialDatumsTask)(nil),         // 2: pachyderm.worker.pipeline.transform.CreateSerialDatumsTask
	(*CreateSerialDatumsTaskResult)(nil),   // 3: pachyderm.worker.pipeline.transform.CreateSerialDatumsTaskResult
	(*CreateDatumSetsTask)(nil),            // 4: pachyderm.worker.pipeline.transform.CreateDatumSetsTask
	(*CreateDatumSetsTaskResult)(nil),      // 5: pachyderm.worker.pipeline.transform.CreateDatumSetsTaskResult
	(*DatumSetTask)(nil),                   // 6: pachyderm.worker.pipeline.transform.DatumSetTask
	(*DatumSetTaskResult)(nil),             // 7: pachyderm.worker.pipeline.transform.DatumSetTaskResult
	(*pps.Job)(nil),                        // 8: pps_v2.Job
	(*pfs.PathRange)(nil),                  // 9: pfs_v2.PathRange
	(*datum.Stats)(nil),                    // 10: datum.Stats
	(*pfs.Commit)(nil),                     // 11: pfs_v2.Commit
	(*datum.SetSpec)(nil),                  // 12: datum.SetSpec
}
var file_server_worker_pipeline_transform_transform_proto_depIdxs = []int32{
	8,  // 0: pachyderm.worker.pipeline.transform.CreateParallelDatumsTask.job:type_name -> pps_v2.Job
	9,  // 1: pachyderm.worker.pipeline.transform.CreateParallelDatumsTask.path_range:type_name -> pfs_v2.PathRange
	10, // 2: pachyderm.worker.pipeline.transform.CreateParallelDatumsTaskResult.stats:type_name -> datum.Stats
	8,  // 3: pachyderm.worker.pipeline.transform.CreateSerialDatumsTask.job:type_name -> pps_v2.Job
	11, // 4: pachyderm.worker.pipeline.transform.CreateSerialDatumsTask.base_meta_commit:type_name -> pfs_v2.Commit
	9,  // 5: pachyderm.worker.pipeline.transform.CreateSerialDatumsTask.path_range:type_name -> pfs_v2.PathRange
	10, // 6: pachyderm.worker.pipeline.transform.CreateSerialDatumsTaskResult.stats:type_name -> datum.Stats
	9,  // 7: pachyderm.worker.pipeline.transform.CreateDatumSetsTask.path_range:type_name -> pfs_v2.PathRange
	12, // 8: pachyderm.worker.pipeline.transform.CreateDatumSetsTask.set_spec:type_name -> datum.SetSpec
	9,  // 9: pachyderm.worker.pipeline.transform.CreateDatumSetsTaskResult.datum_sets:type_name -> pfs_v2.PathRange
	8,  // 10: pachyderm.worker.pipeline.transform.DatumSetTask.job:type_name -> pps_v2.Job
	9,  // 11: pachyderm.worker.pipeline.transform.DatumSetTask.path_range:type_name -> pfs_v2.PathRange
	11, // 12: pachyderm.worker.pipeline.transform.DatumSetTask.output_commit:type_name -> pfs_v2.Commit
	10, // 13: pachyderm.worker.pipeline.transform.DatumSetTaskResult.stats:type_name -> datum.Stats
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_server_worker_pipeline_transform_transform_proto_init() }
func file_server_worker_pipeline_transform_transform_proto_init() {
	if File_server_worker_pipeline_transform_transform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_worker_pipeline_transform_transform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParallelDatumsTask); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParallelDatumsTaskResult); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSerialDatumsTask); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSerialDatumsTaskResult); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatumSetsTask); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatumSetsTaskResult); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumSetTask); i {
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
		file_server_worker_pipeline_transform_transform_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatumSetTaskResult); i {
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
			RawDescriptor: file_server_worker_pipeline_transform_transform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_worker_pipeline_transform_transform_proto_goTypes,
		DependencyIndexes: file_server_worker_pipeline_transform_transform_proto_depIdxs,
		MessageInfos:      file_server_worker_pipeline_transform_transform_proto_msgTypes,
	}.Build()
	File_server_worker_pipeline_transform_transform_proto = out.File
	file_server_worker_pipeline_transform_transform_proto_rawDesc = nil
	file_server_worker_pipeline_transform_transform_proto_goTypes = nil
	file_server_worker_pipeline_transform_transform_proto_depIdxs = nil
}
