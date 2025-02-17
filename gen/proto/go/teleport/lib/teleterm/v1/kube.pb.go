//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: teleport/lib/teleterm/v1/kube.proto

package teletermv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Kube describes connected Kubernetes cluster
type Kube struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// uri is the kube resource URI
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// name is the kube name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// labels is the kube labels
	Labels        []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Kube) Reset() {
	*x = Kube{}
	mi := &file_teleport_lib_teleterm_v1_kube_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Kube) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kube) ProtoMessage() {}

func (x *Kube) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_kube_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kube.ProtoReflect.Descriptor instead.
func (*Kube) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_kube_proto_rawDescGZIP(), []int{0}
}

func (x *Kube) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Kube) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Kube) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

// KubeResource describes a kube_cluster's subresource eg: pods, namespaces, etc.
type KubeResource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// uri is the kube resource URI
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// kind is the kube subresource kind eg: pods, namespace
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// name is the kube resource name eg: pod name, namespace name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// labels is the kube resource labels
	Labels []*Label `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	// cluster is the kube cluster name that this kube resource belongs to
	// eg: the kube cluster that a namespace belongs to
	Cluster string `protobuf:"bytes,5,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// namespace is the kube namespace where the resource is located
	// note: this field will be blank if this resource "kind" is "namespace",
	// refer to field "name" for the name of namespace
	Namespace     string `protobuf:"bytes,6,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubeResource) Reset() {
	*x = KubeResource{}
	mi := &file_teleport_lib_teleterm_v1_kube_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubeResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeResource) ProtoMessage() {}

func (x *KubeResource) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_lib_teleterm_v1_kube_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeResource.ProtoReflect.Descriptor instead.
func (*KubeResource) Descriptor() ([]byte, []int) {
	return file_teleport_lib_teleterm_v1_kube_proto_rawDescGZIP(), []int{1}
}

func (x *KubeResource) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *KubeResource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *KubeResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubeResource) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *KubeResource) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *KubeResource) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_teleport_lib_teleterm_v1_kube_proto protoreflect.FileDescriptor

var file_teleport_lib_teleterm_v1_kube_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x24, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x04, 0x4b, 0x75, 0x62, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6c,
	0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xb9, 0x01, 0x0a,
	0x0c, 0x4b, 0x75, 0x62, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x6c, 0x69, 0x62, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_lib_teleterm_v1_kube_proto_rawDescOnce sync.Once
	file_teleport_lib_teleterm_v1_kube_proto_rawDescData []byte
)

func file_teleport_lib_teleterm_v1_kube_proto_rawDescGZIP() []byte {
	file_teleport_lib_teleterm_v1_kube_proto_rawDescOnce.Do(func() {
		file_teleport_lib_teleterm_v1_kube_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_lib_teleterm_v1_kube_proto_rawDesc), len(file_teleport_lib_teleterm_v1_kube_proto_rawDesc)))
	})
	return file_teleport_lib_teleterm_v1_kube_proto_rawDescData
}

var file_teleport_lib_teleterm_v1_kube_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_lib_teleterm_v1_kube_proto_goTypes = []any{
	(*Kube)(nil),         // 0: teleport.lib.teleterm.v1.Kube
	(*KubeResource)(nil), // 1: teleport.lib.teleterm.v1.KubeResource
	(*Label)(nil),        // 2: teleport.lib.teleterm.v1.Label
}
var file_teleport_lib_teleterm_v1_kube_proto_depIdxs = []int32{
	2, // 0: teleport.lib.teleterm.v1.Kube.labels:type_name -> teleport.lib.teleterm.v1.Label
	2, // 1: teleport.lib.teleterm.v1.KubeResource.labels:type_name -> teleport.lib.teleterm.v1.Label
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teleport_lib_teleterm_v1_kube_proto_init() }
func file_teleport_lib_teleterm_v1_kube_proto_init() {
	if File_teleport_lib_teleterm_v1_kube_proto != nil {
		return
	}
	file_teleport_lib_teleterm_v1_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_lib_teleterm_v1_kube_proto_rawDesc), len(file_teleport_lib_teleterm_v1_kube_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_lib_teleterm_v1_kube_proto_goTypes,
		DependencyIndexes: file_teleport_lib_teleterm_v1_kube_proto_depIdxs,
		MessageInfos:      file_teleport_lib_teleterm_v1_kube_proto_msgTypes,
	}.Build()
	File_teleport_lib_teleterm_v1_kube_proto = out.File
	file_teleport_lib_teleterm_v1_kube_proto_goTypes = nil
	file_teleport_lib_teleterm_v1_kube_proto_depIdxs = nil
}
