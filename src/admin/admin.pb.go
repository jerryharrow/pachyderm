// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: admin/admin.proto

package admin

import (
	pfs "github.com/pachyderm/pachyderm/v2/src/pfs"
	versionpb "github.com/pachyderm/pachyderm/v2/src/version/versionpb"
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

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DeploymentId string `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	// True if the server is capable of generating warnings.
	WarningsOk bool `protobuf:"varint,3,opt,name=warnings_ok,json=warningsOk,proto3" json:"warnings_ok,omitempty"`
	// Warnings about the client configuration.
	Warnings []string `protobuf:"bytes,4,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// The configured public URL of Pachyderm.
	ProxyHost string `protobuf:"bytes,5,opt,name=proxy_host,json=proxyHost,proto3" json:"proxy_host,omitempty"`
	// True if Pachyderm is served over TLS (HTTPS).
	ProxyTls bool `protobuf:"varint,6,opt,name=proxy_tls,json=proxyTls,proto3" json:"proxy_tls,omitempty"`
	// True if this pachd is in "paused" mode.
	Paused bool `protobuf:"varint,7,opt,name=paused,proto3" json:"paused,omitempty"`
	// Any HTTP links that the client might want to be aware of.
	WebResources *WebResource `protobuf:"bytes,8,opt,name=web_resources,json=webResources,proto3" json:"web_resources,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClusterInfo) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ClusterInfo) GetWarningsOk() bool {
	if x != nil {
		return x.WarningsOk
	}
	return false
}

func (x *ClusterInfo) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *ClusterInfo) GetProxyHost() string {
	if x != nil {
		return x.ProxyHost
	}
	return ""
}

func (x *ClusterInfo) GetProxyTls() bool {
	if x != nil {
		return x.ProxyTls
	}
	return false
}

func (x *ClusterInfo) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *ClusterInfo) GetWebResources() *WebResource {
	if x != nil {
		return x.WebResources
	}
	return nil
}

type InspectClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the client that's connecting; used by the server to warn about too-old (or
	// too-new!) clients.
	ClientVersion *versionpb.Version `protobuf:"bytes,1,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	// If CurrentProject is set, then InspectCluster will return an error if the
	// project does not exist.
	CurrentProject *pfs.Project `protobuf:"bytes,2,opt,name=current_project,json=currentProject,proto3" json:"current_project,omitempty"`
}

func (x *InspectClusterRequest) Reset() {
	*x = InspectClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InspectClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InspectClusterRequest) ProtoMessage() {}

func (x *InspectClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InspectClusterRequest.ProtoReflect.Descriptor instead.
func (*InspectClusterRequest) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *InspectClusterRequest) GetClientVersion() *versionpb.Version {
	if x != nil {
		return x.ClientVersion
	}
	return nil
}

func (x *InspectClusterRequest) GetCurrentProject() *pfs.Project {
	if x != nil {
		return x.CurrentProject
	}
	return nil
}

// WebResource contains URL prefixes of common HTTP functions.
type WebResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base URL of the archive server; append a filename to this.  Empty if the archive server is
	// not exposed.
	ArchiveDownloadBaseUrl string `protobuf:"bytes,1,opt,name=archive_download_base_url,json=archiveDownloadBaseUrl,proto3" json:"archive_download_base_url,omitempty"`
	// Where to find the CreatePipelineRequest JSON schema; if this server is not accessible via a
	// URL, then a link to Github is provided based on the baked-in version of the server.
	CreatePipelineRequestJsonSchemaUrl string `protobuf:"bytes,2,opt,name=create_pipeline_request_json_schema_url,json=createPipelineRequestJsonSchemaUrl,proto3" json:"create_pipeline_request_json_schema_url,omitempty"`
}

func (x *WebResource) Reset() {
	*x = WebResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebResource) ProtoMessage() {}

func (x *WebResource) ProtoReflect() protoreflect.Message {
	mi := &file_admin_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebResource.ProtoReflect.Descriptor instead.
func (*WebResource) Descriptor() ([]byte, []int) {
	return file_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *WebResource) GetArchiveDownloadBaseUrl() string {
	if x != nil {
		return x.ArchiveDownloadBaseUrl
	}
	return ""
}

func (x *WebResource) GetCreatePipelineRequestJsonSchemaUrl() string {
	if x != nil {
		return x.CreatePipelineRequestJsonSchemaUrl
	}
	return ""
}

var File_admin_admin_proto protoreflect.FileDescriptor

var file_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x32, 0x1a, 0x1f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x70, 0x66, 0x73, 0x2f, 0x70, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02,
	0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x6f,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x73, 0x4f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75,
	0x73, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x76, 0x32, 0x2e, 0x57, 0x65, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22,
	0x8f, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x5f, 0x76, 0x32,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x66, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x57, 0x65, 0x62, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x19, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x53, 0x0a, 0x27,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x72,
	0x6c, 0x32, 0x51, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x4a, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x63,
	0x68, 0x79, 0x64, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_admin_proto_rawDescOnce sync.Once
	file_admin_admin_proto_rawDescData = file_admin_admin_proto_rawDesc
)

func file_admin_admin_proto_rawDescGZIP() []byte {
	file_admin_admin_proto_rawDescOnce.Do(func() {
		file_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_admin_proto_rawDescData)
	})
	return file_admin_admin_proto_rawDescData
}

var file_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_admin_admin_proto_goTypes = []interface{}{
	(*ClusterInfo)(nil),           // 0: admin_v2.ClusterInfo
	(*InspectClusterRequest)(nil), // 1: admin_v2.InspectClusterRequest
	(*WebResource)(nil),           // 2: admin_v2.WebResource
	(*versionpb.Version)(nil),     // 3: versionpb_v2.Version
	(*pfs.Project)(nil),           // 4: pfs_v2.Project
}
var file_admin_admin_proto_depIdxs = []int32{
	2, // 0: admin_v2.ClusterInfo.web_resources:type_name -> admin_v2.WebResource
	3, // 1: admin_v2.InspectClusterRequest.client_version:type_name -> versionpb_v2.Version
	4, // 2: admin_v2.InspectClusterRequest.current_project:type_name -> pfs_v2.Project
	1, // 3: admin_v2.API.InspectCluster:input_type -> admin_v2.InspectClusterRequest
	0, // 4: admin_v2.API.InspectCluster:output_type -> admin_v2.ClusterInfo
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_admin_proto_init() }
func file_admin_admin_proto_init() {
	if File_admin_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
		file_admin_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InspectClusterRequest); i {
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
		file_admin_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebResource); i {
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
			RawDescriptor: file_admin_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_admin_proto_goTypes,
		DependencyIndexes: file_admin_admin_proto_depIdxs,
		MessageInfos:      file_admin_admin_proto_msgTypes,
	}.Build()
	File_admin_admin_proto = out.File
	file_admin_admin_proto_rawDesc = nil
	file_admin_admin_proto_goTypes = nil
	file_admin_admin_proto_depIdxs = nil
}
