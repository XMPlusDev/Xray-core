// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: transport/internet/splithttp/config.proto

package splithttp

import (
	internet "github.com/xmplusdev/xray-core/v24/transport/internet"
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

type RangeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   int32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *RangeConfig) Reset() {
	*x = RangeConfig{}
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RangeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeConfig) ProtoMessage() {}

func (x *RangeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeConfig.ProtoReflect.Descriptor instead.
func (*RangeConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_splithttp_config_proto_rawDescGZIP(), []int{0}
}

func (x *RangeConfig) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *RangeConfig) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

type XmuxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxConcurrency   *RangeConfig `protobuf:"bytes,1,opt,name=maxConcurrency,proto3" json:"maxConcurrency,omitempty"`
	MaxConnections   *RangeConfig `protobuf:"bytes,2,opt,name=maxConnections,proto3" json:"maxConnections,omitempty"`
	CMaxReuseTimes   *RangeConfig `protobuf:"bytes,3,opt,name=cMaxReuseTimes,proto3" json:"cMaxReuseTimes,omitempty"`
	CMaxLifetimeMs   *RangeConfig `protobuf:"bytes,4,opt,name=cMaxLifetimeMs,proto3" json:"cMaxLifetimeMs,omitempty"`
	HMaxRequestTimes *RangeConfig `protobuf:"bytes,5,opt,name=hMaxRequestTimes,proto3" json:"hMaxRequestTimes,omitempty"`
	HKeepAlivePeriod int64        `protobuf:"varint,6,opt,name=hKeepAlivePeriod,proto3" json:"hKeepAlivePeriod,omitempty"`
}

func (x *XmuxConfig) Reset() {
	*x = XmuxConfig{}
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *XmuxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XmuxConfig) ProtoMessage() {}

func (x *XmuxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XmuxConfig.ProtoReflect.Descriptor instead.
func (*XmuxConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_splithttp_config_proto_rawDescGZIP(), []int{1}
}

func (x *XmuxConfig) GetMaxConcurrency() *RangeConfig {
	if x != nil {
		return x.MaxConcurrency
	}
	return nil
}

func (x *XmuxConfig) GetMaxConnections() *RangeConfig {
	if x != nil {
		return x.MaxConnections
	}
	return nil
}

func (x *XmuxConfig) GetCMaxReuseTimes() *RangeConfig {
	if x != nil {
		return x.CMaxReuseTimes
	}
	return nil
}

func (x *XmuxConfig) GetCMaxLifetimeMs() *RangeConfig {
	if x != nil {
		return x.CMaxLifetimeMs
	}
	return nil
}

func (x *XmuxConfig) GetHMaxRequestTimes() *RangeConfig {
	if x != nil {
		return x.HMaxRequestTimes
	}
	return nil
}

func (x *XmuxConfig) GetHKeepAlivePeriod() int64 {
	if x != nil {
		return x.HKeepAlivePeriod
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                 string                 `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                 string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Mode                 string                 `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Headers              map[string]string      `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XPaddingBytes        *RangeConfig           `protobuf:"bytes,5,opt,name=xPaddingBytes,proto3" json:"xPaddingBytes,omitempty"`
	NoGRPCHeader         bool                   `protobuf:"varint,6,opt,name=noGRPCHeader,proto3" json:"noGRPCHeader,omitempty"`
	NoSSEHeader          bool                   `protobuf:"varint,7,opt,name=noSSEHeader,proto3" json:"noSSEHeader,omitempty"`
	ScMaxEachPostBytes   *RangeConfig           `protobuf:"bytes,8,opt,name=scMaxEachPostBytes,proto3" json:"scMaxEachPostBytes,omitempty"`
	ScMinPostsIntervalMs *RangeConfig           `protobuf:"bytes,9,opt,name=scMinPostsIntervalMs,proto3" json:"scMinPostsIntervalMs,omitempty"`
	ScMaxBufferedPosts   int64                  `protobuf:"varint,10,opt,name=scMaxBufferedPosts,proto3" json:"scMaxBufferedPosts,omitempty"`
	Xmux                 *XmuxConfig            `protobuf:"bytes,11,opt,name=xmux,proto3" json:"xmux,omitempty"`
	DownloadSettings     *internet.StreamConfig `protobuf:"bytes,12,opt,name=downloadSettings,proto3" json:"downloadSettings,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_splithttp_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_splithttp_config_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Config) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Config) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Config) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Config) GetXPaddingBytes() *RangeConfig {
	if x != nil {
		return x.XPaddingBytes
	}
	return nil
}

func (x *Config) GetNoGRPCHeader() bool {
	if x != nil {
		return x.NoGRPCHeader
	}
	return false
}

func (x *Config) GetNoSSEHeader() bool {
	if x != nil {
		return x.NoSSEHeader
	}
	return false
}

func (x *Config) GetScMaxEachPostBytes() *RangeConfig {
	if x != nil {
		return x.ScMaxEachPostBytes
	}
	return nil
}

func (x *Config) GetScMinPostsIntervalMs() *RangeConfig {
	if x != nil {
		return x.ScMinPostsIntervalMs
	}
	return nil
}

func (x *Config) GetScMaxBufferedPosts() int64 {
	if x != nil {
		return x.ScMaxBufferedPosts
	}
	return 0
}

func (x *Config) GetXmux() *XmuxConfig {
	if x != nil {
		return x.Xmux
	}
	return nil
}

func (x *Config) GetDownloadSettings() *internet.StreamConfig {
	if x != nil {
		return x.DownloadSettings
	}
	return nil
}

var File_transport_internet_splithttp_config_proto protoreflect.FileDescriptor

var file_transport_internet_splithttp_config_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x1f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x74, 0x6f, 0x22, 0xf4, 0x03, 0x0a, 0x0a, 0x58, 0x6d, 0x75, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x56, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x56, 0x0a, 0x0e, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x75, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x63, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0e, 0x63, 0x4d, 0x61,
	0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0e, 0x63, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x4d,
	0x73, 0x12, 0x5a, 0x0a, 0x10, 0x68, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72,
	0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x68, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x10, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xf8, 0x05, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x50, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x78, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x78, 0x50, 0x61, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x47, 0x52,
	0x50, 0x43, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x6e, 0x6f, 0x47, 0x52, 0x50, 0x43, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x6f, 0x53, 0x53, 0x45, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x53, 0x53, 0x45, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x5e,
	0x0a, 0x12, 0x73, 0x63, 0x4d, 0x61, 0x78, 0x45, 0x61, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x73, 0x63, 0x4d, 0x61,
	0x78, 0x45, 0x61, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x62,
	0x0a, 0x14, 0x73, 0x63, 0x4d, 0x69, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x73, 0x63,
	0x4d, 0x69, 0x6e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x4d, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x63, 0x4d, 0x61, 0x78, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12,
	0x73, 0x63, 0x4d, 0x61, 0x78, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x41, 0x0a, 0x04, 0x78, 0x6d, 0x75, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x58, 0x6d, 0x75, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x04, 0x78, 0x6d, 0x75, 0x78, 0x12, 0x51, 0x0a, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x85, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0x50, 0x01,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c,
	0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x73,
	0x70, 0x6c, 0x69, 0x74, 0x68, 0x74, 0x74, 0x70, 0xaa, 0x02, 0x21, 0x58, 0x72, 0x61, 0x79, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_splithttp_config_proto_rawDescOnce sync.Once
	file_transport_internet_splithttp_config_proto_rawDescData = file_transport_internet_splithttp_config_proto_rawDesc
)

func file_transport_internet_splithttp_config_proto_rawDescGZIP() []byte {
	file_transport_internet_splithttp_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_splithttp_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_splithttp_config_proto_rawDescData)
	})
	return file_transport_internet_splithttp_config_proto_rawDescData
}

var file_transport_internet_splithttp_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_transport_internet_splithttp_config_proto_goTypes = []any{
	(*RangeConfig)(nil),           // 0: xray.transport.internet.splithttp.RangeConfig
	(*XmuxConfig)(nil),            // 1: xray.transport.internet.splithttp.XmuxConfig
	(*Config)(nil),                // 2: xray.transport.internet.splithttp.Config
	nil,                           // 3: xray.transport.internet.splithttp.Config.HeadersEntry
	(*internet.StreamConfig)(nil), // 4: xray.transport.internet.StreamConfig
}
var file_transport_internet_splithttp_config_proto_depIdxs = []int32{
	0,  // 0: xray.transport.internet.splithttp.XmuxConfig.maxConcurrency:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 1: xray.transport.internet.splithttp.XmuxConfig.maxConnections:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 2: xray.transport.internet.splithttp.XmuxConfig.cMaxReuseTimes:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 3: xray.transport.internet.splithttp.XmuxConfig.cMaxLifetimeMs:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 4: xray.transport.internet.splithttp.XmuxConfig.hMaxRequestTimes:type_name -> xray.transport.internet.splithttp.RangeConfig
	3,  // 5: xray.transport.internet.splithttp.Config.headers:type_name -> xray.transport.internet.splithttp.Config.HeadersEntry
	0,  // 6: xray.transport.internet.splithttp.Config.xPaddingBytes:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 7: xray.transport.internet.splithttp.Config.scMaxEachPostBytes:type_name -> xray.transport.internet.splithttp.RangeConfig
	0,  // 8: xray.transport.internet.splithttp.Config.scMinPostsIntervalMs:type_name -> xray.transport.internet.splithttp.RangeConfig
	1,  // 9: xray.transport.internet.splithttp.Config.xmux:type_name -> xray.transport.internet.splithttp.XmuxConfig
	4,  // 10: xray.transport.internet.splithttp.Config.downloadSettings:type_name -> xray.transport.internet.StreamConfig
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_transport_internet_splithttp_config_proto_init() }
func file_transport_internet_splithttp_config_proto_init() {
	if File_transport_internet_splithttp_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_internet_splithttp_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_splithttp_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_splithttp_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_splithttp_config_proto_msgTypes,
	}.Build()
	File_transport_internet_splithttp_config_proto = out.File
	file_transport_internet_splithttp_config_proto_rawDesc = nil
	file_transport_internet_splithttp_config_proto_goTypes = nil
	file_transport_internet_splithttp_config_proto_depIdxs = nil
}
