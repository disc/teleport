// Copyright 2024 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: teleport/kubewaitingcontainer/v1/kubewaitingcontainer.proto

package kubewaitingcontainerv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
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

// KubernetesWaitingContainer is a Kubernetes pod that has ephemeral containers
// waiting to be created until moderated session requirements are met.
type KubernetesWaitingContainer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// kind is a resource kind
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// sub_kind is an optional resource sub kind, used in some resources
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// version is the resource version. It must be specified.
	// Supported values are: `v1`.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// metadata is resource metadata
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec is the Kubernetes waiting container spec.
	Spec          *KubernetesWaitingContainerSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubernetesWaitingContainer) Reset() {
	*x = KubernetesWaitingContainer{}
	mi := &file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubernetesWaitingContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesWaitingContainer) ProtoMessage() {}

func (x *KubernetesWaitingContainer) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesWaitingContainer.ProtoReflect.Descriptor instead.
func (*KubernetesWaitingContainer) Descriptor() ([]byte, []int) {
	return file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesWaitingContainer) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *KubernetesWaitingContainer) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *KubernetesWaitingContainer) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *KubernetesWaitingContainer) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *KubernetesWaitingContainer) GetSpec() *KubernetesWaitingContainerSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// KubernetesWaitingContainerSpec is the Kubernetes waiting ephemeral container spec.
type KubernetesWaitingContainerSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the Teleport user that attempted to create the container
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// cluster is the Kubernetes cluster of this container
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// namespace is the Kubernetes namespace of this container
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// pod_name is the name of the parent pod
	PodName string `protobuf:"bytes,4,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	// container_name is the name of the ephemeral container
	ContainerName string `protobuf:"bytes,5,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// patch is the patch that should be applied to the parent pod
	// to create this ephemeral container
	Patch []byte `protobuf:"bytes,6,opt,name=patch,proto3" json:"patch,omitempty"`
	// patch_type identifies the patch model to be applied.
	PatchType     string `protobuf:"bytes,7,opt,name=patch_type,json=patchType,proto3" json:"patch_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubernetesWaitingContainerSpec) Reset() {
	*x = KubernetesWaitingContainerSpec{}
	mi := &file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubernetesWaitingContainerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesWaitingContainerSpec) ProtoMessage() {}

func (x *KubernetesWaitingContainerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesWaitingContainerSpec.ProtoReflect.Descriptor instead.
func (*KubernetesWaitingContainerSpec) Descriptor() ([]byte, []int) {
	return file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesWaitingContainerSpec) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *KubernetesWaitingContainerSpec) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *KubernetesWaitingContainerSpec) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesWaitingContainerSpec) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *KubernetesWaitingContainerSpec) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *KubernetesWaitingContainerSpec) GetPatch() []byte {
	if x != nil {
		return x.Patch
	}
	return nil
}

func (x *KubernetesWaitingContainerSpec) GetPatchType() string {
	if x != nil {
		return x.PatchType
	}
	return ""
}

var File_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto protoreflect.FileDescriptor

var file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x77,
	0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x77, 0x61, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x1a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x40, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xeb, 0x01, 0x0a, 0x1e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x42, 0x6c, 0x5a, 0x6a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x77, 0x61, 0x69, 0x74,
	0x69, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x6b, 0x75, 0x62, 0x65, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescOnce sync.Once
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescData = file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDesc
)

func file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescGZIP() []byte {
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescOnce.Do(func() {
		file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescData)
	})
	return file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDescData
}

var file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_goTypes = []any{
	(*KubernetesWaitingContainer)(nil),     // 0: teleport.kubewaitingcontainer.v1.KubernetesWaitingContainer
	(*KubernetesWaitingContainerSpec)(nil), // 1: teleport.kubewaitingcontainer.v1.KubernetesWaitingContainerSpec
	(*v1.Metadata)(nil),                    // 2: teleport.header.v1.Metadata
}
var file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_depIdxs = []int32{
	2, // 0: teleport.kubewaitingcontainer.v1.KubernetesWaitingContainer.metadata:type_name -> teleport.header.v1.Metadata
	1, // 1: teleport.kubewaitingcontainer.v1.KubernetesWaitingContainer.spec:type_name -> teleport.kubewaitingcontainer.v1.KubernetesWaitingContainerSpec
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_init() }
func file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_init() {
	if File_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_goTypes,
		DependencyIndexes: file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_depIdxs,
		MessageInfos:      file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_msgTypes,
	}.Build()
	File_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto = out.File
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_rawDesc = nil
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_goTypes = nil
	file_teleport_kubewaitingcontainer_v1_kubewaitingcontainer_proto_depIdxs = nil
}
