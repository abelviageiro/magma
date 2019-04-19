// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feg_config.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GyInitMethod int32

const (
	GyInitMethod_RESERVED    GyInitMethod = 0
	GyInitMethod_PER_SESSION GyInitMethod = 1
	GyInitMethod_PER_KEY     GyInitMethod = 2
)

var GyInitMethod_name = map[int32]string{
	0: "RESERVED",
	1: "PER_SESSION",
	2: "PER_KEY",
}
var GyInitMethod_value = map[string]int32{
	"RESERVED":    0,
	"PER_SESSION": 1,
	"PER_KEY":     2,
}

func (x GyInitMethod) String() string {
	return proto.EnumName(GyInitMethod_name, int32(x))
}
func (GyInitMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{0}
}

type DiamClientConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Retransmits          uint32   `protobuf:"varint,3,opt,name=retransmits,proto3" json:"retransmits,omitempty"`
	WatchdogInterval     uint32   `protobuf:"varint,4,opt,name=watchdog_interval,json=watchdogInterval,proto3" json:"watchdog_interval,omitempty"`
	RetryCount           uint32   `protobuf:"varint,5,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	LocalAddress         string   `protobuf:"bytes,6,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ProductName          string   `protobuf:"bytes,7,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Realm                string   `protobuf:"bytes,8,opt,name=realm,proto3" json:"realm,omitempty"`
	Host                 string   `protobuf:"bytes,9,opt,name=host,proto3" json:"host,omitempty"`
	DestRealm            string   `protobuf:"bytes,10,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	DestHost             string   `protobuf:"bytes,11,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamClientConfig) Reset()         { *m = DiamClientConfig{} }
func (m *DiamClientConfig) String() string { return proto.CompactTextString(m) }
func (*DiamClientConfig) ProtoMessage()    {}
func (*DiamClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{0}
}
func (m *DiamClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamClientConfig.Unmarshal(m, b)
}
func (m *DiamClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamClientConfig.Marshal(b, m, deterministic)
}
func (dst *DiamClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamClientConfig.Merge(dst, src)
}
func (m *DiamClientConfig) XXX_Size() int {
	return xxx_messageInfo_DiamClientConfig.Size(m)
}
func (m *DiamClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamClientConfig proto.InternalMessageInfo

func (m *DiamClientConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamClientConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamClientConfig) GetRetransmits() uint32 {
	if m != nil {
		return m.Retransmits
	}
	return 0
}

func (m *DiamClientConfig) GetWatchdogInterval() uint32 {
	if m != nil {
		return m.WatchdogInterval
	}
	return 0
}

func (m *DiamClientConfig) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *DiamClientConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamClientConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *DiamClientConfig) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *DiamClientConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *DiamClientConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

func (m *DiamClientConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

type DiamServerConfig struct {
	Protocol             string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	LocalAddress         string   `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	DestHost             string   `protobuf:"bytes,4,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`
	DestRealm            string   `protobuf:"bytes,5,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiamServerConfig) Reset()         { *m = DiamServerConfig{} }
func (m *DiamServerConfig) String() string { return proto.CompactTextString(m) }
func (*DiamServerConfig) ProtoMessage()    {}
func (*DiamServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{1}
}
func (m *DiamServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiamServerConfig.Unmarshal(m, b)
}
func (m *DiamServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiamServerConfig.Marshal(b, m, deterministic)
}
func (dst *DiamServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiamServerConfig.Merge(dst, src)
}
func (m *DiamServerConfig) XXX_Size() int {
	return xxx_messageInfo_DiamServerConfig.Size(m)
}
func (m *DiamServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DiamServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DiamServerConfig proto.InternalMessageInfo

func (m *DiamServerConfig) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *DiamServerConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DiamServerConfig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *DiamServerConfig) GetDestHost() string {
	if m != nil {
		return m.DestHost
	}
	return ""
}

func (m *DiamServerConfig) GetDestRealm() string {
	if m != nil {
		return m.DestRealm
	}
	return ""
}

type S6AConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *S6AConfig) Reset()         { *m = S6AConfig{} }
func (m *S6AConfig) String() string { return proto.CompactTextString(m) }
func (*S6AConfig) ProtoMessage()    {}
func (*S6AConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{2}
}
func (m *S6AConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S6AConfig.Unmarshal(m, b)
}
func (m *S6AConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S6AConfig.Marshal(b, m, deterministic)
}
func (dst *S6AConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S6AConfig.Merge(dst, src)
}
func (m *S6AConfig) XXX_Size() int {
	return xxx_messageInfo_S6AConfig.Size(m)
}
func (m *S6AConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_S6AConfig.DiscardUnknown(m)
}

var xxx_messageInfo_S6AConfig proto.InternalMessageInfo

func (m *S6AConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

type GxConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GxConfig) Reset()         { *m = GxConfig{} }
func (m *GxConfig) String() string { return proto.CompactTextString(m) }
func (*GxConfig) ProtoMessage()    {}
func (*GxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{3}
}
func (m *GxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GxConfig.Unmarshal(m, b)
}
func (m *GxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GxConfig.Marshal(b, m, deterministic)
}
func (dst *GxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GxConfig.Merge(dst, src)
}
func (m *GxConfig) XXX_Size() int {
	return xxx_messageInfo_GxConfig.Size(m)
}
func (m *GxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GxConfig proto.InternalMessageInfo

func (m *GxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

type GyConfig struct {
	Server               *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	InitMethod           GyInitMethod      `protobuf:"varint,2,opt,name=init_method,json=initMethod,proto3,enum=feg.GyInitMethod" json:"init_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GyConfig) Reset()         { *m = GyConfig{} }
func (m *GyConfig) String() string { return proto.CompactTextString(m) }
func (*GyConfig) ProtoMessage()    {}
func (*GyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{4}
}
func (m *GyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GyConfig.Unmarshal(m, b)
}
func (m *GyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GyConfig.Marshal(b, m, deterministic)
}
func (dst *GyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GyConfig.Merge(dst, src)
}
func (m *GyConfig) XXX_Size() int {
	return xxx_messageInfo_GyConfig.Size(m)
}
func (m *GyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GyConfig proto.InternalMessageInfo

func (m *GyConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *GyConfig) GetInitMethod() GyInitMethod {
	if m != nil {
		return m.InitMethod
	}
	return GyInitMethod_RESERVED
}

type SwxConfig struct {
	Server *DiamClientConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// After auth, verify Non-3GPP IP Access enabled
	VerifyAuthorization  bool     `protobuf:"varint,2,opt,name=verify_authorization,json=verifyAuthorization,proto3" json:"verify_authorization,omitempty"`
	CacheTTLSeconds      uint32   `protobuf:"varint,3,opt,name=CacheTTLSeconds,proto3" json:"CacheTTLSeconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwxConfig) Reset()         { *m = SwxConfig{} }
func (m *SwxConfig) String() string { return proto.CompactTextString(m) }
func (*SwxConfig) ProtoMessage()    {}
func (*SwxConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{5}
}
func (m *SwxConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwxConfig.Unmarshal(m, b)
}
func (m *SwxConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwxConfig.Marshal(b, m, deterministic)
}
func (dst *SwxConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwxConfig.Merge(dst, src)
}
func (m *SwxConfig) XXX_Size() int {
	return xxx_messageInfo_SwxConfig.Size(m)
}
func (m *SwxConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SwxConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SwxConfig proto.InternalMessageInfo

func (m *SwxConfig) GetServer() *DiamClientConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *SwxConfig) GetVerifyAuthorization() bool {
	if m != nil {
		return m.VerifyAuthorization
	}
	return false
}

func (m *SwxConfig) GetCacheTTLSeconds() uint32 {
	if m != nil {
		return m.CacheTTLSeconds
	}
	return 0
}

type HSSConfig struct {
	Server *DiamServerConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Operator configuration field for LTE.
	LteAuthOp []byte `protobuf:"bytes,2,opt,name=lte_auth_op,json=lteAuthOp,proto3" json:"lte_auth_op,omitempty"`
	// Authentication management field for LTE.
	LteAuthAmf []byte `protobuf:"bytes,3,opt,name=lte_auth_amf,json=lteAuthAmf,proto3" json:"lte_auth_amf,omitempty"`
	// Maps from IMSI to SubscriptionProfile.
	SubProfiles map[string]*HSSConfig_SubscriptionProfile `protobuf:"bytes,4,rep,name=sub_profiles,json=subProfiles,proto3" json:"sub_profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If an IMSI if not found in sub_profiles, the default profile is used instead.
	DefaultSubProfile    *HSSConfig_SubscriptionProfile `protobuf:"bytes,5,opt,name=default_sub_profile,json=defaultSubProfile,proto3" json:"default_sub_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *HSSConfig) Reset()         { *m = HSSConfig{} }
func (m *HSSConfig) String() string { return proto.CompactTextString(m) }
func (*HSSConfig) ProtoMessage()    {}
func (*HSSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{6}
}
func (m *HSSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig.Unmarshal(m, b)
}
func (m *HSSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig.Marshal(b, m, deterministic)
}
func (dst *HSSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig.Merge(dst, src)
}
func (m *HSSConfig) XXX_Size() int {
	return xxx_messageInfo_HSSConfig.Size(m)
}
func (m *HSSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig proto.InternalMessageInfo

func (m *HSSConfig) GetServer() *DiamServerConfig {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *HSSConfig) GetLteAuthOp() []byte {
	if m != nil {
		return m.LteAuthOp
	}
	return nil
}

func (m *HSSConfig) GetLteAuthAmf() []byte {
	if m != nil {
		return m.LteAuthAmf
	}
	return nil
}

func (m *HSSConfig) GetSubProfiles() map[string]*HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.SubProfiles
	}
	return nil
}

func (m *HSSConfig) GetDefaultSubProfile() *HSSConfig_SubscriptionProfile {
	if m != nil {
		return m.DefaultSubProfile
	}
	return nil
}

type HSSConfig_SubscriptionProfile struct {
	// Maximum uplink bit rate (AMBR-UL)
	MaxUlBitRate uint64 `protobuf:"varint,1,opt,name=max_ul_bit_rate,json=maxUlBitRate,proto3" json:"max_ul_bit_rate,omitempty"`
	// Maximum downlink bit rate (AMBR-DL)
	MaxDlBitRate         uint64   `protobuf:"varint,2,opt,name=max_dl_bit_rate,json=maxDlBitRate,proto3" json:"max_dl_bit_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HSSConfig_SubscriptionProfile) Reset()         { *m = HSSConfig_SubscriptionProfile{} }
func (m *HSSConfig_SubscriptionProfile) String() string { return proto.CompactTextString(m) }
func (*HSSConfig_SubscriptionProfile) ProtoMessage()    {}
func (*HSSConfig_SubscriptionProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{6, 0}
}
func (m *HSSConfig_SubscriptionProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Unmarshal(m, b)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Marshal(b, m, deterministic)
}
func (dst *HSSConfig_SubscriptionProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.Merge(dst, src)
}
func (m *HSSConfig_SubscriptionProfile) XXX_Size() int {
	return xxx_messageInfo_HSSConfig_SubscriptionProfile.Size(m)
}
func (m *HSSConfig_SubscriptionProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_HSSConfig_SubscriptionProfile.DiscardUnknown(m)
}

var xxx_messageInfo_HSSConfig_SubscriptionProfile proto.InternalMessageInfo

func (m *HSSConfig_SubscriptionProfile) GetMaxUlBitRate() uint64 {
	if m != nil {
		return m.MaxUlBitRate
	}
	return 0
}

func (m *HSSConfig_SubscriptionProfile) GetMaxDlBitRate() uint64 {
	if m != nil {
		return m.MaxDlBitRate
	}
	return 0
}

type HealthConfig struct {
	// Services the health service is responsible for tracking
	HealthServices []string `protobuf:"bytes,1,rep,name=health_services,json=healthServices,proto3" json:"health_services,omitempty"`
	// Frequency of FeG health manager updates to the cloud
	UpdateIntervalSecs uint32 `protobuf:"varint,2,opt,name=update_interval_secs,json=updateIntervalSecs,proto3" json:"update_interval_secs,omitempty"`
	// Period to disable connection creation when requested to from cloud
	CloudDisablePeriodSecs uint32 `protobuf:"varint,3,opt,name=cloud_disable_period_secs,json=cloudDisablePeriodSecs,proto3" json:"cloud_disable_period_secs,omitempty"`
	// Period to disable connection creation when locally determined
	LocalDisablePeriodSecs uint32 `protobuf:"varint,4,opt,name=local_disable_period_secs,json=localDisablePeriodSecs,proto3" json:"local_disable_period_secs,omitempty"`
	// The number of consecutive health update failures before locally disabling
	UpdateFailureThreshold uint32 `protobuf:"varint,5,opt,name=update_failure_threshold,json=updateFailureThreshold,proto3" json:"update_failure_threshold,omitempty"`
	// Percentage of request failures considered to be unhealthy
	RequestFailureThreshold float32 `protobuf:"fixed32,6,opt,name=request_failure_threshold,json=requestFailureThreshold,proto3" json:"request_failure_threshold,omitempty"`
	// Minimum number of requests necessary to consider a metrics interval valid
	MinimumRequestThreshold uint32 `protobuf:"varint,7,opt,name=minimum_request_threshold,json=minimumRequestThreshold,proto3" json:"minimum_request_threshold,omitempty"`
	// Cpu utilization healthy threshold
	CpuUtilizationThreshold float32 `protobuf:"fixed32,8,opt,name=cpu_utilization_threshold,json=cpuUtilizationThreshold,proto3" json:"cpu_utilization_threshold,omitempty"`
	// Available memory healthy threshold
	MemoryAvailableThreshold float32  `protobuf:"fixed32,9,opt,name=memory_available_threshold,json=memoryAvailableThreshold,proto3" json:"memory_available_threshold,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *HealthConfig) Reset()         { *m = HealthConfig{} }
func (m *HealthConfig) String() string { return proto.CompactTextString(m) }
func (*HealthConfig) ProtoMessage()    {}
func (*HealthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{7}
}
func (m *HealthConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthConfig.Unmarshal(m, b)
}
func (m *HealthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthConfig.Marshal(b, m, deterministic)
}
func (dst *HealthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthConfig.Merge(dst, src)
}
func (m *HealthConfig) XXX_Size() int {
	return xxx_messageInfo_HealthConfig.Size(m)
}
func (m *HealthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HealthConfig proto.InternalMessageInfo

func (m *HealthConfig) GetHealthServices() []string {
	if m != nil {
		return m.HealthServices
	}
	return nil
}

func (m *HealthConfig) GetUpdateIntervalSecs() uint32 {
	if m != nil {
		return m.UpdateIntervalSecs
	}
	return 0
}

func (m *HealthConfig) GetCloudDisablePeriodSecs() uint32 {
	if m != nil {
		return m.CloudDisablePeriodSecs
	}
	return 0
}

func (m *HealthConfig) GetLocalDisablePeriodSecs() uint32 {
	if m != nil {
		return m.LocalDisablePeriodSecs
	}
	return 0
}

func (m *HealthConfig) GetUpdateFailureThreshold() uint32 {
	if m != nil {
		return m.UpdateFailureThreshold
	}
	return 0
}

func (m *HealthConfig) GetRequestFailureThreshold() float32 {
	if m != nil {
		return m.RequestFailureThreshold
	}
	return 0
}

func (m *HealthConfig) GetMinimumRequestThreshold() uint32 {
	if m != nil {
		return m.MinimumRequestThreshold
	}
	return 0
}

func (m *HealthConfig) GetCpuUtilizationThreshold() float32 {
	if m != nil {
		return m.CpuUtilizationThreshold
	}
	return 0
}

func (m *HealthConfig) GetMemoryAvailableThreshold() float32 {
	if m != nil {
		return m.MemoryAvailableThreshold
	}
	return 0
}

type Config struct {
	// FeG config params
	S6A                  *S6AConfig    `protobuf:"bytes,4,opt,name=s6a,proto3" json:"s6a,omitempty"`
	Gx                   *GxConfig     `protobuf:"bytes,5,opt,name=gx,proto3" json:"gx,omitempty"`
	Gy                   *GyConfig     `protobuf:"bytes,6,opt,name=gy,proto3" json:"gy,omitempty"`
	ServedNetworkIds     []string      `protobuf:"bytes,7,rep,name=served_network_ids,json=servedNetworkIds,proto3" json:"served_network_ids,omitempty"`
	Hss                  *HSSConfig    `protobuf:"bytes,8,opt,name=hss,proto3" json:"hss,omitempty"`
	Swx                  *SwxConfig    `protobuf:"bytes,9,opt,name=swx,proto3" json:"swx,omitempty"`
	Health               *HealthConfig `protobuf:"bytes,10,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_feg_config_62d074f461291ea7, []int{8}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetS6A() *S6AConfig {
	if m != nil {
		return m.S6A
	}
	return nil
}

func (m *Config) GetGx() *GxConfig {
	if m != nil {
		return m.Gx
	}
	return nil
}

func (m *Config) GetGy() *GyConfig {
	if m != nil {
		return m.Gy
	}
	return nil
}

func (m *Config) GetServedNetworkIds() []string {
	if m != nil {
		return m.ServedNetworkIds
	}
	return nil
}

func (m *Config) GetHss() *HSSConfig {
	if m != nil {
		return m.Hss
	}
	return nil
}

func (m *Config) GetSwx() *SwxConfig {
	if m != nil {
		return m.Swx
	}
	return nil
}

func (m *Config) GetHealth() *HealthConfig {
	if m != nil {
		return m.Health
	}
	return nil
}

func init() {
	proto.RegisterType((*DiamClientConfig)(nil), "feg.DiamClientConfig")
	proto.RegisterType((*DiamServerConfig)(nil), "feg.DiamServerConfig")
	proto.RegisterType((*S6AConfig)(nil), "feg.S6aConfig")
	proto.RegisterType((*GxConfig)(nil), "feg.GxConfig")
	proto.RegisterType((*GyConfig)(nil), "feg.GyConfig")
	proto.RegisterType((*SwxConfig)(nil), "feg.SwxConfig")
	proto.RegisterType((*HSSConfig)(nil), "feg.HSSConfig")
	proto.RegisterMapType((map[string]*HSSConfig_SubscriptionProfile)(nil), "feg.HSSConfig.SubProfilesEntry")
	proto.RegisterType((*HSSConfig_SubscriptionProfile)(nil), "feg.HSSConfig.SubscriptionProfile")
	proto.RegisterType((*HealthConfig)(nil), "feg.HealthConfig")
	proto.RegisterType((*Config)(nil), "feg.Config")
	proto.RegisterEnum("feg.GyInitMethod", GyInitMethod_name, GyInitMethod_value)
}

func init() { proto.RegisterFile("feg_config.proto", fileDescriptor_feg_config_62d074f461291ea7) }

var fileDescriptor_feg_config_62d074f461291ea7 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0x25, 0x1f, 0x6d, 0x93, 0xeb, 0x74, 0x9b, 0x4e, 0x4b, 0xd7, 0x5b, 0xb4, 0x6c, 0x08, 0x42,
	0x94, 0xaf, 0x0a, 0x82, 0xb4, 0xea, 0x56, 0xbc, 0xf4, 0x23, 0xb4, 0x15, 0xd0, 0xad, 0xc6, 0x5d,
	0x24, 0x78, 0x19, 0x4d, 0xec, 0x49, 0x32, 0x5a, 0x7f, 0x04, 0xcf, 0xb8, 0x4d, 0x78, 0xe7, 0x89,
	0x1f, 0xc0, 0x6f, 0xe0, 0x47, 0xf1, 0xca, 0xef, 0x40, 0xbe, 0x33, 0x76, 0xd2, 0xb4, 0x12, 0xa8,
	0x3c, 0xc5, 0x73, 0xcf, 0x39, 0x33, 0x67, 0x6e, 0xae, 0x8f, 0xa1, 0x3d, 0x14, 0x23, 0xe6, 0x27,
	0xf1, 0x50, 0x8e, 0xf6, 0x27, 0x69, 0xa2, 0x13, 0x52, 0x1b, 0x8a, 0x51, 0xf7, 0xef, 0x2a, 0xb4,
	0x4f, 0x25, 0x8f, 0x4e, 0x42, 0x29, 0x62, 0x7d, 0x82, 0x38, 0xd9, 0x85, 0x06, 0x52, 0xfc, 0x24,
	0x74, 0x2b, 0x9d, 0xca, 0x5e, 0x93, 0x96, 0x6b, 0xe2, 0xc2, 0x1a, 0x0f, 0x82, 0x54, 0x28, 0xe5,
	0x56, 0x11, 0x2a, 0x96, 0xa4, 0x03, 0x4e, 0x2a, 0x74, 0xca, 0x63, 0x15, 0x49, 0xad, 0xdc, 0x5a,
	0xa7, 0xb2, 0xb7, 0x4e, 0x17, 0x4b, 0xe4, 0x33, 0xd8, 0xbc, 0xe5, 0xda, 0x1f, 0x07, 0xc9, 0x88,
	0xc9, 0x58, 0x8b, 0xf4, 0x86, 0x87, 0x6e, 0x1d, 0x79, 0xed, 0x02, 0xb8, 0xb0, 0x75, 0xf2, 0xc2,
	0x6c, 0x37, 0x63, 0x7e, 0x92, 0xc5, 0xda, 0x5d, 0x41, 0x1a, 0x60, 0xe9, 0x24, 0xaf, 0x90, 0x0f,
	0x61, 0x3d, 0x4c, 0x7c, 0x1e, 0xb2, 0xc2, 0xcf, 0x2a, 0xfa, 0x69, 0x61, 0xf1, 0xc8, 0x9a, 0xfa,
	0x00, 0x5a, 0x93, 0x34, 0x09, 0x32, 0x5f, 0xb3, 0x98, 0x47, 0xc2, 0x5d, 0x43, 0x8e, 0x63, 0x6b,
	0x97, 0x3c, 0x12, 0x64, 0x1b, 0x56, 0x52, 0xc1, 0xc3, 0xc8, 0x6d, 0x20, 0x66, 0x16, 0x84, 0x40,
	0x7d, 0x9c, 0x28, 0xed, 0x36, 0xb1, 0x88, 0xcf, 0xe4, 0x39, 0x40, 0x20, 0x94, 0x66, 0x86, 0x0e,
	0x88, 0x34, 0xf3, 0x0a, 0x45, 0xc9, 0x7b, 0x80, 0x0b, 0x86, 0x3a, 0xc7, 0xf4, 0x2d, 0x2f, 0x9c,
	0x27, 0x4a, 0x77, 0xff, 0xac, 0x98, 0x46, 0x7b, 0x22, 0xbd, 0x11, 0xe9, 0xff, 0x6a, 0xf4, 0xbd,
	0x8b, 0xd7, 0x1e, 0xb8, 0xf8, 0x1d, 0x33, 0xf5, 0xbb, 0x66, 0x96, 0x2e, 0xb2, 0xb2, 0x74, 0x91,
	0xee, 0x21, 0x34, 0xbd, 0x97, 0xdc, 0x7a, 0xfc, 0x02, 0x56, 0x15, 0x7a, 0x46, 0x87, 0x4e, 0xef,
	0xdd, 0xfd, 0xa1, 0x18, 0xed, 0x2f, 0xcf, 0x0c, 0xb5, 0xa4, 0xee, 0x2b, 0x68, 0x9c, 0x4d, 0x1f,
	0x27, 0x8d, 0xa0, 0x71, 0x36, 0x7b, 0x94, 0x94, 0xf4, 0xc0, 0x91, 0xb1, 0xd4, 0x2c, 0x12, 0x7a,
	0x9c, 0x04, 0xd8, 0xb0, 0x27, 0xbd, 0x4d, 0xd4, 0x9c, 0xcd, 0x2e, 0x62, 0xa9, 0x7f, 0x40, 0x80,
	0x82, 0x2c, 0x9f, 0xbb, 0x7f, 0x54, 0xa0, 0xe9, 0xdd, 0x3e, 0xce, 0x2b, 0xf9, 0x0a, 0xb6, 0x6f,
	0x44, 0x2a, 0x87, 0x33, 0xc6, 0x33, 0x3d, 0x4e, 0x52, 0xf9, 0x2b, 0xd7, 0x32, 0x89, 0xf1, 0xe4,
	0x06, 0xdd, 0x32, 0xd8, 0xd1, 0x22, 0x44, 0xf6, 0x60, 0xe3, 0x84, 0xfb, 0x63, 0x71, 0x7d, 0xfd,
	0xbd, 0x27, 0xfc, 0x24, 0x0e, 0x8a, 0x77, 0x64, 0xb9, 0xdc, 0xfd, 0xab, 0x06, 0xcd, 0x73, 0xcf,
	0xfb, 0x57, 0x67, 0x8b, 0xb3, 0x54, 0x3a, 0x7b, 0x1f, 0x9c, 0x50, 0x0b, 0xb4, 0xc5, 0x92, 0x09,
	0x1a, 0x6a, 0xd1, 0x66, 0xa8, 0x45, 0xee, 0xe6, 0xf5, 0x84, 0x74, 0xa0, 0x55, 0xe2, 0x3c, 0x1a,
	0xa2, 0x87, 0x16, 0x05, 0x4b, 0x38, 0x8a, 0x86, 0xe4, 0x18, 0x5a, 0x2a, 0x1b, 0xb0, 0x49, 0x9a,
	0x0c, 0x65, 0x28, 0x94, 0x5b, 0xef, 0xd4, 0xf6, 0x9c, 0xde, 0x0b, 0x3c, 0xb6, 0xb4, 0xb5, 0xef,
	0x65, 0x83, 0x2b, 0xcb, 0xe8, 0xc7, 0x3a, 0x9d, 0x51, 0x47, 0xcd, 0x2b, 0x84, 0xc2, 0x56, 0x20,
	0x86, 0x3c, 0x0b, 0x35, 0x5b, 0xd8, 0x0b, 0x47, 0xcd, 0xe9, 0x75, 0xef, 0x6f, 0xa5, 0xfc, 0x54,
	0x4e, 0xf2, 0x36, 0xd9, 0x1d, 0xe8, 0xa6, 0x95, 0xcf, 0x8f, 0xd9, 0xf5, 0x61, 0xeb, 0x01, 0x26,
	0xf9, 0x08, 0x36, 0x22, 0x3e, 0x65, 0x59, 0xc8, 0x06, 0x52, 0xb3, 0x94, 0x6b, 0x81, 0x8d, 0xaa,
	0xd3, 0x56, 0xc4, 0xa7, 0x6f, 0xc2, 0x63, 0xa9, 0x29, 0xd7, 0x25, 0x2d, 0x58, 0xa0, 0x55, 0x4b,
	0xda, 0x69, 0x41, 0xdb, 0x1d, 0x40, 0x7b, 0xf9, 0x66, 0xa4, 0x0d, 0xb5, 0xb7, 0x62, 0x66, 0xdf,
	0xd0, 0xfc, 0x91, 0x1c, 0xc0, 0xca, 0x0d, 0x0f, 0x33, 0xb3, 0xc5, 0x7f, 0xbb, 0x90, 0x11, 0x1c,
	0x56, 0x0f, 0x2a, 0xdd, 0xdf, 0xeb, 0xd0, 0x3a, 0x17, 0x3c, 0xd4, 0x63, 0xfb, 0x17, 0x7f, 0x0c,
	0x1b, 0x63, 0x5c, 0xb3, 0xfc, 0x4f, 0x94, 0xbe, 0x50, 0x6e, 0xa5, 0x53, 0xdb, 0x6b, 0xd2, 0x27,
	0xa6, 0xec, 0xd9, 0x2a, 0xf9, 0x12, 0xb6, 0xb3, 0x49, 0xc0, 0xb5, 0x28, 0xf3, 0x93, 0x29, 0xe1,
	0x9b, 0x84, 0x58, 0xa7, 0xc4, 0x60, 0x45, 0x84, 0x7a, 0xc2, 0x57, 0xe4, 0x15, 0x3c, 0xf3, 0xc3,
	0x24, 0x0b, 0x58, 0x20, 0x15, 0x1f, 0x84, 0x82, 0x4d, 0x44, 0x2a, 0x93, 0xc0, 0xc8, 0xcc, 0xfc,
	0xed, 0x20, 0xe1, 0xd4, 0xe0, 0x57, 0x08, 0x17, 0x52, 0x93, 0x33, 0x0f, 0x49, 0x4d, 0x6c, 0xef,
	0x20, 0xe1, 0xbe, 0xf4, 0x00, 0x5c, 0xeb, 0x73, 0xc8, 0x65, 0x98, 0xa5, 0x82, 0xe9, 0x71, 0x2a,
	0xd4, 0x38, 0x09, 0x03, 0x9b, 0xe4, 0x3b, 0x06, 0xff, 0xd6, 0xc0, 0xd7, 0x05, 0x4a, 0x0e, 0xe1,
	0x59, 0x2a, 0x7e, 0xc9, 0xf2, 0x74, 0xba, 0x2f, 0xcd, 0x13, 0xbe, 0x4a, 0x9f, 0x5a, 0xc2, 0x43,
	0xda, 0x48, 0xc6, 0x32, 0xca, 0x22, 0x56, 0xec, 0x31, 0xd7, 0xae, 0xe1, 0xb1, 0x4f, 0x2d, 0x81,
	0x1a, 0xfc, 0x8e, 0xd6, 0x9f, 0x64, 0x2c, 0xd3, 0x32, 0xb4, 0x2f, 0xec, 0x82, 0xb6, 0x61, 0xce,
	0xf5, 0x27, 0xd9, 0x9b, 0x39, 0x3e, 0xd7, 0x7e, 0x03, 0xbb, 0x91, 0x88, 0x92, 0x74, 0xc6, 0xf8,
	0x0d, 0x97, 0x21, 0xf6, 0x6a, 0x2e, 0x6e, 0xa2, 0xd8, 0x35, 0x8c, 0xa3, 0x82, 0x50, 0xaa, 0xbb,
	0xbf, 0x55, 0x61, 0xd5, 0xce, 0x41, 0x07, 0x6a, 0xea, 0x25, 0xc7, 0xde, 0x3a, 0xbd, 0x27, 0x38,
	0x54, 0x65, 0x10, 0xd3, 0x1c, 0x22, 0xcf, 0xa1, 0x3a, 0x9a, 0xda, 0xd7, 0x68, 0xdd, 0xe4, 0x9b,
	0x4d, 0x30, 0x5a, 0x1d, 0x4d, 0x11, 0x9e, 0x61, 0x9b, 0x4a, 0x78, 0x56, 0xc2, 0x33, 0xf2, 0x39,
	0x10, 0x4c, 0x89, 0x80, 0xc5, 0x42, 0xdf, 0x26, 0xe9, 0x5b, 0x26, 0x03, 0xe5, 0xae, 0xe1, 0xa8,
	0xb5, 0x0d, 0x72, 0x69, 0x80, 0x8b, 0x20, 0xff, 0xa0, 0xd7, 0xc6, 0x4a, 0xe1, 0xe5, 0x0b, 0x37,
	0xe5, 0x88, 0xd3, 0x1c, 0x42, 0xbf, 0xb7, 0x53, 0xbc, 0x61, 0xe9, 0xb7, 0x48, 0x54, 0x9a, 0x43,
	0xe4, 0x13, 0x58, 0x35, 0x23, 0x8c, 0x9f, 0x4b, 0xc7, 0x66, 0xf2, 0xe2, 0xf0, 0x53, 0x4b, 0xf8,
	0xf4, 0x10, 0x5a, 0x8b, 0x59, 0x4d, 0x5a, 0xd0, 0xa0, 0x7d, 0xaf, 0x4f, 0x7f, 0xec, 0x9f, 0xb6,
	0xdf, 0x21, 0x1b, 0xe0, 0x5c, 0xf5, 0x29, 0xf3, 0xfa, 0x9e, 0x77, 0xf1, 0xfa, 0xb2, 0x5d, 0x21,
	0x0e, 0xac, 0xe5, 0x85, 0xef, 0xfa, 0x3f, 0xb5, 0xab, 0xc7, 0x8d, 0x9f, 0x57, 0xf1, 0xc3, 0xa9,
	0x06, 0xe6, 0xf7, 0xeb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x62, 0xfe, 0xbe, 0xf0, 0x08,
	0x00, 0x00,
}