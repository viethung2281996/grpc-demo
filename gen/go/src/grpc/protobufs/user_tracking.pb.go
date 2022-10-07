// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: src/grpc/protobufs/user_tracking.proto

package v1

import (
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

type CompressionType int32

const (
	CompressionType_GZIP   CompressionType = 0
	CompressionType_SNAPPY CompressionType = 1
)

// Enum value maps for CompressionType.
var (
	CompressionType_name = map[int32]string{
		0: "GZIP",
		1: "SNAPPY",
	}
	CompressionType_value = map[string]int32{
		"GZIP":   0,
		"SNAPPY": 1,
	}
)

func (x CompressionType) Enum() *CompressionType {
	p := new(CompressionType)
	*p = x
	return p
}

func (x CompressionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_grpc_protobufs_user_tracking_proto_enumTypes[0].Descriptor()
}

func (CompressionType) Type() protoreflect.EnumType {
	return &file_src_grpc_protobufs_user_tracking_proto_enumTypes[0]
}

func (x CompressionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionType.Descriptor instead.
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{0}
}

type RequestCompress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompressionType CompressionType `protobuf:"varint,1,opt,name=compression_type,json=compressionType,proto3,enum=compress.service.v1.CompressionType" json:"compression_type,omitempty"`
	Data            []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestCompress) Reset() {
	*x = RequestCompress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCompress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCompress) ProtoMessage() {}

func (x *RequestCompress) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCompress.ProtoReflect.Descriptor instead.
func (*RequestCompress) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *RequestCompress) GetCompressionType() CompressionType {
	if x != nil {
		return x.CompressionType
	}
	return CompressionType_GZIP
}

func (x *RequestCompress) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseCompress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseCompress) Reset() {
	*x = ResponseCompress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCompress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCompress) ProtoMessage() {}

func (x *ResponseCompress) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCompress.ProtoReflect.Descriptor instead.
func (*ResponseCompress) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseCompress) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Sdk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Source  string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Sdk) Reset() {
	*x = Sdk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sdk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sdk) ProtoMessage() {}

func (x *Sdk) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sdk.ProtoReflect.Descriptor instead.
func (*Sdk) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{2}
}

func (x *Sdk) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Sdk) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Sdk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sdk) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Os struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Os) Reset() {
	*x = Os{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Os) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Os) ProtoMessage() {}

func (x *Os) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Os.ProtoReflect.Descriptor instead.
func (*Os) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{3}
}

func (x *Os) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Os) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cellular  bool   `protobuf:"varint,1,opt,name=cellular,proto3" json:"cellular,omitempty"`
	Bluetooth bool   `protobuf:"varint,2,opt,name=bluetooth,proto3" json:"bluetooth,omitempty"`
	Wifi      bool   `protobuf:"varint,3,opt,name=wifi,proto3" json:"wifi,omitempty"`
	Address   string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{4}
}

func (x *Network) GetCellular() bool {
	if x != nil {
		return x.Cellular
	}
	return false
}

func (x *Network) GetBluetooth() bool {
	if x != nil {
		return x.Bluetooth
	}
	return false
}

func (x *Network) GetWifi() bool {
	if x != nil {
		return x.Wifi
	}
	return false
}

func (x *Network) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Screen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Screen) Reset() {
	*x = Screen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Screen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Screen) ProtoMessage() {}

func (x *Screen) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Screen.ProtoReflect.Descriptor instead.
func (*Screen) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{5}
}

func (x *Screen) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Screen) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Web struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Browser     string `protobuf:"bytes,1,opt,name=browser,proto3" json:"browser,omitempty"`
	Device      string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Version     string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Orientation string `protobuf:"bytes,4,opt,name=orientation,proto3" json:"orientation,omitempty"`
}

func (x *Web) Reset() {
	*x = Web{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Web) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Web) ProtoMessage() {}

func (x *Web) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Web.ProtoReflect.Descriptor instead.
func (*Web) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{6}
}

func (x *Web) GetBrowser() string {
	if x != nil {
		return x.Browser
	}
	return ""
}

func (x *Web) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Web) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Web) GetOrientation() string {
	if x != nil {
		return x.Orientation
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DId       string   `protobuf:"bytes,1,opt,name=d_id,json=dId,proto3" json:"d_id,omitempty"`
	TId       string   `protobuf:"bytes,2,opt,name=t_id,json=tId,proto3" json:"t_id,omitempty"`
	UId       string   `protobuf:"bytes,3,opt,name=u_id,json=uId,proto3" json:"u_id,omitempty"`
	DeviceId  string   `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ProfileId string   `protobuf:"bytes,5,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Type      string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Channel   string   `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	Os        *Os      `protobuf:"bytes,8,opt,name=os,proto3" json:"os,omitempty"`
	Network   *Network `protobuf:"bytes,9,opt,name=network,proto3" json:"network,omitempty"`
	Screen    *Screen  `protobuf:"bytes,10,opt,name=screen,proto3" json:"screen,omitempty"`
	Locale    string   `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale,omitempty"`
	Timezone  string   `protobuf:"bytes,12,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Web       *Web     `protobuf:"bytes,13,opt,name=web,proto3" json:"web,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{7}
}

func (x *Device) GetDId() string {
	if x != nil {
		return x.DId
	}
	return ""
}

func (x *Device) GetTId() string {
	if x != nil {
		return x.TId
	}
	return ""
}

func (x *Device) GetUId() string {
	if x != nil {
		return x.UId
	}
	return ""
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Device) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *Device) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Device) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Device) GetOs() *Os {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *Device) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *Device) GetScreen() *Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *Device) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Device) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Device) GetWeb() *Web {
	if x != nil {
		return x.Web
	}
	return nil
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *DeviceInfo_Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{8}
}

func (x *DeviceInfo) GetIdentity() *DeviceInfo_Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

type DeviceInfo_Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionTime uint64  `protobuf:"varint,1,opt,name=action_time,json=actionTime,proto3" json:"action_time,omitempty"`
	Sdk        *Sdk    `protobuf:"bytes,2,opt,name=sdk,proto3" json:"sdk,omitempty"`
	Device     *Device `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *DeviceInfo_Identity) Reset() {
	*x = DeviceInfo_Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo_Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo_Identity) ProtoMessage() {}

func (x *DeviceInfo_Identity) ProtoReflect() protoreflect.Message {
	mi := &file_src_grpc_protobufs_user_tracking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo_Identity.ProtoReflect.Descriptor instead.
func (*DeviceInfo_Identity) Descriptor() ([]byte, []int) {
	return file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP(), []int{8, 0}
}

func (x *DeviceInfo_Identity) GetActionTime() uint64 {
	if x != nil {
		return x.ActionTime
	}
	return 0
}

func (x *DeviceInfo_Identity) GetSdk() *Sdk {
	if x != nil {
		return x.Sdk
	}
	return nil
}

func (x *DeviceInfo_Identity) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_src_grpc_protobufs_user_tracking_proto protoreflect.FileDescriptor

var file_src_grpc_protobufs_user_tracking_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x76, 0x0a,
	0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x4f, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a,
	0x03, 0x53, 0x64, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x32,
	0x0a, 0x02, 0x4f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x75,
	0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c,
	0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x66, 0x69, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x77, 0x69, 0x66, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x73, 0x0a,
	0x03, 0x57, 0x65, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xa1, 0x03, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x11, 0x0a,
	0x04, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x49, 0x64,
	0x12, 0x11, 0x0a, 0x04, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x49, 0x64, 0x12, 0x11, 0x0a, 0x04, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x27, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x73, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x06,
	0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x77, 0x65,
	0x62, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65,
	0x62, 0x52, 0x03, 0x77, 0x65, 0x62, 0x22, 0xe1, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x8c, 0x01, 0x0a, 0x08,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x73, 0x64, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x64, 0x6b,
	0x52, 0x03, 0x73, 0x64, 0x6b, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2a, 0x27, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4e, 0x41, 0x50, 0x50,
	0x59, 0x10, 0x01, 0x32, 0x6a, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x32,
	0x5f, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4e, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_src_grpc_protobufs_user_tracking_proto_rawDescOnce sync.Once
	file_src_grpc_protobufs_user_tracking_proto_rawDescData = file_src_grpc_protobufs_user_tracking_proto_rawDesc
)

func file_src_grpc_protobufs_user_tracking_proto_rawDescGZIP() []byte {
	file_src_grpc_protobufs_user_tracking_proto_rawDescOnce.Do(func() {
		file_src_grpc_protobufs_user_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_grpc_protobufs_user_tracking_proto_rawDescData)
	})
	return file_src_grpc_protobufs_user_tracking_proto_rawDescData
}

var file_src_grpc_protobufs_user_tracking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_grpc_protobufs_user_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_src_grpc_protobufs_user_tracking_proto_goTypes = []interface{}{
	(CompressionType)(0),        // 0: compress.service.v1.CompressionType
	(*RequestCompress)(nil),     // 1: compress.service.v1.RequestCompress
	(*ResponseCompress)(nil),    // 2: compress.service.v1.ResponseCompress
	(*Sdk)(nil),                 // 3: compress.service.v1.Sdk
	(*Os)(nil),                  // 4: compress.service.v1.Os
	(*Network)(nil),             // 5: compress.service.v1.Network
	(*Screen)(nil),              // 6: compress.service.v1.Screen
	(*Web)(nil),                 // 7: compress.service.v1.Web
	(*Device)(nil),              // 8: compress.service.v1.Device
	(*DeviceInfo)(nil),          // 9: compress.service.v1.DeviceInfo
	(*DeviceInfo_Identity)(nil), // 10: compress.service.v1.DeviceInfo.Identity
}
var file_src_grpc_protobufs_user_tracking_proto_depIdxs = []int32{
	0,  // 0: compress.service.v1.RequestCompress.compression_type:type_name -> compress.service.v1.CompressionType
	4,  // 1: compress.service.v1.Device.os:type_name -> compress.service.v1.Os
	5,  // 2: compress.service.v1.Device.network:type_name -> compress.service.v1.Network
	6,  // 3: compress.service.v1.Device.screen:type_name -> compress.service.v1.Screen
	7,  // 4: compress.service.v1.Device.web:type_name -> compress.service.v1.Web
	10, // 5: compress.service.v1.DeviceInfo.identity:type_name -> compress.service.v1.DeviceInfo.Identity
	3,  // 6: compress.service.v1.DeviceInfo.Identity.sdk:type_name -> compress.service.v1.Sdk
	8,  // 7: compress.service.v1.DeviceInfo.Identity.device:type_name -> compress.service.v1.Device
	1,  // 8: compress.service.v1.CompressService.Compress:input_type -> compress.service.v1.RequestCompress
	9,  // 9: compress.service.v1.DeviceService.RegisterDevice:input_type -> compress.service.v1.DeviceInfo
	2,  // 10: compress.service.v1.CompressService.Compress:output_type -> compress.service.v1.ResponseCompress
	8,  // 11: compress.service.v1.DeviceService.RegisterDevice:output_type -> compress.service.v1.Device
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_src_grpc_protobufs_user_tracking_proto_init() }
func file_src_grpc_protobufs_user_tracking_proto_init() {
	if File_src_grpc_protobufs_user_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCompress); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCompress); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sdk); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Os); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Screen); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Web); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
		file_src_grpc_protobufs_user_tracking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo_Identity); i {
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
			RawDescriptor: file_src_grpc_protobufs_user_tracking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_src_grpc_protobufs_user_tracking_proto_goTypes,
		DependencyIndexes: file_src_grpc_protobufs_user_tracking_proto_depIdxs,
		EnumInfos:         file_src_grpc_protobufs_user_tracking_proto_enumTypes,
		MessageInfos:      file_src_grpc_protobufs_user_tracking_proto_msgTypes,
	}.Build()
	File_src_grpc_protobufs_user_tracking_proto = out.File
	file_src_grpc_protobufs_user_tracking_proto_rawDesc = nil
	file_src_grpc_protobufs_user_tracking_proto_goTypes = nil
	file_src_grpc_protobufs_user_tracking_proto_depIdxs = nil
}