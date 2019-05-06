// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSecuritySnapshot.proto

package Qot_GetSecuritySnapshot

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/x6doooo/go-futu-api/pb/Qot_Common"
	"math"
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

type C2S struct {
	SecurityList         []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{0}
}

func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

// 正股类型额外数据
type EquitySnapshotExData struct {
	IssuedShares         *int64   `protobuf:"varint,1,req,name=issuedShares" json:"issuedShares,omitempty"`
	IssuedMarketVal      *float64 `protobuf:"fixed64,2,req,name=issuedMarketVal" json:"issuedMarketVal,omitempty"`
	NetAsset             *float64 `protobuf:"fixed64,3,req,name=netAsset" json:"netAsset,omitempty"`
	NetProfit            *float64 `protobuf:"fixed64,4,req,name=netProfit" json:"netProfit,omitempty"`
	EarningsPershare     *float64 `protobuf:"fixed64,5,req,name=earningsPershare" json:"earningsPershare,omitempty"`
	OutstandingShares    *int64   `protobuf:"varint,6,req,name=outstandingShares" json:"outstandingShares,omitempty"`
	OutstandingMarketVal *float64 `protobuf:"fixed64,7,req,name=outstandingMarketVal" json:"outstandingMarketVal,omitempty"`
	NetAssetPershare     *float64 `protobuf:"fixed64,8,req,name=netAssetPershare" json:"netAssetPershare,omitempty"`
	EyRate               *float64 `protobuf:"fixed64,9,req,name=eyRate" json:"eyRate,omitempty"`
	PeRate               *float64 `protobuf:"fixed64,10,req,name=peRate" json:"peRate,omitempty"`
	PbRate               *float64 `protobuf:"fixed64,11,req,name=pbRate" json:"pbRate,omitempty"`
	PeTTMRate            *float64 `protobuf:"fixed64,12,req,name=peTTMRate" json:"peTTMRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquitySnapshotExData) Reset()         { *m = EquitySnapshotExData{} }
func (m *EquitySnapshotExData) String() string { return proto.CompactTextString(m) }
func (*EquitySnapshotExData) ProtoMessage()    {}
func (*EquitySnapshotExData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{1}
}

func (m *EquitySnapshotExData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquitySnapshotExData.Unmarshal(m, b)
}
func (m *EquitySnapshotExData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquitySnapshotExData.Marshal(b, m, deterministic)
}
func (m *EquitySnapshotExData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquitySnapshotExData.Merge(m, src)
}
func (m *EquitySnapshotExData) XXX_Size() int {
	return xxx_messageInfo_EquitySnapshotExData.Size(m)
}
func (m *EquitySnapshotExData) XXX_DiscardUnknown() {
	xxx_messageInfo_EquitySnapshotExData.DiscardUnknown(m)
}

var xxx_messageInfo_EquitySnapshotExData proto.InternalMessageInfo

func (m *EquitySnapshotExData) GetIssuedShares() int64 {
	if m != nil && m.IssuedShares != nil {
		return *m.IssuedShares
	}
	return 0
}

func (m *EquitySnapshotExData) GetIssuedMarketVal() float64 {
	if m != nil && m.IssuedMarketVal != nil {
		return *m.IssuedMarketVal
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetAsset() float64 {
	if m != nil && m.NetAsset != nil {
		return *m.NetAsset
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetProfit() float64 {
	if m != nil && m.NetProfit != nil {
		return *m.NetProfit
	}
	return 0
}

func (m *EquitySnapshotExData) GetEarningsPershare() float64 {
	if m != nil && m.EarningsPershare != nil {
		return *m.EarningsPershare
	}
	return 0
}

func (m *EquitySnapshotExData) GetOutstandingShares() int64 {
	if m != nil && m.OutstandingShares != nil {
		return *m.OutstandingShares
	}
	return 0
}

func (m *EquitySnapshotExData) GetOutstandingMarketVal() float64 {
	if m != nil && m.OutstandingMarketVal != nil {
		return *m.OutstandingMarketVal
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetAssetPershare() float64 {
	if m != nil && m.NetAssetPershare != nil {
		return *m.NetAssetPershare
	}
	return 0
}

func (m *EquitySnapshotExData) GetEyRate() float64 {
	if m != nil && m.EyRate != nil {
		return *m.EyRate
	}
	return 0
}

func (m *EquitySnapshotExData) GetPeRate() float64 {
	if m != nil && m.PeRate != nil {
		return *m.PeRate
	}
	return 0
}

func (m *EquitySnapshotExData) GetPbRate() float64 {
	if m != nil && m.PbRate != nil {
		return *m.PbRate
	}
	return 0
}

func (m *EquitySnapshotExData) GetPeTTMRate() float64 {
	if m != nil && m.PeTTMRate != nil {
		return *m.PeTTMRate
	}
	return 0
}

// 涡轮类型额外数据
type WarrantSnapshotExData struct {
	ConversionRate       *float64             `protobuf:"fixed64,1,req,name=conversionRate" json:"conversionRate,omitempty"`
	WarrantType          *int32               `protobuf:"varint,2,req,name=warrantType" json:"warrantType,omitempty"`
	StrikePrice          *float64             `protobuf:"fixed64,3,req,name=strikePrice" json:"strikePrice,omitempty"`
	MaturityTime         *string              `protobuf:"bytes,4,req,name=maturityTime" json:"maturityTime,omitempty"`
	EndTradeTime         *string              `protobuf:"bytes,5,req,name=endTradeTime" json:"endTradeTime,omitempty"`
	Owner                *Qot_Common.Security `protobuf:"bytes,6,req,name=owner" json:"owner,omitempty"`
	RecoveryPrice        *float64             `protobuf:"fixed64,7,req,name=recoveryPrice" json:"recoveryPrice,omitempty"`
	StreetVolumn         *int64               `protobuf:"varint,8,req,name=streetVolumn" json:"streetVolumn,omitempty"`
	IssueVolumn          *int64               `protobuf:"varint,9,req,name=issueVolumn" json:"issueVolumn,omitempty"`
	StreetRate           *float64             `protobuf:"fixed64,10,req,name=streetRate" json:"streetRate,omitempty"`
	Delta                *float64             `protobuf:"fixed64,11,req,name=delta" json:"delta,omitempty"`
	ImpliedVolatility    *float64             `protobuf:"fixed64,12,req,name=impliedVolatility" json:"impliedVolatility,omitempty"`
	Premium              *float64             `protobuf:"fixed64,13,req,name=premium" json:"premium,omitempty"`
	MaturityTimestamp    *float64             `protobuf:"fixed64,14,opt,name=maturityTimestamp" json:"maturityTimestamp,omitempty"`
	EndTradeTimestamp    *float64             `protobuf:"fixed64,15,opt,name=endTradeTimestamp" json:"endTradeTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WarrantSnapshotExData) Reset()         { *m = WarrantSnapshotExData{} }
func (m *WarrantSnapshotExData) String() string { return proto.CompactTextString(m) }
func (*WarrantSnapshotExData) ProtoMessage()    {}
func (*WarrantSnapshotExData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{2}
}

func (m *WarrantSnapshotExData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarrantSnapshotExData.Unmarshal(m, b)
}
func (m *WarrantSnapshotExData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarrantSnapshotExData.Marshal(b, m, deterministic)
}
func (m *WarrantSnapshotExData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarrantSnapshotExData.Merge(m, src)
}
func (m *WarrantSnapshotExData) XXX_Size() int {
	return xxx_messageInfo_WarrantSnapshotExData.Size(m)
}
func (m *WarrantSnapshotExData) XXX_DiscardUnknown() {
	xxx_messageInfo_WarrantSnapshotExData.DiscardUnknown(m)
}

var xxx_messageInfo_WarrantSnapshotExData proto.InternalMessageInfo

func (m *WarrantSnapshotExData) GetConversionRate() float64 {
	if m != nil && m.ConversionRate != nil {
		return *m.ConversionRate
	}
	return 0
}

func (m *WarrantSnapshotExData) GetWarrantType() int32 {
	if m != nil && m.WarrantType != nil {
		return *m.WarrantType
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStrikePrice() float64 {
	if m != nil && m.StrikePrice != nil {
		return *m.StrikePrice
	}
	return 0
}

func (m *WarrantSnapshotExData) GetMaturityTime() string {
	if m != nil && m.MaturityTime != nil {
		return *m.MaturityTime
	}
	return ""
}

func (m *WarrantSnapshotExData) GetEndTradeTime() string {
	if m != nil && m.EndTradeTime != nil {
		return *m.EndTradeTime
	}
	return ""
}

func (m *WarrantSnapshotExData) GetOwner() *Qot_Common.Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *WarrantSnapshotExData) GetRecoveryPrice() float64 {
	if m != nil && m.RecoveryPrice != nil {
		return *m.RecoveryPrice
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStreetVolumn() int64 {
	if m != nil && m.StreetVolumn != nil {
		return *m.StreetVolumn
	}
	return 0
}

func (m *WarrantSnapshotExData) GetIssueVolumn() int64 {
	if m != nil && m.IssueVolumn != nil {
		return *m.IssueVolumn
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStreetRate() float64 {
	if m != nil && m.StreetRate != nil {
		return *m.StreetRate
	}
	return 0
}

func (m *WarrantSnapshotExData) GetDelta() float64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *WarrantSnapshotExData) GetImpliedVolatility() float64 {
	if m != nil && m.ImpliedVolatility != nil {
		return *m.ImpliedVolatility
	}
	return 0
}

func (m *WarrantSnapshotExData) GetPremium() float64 {
	if m != nil && m.Premium != nil {
		return *m.Premium
	}
	return 0
}

func (m *WarrantSnapshotExData) GetMaturityTimestamp() float64 {
	if m != nil && m.MaturityTimestamp != nil {
		return *m.MaturityTimestamp
	}
	return 0
}

func (m *WarrantSnapshotExData) GetEndTradeTimestamp() float64 {
	if m != nil && m.EndTradeTimestamp != nil {
		return *m.EndTradeTimestamp
	}
	return 0
}

// 期权类型额外数据
type OptionSnapshotExData struct {
	Type                 *int32               `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Owner                *Qot_Common.Security `protobuf:"bytes,2,req,name=owner" json:"owner,omitempty"`
	StrikeTime           *string              `protobuf:"bytes,3,req,name=strikeTime" json:"strikeTime,omitempty"`
	StrikePrice          *float64             `protobuf:"fixed64,4,req,name=strikePrice" json:"strikePrice,omitempty"`
	ContractSize         *int32               `protobuf:"varint,5,req,name=contractSize" json:"contractSize,omitempty"`
	OpenInterest         *int32               `protobuf:"varint,6,req,name=openInterest" json:"openInterest,omitempty"`
	ImpliedVolatility    *float64             `protobuf:"fixed64,7,req,name=impliedVolatility" json:"impliedVolatility,omitempty"`
	Premium              *float64             `protobuf:"fixed64,8,req,name=premium" json:"premium,omitempty"`
	Delta                *float64             `protobuf:"fixed64,9,req,name=delta" json:"delta,omitempty"`
	Gamma                *float64             `protobuf:"fixed64,10,req,name=gamma" json:"gamma,omitempty"`
	Vega                 *float64             `protobuf:"fixed64,11,req,name=vega" json:"vega,omitempty"`
	Theta                *float64             `protobuf:"fixed64,12,req,name=theta" json:"theta,omitempty"`
	Rho                  *float64             `protobuf:"fixed64,13,req,name=rho" json:"rho,omitempty"`
	StrikeTimestamp      *float64             `protobuf:"fixed64,14,opt,name=strikeTimestamp" json:"strikeTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OptionSnapshotExData) Reset()         { *m = OptionSnapshotExData{} }
func (m *OptionSnapshotExData) String() string { return proto.CompactTextString(m) }
func (*OptionSnapshotExData) ProtoMessage()    {}
func (*OptionSnapshotExData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{3}
}

func (m *OptionSnapshotExData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionSnapshotExData.Unmarshal(m, b)
}
func (m *OptionSnapshotExData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionSnapshotExData.Marshal(b, m, deterministic)
}
func (m *OptionSnapshotExData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionSnapshotExData.Merge(m, src)
}
func (m *OptionSnapshotExData) XXX_Size() int {
	return xxx_messageInfo_OptionSnapshotExData.Size(m)
}
func (m *OptionSnapshotExData) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionSnapshotExData.DiscardUnknown(m)
}

var xxx_messageInfo_OptionSnapshotExData proto.InternalMessageInfo

func (m *OptionSnapshotExData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *OptionSnapshotExData) GetOwner() *Qot_Common.Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *OptionSnapshotExData) GetStrikeTime() string {
	if m != nil && m.StrikeTime != nil {
		return *m.StrikeTime
	}
	return ""
}

func (m *OptionSnapshotExData) GetStrikePrice() float64 {
	if m != nil && m.StrikePrice != nil {
		return *m.StrikePrice
	}
	return 0
}

func (m *OptionSnapshotExData) GetContractSize() int32 {
	if m != nil && m.ContractSize != nil {
		return *m.ContractSize
	}
	return 0
}

func (m *OptionSnapshotExData) GetOpenInterest() int32 {
	if m != nil && m.OpenInterest != nil {
		return *m.OpenInterest
	}
	return 0
}

func (m *OptionSnapshotExData) GetImpliedVolatility() float64 {
	if m != nil && m.ImpliedVolatility != nil {
		return *m.ImpliedVolatility
	}
	return 0
}

func (m *OptionSnapshotExData) GetPremium() float64 {
	if m != nil && m.Premium != nil {
		return *m.Premium
	}
	return 0
}

func (m *OptionSnapshotExData) GetDelta() float64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *OptionSnapshotExData) GetGamma() float64 {
	if m != nil && m.Gamma != nil {
		return *m.Gamma
	}
	return 0
}

func (m *OptionSnapshotExData) GetVega() float64 {
	if m != nil && m.Vega != nil {
		return *m.Vega
	}
	return 0
}

func (m *OptionSnapshotExData) GetTheta() float64 {
	if m != nil && m.Theta != nil {
		return *m.Theta
	}
	return 0
}

func (m *OptionSnapshotExData) GetRho() float64 {
	if m != nil && m.Rho != nil {
		return *m.Rho
	}
	return 0
}

func (m *OptionSnapshotExData) GetStrikeTimestamp() float64 {
	if m != nil && m.StrikeTimestamp != nil {
		return *m.StrikeTimestamp
	}
	return 0
}

//基本快照数据
type SnapshotBasicData struct {
	Security                *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	Type                    *int32    `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	IsSuspend               *bool     `protobuf:"varint,3,req,name=isSuspend" json:"isSuspend,omitempty"`
	ListTime                *string   `protobuf:"bytes,4,req,name=listTime" json:"listTime,omitempty"`
	LotSize                 *int32    `protobuf:"varint,5,req,name=lotSize" json:"lotSize,omitempty"`
	PriceSpread             *float64  `protobuf:"fixed64,6,req,name=priceSpread" json:"priceSpread,omitempty"`
	UpdateTime              *string   `protobuf:"bytes,7,req,name=updateTime" json:"updateTime,omitempty"`
	HighPrice               *float64  `protobuf:"fixed64,8,req,name=highPrice" json:"highPrice,omitempty"`
	OpenPrice               *float64  `protobuf:"fixed64,9,req,name=openPrice" json:"openPrice,omitempty"`
	LowPrice                *float64  `protobuf:"fixed64,10,req,name=lowPrice" json:"lowPrice,omitempty"`
	LastClosePrice          *float64  `protobuf:"fixed64,11,req,name=lastClosePrice" json:"lastClosePrice,omitempty"`
	CurPrice                *float64  `protobuf:"fixed64,12,req,name=curPrice" json:"curPrice,omitempty"`
	Volume                  *int64    `protobuf:"varint,13,req,name=volume" json:"volume,omitempty"`
	Turnover                *float64  `protobuf:"fixed64,14,req,name=turnover" json:"turnover,omitempty"`
	TurnoverRate            *float64  `protobuf:"fixed64,15,req,name=turnoverRate" json:"turnoverRate,omitempty"`
	ListTimestamp           *float64  `protobuf:"fixed64,16,opt,name=listTimestamp" json:"listTimestamp,omitempty"`
	UpdateTimestamp         *float64  `protobuf:"fixed64,17,opt,name=updateTimestamp" json:"updateTimestamp,omitempty"`
	AskPrice                *float64  `protobuf:"fixed64,18,opt,name=askPrice" json:"askPrice,omitempty"`
	BidPrice                *float64  `protobuf:"fixed64,19,opt,name=bidPrice" json:"bidPrice,omitempty"`
	AskVol                  *int64    `protobuf:"varint,20,opt,name=askVol" json:"askVol,omitempty"`
	BidVol                  *int64    `protobuf:"varint,21,opt,name=bidVol" json:"bidVol,omitempty"`
	EnableMargin            *bool     `protobuf:"varint,22,opt,name=enableMargin" json:"enableMargin,omitempty"`
	MortgageRatio           *float64  `protobuf:"fixed64,23,opt,name=mortgageRatio" json:"mortgageRatio,omitempty"`
	LongMarginInitialRatio  *float64  `protobuf:"fixed64,24,opt,name=longMarginInitialRatio" json:"longMarginInitialRatio,omitempty"`
	EnableShortSell         *bool     `protobuf:"varint,25,opt,name=enableShortSell" json:"enableShortSell,omitempty"`
	ShortSellRate           *float64  `protobuf:"fixed64,26,opt,name=shortSellRate" json:"shortSellRate,omitempty"`
	ShortAvailableVolume    *int64    `protobuf:"varint,27,opt,name=shortAvailableVolume" json:"shortAvailableVolume,omitempty"`
	ShortMarginInitialRatio *float64  `protobuf:"fixed64,28,opt,name=shortMarginInitialRatio" json:"shortMarginInitialRatio,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}  `json:"-"`
	XXX_unrecognized        []byte    `json:"-"`
	XXX_sizecache           int32     `json:"-"`
}

func (m *SnapshotBasicData) Reset()         { *m = SnapshotBasicData{} }
func (m *SnapshotBasicData) String() string { return proto.CompactTextString(m) }
func (*SnapshotBasicData) ProtoMessage()    {}
func (*SnapshotBasicData) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{4}
}

func (m *SnapshotBasicData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBasicData.Unmarshal(m, b)
}
func (m *SnapshotBasicData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBasicData.Marshal(b, m, deterministic)
}
func (m *SnapshotBasicData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBasicData.Merge(m, src)
}
func (m *SnapshotBasicData) XXX_Size() int {
	return xxx_messageInfo_SnapshotBasicData.Size(m)
}
func (m *SnapshotBasicData) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBasicData.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBasicData proto.InternalMessageInfo

func (m *SnapshotBasicData) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SnapshotBasicData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SnapshotBasicData) GetIsSuspend() bool {
	if m != nil && m.IsSuspend != nil {
		return *m.IsSuspend
	}
	return false
}

func (m *SnapshotBasicData) GetListTime() string {
	if m != nil && m.ListTime != nil {
		return *m.ListTime
	}
	return ""
}

func (m *SnapshotBasicData) GetLotSize() int32 {
	if m != nil && m.LotSize != nil {
		return *m.LotSize
	}
	return 0
}

func (m *SnapshotBasicData) GetPriceSpread() float64 {
	if m != nil && m.PriceSpread != nil {
		return *m.PriceSpread
	}
	return 0
}

func (m *SnapshotBasicData) GetUpdateTime() string {
	if m != nil && m.UpdateTime != nil {
		return *m.UpdateTime
	}
	return ""
}

func (m *SnapshotBasicData) GetHighPrice() float64 {
	if m != nil && m.HighPrice != nil {
		return *m.HighPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetOpenPrice() float64 {
	if m != nil && m.OpenPrice != nil {
		return *m.OpenPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetLowPrice() float64 {
	if m != nil && m.LowPrice != nil {
		return *m.LowPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetLastClosePrice() float64 {
	if m != nil && m.LastClosePrice != nil {
		return *m.LastClosePrice
	}
	return 0
}

func (m *SnapshotBasicData) GetCurPrice() float64 {
	if m != nil && m.CurPrice != nil {
		return *m.CurPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetVolume() int64 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *SnapshotBasicData) GetTurnover() float64 {
	if m != nil && m.Turnover != nil {
		return *m.Turnover
	}
	return 0
}

func (m *SnapshotBasicData) GetTurnoverRate() float64 {
	if m != nil && m.TurnoverRate != nil {
		return *m.TurnoverRate
	}
	return 0
}

func (m *SnapshotBasicData) GetListTimestamp() float64 {
	if m != nil && m.ListTimestamp != nil {
		return *m.ListTimestamp
	}
	return 0
}

func (m *SnapshotBasicData) GetUpdateTimestamp() float64 {
	if m != nil && m.UpdateTimestamp != nil {
		return *m.UpdateTimestamp
	}
	return 0
}

func (m *SnapshotBasicData) GetAskPrice() float64 {
	if m != nil && m.AskPrice != nil {
		return *m.AskPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetBidPrice() float64 {
	if m != nil && m.BidPrice != nil {
		return *m.BidPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetAskVol() int64 {
	if m != nil && m.AskVol != nil {
		return *m.AskVol
	}
	return 0
}

func (m *SnapshotBasicData) GetBidVol() int64 {
	if m != nil && m.BidVol != nil {
		return *m.BidVol
	}
	return 0
}

func (m *SnapshotBasicData) GetEnableMargin() bool {
	if m != nil && m.EnableMargin != nil {
		return *m.EnableMargin
	}
	return false
}

func (m *SnapshotBasicData) GetMortgageRatio() float64 {
	if m != nil && m.MortgageRatio != nil {
		return *m.MortgageRatio
	}
	return 0
}

func (m *SnapshotBasicData) GetLongMarginInitialRatio() float64 {
	if m != nil && m.LongMarginInitialRatio != nil {
		return *m.LongMarginInitialRatio
	}
	return 0
}

func (m *SnapshotBasicData) GetEnableShortSell() bool {
	if m != nil && m.EnableShortSell != nil {
		return *m.EnableShortSell
	}
	return false
}

func (m *SnapshotBasicData) GetShortSellRate() float64 {
	if m != nil && m.ShortSellRate != nil {
		return *m.ShortSellRate
	}
	return 0
}

func (m *SnapshotBasicData) GetShortAvailableVolume() int64 {
	if m != nil && m.ShortAvailableVolume != nil {
		return *m.ShortAvailableVolume
	}
	return 0
}

func (m *SnapshotBasicData) GetShortMarginInitialRatio() float64 {
	if m != nil && m.ShortMarginInitialRatio != nil {
		return *m.ShortMarginInitialRatio
	}
	return 0
}

type Snapshot struct {
	Basic                *SnapshotBasicData     `protobuf:"bytes,1,req,name=basic" json:"basic,omitempty"`
	EquityExData         *EquitySnapshotExData  `protobuf:"bytes,2,opt,name=equityExData" json:"equityExData,omitempty"`
	WarrantExData        *WarrantSnapshotExData `protobuf:"bytes,3,opt,name=warrantExData" json:"warrantExData,omitempty"`
	OptionExData         *OptionSnapshotExData  `protobuf:"bytes,4,opt,name=optionExData" json:"optionExData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{5}
}

func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetBasic() *SnapshotBasicData {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *Snapshot) GetEquityExData() *EquitySnapshotExData {
	if m != nil {
		return m.EquityExData
	}
	return nil
}

func (m *Snapshot) GetWarrantExData() *WarrantSnapshotExData {
	if m != nil {
		return m.WarrantExData
	}
	return nil
}

func (m *Snapshot) GetOptionExData() *OptionSnapshotExData {
	if m != nil {
		return m.OptionExData
	}
	return nil
}

type S2C struct {
	SnapshotList         []*Snapshot `protobuf:"bytes,1,rep,name=snapshotList" json:"snapshotList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{6}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetSnapshotList() []*Snapshot {
	if m != nil {
		return m.SnapshotList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{7}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_39e6284ba27d5399, []int{8}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetSecuritySnapshot.C2S")
	proto.RegisterType((*EquitySnapshotExData)(nil), "Qot_GetSecuritySnapshot.EquitySnapshotExData")
	proto.RegisterType((*WarrantSnapshotExData)(nil), "Qot_GetSecuritySnapshot.WarrantSnapshotExData")
	proto.RegisterType((*OptionSnapshotExData)(nil), "Qot_GetSecuritySnapshot.OptionSnapshotExData")
	proto.RegisterType((*SnapshotBasicData)(nil), "Qot_GetSecuritySnapshot.SnapshotBasicData")
	proto.RegisterType((*Snapshot)(nil), "Qot_GetSecuritySnapshot.Snapshot")
	proto.RegisterType((*S2C)(nil), "Qot_GetSecuritySnapshot.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSecuritySnapshot.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSecuritySnapshot.Response")
}

func init() { proto.RegisterFile("Qot_GetSecuritySnapshot.proto", fileDescriptor_39e6284ba27d5399) }

var fileDescriptor_39e6284ba27d5399 = []byte{
	// 1208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x6f, 0x8f, 0x1b, 0xb5,
	0x13, 0x56, 0xb2, 0x97, 0x26, 0xf1, 0x5d, 0xdb, 0xab, 0x7f, 0x69, 0xbb, 0xbf, 0xe3, 0xa8, 0x42,
	0x84, 0x50, 0x54, 0xd1, 0x53, 0x15, 0x21, 0x54, 0x78, 0x03, 0x25, 0x54, 0xa8, 0x52, 0x4f, 0xb4,
	0xde, 0xe8, 0x78, 0x89, 0x7c, 0x89, 0x49, 0xac, 0xdb, 0xb5, 0xb7, 0xb6, 0x73, 0xe5, 0xf8, 0x0c,
	0x7c, 0x22, 0x24, 0x78, 0xcd, 0x77, 0xe0, 0xcb, 0xa0, 0x99, 0xd9, 0x4d, 0x76, 0xf3, 0x47, 0xc7,
	0xbb, 0x9d, 0x67, 0x1e, 0xcf, 0x8e, 0xe7, 0x79, 0xd6, 0x6b, 0xf6, 0xf1, 0x3b, 0x1b, 0x7e, 0xfe,
	0x41, 0x85, 0x44, 0x4d, 0x97, 0x4e, 0x87, 0x9b, 0xc4, 0xc8, 0xdc, 0x2f, 0x6c, 0x38, 0xcb, 0x9d,
	0x0d, 0x96, 0x3f, 0xde, 0x93, 0x3e, 0x39, 0x1a, 0xdb, 0x2c, 0xb3, 0x86, 0x68, 0x27, 0xc7, 0x40,
	0xab, 0x22, 0x83, 0x6f, 0x58, 0x34, 0x1e, 0x25, 0xfc, 0x05, 0x3b, 0xf2, 0xc5, 0xd2, 0x37, 0xda,
	0x87, 0xb8, 0xd1, 0x8f, 0x86, 0x87, 0xa3, 0xde, 0x59, 0x85, 0x5f, 0x96, 0x16, 0x35, 0xe6, 0xe0,
	0xaf, 0x88, 0xf5, 0x5e, 0xbd, 0x5f, 0x56, 0xde, 0xf9, 0xea, 0xd7, 0xef, 0x65, 0x90, 0x7c, 0xc0,
	0x8e, 0xb4, 0xf7, 0x4b, 0x35, 0x4b, 0x16, 0xd2, 0x29, 0x1f, 0x37, 0xfa, 0xcd, 0x61, 0x24, 0x6a,
	0x18, 0x1f, 0xb2, 0xfb, 0x14, 0x9f, 0x4b, 0x77, 0xa5, 0xc2, 0x85, 0x4c, 0xe3, 0x66, 0xbf, 0x39,
	0x6c, 0x88, 0x4d, 0x98, 0x9f, 0xb0, 0x8e, 0x51, 0xe1, 0xa5, 0xf7, 0x2a, 0xc4, 0x11, 0x52, 0x56,
	0x31, 0x3f, 0x65, 0x5d, 0xa3, 0xc2, 0x5b, 0x67, 0x7f, 0xd1, 0x21, 0x3e, 0xc0, 0xe4, 0x1a, 0xe0,
	0x4f, 0xd9, 0xb1, 0x92, 0xce, 0x68, 0x33, 0xf7, 0x6f, 0x95, 0xf3, 0xf0, 0xe2, 0xb8, 0x85, 0xa4,
	0x2d, 0x9c, 0x7f, 0xce, 0x1e, 0xd8, 0x65, 0xf0, 0x41, 0x9a, 0x99, 0x36, 0xf3, 0xa2, 0xf1, 0x3b,
	0xd8, 0xf8, 0x76, 0x82, 0x8f, 0x58, 0xaf, 0x02, 0xae, 0xb7, 0xd0, 0xc6, 0xea, 0x3b, 0x73, 0xd0,
	0x4d, 0xd9, 0xf7, 0xaa, 0x9b, 0x0e, 0x75, 0xb3, 0x89, 0xf3, 0x47, 0xec, 0x8e, 0xba, 0x11, 0x32,
	0xa8, 0xb8, 0x8b, 0x8c, 0x22, 0x02, 0x3c, 0x57, 0x88, 0x33, 0xc2, 0x29, 0x42, 0xfc, 0x12, 0xf1,
	0xc3, 0x02, 0xc7, 0x08, 0xe6, 0x93, 0xab, 0xc9, 0xe4, 0x1c, 0x53, 0x47, 0x34, 0x9f, 0x15, 0x30,
	0xf8, 0xf3, 0x80, 0x3d, 0xfc, 0x49, 0x3a, 0x27, 0x4d, 0xd8, 0x50, 0xf0, 0x33, 0x76, 0x6f, 0x6a,
	0xcd, 0xb5, 0x72, 0x5e, 0x5b, 0x83, 0x8b, 0x1b, 0xb8, 0x78, 0x03, 0xe5, 0x7d, 0x76, 0xf8, 0x81,
	0x0a, 0x4c, 0x6e, 0x72, 0x85, 0x0a, 0xb6, 0x44, 0x15, 0x02, 0x86, 0x0f, 0x4e, 0x5f, 0xa9, 0xb7,
	0x4e, 0x4f, 0x55, 0x21, 0x60, 0x15, 0x02, 0xb7, 0x64, 0x32, 0xa0, 0xad, 0x26, 0x3a, 0x53, 0x28,
	0x63, 0x57, 0xd4, 0x30, 0xe0, 0x28, 0x33, 0x9b, 0x38, 0x39, 0x53, 0xc8, 0x69, 0x11, 0xa7, 0x8a,
	0xf1, 0xa7, 0xac, 0x65, 0x3f, 0x18, 0xe5, 0x50, 0xb5, 0x7d, 0x0e, 0x26, 0x0a, 0xff, 0x94, 0xdd,
	0x75, 0x6a, 0x6a, 0xaf, 0x95, 0xbb, 0xa1, 0xbe, 0x48, 0xb8, 0x3a, 0x08, 0x6f, 0xf5, 0xc1, 0x29,
	0x15, 0x2e, 0x6c, 0xba, 0xcc, 0x0c, 0xaa, 0x15, 0x89, 0x1a, 0x06, 0xfb, 0x43, 0xc3, 0x16, 0x94,
	0x2e, 0x52, 0xaa, 0x10, 0x7f, 0xc2, 0x18, 0xad, 0xa8, 0xe8, 0x56, 0x41, 0x78, 0x8f, 0xb5, 0x66,
	0x2a, 0x0d, 0xb2, 0x90, 0x8e, 0x02, 0xf0, 0xa3, 0xce, 0xf2, 0x54, 0xab, 0xd9, 0x85, 0x4d, 0x65,
	0xd0, 0xa9, 0x0e, 0x37, 0x85, 0x82, 0xdb, 0x09, 0x1e, 0xb3, 0x76, 0xee, 0x54, 0xa6, 0x97, 0x59,
	0x7c, 0x17, 0x39, 0x65, 0x08, 0x75, 0xaa, 0x93, 0xf4, 0x41, 0x66, 0x79, 0x7c, 0xaf, 0xdf, 0x80,
	0x3a, 0x5b, 0x09, 0x60, 0x57, 0x67, 0x4a, 0xec, 0xfb, 0xc4, 0xde, 0x4a, 0x0c, 0xfe, 0x88, 0x58,
	0xef, 0xc7, 0x3c, 0x68, 0x6b, 0x36, 0xec, 0xc3, 0xd9, 0x41, 0x00, 0x3f, 0x34, 0xd0, 0x0f, 0xf8,
	0xbc, 0x96, 0xa7, 0x79, 0xbb, 0x3c, 0x34, 0x32, 0x7d, 0x45, 0x62, 0x47, 0x28, 0x76, 0x05, 0xd9,
	0x34, 0xd5, 0xc1, 0x4e, 0x53, 0x4d, 0xad, 0x09, 0x4e, 0x4e, 0x43, 0xa2, 0x7f, 0x23, 0xc3, 0xb4,
	0x44, 0x0d, 0x03, 0x8e, 0xcd, 0x95, 0x79, 0x6d, 0x82, 0x72, 0xca, 0x07, 0xf4, 0x4d, 0x4b, 0xd4,
	0xb0, 0xdd, 0x32, 0xb4, 0xff, 0x83, 0x0c, 0x9d, 0xba, 0x0c, 0x2b, 0x91, 0xbb, 0x55, 0x91, 0x7b,
	0xac, 0x35, 0x97, 0x59, 0x26, 0x0b, 0x57, 0x50, 0x00, 0xd3, 0xbb, 0x56, 0xf3, 0xd2, 0x0f, 0xf8,
	0x0c, 0xcc, 0xb0, 0x50, 0x41, 0x16, 0x16, 0xa0, 0x80, 0x1f, 0xb3, 0xc8, 0x2d, 0x6c, 0x21, 0x39,
	0x3c, 0xc2, 0xb1, 0xba, 0x9e, 0x53, 0x55, 0xec, 0x4d, 0x78, 0xf0, 0x4f, 0x9b, 0x3d, 0x28, 0x65,
	0xfb, 0x4e, 0x7a, 0x3d, 0x45, 0xe5, 0x9e, 0xb3, 0x4e, 0x79, 0xc6, 0xa3, 0x7a, 0xfb, 0x84, 0x5a,
	0xb1, 0x56, 0x5a, 0x37, 0x2b, 0x5a, 0x9f, 0xb2, 0xae, 0xf6, 0xc9, 0xd2, 0xe7, 0xca, 0xcc, 0x50,
	0xbe, 0x8e, 0x58, 0x03, 0x70, 0xa0, 0xa7, 0xda, 0x87, 0xca, 0xc7, 0xbe, 0x8a, 0x61, 0x82, 0xa9,
	0xad, 0x4a, 0x56, 0x86, 0xa0, 0x79, 0x0e, 0xd2, 0x26, 0xb9, 0x53, 0x72, 0x86, 0x62, 0x35, 0x44,
	0x15, 0x02, 0xd7, 0x2c, 0xf3, 0x99, 0x0c, 0xe4, 0x9a, 0x36, 0xb9, 0x66, 0x8d, 0x40, 0x57, 0x0b,
	0x3d, 0x5f, 0x90, 0x67, 0x48, 0x9f, 0x35, 0x00, 0x59, 0x50, 0x9e, 0xb2, 0xa4, 0xd2, 0x1a, 0xc0,
	0x9e, 0xed, 0x07, 0x4a, 0x92, 0x58, 0xab, 0x18, 0x0e, 0xcb, 0x54, 0xfa, 0x30, 0x4e, 0xad, 0x2f,
	0x0c, 0x49, 0xca, 0x6d, 0xa0, 0x50, 0x63, 0xba, 0x74, 0xc4, 0x20, 0x19, 0x57, 0x31, 0x1c, 0xe0,
	0xd7, 0x70, 0x5c, 0x28, 0x14, 0x33, 0x12, 0x45, 0x04, 0x6b, 0xc2, 0xd2, 0x19, 0x38, 0x94, 0xe2,
	0x7b, 0xb4, 0xa6, 0x8c, 0xc1, 0xbf, 0xe5, 0x33, 0x1e, 0x2d, 0xf7, 0x31, 0x5f, 0xc3, 0xe0, 0xa0,
	0x2b, 0x67, 0x4b, 0x6e, 0x38, 0x46, 0x37, 0xd4, 0x41, 0x70, 0xcd, 0x7a, 0x4e, 0xc4, 0x7b, 0x40,
	0xae, 0xd9, 0x80, 0xa1, 0x1f, 0xe9, 0xaf, 0x68, 0x0f, 0x1c, 0x29, 0xab, 0x18, 0x72, 0x97, 0x7a,
	0x46, 0xb9, 0xff, 0x51, 0xae, 0x8c, 0x61, 0x7f, 0xd2, 0x5f, 0x5d, 0xd8, 0x34, 0xee, 0xf5, 0x1b,
	0xb0, 0x3f, 0x8a, 0x00, 0xbf, 0xd4, 0xf0, 0x09, 0xc5, 0x0f, 0x09, 0xa7, 0x88, 0x0e, 0x7c, 0x79,
	0x99, 0xaa, 0x73, 0xe9, 0xe6, 0xda, 0xc4, 0x8f, 0xfa, 0x8d, 0x61, 0x47, 0xd4, 0x30, 0xd8, 0x5b,
	0x66, 0x5d, 0x98, 0xcb, 0x39, 0xfc, 0x04, 0xb5, 0x8d, 0x1f, 0xd3, 0xde, 0x6a, 0x20, 0xff, 0x92,
	0x3d, 0x4a, 0x2d, 0xfe, 0x87, 0xe7, 0xda, 0xbc, 0x36, 0x3a, 0x68, 0x99, 0x12, 0x3d, 0x46, 0xfa,
	0x9e, 0x2c, 0xcc, 0x84, 0xde, 0x96, 0x2c, 0xac, 0x0b, 0x89, 0x4a, 0xd3, 0xf8, 0xff, 0xd8, 0xc4,
	0x26, 0x0c, 0x7d, 0xf8, 0x32, 0x40, 0x21, 0x4e, 0xa8, 0x8f, 0x1a, 0x08, 0x57, 0x06, 0x04, 0x5e,
	0x5e, 0x4b, 0x9d, 0x42, 0x81, 0x0b, 0xd2, 0xfb, 0x23, 0xdc, 0xf7, 0xce, 0x1c, 0x7f, 0xc1, 0x1e,
	0x23, 0xbe, 0xa3, 0xf9, 0x53, 0x7c, 0xc7, 0xbe, 0xf4, 0xe0, 0xef, 0x26, 0xeb, 0x94, 0x5f, 0x37,
	0xff, 0x96, 0xb5, 0x2e, 0xe1, 0x0b, 0x2f, 0xbe, 0xe8, 0xa7, 0x67, 0xfb, 0x6e, 0x94, 0x5b, 0xe7,
	0x81, 0xa0, 0x85, 0xfc, 0x1d, 0x3b, 0x52, 0x78, 0xd3, 0xa3, 0x03, 0x3e, 0x6e, 0xf6, 0x1b, 0xc3,
	0xc3, 0xd1, 0xb3, 0xbd, 0x85, 0x76, 0x5d, 0x0b, 0x45, 0xad, 0x04, 0x9f, 0xb0, 0xbb, 0xc5, 0x3d,
	0xa1, 0xa8, 0x19, 0x61, 0xcd, 0xb3, 0xbd, 0x35, 0x77, 0xde, 0x54, 0x44, 0xbd, 0x08, 0x34, 0x6a,
	0xf1, 0x8f, 0x54, 0x14, 0x3d, 0xb8, 0xa5, 0xd1, 0x5d, 0xbf, 0x2f, 0x51, 0x2b, 0x31, 0x78, 0xc3,
	0xa2, 0x64, 0x34, 0xe6, 0xaf, 0xd8, 0x91, 0x2f, 0x68, 0x95, 0x7b, 0xf2, 0x27, 0xb7, 0xce, 0x52,
	0xd4, 0x96, 0x0d, 0xbe, 0x62, 0x6d, 0xa1, 0xde, 0x2f, 0xe1, 0xdf, 0x72, 0xc6, 0xa2, 0xe9, 0xc8,
	0x17, 0xa2, 0x9c, 0xee, 0x2d, 0x34, 0x1e, 0x25, 0x02, 0x88, 0x83, 0xdf, 0x1b, 0xac, 0x23, 0x94,
	0xcf, 0xad, 0xf1, 0x8a, 0x3f, 0x61, 0x6d, 0xa7, 0xe8, 0xd6, 0x85, 0x7f, 0xd9, 0xaf, 0x0f, 0x9e,
	0x7d, 0xf1, 0xfc, 0xb9, 0x28, 0x41, 0xf8, 0xb0, 0x9c, 0x0a, 0xe7, 0x7e, 0x8e, 0x5a, 0x75, 0x45,
	0x11, 0xc1, 0x01, 0xab, 0x9c, 0x1b, 0xdb, 0x99, 0xc2, 0x81, 0xb7, 0x44, 0x19, 0x42, 0x3b, 0x7e,
	0x34, 0x2d, 0x26, 0xb6, 0xbf, 0x9d, 0x64, 0x34, 0x16, 0x40, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x43, 0x4e, 0x4c, 0xe5, 0x98, 0x0c, 0x00, 0x00,
}
