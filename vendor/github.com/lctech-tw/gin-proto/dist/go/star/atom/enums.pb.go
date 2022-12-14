// Code generated by protoc-gen-go. DO NOT EDIT.
// source: star/atom/enums.proto

package atom

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

type ContentType int32

const (
	ContentType_BANNER        ContentType = 0
	ContentType_MODEL_CARD    ContentType = 1
	ContentType_IDENTITY_CARD ContentType = 2
	ContentType_BANK_ACCOUNT  ContentType = 3
	ContentType_LOGO          ContentType = 4
)

var ContentType_name = map[int32]string{
	0: "BANNER",
	1: "MODEL_CARD",
	2: "IDENTITY_CARD",
	3: "BANK_ACCOUNT",
	4: "LOGO",
}

var ContentType_value = map[string]int32{
	"BANNER":        0,
	"MODEL_CARD":    1,
	"IDENTITY_CARD": 2,
	"BANK_ACCOUNT":  3,
	"LOGO":          4,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{0}
}

type GinRole int32

const (
	GinRole_NONE    GinRole = 0
	GinRole_COMPANY GinRole = 1
	GinRole_AGENT   GinRole = 2
	GinRole_ADMIN   GinRole = 3
)

var GinRole_name = map[int32]string{
	0: "NONE",
	1: "COMPANY",
	2: "AGENT",
	3: "ADMIN",
}

var GinRole_value = map[string]int32{
	"NONE":    0,
	"COMPANY": 1,
	"AGENT":   2,
	"ADMIN":   3,
}

func (x GinRole) String() string {
	return proto.EnumName(GinRole_name, int32(x))
}

func (GinRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{1}
}

type PhoneStatus int32

const (
	PhoneStatus_EXIST PhoneStatus = 0
	PhoneStatus_EMPTY PhoneStatus = 1
)

var PhoneStatus_name = map[int32]string{
	0: "EXIST",
	1: "EMPTY",
}

var PhoneStatus_value = map[string]int32{
	"EXIST": 0,
	"EMPTY": 1,
}

func (x PhoneStatus) String() string {
	return proto.EnumName(PhoneStatus_name, int32(x))
}

func (PhoneStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{2}
}

type RegisterStatus int32

const (
	RegisterStatus_DEFAULT     RegisterStatus = 0
	RegisterStatus_WITHDRAW    RegisterStatus = 1
	RegisterStatus_WAIT_VERIFY RegisterStatus = 2
	RegisterStatus_CONFIRM     RegisterStatus = 3
	RegisterStatus_TRASH       RegisterStatus = 4
)

var RegisterStatus_name = map[int32]string{
	0: "DEFAULT",
	1: "WITHDRAW",
	2: "WAIT_VERIFY",
	3: "CONFIRM",
	4: "TRASH",
}

var RegisterStatus_value = map[string]int32{
	"DEFAULT":     0,
	"WITHDRAW":    1,
	"WAIT_VERIFY": 2,
	"CONFIRM":     3,
	"TRASH":       4,
}

func (x RegisterStatus) String() string {
	return proto.EnumName(RegisterStatus_name, int32(x))
}

func (RegisterStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{3}
}

type BoundaryType int32

const (
	BoundaryType_BT_NORMAL BoundaryType = 0
	BoundaryType_BT_SMALL  BoundaryType = 1
	BoundaryType_BT_MEDIUM BoundaryType = 2
	BoundaryType_BT_LARGE  BoundaryType = 3
)

var BoundaryType_name = map[int32]string{
	0: "BT_NORMAL",
	1: "BT_SMALL",
	2: "BT_MEDIUM",
	3: "BT_LARGE",
}

var BoundaryType_value = map[string]int32{
	"BT_NORMAL": 0,
	"BT_SMALL":  1,
	"BT_MEDIUM": 2,
	"BT_LARGE":  3,
}

func (x BoundaryType) String() string {
	return proto.EnumName(BoundaryType_name, int32(x))
}

func (BoundaryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{4}
}

type SmsContentType int32

const (
	SmsContentType_SMS_VERIFY_CODE SmsContentType = 0
	SmsContentType_SMM_APP_LINK    SmsContentType = 1
	SmsContentType_SMM_REVOKE_FILE SmsContentType = 2
)

var SmsContentType_name = map[int32]string{
	0: "SMS_VERIFY_CODE",
	1: "SMM_APP_LINK",
	2: "SMM_REVOKE_FILE",
}

var SmsContentType_value = map[string]int32{
	"SMS_VERIFY_CODE": 0,
	"SMM_APP_LINK":    1,
	"SMM_REVOKE_FILE": 2,
}

func (x SmsContentType) String() string {
	return proto.EnumName(SmsContentType_name, int32(x))
}

func (SmsContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{5}
}

type RevenueTargetType int32

const (
	RevenueTargetType_RT_AGENT RevenueTargetType = 0
	RevenueTargetType_RT_STAR  RevenueTargetType = 1
)

var RevenueTargetType_name = map[int32]string{
	0: "RT_AGENT",
	1: "RT_STAR",
}

var RevenueTargetType_value = map[string]int32{
	"RT_AGENT": 0,
	"RT_STAR":  1,
}

func (x RevenueTargetType) String() string {
	return proto.EnumName(RevenueTargetType_name, int32(x))
}

func (RevenueTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{6}
}

type RevenueCategoryType int32

const (
	RevenueCategoryType_RC_ALL  RevenueCategoryType = 0
	RevenueCategoryType_RC_GIFT RevenueCategoryType = 1
	RevenueCategoryType_RC_POST RevenueCategoryType = 2
)

var RevenueCategoryType_name = map[int32]string{
	0: "RC_ALL",
	1: "RC_GIFT",
	2: "RC_POST",
}

var RevenueCategoryType_value = map[string]int32{
	"RC_ALL":  0,
	"RC_GIFT": 1,
	"RC_POST": 2,
}

func (x RevenueCategoryType) String() string {
	return proto.EnumName(RevenueCategoryType_name, int32(x))
}

func (RevenueCategoryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{7}
}

type StarRole int32

const (
	StarRole_SR_DEFAULT StarRole = 0
	StarRole_SR_LADY    StarRole = 1
)

var StarRole_name = map[int32]string{
	0: "SR_DEFAULT",
	1: "SR_LADY",
}

var StarRole_value = map[string]int32{
	"SR_DEFAULT": 0,
	"SR_LADY":    1,
}

func (x StarRole) String() string {
	return proto.EnumName(StarRole_name, int32(x))
}

func (StarRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{8}
}

type StarStatus int32

const (
	StarStatus_SS_NORMAL   StarStatus = 0
	StarStatus_SS_GRADUATE StarStatus = 1
	StarStatus_SS_ALL      StarStatus = 2
)

var StarStatus_name = map[int32]string{
	0: "SS_NORMAL",
	1: "SS_GRADUATE",
	2: "SS_ALL",
}

var StarStatus_value = map[string]int32{
	"SS_NORMAL":   0,
	"SS_GRADUATE": 1,
	"SS_ALL":      2,
}

func (x StarStatus) String() string {
	return proto.EnumName(StarStatus_name, int32(x))
}

func (StarStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{9}
}

type Device int32

const (
	Device_D_DEFAULT Device = 0
	Device_D_DESKTOP Device = 1
	Device_D_MOBILE  Device = 2
)

var Device_name = map[int32]string{
	0: "D_DEFAULT",
	1: "D_DESKTOP",
	2: "D_MOBILE",
}

var Device_value = map[string]int32{
	"D_DEFAULT": 0,
	"D_DESKTOP": 1,
	"D_MOBILE":  2,
}

func (x Device) String() string {
	return proto.EnumName(Device_name, int32(x))
}

func (Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{10}
}

type BloodType int32

const (
	BloodType_BLOOD_TYPE_EMPTY   BloodType = 0
	BloodType_BLOOD_TYPE_A       BloodType = 1
	BloodType_BLOOD_TYPE_B       BloodType = 2
	BloodType_BLOOD_TYPE_AB      BloodType = 3
	BloodType_BLOOD_TYPE_O       BloodType = 4
	BloodType_BLOOD_TYPE_UNKNOWN BloodType = 5
)

var BloodType_name = map[int32]string{
	0: "BLOOD_TYPE_EMPTY",
	1: "BLOOD_TYPE_A",
	2: "BLOOD_TYPE_B",
	3: "BLOOD_TYPE_AB",
	4: "BLOOD_TYPE_O",
	5: "BLOOD_TYPE_UNKNOWN",
}

var BloodType_value = map[string]int32{
	"BLOOD_TYPE_EMPTY":   0,
	"BLOOD_TYPE_A":       1,
	"BLOOD_TYPE_B":       2,
	"BLOOD_TYPE_AB":      3,
	"BLOOD_TYPE_O":       4,
	"BLOOD_TYPE_UNKNOWN": 5,
}

func (x BloodType) String() string {
	return proto.EnumName(BloodType_name, int32(x))
}

func (BloodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d391e2555b1794fd, []int{11}
}

func init() {
	proto.RegisterEnum("star.atom.enums.ContentType", ContentType_name, ContentType_value)
	proto.RegisterEnum("star.atom.enums.GinRole", GinRole_name, GinRole_value)
	proto.RegisterEnum("star.atom.enums.PhoneStatus", PhoneStatus_name, PhoneStatus_value)
	proto.RegisterEnum("star.atom.enums.RegisterStatus", RegisterStatus_name, RegisterStatus_value)
	proto.RegisterEnum("star.atom.enums.BoundaryType", BoundaryType_name, BoundaryType_value)
	proto.RegisterEnum("star.atom.enums.SmsContentType", SmsContentType_name, SmsContentType_value)
	proto.RegisterEnum("star.atom.enums.RevenueTargetType", RevenueTargetType_name, RevenueTargetType_value)
	proto.RegisterEnum("star.atom.enums.RevenueCategoryType", RevenueCategoryType_name, RevenueCategoryType_value)
	proto.RegisterEnum("star.atom.enums.StarRole", StarRole_name, StarRole_value)
	proto.RegisterEnum("star.atom.enums.StarStatus", StarStatus_name, StarStatus_value)
	proto.RegisterEnum("star.atom.enums.Device", Device_name, Device_value)
	proto.RegisterEnum("star.atom.enums.BloodType", BloodType_name, BloodType_value)
}

func init() { proto.RegisterFile("star/atom/enums.proto", fileDescriptor_d391e2555b1794fd) }

var fileDescriptor_d391e2555b1794fd = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x5d, 0x6f, 0xda, 0x3e,
	0x14, 0xc6, 0x9b, 0x40, 0x79, 0x39, 0x50, 0x38, 0x4d, 0xff, 0xff, 0x7d, 0x81, 0x5d, 0x4c, 0x8a,
	0x54, 0x98, 0xb6, 0x69, 0x9a, 0xb4, 0x2b, 0x27, 0x31, 0x34, 0x22, 0xb1, 0x23, 0xdb, 0x94, 0xb2,
	0x1b, 0x8b, 0xd2, 0x88, 0x22, 0x95, 0xa4, 0x02, 0xd3, 0xa9, 0xb7, 0xfb, 0xe4, 0x53, 0x5e, 0xb4,
	0xa2, 0xde, 0xd9, 0xcf, 0x79, 0xf4, 0xf8, 0x9c, 0x9f, 0x6d, 0xf8, 0xff, 0x60, 0x56, 0xfb, 0xf1,
	0xca, 0xe4, 0xbb, 0x71, 0x9a, 0x1d, 0x77, 0x87, 0xd1, 0xf3, 0x3e, 0x37, 0xb9, 0x33, 0x2c, 0xe4,
	0x51, 0x21, 0x8f, 0x4a, 0xd9, 0xbd, 0x83, 0x9e, 0x9f, 0x67, 0x26, 0xcd, 0x8c, 0x7a, 0x7d, 0x4e,
	0x1d, 0x80, 0x96, 0x47, 0x18, 0xa3, 0x02, 0xcf, 0x9c, 0x01, 0x40, 0xcc, 0x03, 0x1a, 0x69, 0x9f,
	0x88, 0x00, 0x2d, 0xe7, 0x12, 0x2e, 0xc2, 0x80, 0x32, 0x15, 0xaa, 0x65, 0x25, 0xd9, 0x0e, 0x42,
	0xdf, 0x23, 0x6c, 0xa6, 0x89, 0xef, 0xf3, 0x39, 0x53, 0xd8, 0x70, 0x3a, 0xd0, 0x8c, 0xf8, 0x94,
	0x63, 0xd3, 0xfd, 0x0e, 0xed, 0xe9, 0x36, 0x13, 0xf9, 0x53, 0x5a, 0x88, 0x8c, 0x33, 0x8a, 0x67,
	0x4e, 0x0f, 0xda, 0x3e, 0x8f, 0x13, 0xc2, 0x96, 0x68, 0x39, 0x5d, 0x38, 0x27, 0x53, 0xca, 0x14,
	0xda, 0xe5, 0x32, 0x88, 0x43, 0x86, 0x0d, 0xf7, 0x23, 0xf4, 0x92, 0xc7, 0x3c, 0x4b, 0xa5, 0x59,
	0x99, 0xe3, 0xa1, 0xa8, 0xd0, 0xbb, 0x50, 0x2a, 0x3c, 0x2b, 0x97, 0x71, 0xa2, 0x96, 0x68, 0xb9,
	0x0a, 0x06, 0x22, 0xdd, 0x6c, 0x0f, 0x26, 0xdd, 0xd7, 0xbe, 0x1e, 0xb4, 0x03, 0x3a, 0x21, 0xf3,
	0xa8, 0x70, 0xf6, 0xa1, 0xb3, 0x08, 0xd5, 0x4d, 0x20, 0xc8, 0x02, 0x2d, 0x67, 0x08, 0xbd, 0x05,
	0x09, 0x95, 0xbe, 0xa5, 0x22, 0x9c, 0x2c, 0xd1, 0xae, 0xba, 0x60, 0x93, 0x50, 0xc4, 0xd8, 0x28,
	0x52, 0x95, 0x20, 0xf2, 0x06, 0x9b, 0xee, 0x0d, 0xf4, 0xbd, 0xfc, 0x98, 0x3d, 0xac, 0xf6, 0xaf,
	0x25, 0x8d, 0x0b, 0xe8, 0x7a, 0x4a, 0x33, 0x2e, 0x62, 0x12, 0x55, 0xa9, 0x9e, 0xd2, 0x32, 0x26,
	0x51, 0x84, 0x56, 0x5d, 0x8c, 0x69, 0x10, 0xce, 0x63, 0xb4, 0xeb, 0x62, 0x44, 0xc4, 0x94, 0x62,
	0xc3, 0x8d, 0x60, 0x20, 0x77, 0x87, 0x53, 0xb2, 0x57, 0x30, 0x94, 0xb1, 0xac, 0x7b, 0xd0, 0x3e,
	0x0f, 0x0a, 0x1c, 0x08, 0x7d, 0x19, 0xc7, 0x9a, 0x24, 0x89, 0x8e, 0x42, 0x36, 0x43, 0xab, 0xb2,
	0xc5, 0x5a, 0xd0, 0x5b, 0x3e, 0xa3, 0x7a, 0x12, 0x46, 0x14, 0x6d, 0x77, 0x04, 0x97, 0x22, 0x7d,
	0x49, 0xb3, 0x63, 0xaa, 0x56, 0xfb, 0x4d, 0x5a, 0x05, 0xf6, 0xa1, 0x23, 0x94, 0xae, 0x00, 0x96,
	0x60, 0x85, 0xd2, 0x52, 0x11, 0x81, 0x96, 0xfb, 0x13, 0xae, 0x6a, 0xbf, 0xbf, 0x32, 0xe9, 0x26,
	0xaf, 0xc7, 0x01, 0x68, 0x09, 0x5f, 0x17, 0xdd, 0x57, 0x7e, 0x5f, 0x4f, 0xc3, 0x89, 0x42, 0xab,
	0xde, 0x24, 0x5c, 0x2a, 0xb4, 0xdd, 0x4f, 0xd0, 0x91, 0x66, 0xb5, 0x2f, 0x2f, 0x6e, 0x00, 0x20,
	0x85, 0x7e, 0xe3, 0xda, 0x83, 0xb6, 0x14, 0x3a, 0x22, 0x41, 0x71, 0x07, 0x3f, 0x00, 0x0a, 0x63,
	0xcd, 0xff, 0x02, 0xba, 0x52, 0xbe, 0xb1, 0x1a, 0x42, 0x4f, 0x4a, 0x3d, 0x15, 0x24, 0x98, 0x13,
	0x45, 0xd1, 0x2a, 0x0e, 0x97, 0xb2, 0x3c, 0xdc, 0x76, 0xbf, 0x41, 0x2b, 0x48, 0x5f, 0xb6, 0xeb,
	0x92, 0x70, 0x70, 0x92, 0x5f, 0x6f, 0xe5, 0x4c, 0xf1, 0x04, 0xad, 0x62, 0xc4, 0x40, 0xc7, 0xdc,
	0xab, 0x28, 0xfc, 0xb1, 0xa0, 0xeb, 0x3d, 0xe5, 0xf9, 0x43, 0x39, 0xcc, 0x7f, 0x80, 0x5e, 0xc4,
	0x79, 0xa0, 0xd5, 0x32, 0xa1, 0xba, 0x7a, 0x17, 0x25, 0xd0, 0x13, 0x95, 0xa0, 0xf5, 0x4e, 0xf1,
	0xd0, 0x2e, 0xde, 0xf1, 0xa9, 0xc7, 0xc3, 0xc6, 0x3b, 0x13, 0xc7, 0xa6, 0xf3, 0x01, 0x9c, 0x13,
	0x65, 0xce, 0x66, 0x8c, 0x2f, 0x18, 0x9e, 0x7b, 0x5f, 0x7e, 0x7d, 0xde, 0x6c, 0xcd, 0xe3, 0xf1,
	0x7e, 0xb4, 0xce, 0x77, 0xe3, 0xa7, 0xb5, 0x49, 0xd7, 0x8f, 0xd7, 0xe6, 0xf7, 0x78, 0xb3, 0xcd,
	0xae, 0xcb, 0x2f, 0x36, 0x7e, 0xd8, 0x1e, 0xcc, 0x78, 0x93, 0x8f, 0xff, 0x7d, 0xc0, 0xfb, 0x56,
	0x59, 0xf8, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x1d, 0x08, 0x47, 0x94, 0x03, 0x00, 0x00,
}
