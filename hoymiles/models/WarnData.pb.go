// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/WarnData.proto

package models

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

type WarnReqDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DtuSn      string    `protobuf:"bytes,1,opt,name=dtu_sn,json=dtuSn,proto3" json:"dtu_sn,omitempty"`
	Time       int32     `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	PackageNub int32     `protobuf:"varint,3,opt,name=package_nub,json=packageNub,proto3" json:"package_nub,omitempty"`
	PackageNow int32     `protobuf:"varint,4,opt,name=package_now,json=packageNow,proto3" json:"package_now,omitempty"`
	WarnDevice int32     `protobuf:"varint,5,opt,name=warn_device,json=warnDevice,proto3" json:"warn_device,omitempty"`
	Warns      []*WarnMO `protobuf:"bytes,6,rep,name=warns,proto3" json:"warns,omitempty"`
}

func (x *WarnReqDTO) Reset() {
	*x = WarnReqDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_WarnData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarnReqDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarnReqDTO) ProtoMessage() {}

func (x *WarnReqDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_WarnData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarnReqDTO.ProtoReflect.Descriptor instead.
func (*WarnReqDTO) Descriptor() ([]byte, []int) {
	return file_protos_WarnData_proto_rawDescGZIP(), []int{0}
}

func (x *WarnReqDTO) GetDtuSn() string {
	if x != nil {
		return x.DtuSn
	}
	return ""
}

func (x *WarnReqDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *WarnReqDTO) GetPackageNub() int32 {
	if x != nil {
		return x.PackageNub
	}
	return 0
}

func (x *WarnReqDTO) GetPackageNow() int32 {
	if x != nil {
		return x.PackageNow
	}
	return 0
}

func (x *WarnReqDTO) GetWarnDevice() int32 {
	if x != nil {
		return x.WarnDevice
	}
	return 0
}

func (x *WarnReqDTO) GetWarns() []*WarnMO {
	if x != nil {
		return x.Warns
	}
	return nil
}

type WarnMO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PvSn   int64 `protobuf:"varint,1,opt,name=pv_sn,json=pvSn,proto3" json:"pv_sn,omitempty"`
	Code   int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Num    int32 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	STime  int32 `protobuf:"varint,4,opt,name=s_time,json=sTime,proto3" json:"s_time,omitempty"`
	ETime  int32 `protobuf:"varint,5,opt,name=e_time,json=eTime,proto3" json:"e_time,omitempty"`
	WData1 int32 `protobuf:"varint,6,opt,name=w_data1,json=wData1,proto3" json:"w_data1,omitempty"`
	WData2 int32 `protobuf:"varint,7,opt,name=w_data2,json=wData2,proto3" json:"w_data2,omitempty"`
}

func (x *WarnMO) Reset() {
	*x = WarnMO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_WarnData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarnMO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarnMO) ProtoMessage() {}

func (x *WarnMO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_WarnData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarnMO.ProtoReflect.Descriptor instead.
func (*WarnMO) Descriptor() ([]byte, []int) {
	return file_protos_WarnData_proto_rawDescGZIP(), []int{1}
}

func (x *WarnMO) GetPvSn() int64 {
	if x != nil {
		return x.PvSn
	}
	return 0
}

func (x *WarnMO) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WarnMO) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *WarnMO) GetSTime() int32 {
	if x != nil {
		return x.STime
	}
	return 0
}

func (x *WarnMO) GetETime() int32 {
	if x != nil {
		return x.ETime
	}
	return 0
}

func (x *WarnMO) GetWData1() int32 {
	if x != nil {
		return x.WData1
	}
	return 0
}

func (x *WarnMO) GetWData2() int32 {
	if x != nil {
		return x.WData2
	}
	return 0
}

type WarnResDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YmdHms     string `protobuf:"bytes,1,opt,name=ymd_hms,json=ymdHms,proto3" json:"ymd_hms,omitempty"`
	PackageNow int32  `protobuf:"varint,2,opt,name=package_now,json=packageNow,proto3" json:"package_now,omitempty"`
	ErrCode    int32  `protobuf:"varint,3,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Offset     int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Time       int32  `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *WarnResDTO) Reset() {
	*x = WarnResDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_WarnData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WarnResDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WarnResDTO) ProtoMessage() {}

func (x *WarnResDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_WarnData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WarnResDTO.ProtoReflect.Descriptor instead.
func (*WarnResDTO) Descriptor() ([]byte, []int) {
	return file_protos_WarnData_proto_rawDescGZIP(), []int{2}
}

func (x *WarnResDTO) GetYmdHms() string {
	if x != nil {
		return x.YmdHms
	}
	return ""
}

func (x *WarnResDTO) GetPackageNow() int32 {
	if x != nil {
		return x.PackageNow
	}
	return 0
}

func (x *WarnResDTO) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *WarnResDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *WarnResDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type WaveReqDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DtuSn      string `protobuf:"bytes,1,opt,name=dtu_sn,json=dtuSn,proto3" json:"dtu_sn,omitempty"`
	Time       int32  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	PackageNub int32  `protobuf:"varint,3,opt,name=package_nub,json=packageNub,proto3" json:"package_nub,omitempty"`
	PackageNow int32  `protobuf:"varint,4,opt,name=package_now,json=packageNow,proto3" json:"package_now,omitempty"`
	PvSn       int64  `protobuf:"varint,5,opt,name=pv_sn,json=pvSn,proto3" json:"pv_sn,omitempty"`
	Code       int32  `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
	Num        int32  `protobuf:"varint,7,opt,name=num,proto3" json:"num,omitempty"`
	WarnTime   int32  `protobuf:"varint,8,opt,name=warn_time,json=warnTime,proto3" json:"warn_time,omitempty"`
	DataLen    int32  `protobuf:"varint,9,opt,name=data_len,json=dataLen,proto3" json:"data_len,omitempty"`
	Pos        int32  `protobuf:"varint,10,opt,name=pos,proto3" json:"pos,omitempty"`
	WarnData   string `protobuf:"bytes,11,opt,name=warn_data,json=warnData,proto3" json:"warn_data,omitempty"`
}

func (x *WaveReqDTO) Reset() {
	*x = WaveReqDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_WarnData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaveReqDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaveReqDTO) ProtoMessage() {}

func (x *WaveReqDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_WarnData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaveReqDTO.ProtoReflect.Descriptor instead.
func (*WaveReqDTO) Descriptor() ([]byte, []int) {
	return file_protos_WarnData_proto_rawDescGZIP(), []int{3}
}

func (x *WaveReqDTO) GetDtuSn() string {
	if x != nil {
		return x.DtuSn
	}
	return ""
}

func (x *WaveReqDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *WaveReqDTO) GetPackageNub() int32 {
	if x != nil {
		return x.PackageNub
	}
	return 0
}

func (x *WaveReqDTO) GetPackageNow() int32 {
	if x != nil {
		return x.PackageNow
	}
	return 0
}

func (x *WaveReqDTO) GetPvSn() int64 {
	if x != nil {
		return x.PvSn
	}
	return 0
}

func (x *WaveReqDTO) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WaveReqDTO) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *WaveReqDTO) GetWarnTime() int32 {
	if x != nil {
		return x.WarnTime
	}
	return 0
}

func (x *WaveReqDTO) GetDataLen() int32 {
	if x != nil {
		return x.DataLen
	}
	return 0
}

func (x *WaveReqDTO) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *WaveReqDTO) GetWarnData() string {
	if x != nil {
		return x.WarnData
	}
	return ""
}

type WaveResDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YmdHms     string `protobuf:"bytes,1,opt,name=ymd_hms,json=ymdHms,proto3" json:"ymd_hms,omitempty"`
	PackageNow int32  `protobuf:"varint,2,opt,name=package_now,json=packageNow,proto3" json:"package_now,omitempty"`
	ErrCode    int32  `protobuf:"varint,3,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Offset     int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Time       int32  `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *WaveResDTO) Reset() {
	*x = WaveResDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_WarnData_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaveResDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaveResDTO) ProtoMessage() {}

func (x *WaveResDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_WarnData_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaveResDTO.ProtoReflect.Descriptor instead.
func (*WaveResDTO) Descriptor() ([]byte, []int) {
	return file_protos_WarnData_proto_rawDescGZIP(), []int{4}
}

func (x *WaveResDTO) GetYmdHms() string {
	if x != nil {
		return x.YmdHms
	}
	return ""
}

func (x *WaveResDTO) GetPackageNow() int32 {
	if x != nil {
		return x.PackageNow
	}
	return 0
}

func (x *WaveResDTO) GetErrCode() int32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *WaveResDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *WaveResDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_protos_WarnData_proto protoreflect.FileDescriptor

var file_protos_WarnData_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x57, 0x61, 0x72, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0a, 0x57, 0x61, 0x72, 0x6e,
	0x52, 0x65, 0x71, 0x44, 0x54, 0x4f, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x74, 0x75, 0x5f, 0x73, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x74, 0x75, 0x53, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f,
	0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x77, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x61, 0x72, 0x6e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x77, 0x61, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x4d, 0x4f, 0x52, 0x05, 0x77, 0x61,
	0x72, 0x6e, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x06, 0x57, 0x61, 0x72, 0x6e, 0x4d, 0x4f, 0x12, 0x13,
	0x0a, 0x05, 0x70, 0x76, 0x5f, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x76, 0x53, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x44, 0x61, 0x74, 0x61, 0x31,
	0x12, 0x17, 0x0a, 0x07, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x77, 0x44, 0x61, 0x74, 0x61, 0x32, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x57, 0x61,
	0x72, 0x6e, 0x52, 0x65, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x6d, 0x64, 0x5f,
	0x68, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x79, 0x6d, 0x64, 0x48, 0x6d,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x77,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x9b, 0x02, 0x0a, 0x0a, 0x57, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x44, 0x54, 0x4f, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x74, 0x75, 0x5f,
	0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x74, 0x75, 0x53, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x4e, 0x6f, 0x77, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x76, 0x5f, 0x73, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x76, 0x53, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x72, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61,
	0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x61, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x57, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x6d, 0x64, 0x5f, 0x68, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x79, 0x6d, 0x64, 0x48, 0x6d, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x77,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x23, 0x5a, 0x0f, 0x68, 0x6f, 0x79, 0x6d, 0x69,
	0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x0f, 0x48, 0x6f, 0x79,
	0x6d, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_WarnData_proto_rawDescOnce sync.Once
	file_protos_WarnData_proto_rawDescData = file_protos_WarnData_proto_rawDesc
)

func file_protos_WarnData_proto_rawDescGZIP() []byte {
	file_protos_WarnData_proto_rawDescOnce.Do(func() {
		file_protos_WarnData_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_WarnData_proto_rawDescData)
	})
	return file_protos_WarnData_proto_rawDescData
}

var file_protos_WarnData_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_WarnData_proto_goTypes = []any{
	(*WarnReqDTO)(nil), // 0: WarnReqDTO
	(*WarnMO)(nil),     // 1: WarnMO
	(*WarnResDTO)(nil), // 2: WarnResDTO
	(*WaveReqDTO)(nil), // 3: WaveReqDTO
	(*WaveResDTO)(nil), // 4: WaveResDTO
}
var file_protos_WarnData_proto_depIdxs = []int32{
	1, // 0: WarnReqDTO.warns:type_name -> WarnMO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_WarnData_proto_init() }
func file_protos_WarnData_proto_init() {
	if File_protos_WarnData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_WarnData_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WarnReqDTO); i {
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
		file_protos_WarnData_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WarnMO); i {
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
		file_protos_WarnData_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WarnResDTO); i {
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
		file_protos_WarnData_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WaveReqDTO); i {
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
		file_protos_WarnData_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WaveResDTO); i {
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
			RawDescriptor: file_protos_WarnData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_WarnData_proto_goTypes,
		DependencyIndexes: file_protos_WarnData_proto_depIdxs,
		MessageInfos:      file_protos_WarnData_proto_msgTypes,
	}.Build()
	File_protos_WarnData_proto = out.File
	file_protos_WarnData_proto_rawDesc = nil
	file_protos_WarnData_proto_goTypes = nil
	file_protos_WarnData_proto_depIdxs = nil
}
