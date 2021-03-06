// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/metadata/v3/metadata.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
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

// MetadataKey provides a general interface using `key` and `path` to retrieve value from
// :ref:`Metadata <envoy_api_msg_config.core.v3.Metadata>`.
//
// For example, for the following Metadata:
//
// .. code-block:: yaml
//
//    filter_metadata:
//      envoy.xxx:
//        prop:
//          foo: bar
//          xyz:
//            hello: envoy
//
// The following MetadataKey will retrieve a string value "bar" from the Metadata.
//
// .. code-block:: yaml
//
//    key: envoy.xxx
//    path:
//    - key: prop
//    - key: foo
//
type MetadataKey struct {
	// The key name of Metadata to retrieve the Struct from the metadata.
	// Typically, it represents a builtin subsystem or custom extension.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The path to retrieve the Value from the Struct. It can be a prefix or a full path,
	// e.g. ``[prop, xyz]`` for a struct or ``[prop, foo]`` for a string in the example,
	// which depends on the particular scenario.
	//
	// Note: Due to that only the key type segment is supported, the path can not specify a list
	// unless the list is the last segment.
	Path                 []*MetadataKey_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MetadataKey) Reset()         { *m = MetadataKey{} }
func (m *MetadataKey) String() string { return proto.CompactTextString(m) }
func (*MetadataKey) ProtoMessage()    {}
func (*MetadataKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{0}
}
func (m *MetadataKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKey.Unmarshal(m, b)
}
func (m *MetadataKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKey.Marshal(b, m, deterministic)
}
func (m *MetadataKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKey.Merge(m, src)
}
func (m *MetadataKey) XXX_Size() int {
	return xxx_messageInfo_MetadataKey.Size(m)
}
func (m *MetadataKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKey.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKey proto.InternalMessageInfo

func (m *MetadataKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MetadataKey) GetPath() []*MetadataKey_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Metadata.
// Currently it is only supported to specify the key, i.e. field name, as one segment of a path.
type MetadataKey_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataKey_PathSegment_Key
	Segment              isMetadataKey_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MetadataKey_PathSegment) Reset()         { *m = MetadataKey_PathSegment{} }
func (m *MetadataKey_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataKey_PathSegment) ProtoMessage()    {}
func (*MetadataKey_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{0, 0}
}
func (m *MetadataKey_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKey_PathSegment.Unmarshal(m, b)
}
func (m *MetadataKey_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKey_PathSegment.Marshal(b, m, deterministic)
}
func (m *MetadataKey_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKey_PathSegment.Merge(m, src)
}
func (m *MetadataKey_PathSegment) XXX_Size() int {
	return xxx_messageInfo_MetadataKey_PathSegment.Size(m)
}
func (m *MetadataKey_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKey_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKey_PathSegment proto.InternalMessageInfo

type isMetadataKey_PathSegment_Segment interface {
	isMetadataKey_PathSegment_Segment()
	Equal(interface{}) bool
}

type MetadataKey_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
}

func (*MetadataKey_PathSegment_Key) isMetadataKey_PathSegment_Segment() {}

func (m *MetadataKey_PathSegment) GetSegment() isMetadataKey_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataKey_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataKey_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataKey_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataKey_PathSegment_Key)(nil),
	}
}

// Describes what kind of metadata.
type MetadataKind struct {
	// Types that are valid to be assigned to Kind:
	//	*MetadataKind_Request_
	//	*MetadataKind_Route_
	//	*MetadataKind_Cluster_
	//	*MetadataKind_Host_
	Kind                 isMetadataKind_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MetadataKind) Reset()         { *m = MetadataKind{} }
func (m *MetadataKind) String() string { return proto.CompactTextString(m) }
func (*MetadataKind) ProtoMessage()    {}
func (*MetadataKind) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{1}
}
func (m *MetadataKind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKind.Unmarshal(m, b)
}
func (m *MetadataKind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKind.Marshal(b, m, deterministic)
}
func (m *MetadataKind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKind.Merge(m, src)
}
func (m *MetadataKind) XXX_Size() int {
	return xxx_messageInfo_MetadataKind.Size(m)
}
func (m *MetadataKind) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKind.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKind proto.InternalMessageInfo

type isMetadataKind_Kind interface {
	isMetadataKind_Kind()
	Equal(interface{}) bool
}

type MetadataKind_Request_ struct {
	Request *MetadataKind_Request `protobuf:"bytes,1,opt,name=request,proto3,oneof" json:"request,omitempty"`
}
type MetadataKind_Route_ struct {
	Route *MetadataKind_Route `protobuf:"bytes,2,opt,name=route,proto3,oneof" json:"route,omitempty"`
}
type MetadataKind_Cluster_ struct {
	Cluster *MetadataKind_Cluster `protobuf:"bytes,3,opt,name=cluster,proto3,oneof" json:"cluster,omitempty"`
}
type MetadataKind_Host_ struct {
	Host *MetadataKind_Host `protobuf:"bytes,4,opt,name=host,proto3,oneof" json:"host,omitempty"`
}

func (*MetadataKind_Request_) isMetadataKind_Kind() {}
func (*MetadataKind_Route_) isMetadataKind_Kind()   {}
func (*MetadataKind_Cluster_) isMetadataKind_Kind() {}
func (*MetadataKind_Host_) isMetadataKind_Kind()    {}

func (m *MetadataKind) GetKind() isMetadataKind_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *MetadataKind) GetRequest() *MetadataKind_Request {
	if x, ok := m.GetKind().(*MetadataKind_Request_); ok {
		return x.Request
	}
	return nil
}

func (m *MetadataKind) GetRoute() *MetadataKind_Route {
	if x, ok := m.GetKind().(*MetadataKind_Route_); ok {
		return x.Route
	}
	return nil
}

func (m *MetadataKind) GetCluster() *MetadataKind_Cluster {
	if x, ok := m.GetKind().(*MetadataKind_Cluster_); ok {
		return x.Cluster
	}
	return nil
}

func (m *MetadataKind) GetHost() *MetadataKind_Host {
	if x, ok := m.GetKind().(*MetadataKind_Host_); ok {
		return x.Host
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataKind) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataKind_Request_)(nil),
		(*MetadataKind_Route_)(nil),
		(*MetadataKind_Cluster_)(nil),
		(*MetadataKind_Host_)(nil),
	}
}

// Represents dynamic metadata associated with the request.
type MetadataKind_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataKind_Request) Reset()         { *m = MetadataKind_Request{} }
func (m *MetadataKind_Request) String() string { return proto.CompactTextString(m) }
func (*MetadataKind_Request) ProtoMessage()    {}
func (*MetadataKind_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{1, 0}
}
func (m *MetadataKind_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKind_Request.Unmarshal(m, b)
}
func (m *MetadataKind_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKind_Request.Marshal(b, m, deterministic)
}
func (m *MetadataKind_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKind_Request.Merge(m, src)
}
func (m *MetadataKind_Request) XXX_Size() int {
	return xxx_messageInfo_MetadataKind_Request.Size(m)
}
func (m *MetadataKind_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKind_Request.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKind_Request proto.InternalMessageInfo

// Represents metadata from :ref:`the route<envoy_api_field_config.route.v3.Route.metadata>`.
type MetadataKind_Route struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataKind_Route) Reset()         { *m = MetadataKind_Route{} }
func (m *MetadataKind_Route) String() string { return proto.CompactTextString(m) }
func (*MetadataKind_Route) ProtoMessage()    {}
func (*MetadataKind_Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{1, 1}
}
func (m *MetadataKind_Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKind_Route.Unmarshal(m, b)
}
func (m *MetadataKind_Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKind_Route.Marshal(b, m, deterministic)
}
func (m *MetadataKind_Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKind_Route.Merge(m, src)
}
func (m *MetadataKind_Route) XXX_Size() int {
	return xxx_messageInfo_MetadataKind_Route.Size(m)
}
func (m *MetadataKind_Route) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKind_Route.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKind_Route proto.InternalMessageInfo

// Represents metadata from :ref:`the upstream cluster<envoy_api_field_config.cluster.v3.Cluster.metadata>`.
type MetadataKind_Cluster struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataKind_Cluster) Reset()         { *m = MetadataKind_Cluster{} }
func (m *MetadataKind_Cluster) String() string { return proto.CompactTextString(m) }
func (*MetadataKind_Cluster) ProtoMessage()    {}
func (*MetadataKind_Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{1, 2}
}
func (m *MetadataKind_Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKind_Cluster.Unmarshal(m, b)
}
func (m *MetadataKind_Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKind_Cluster.Marshal(b, m, deterministic)
}
func (m *MetadataKind_Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKind_Cluster.Merge(m, src)
}
func (m *MetadataKind_Cluster) XXX_Size() int {
	return xxx_messageInfo_MetadataKind_Cluster.Size(m)
}
func (m *MetadataKind_Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKind_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKind_Cluster proto.InternalMessageInfo

// Represents metadata from :ref:`the upstream
// host<envoy_api_field_config.endpoint.v3.LbEndpoint.metadata>`.
type MetadataKind_Host struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetadataKind_Host) Reset()         { *m = MetadataKind_Host{} }
func (m *MetadataKind_Host) String() string { return proto.CompactTextString(m) }
func (*MetadataKind_Host) ProtoMessage()    {}
func (*MetadataKind_Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf0001569456736, []int{1, 3}
}
func (m *MetadataKind_Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataKind_Host.Unmarshal(m, b)
}
func (m *MetadataKind_Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataKind_Host.Marshal(b, m, deterministic)
}
func (m *MetadataKind_Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataKind_Host.Merge(m, src)
}
func (m *MetadataKind_Host) XXX_Size() int {
	return xxx_messageInfo_MetadataKind_Host.Size(m)
}
func (m *MetadataKind_Host) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataKind_Host.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataKind_Host proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetadataKey)(nil), "envoy.type.metadata.v3.MetadataKey")
	proto.RegisterType((*MetadataKey_PathSegment)(nil), "envoy.type.metadata.v3.MetadataKey.PathSegment")
	proto.RegisterType((*MetadataKind)(nil), "envoy.type.metadata.v3.MetadataKind")
	proto.RegisterType((*MetadataKind_Request)(nil), "envoy.type.metadata.v3.MetadataKind.Request")
	proto.RegisterType((*MetadataKind_Route)(nil), "envoy.type.metadata.v3.MetadataKind.Route")
	proto.RegisterType((*MetadataKind_Cluster)(nil), "envoy.type.metadata.v3.MetadataKind.Cluster")
	proto.RegisterType((*MetadataKind_Host)(nil), "envoy.type.metadata.v3.MetadataKind.Host")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/metadata/v3/metadata.proto", fileDescriptor_baf0001569456736)
}

var fileDescriptor_baf0001569456736 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x34, 0xdd, 0xae, 0xaf, 0x2a, 0x12, 0x44, 0x6b, 0xc4, 0xa5, 0xd6, 0x15, 0xda,
	0xae, 0x9b, 0x59, 0x5a, 0x44, 0x2c, 0x88, 0x10, 0x2f, 0x05, 0x29, 0x94, 0x7a, 0x11, 0x6f, 0xb3,
	0xed, 0x90, 0x8e, 0x6d, 0xe7, 0x65, 0x33, 0x93, 0xb2, 0xbd, 0x79, 0xf4, 0xbc, 0xe0, 0xc5, 0x4f,
	0xe0, 0x67, 0xf0, 0x2e, 0x78, 0xd5, 0x8f, 0xe0, 0xa7, 0x90, 0x3d, 0xc9, 0x4c, 0xd2, 0x52, 0xb0,
	0x4b, 0xb3, 0xb7, 0x97, 0x79, 0xf3, 0xfb, 0xbf, 0xf7, 0xfe, 0x4c, 0x1e, 0xbc, 0x0f, 0xb9, 0x9a,
	0x24, 0xa7, 0xfe, 0x08, 0xe7, 0x44, 0xe2, 0x0c, 0x8f, 0x39, 0x92, 0x70, 0x86, 0x48, 0xa2, 0x18,
	0x3f, 0xb2, 0x91, 0x92, 0xe9, 0x17, 0x8d, 0x38, 0x61, 0xe7, 0x8a, 0xc5, 0x82, 0xce, 0x08, 0x13,
	0x0b, 0x5c, 0x12, 0xb5, 0x8c, 0x18, 0x99, 0x33, 0x45, 0xc7, 0x54, 0x51, 0xb2, 0xe8, 0xac, 0x63,
	0x3f, 0x8a, 0x51, 0xa1, 0x7b, 0xcf, 0x5c, 0xf3, 0xf5, 0x35, 0x7f, 0x9d, 0x5a, 0x74, 0xbc, 0x47,
	0xc9, 0x38, 0xa2, 0x84, 0x0a, 0x81, 0x8a, 0x2a, 0x8e, 0x42, 0x12, 0xa9, 0xa8, 0x4a, 0x64, 0x8a,
	0x79, 0x8f, 0xff, 0x4b, 0x2f, 0x58, 0x2c, 0x39, 0x0a, 0x2e, 0xc2, 0xec, 0xca, 0xfd, 0x05, 0x9d,
	0xf1, 0x31, 0x55, 0x8c, 0xac, 0x82, 0x2c, 0x71, 0x37, 0xc4, 0x10, 0x4d, 0x48, 0x74, 0x94, 0x9e,
	0xd6, 0xbf, 0xd8, 0x50, 0xe9, 0x67, 0x0d, 0xbc, 0x65, 0x4b, 0xf7, 0x01, 0x14, 0xa7, 0x6c, 0x59,
	0xb5, 0x6a, 0x56, 0xe3, 0x46, 0x50, 0xbe, 0x0c, 0x9c, 0xd8, 0xae, 0x59, 0x43, 0x7d, 0xe6, 0xf6,
	0xc1, 0x89, 0xa8, 0x9a, 0x54, 0xed, 0x5a, 0xb1, 0x51, 0x69, 0x13, 0x7f, 0xfb, 0x08, 0xfe, 0x86,
	0x9a, 0x3f, 0xa0, 0x6a, 0xf2, 0x8e, 0x85, 0x73, 0x26, 0x54, 0xb0, 0x7f, 0x19, 0x94, 0x2e, 0x2c,
	0x7b, 0xdf, 0x1a, 0x1a, 0x19, 0xef, 0x0c, 0x2a, 0x1b, 0x69, 0xf7, 0xe1, 0xb6, 0xc2, 0xbd, 0x82,
	0x29, 0xdd, 0x7d, 0xfe, 0xf5, 0xc7, 0xe7, 0x83, 0x13, 0xd8, 0x5e, 0xb2, 0x7d, 0x65, 0xc9, 0xdb,
	0x50, 0x96, 0x99, 0x7c, 0xf1, 0x6f, 0x60, 0x75, 0x9b, 0x5a, 0xe6, 0x10, 0xea, 0xbb, 0x65, 0xea,
	0xbf, 0x1d, 0xb8, 0xb9, 0xfe, 0xe6, 0x62, 0xec, 0xf6, 0xa0, 0x1c, 0xb3, 0xb3, 0x84, 0x49, 0x65,
	0x7a, 0xac, 0xb4, 0x9f, 0xed, 0x34, 0x80, 0x8b, 0xb1, 0x3f, 0x4c, 0x99, 0x5e, 0x61, 0xb8, 0xc2,
	0xdd, 0x00, 0x4a, 0x31, 0x26, 0x8a, 0x55, 0x6d, 0xa3, 0xd3, 0xca, 0xa7, 0xa3, 0x89, 0x5e, 0x61,
	0x98, 0xa2, 0xba, 0x9b, 0xd1, 0x2c, 0x91, 0x8a, 0xc5, 0xd5, 0xe2, 0x35, 0xba, 0x79, 0x93, 0x32,
	0xba, 0x9b, 0x0c, 0x77, 0x5f, 0x83, 0x33, 0x41, 0xa9, 0xaa, 0x8e, 0x91, 0x69, 0xe6, 0x92, 0xe9,
	0xa1, 0x99, 0xc8, 0x80, 0xde, 0x2b, 0x28, 0x67, 0x43, 0x76, 0xdb, 0xda, 0xdf, 0x63, 0x38, 0xda,
	0xe5, 0xef, 0x86, 0x31, 0xde, 0x4b, 0x28, 0x99, 0xd9, 0xba, 0x27, 0x1a, 0x3e, 0x82, 0x66, 0x2e,
	0x58, 0x13, 0xba, 0x72, 0x36, 0xd0, 0xb5, 0x2a, 0x67, 0x8c, 0xf7, 0x02, 0x1c, 0x3d, 0x48, 0x97,
	0x68, 0xb6, 0x05, 0x8d, 0x3c, 0xac, 0x01, 0x5a, 0x1a, 0x78, 0x0a, 0x4f, 0x72, 0x00, 0x41, 0x05,
	0x9c, 0xa9, 0x7e, 0x3e, 0xfa, 0xfd, 0x05, 0x17, 0xd6, 0xb7, 0x3f, 0x07, 0xd6, 0xf7, 0x4f, 0x3f,
	0x7f, 0xed, 0xd9, 0x77, 0x6c, 0x38, 0xe4, 0x98, 0xda, 0x1d, 0xc5, 0x78, 0xbe, 0xbc, 0xc2, 0xf9,
	0xe0, 0xd6, 0x4a, 0x6f, 0xa0, 0x7f, 0xd8, 0x81, 0xf5, 0xa1, 0x9f, 0x6f, 0x2b, 0x45, 0xd3, 0x30,
	0xcf, 0x66, 0x3a, 0xdd, 0x33, 0x8b, 0xa0, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x21, 0x78, 0xe1,
	0x1f, 0xed, 0x04, 0x00, 0x00,
}

func (this *MetadataKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKey)
	if !ok {
		that2, ok := that.(MetadataKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if len(this.Path) != len(that1.Path) {
		return false
	}
	for i := range this.Path {
		if !this.Path[i].Equal(that1.Path[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKey_PathSegment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKey_PathSegment)
	if !ok {
		that2, ok := that.(MetadataKey_PathSegment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Segment == nil {
		if this.Segment != nil {
			return false
		}
	} else if this.Segment == nil {
		return false
	} else if !this.Segment.Equal(that1.Segment) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKey_PathSegment_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKey_PathSegment_Key)
	if !ok {
		that2, ok := that.(MetadataKey_PathSegment_Key)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *MetadataKind) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind)
	if !ok {
		that2, ok := that.(MetadataKind)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKind_Request_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Request_)
	if !ok {
		that2, ok := that.(MetadataKind_Request_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Request.Equal(that1.Request) {
		return false
	}
	return true
}
func (this *MetadataKind_Route_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Route_)
	if !ok {
		that2, ok := that.(MetadataKind_Route_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Route.Equal(that1.Route) {
		return false
	}
	return true
}
func (this *MetadataKind_Cluster_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Cluster_)
	if !ok {
		that2, ok := that.(MetadataKind_Cluster_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Cluster.Equal(that1.Cluster) {
		return false
	}
	return true
}
func (this *MetadataKind_Host_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Host_)
	if !ok {
		that2, ok := that.(MetadataKind_Host_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Host.Equal(that1.Host) {
		return false
	}
	return true
}
func (this *MetadataKind_Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Request)
	if !ok {
		that2, ok := that.(MetadataKind_Request)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKind_Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Route)
	if !ok {
		that2, ok := that.(MetadataKind_Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKind_Cluster) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Cluster)
	if !ok {
		that2, ok := that.(MetadataKind_Cluster)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetadataKind_Host) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetadataKind_Host)
	if !ok {
		that2, ok := that.(MetadataKind_Host)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
