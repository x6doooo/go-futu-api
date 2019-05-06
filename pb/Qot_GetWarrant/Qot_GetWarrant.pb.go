// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetWarrant.proto

package Qot_GetWarrant

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

type C2S struct {
	Begin     *int32 `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`
	Num       *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	SortField *int32 `protobuf:"varint,3,req,name=sortField" json:"sortField,omitempty"`
	Ascend    *bool  `protobuf:"varint,4,req,name=ascend" json:"ascend,omitempty"`
	//以下为筛选条件，可选字段，不填表示不过滤
	Owner                 *Security `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	TypeList              []int32   `protobuf:"varint,6,rep,name=typeList" json:"typeList,omitempty"`
	IssuerList            []int32   `protobuf:"varint,7,rep,name=issuerList" json:"issuerList,omitempty"`
	MaturityTimeMin       *string   `protobuf:"bytes,8,opt,name=maturityTimeMin" json:"maturityTimeMin,omitempty"`
	MaturityTimeMax       *string   `protobuf:"bytes,9,opt,name=maturityTimeMax" json:"maturityTimeMax,omitempty"`
	IpoPeriod             *int32    `protobuf:"varint,10,opt,name=ipoPeriod" json:"ipoPeriod,omitempty"`
	PriceType             *int32    `protobuf:"varint,11,opt,name=priceType" json:"priceType,omitempty"`
	Status                *int32    `protobuf:"varint,12,opt,name=status" json:"status,omitempty"`
	CurPriceMin           *float64  `protobuf:"fixed64,13,opt,name=curPriceMin" json:"curPriceMin,omitempty"`
	CurPriceMax           *float64  `protobuf:"fixed64,14,opt,name=curPriceMax" json:"curPriceMax,omitempty"`
	StrikePriceMin        *float64  `protobuf:"fixed64,15,opt,name=strikePriceMin" json:"strikePriceMin,omitempty"`
	StrikePriceMax        *float64  `protobuf:"fixed64,16,opt,name=strikePriceMax" json:"strikePriceMax,omitempty"`
	StreetMin             *float64  `protobuf:"fixed64,17,opt,name=streetMin" json:"streetMin,omitempty"`
	StreetMax             *float64  `protobuf:"fixed64,18,opt,name=streetMax" json:"streetMax,omitempty"`
	ConversionMin         *float64  `protobuf:"fixed64,19,opt,name=conversionMin" json:"conversionMin,omitempty"`
	ConversionMax         *float64  `protobuf:"fixed64,20,opt,name=conversionMax" json:"conversionMax,omitempty"`
	VolMin                *uint64   `protobuf:"varint,21,opt,name=volMin" json:"volMin,omitempty"`
	VolMax                *uint64   `protobuf:"varint,22,opt,name=volMax" json:"volMax,omitempty"`
	PremiumMin            *float64  `protobuf:"fixed64,23,opt,name=premiumMin" json:"premiumMin,omitempty"`
	PremiumMax            *float64  `protobuf:"fixed64,24,opt,name=premiumMax" json:"premiumMax,omitempty"`
	LeverageRatioMin      *float64  `protobuf:"fixed64,25,opt,name=leverageRatioMin" json:"leverageRatioMin,omitempty"`
	LeverageRatioMax      *float64  `protobuf:"fixed64,26,opt,name=leverageRatioMax" json:"leverageRatioMax,omitempty"`
	DeltaMin              *float64  `protobuf:"fixed64,27,opt,name=deltaMin" json:"deltaMin,omitempty"`
	DeltaMax              *float64  `protobuf:"fixed64,28,opt,name=deltaMax" json:"deltaMax,omitempty"`
	ImpliedMin            *float64  `protobuf:"fixed64,29,opt,name=impliedMin" json:"impliedMin,omitempty"`
	ImpliedMax            *float64  `protobuf:"fixed64,30,opt,name=impliedMax" json:"impliedMax,omitempty"`
	RecoveryPriceMin      *float64  `protobuf:"fixed64,31,opt,name=recoveryPriceMin" json:"recoveryPriceMin,omitempty"`
	RecoveryPriceMax      *float64  `protobuf:"fixed64,32,opt,name=recoveryPriceMax" json:"recoveryPriceMax,omitempty"`
	PriceRecoveryRatioMin *float64  `protobuf:"fixed64,33,opt,name=priceRecoveryRatioMin" json:"priceRecoveryRatioMin,omitempty"`
	PriceRecoveryRatioMax *float64  `protobuf:"fixed64,34,opt,name=priceRecoveryRatioMax" json:"priceRecoveryRatioMax,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}  `json:"-"`
	XXX_unrecognized      []byte    `json:"-"`
	XXX_sizecache         int32     `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a92bbe5a0ce691f, []int{0}
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

func (m *C2S) GetBegin() int32 {
	if m != nil && m.Begin != nil {
		return *m.Begin
	}
	return 0
}

func (m *C2S) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *C2S) GetSortField() int32 {
	if m != nil && m.SortField != nil {
		return *m.SortField
	}
	return 0
}

func (m *C2S) GetAscend() bool {
	if m != nil && m.Ascend != nil {
		return *m.Ascend
	}
	return false
}

func (m *C2S) GetOwner() *Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *C2S) GetTypeList() []int32 {
	if m != nil {
		return m.TypeList
	}
	return nil
}

func (m *C2S) GetIssuerList() []int32 {
	if m != nil {
		return m.IssuerList
	}
	return nil
}

func (m *C2S) GetMaturityTimeMin() string {
	if m != nil && m.MaturityTimeMin != nil {
		return *m.MaturityTimeMin
	}
	return ""
}

func (m *C2S) GetMaturityTimeMax() string {
	if m != nil && m.MaturityTimeMax != nil {
		return *m.MaturityTimeMax
	}
	return ""
}

func (m *C2S) GetIpoPeriod() int32 {
	if m != nil && m.IpoPeriod != nil {
		return *m.IpoPeriod
	}
	return 0
}

func (m *C2S) GetPriceType() int32 {
	if m != nil && m.PriceType != nil {
		return *m.PriceType
	}
	return 0
}

func (m *C2S) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *C2S) GetCurPriceMin() float64 {
	if m != nil && m.CurPriceMin != nil {
		return *m.CurPriceMin
	}
	return 0
}

func (m *C2S) GetCurPriceMax() float64 {
	if m != nil && m.CurPriceMax != nil {
		return *m.CurPriceMax
	}
	return 0
}

func (m *C2S) GetStrikePriceMin() float64 {
	if m != nil && m.StrikePriceMin != nil {
		return *m.StrikePriceMin
	}
	return 0
}

func (m *C2S) GetStrikePriceMax() float64 {
	if m != nil && m.StrikePriceMax != nil {
		return *m.StrikePriceMax
	}
	return 0
}

func (m *C2S) GetStreetMin() float64 {
	if m != nil && m.StreetMin != nil {
		return *m.StreetMin
	}
	return 0
}

func (m *C2S) GetStreetMax() float64 {
	if m != nil && m.StreetMax != nil {
		return *m.StreetMax
	}
	return 0
}

func (m *C2S) GetConversionMin() float64 {
	if m != nil && m.ConversionMin != nil {
		return *m.ConversionMin
	}
	return 0
}

func (m *C2S) GetConversionMax() float64 {
	if m != nil && m.ConversionMax != nil {
		return *m.ConversionMax
	}
	return 0
}

func (m *C2S) GetVolMin() uint64 {
	if m != nil && m.VolMin != nil {
		return *m.VolMin
	}
	return 0
}

func (m *C2S) GetVolMax() uint64 {
	if m != nil && m.VolMax != nil {
		return *m.VolMax
	}
	return 0
}

func (m *C2S) GetPremiumMin() float64 {
	if m != nil && m.PremiumMin != nil {
		return *m.PremiumMin
	}
	return 0
}

func (m *C2S) GetPremiumMax() float64 {
	if m != nil && m.PremiumMax != nil {
		return *m.PremiumMax
	}
	return 0
}

func (m *C2S) GetLeverageRatioMin() float64 {
	if m != nil && m.LeverageRatioMin != nil {
		return *m.LeverageRatioMin
	}
	return 0
}

func (m *C2S) GetLeverageRatioMax() float64 {
	if m != nil && m.LeverageRatioMax != nil {
		return *m.LeverageRatioMax
	}
	return 0
}

func (m *C2S) GetDeltaMin() float64 {
	if m != nil && m.DeltaMin != nil {
		return *m.DeltaMin
	}
	return 0
}

func (m *C2S) GetDeltaMax() float64 {
	if m != nil && m.DeltaMax != nil {
		return *m.DeltaMax
	}
	return 0
}

func (m *C2S) GetImpliedMin() float64 {
	if m != nil && m.ImpliedMin != nil {
		return *m.ImpliedMin
	}
	return 0
}

func (m *C2S) GetImpliedMax() float64 {
	if m != nil && m.ImpliedMax != nil {
		return *m.ImpliedMax
	}
	return 0
}

func (m *C2S) GetRecoveryPriceMin() float64 {
	if m != nil && m.RecoveryPriceMin != nil {
		return *m.RecoveryPriceMin
	}
	return 0
}

func (m *C2S) GetRecoveryPriceMax() float64 {
	if m != nil && m.RecoveryPriceMax != nil {
		return *m.RecoveryPriceMax
	}
	return 0
}

func (m *C2S) GetPriceRecoveryRatioMin() float64 {
	if m != nil && m.PriceRecoveryRatioMin != nil {
		return *m.PriceRecoveryRatioMin
	}
	return 0
}

func (m *C2S) GetPriceRecoveryRatioMax() float64 {
	if m != nil && m.PriceRecoveryRatioMax != nil {
		return *m.PriceRecoveryRatioMax
	}
	return 0
}

type WarrantData struct {
	//静态数据项
	Stock              *Security `protobuf:"bytes,1,req,name=stock" json:"stock,omitempty"`
	Owner              *Security `protobuf:"bytes,2,req,name=owner" json:"owner,omitempty"`
	Type               *int32    `protobuf:"varint,3,req,name=type" json:"type,omitempty"`
	Issuer             *int32    `protobuf:"varint,4,req,name=issuer" json:"issuer,omitempty"`
	MaturityTime       *string   `protobuf:"bytes,5,req,name=maturityTime" json:"maturityTime,omitempty"`
	MaturityTimestamp  *float64  `protobuf:"fixed64,6,opt,name=maturityTimestamp" json:"maturityTimestamp,omitempty"`
	ListTime           *string   `protobuf:"bytes,7,req,name=listTime" json:"listTime,omitempty"`
	ListTimestamp      *float64  `protobuf:"fixed64,8,opt,name=listTimestamp" json:"listTimestamp,omitempty"`
	LastTradeTime      *string   `protobuf:"bytes,9,req,name=lastTradeTime" json:"lastTradeTime,omitempty"`
	LastTradeTimestamp *float64  `protobuf:"fixed64,10,opt,name=lastTradeTimestamp" json:"lastTradeTimestamp,omitempty"`
	RecoveryPrice      *float64  `protobuf:"fixed64,11,opt,name=recoveryPrice" json:"recoveryPrice,omitempty"`
	ConversionRatio    *float64  `protobuf:"fixed64,12,req,name=conversionRatio" json:"conversionRatio,omitempty"`
	LotSize            *int32    `protobuf:"varint,13,req,name=lotSize" json:"lotSize,omitempty"`
	StrikePrice        *float64  `protobuf:"fixed64,14,req,name=strikePrice" json:"strikePrice,omitempty"`
	LastClosePrice     *float64  `protobuf:"fixed64,15,req,name=lastClosePrice" json:"lastClosePrice,omitempty"`
	Name               *string   `protobuf:"bytes,16,req,name=name" json:"name,omitempty"`
	//动态数据项
	CurPrice             *float64 `protobuf:"fixed64,17,req,name=curPrice" json:"curPrice,omitempty"`
	PriceChangeVal       *float64 `protobuf:"fixed64,18,req,name=priceChangeVal" json:"priceChangeVal,omitempty"`
	ChangeRate           *float64 `protobuf:"fixed64,19,req,name=changeRate" json:"changeRate,omitempty"`
	Status               *int32   `protobuf:"varint,20,req,name=status" json:"status,omitempty"`
	BidPrice             *float64 `protobuf:"fixed64,21,req,name=bidPrice" json:"bidPrice,omitempty"`
	AskPrice             *float64 `protobuf:"fixed64,22,req,name=askPrice" json:"askPrice,omitempty"`
	BidVol               *int64   `protobuf:"varint,23,req,name=bidVol" json:"bidVol,omitempty"`
	AskVol               *int64   `protobuf:"varint,24,req,name=askVol" json:"askVol,omitempty"`
	Volume               *int64   `protobuf:"varint,25,req,name=volume" json:"volume,omitempty"`
	Turnover             *float64 `protobuf:"fixed64,26,req,name=turnover" json:"turnover,omitempty"`
	Score                *float64 `protobuf:"fixed64,27,req,name=score" json:"score,omitempty"`
	Premium              *float64 `protobuf:"fixed64,28,req,name=premium" json:"premium,omitempty"`
	BreakEvenPoint       *float64 `protobuf:"fixed64,29,req,name=breakEvenPoint" json:"breakEvenPoint,omitempty"`
	Leverage             *float64 `protobuf:"fixed64,30,req,name=leverage" json:"leverage,omitempty"`
	Ipop                 *float64 `protobuf:"fixed64,31,req,name=ipop" json:"ipop,omitempty"`
	PriceRecoveryRatio   *float64 `protobuf:"fixed64,32,opt,name=priceRecoveryRatio" json:"priceRecoveryRatio,omitempty"`
	ConversionPrice      *float64 `protobuf:"fixed64,33,req,name=conversionPrice" json:"conversionPrice,omitempty"`
	StreetRate           *float64 `protobuf:"fixed64,34,req,name=streetRate" json:"streetRate,omitempty"`
	StreetVol            *int64   `protobuf:"varint,35,req,name=streetVol" json:"streetVol,omitempty"`
	Amplitude            *float64 `protobuf:"fixed64,36,req,name=amplitude" json:"amplitude,omitempty"`
	IssueSize            *int64   `protobuf:"varint,37,req,name=issueSize" json:"issueSize,omitempty"`
	HighPrice            *float64 `protobuf:"fixed64,39,req,name=highPrice" json:"highPrice,omitempty"`
	LowPrice             *float64 `protobuf:"fixed64,40,req,name=lowPrice" json:"lowPrice,omitempty"`
	ImpliedVolatility    *float64 `protobuf:"fixed64,41,opt,name=impliedVolatility" json:"impliedVolatility,omitempty"`
	Delta                *float64 `protobuf:"fixed64,42,opt,name=delta" json:"delta,omitempty"`
	EffectiveLeverage    *float64 `protobuf:"fixed64,43,req,name=effectiveLeverage" json:"effectiveLeverage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WarrantData) Reset()         { *m = WarrantData{} }
func (m *WarrantData) String() string { return proto.CompactTextString(m) }
func (*WarrantData) ProtoMessage()    {}
func (*WarrantData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a92bbe5a0ce691f, []int{1}
}

func (m *WarrantData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarrantData.Unmarshal(m, b)
}
func (m *WarrantData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarrantData.Marshal(b, m, deterministic)
}
func (m *WarrantData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarrantData.Merge(m, src)
}
func (m *WarrantData) XXX_Size() int {
	return xxx_messageInfo_WarrantData.Size(m)
}
func (m *WarrantData) XXX_DiscardUnknown() {
	xxx_messageInfo_WarrantData.DiscardUnknown(m)
}

var xxx_messageInfo_WarrantData proto.InternalMessageInfo

func (m *WarrantData) GetStock() *Security {
	if m != nil {
		return m.Stock
	}
	return nil
}

func (m *WarrantData) GetOwner() *Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *WarrantData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *WarrantData) GetIssuer() int32 {
	if m != nil && m.Issuer != nil {
		return *m.Issuer
	}
	return 0
}

func (m *WarrantData) GetMaturityTime() string {
	if m != nil && m.MaturityTime != nil {
		return *m.MaturityTime
	}
	return ""
}

func (m *WarrantData) GetMaturityTimestamp() float64 {
	if m != nil && m.MaturityTimestamp != nil {
		return *m.MaturityTimestamp
	}
	return 0
}

func (m *WarrantData) GetListTime() string {
	if m != nil && m.ListTime != nil {
		return *m.ListTime
	}
	return ""
}

func (m *WarrantData) GetListTimestamp() float64 {
	if m != nil && m.ListTimestamp != nil {
		return *m.ListTimestamp
	}
	return 0
}

func (m *WarrantData) GetLastTradeTime() string {
	if m != nil && m.LastTradeTime != nil {
		return *m.LastTradeTime
	}
	return ""
}

func (m *WarrantData) GetLastTradeTimestamp() float64 {
	if m != nil && m.LastTradeTimestamp != nil {
		return *m.LastTradeTimestamp
	}
	return 0
}

func (m *WarrantData) GetRecoveryPrice() float64 {
	if m != nil && m.RecoveryPrice != nil {
		return *m.RecoveryPrice
	}
	return 0
}

func (m *WarrantData) GetConversionRatio() float64 {
	if m != nil && m.ConversionRatio != nil {
		return *m.ConversionRatio
	}
	return 0
}

func (m *WarrantData) GetLotSize() int32 {
	if m != nil && m.LotSize != nil {
		return *m.LotSize
	}
	return 0
}

func (m *WarrantData) GetStrikePrice() float64 {
	if m != nil && m.StrikePrice != nil {
		return *m.StrikePrice
	}
	return 0
}

func (m *WarrantData) GetLastClosePrice() float64 {
	if m != nil && m.LastClosePrice != nil {
		return *m.LastClosePrice
	}
	return 0
}

func (m *WarrantData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *WarrantData) GetCurPrice() float64 {
	if m != nil && m.CurPrice != nil {
		return *m.CurPrice
	}
	return 0
}

func (m *WarrantData) GetPriceChangeVal() float64 {
	if m != nil && m.PriceChangeVal != nil {
		return *m.PriceChangeVal
	}
	return 0
}

func (m *WarrantData) GetChangeRate() float64 {
	if m != nil && m.ChangeRate != nil {
		return *m.ChangeRate
	}
	return 0
}

func (m *WarrantData) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *WarrantData) GetBidPrice() float64 {
	if m != nil && m.BidPrice != nil {
		return *m.BidPrice
	}
	return 0
}

func (m *WarrantData) GetAskPrice() float64 {
	if m != nil && m.AskPrice != nil {
		return *m.AskPrice
	}
	return 0
}

func (m *WarrantData) GetBidVol() int64 {
	if m != nil && m.BidVol != nil {
		return *m.BidVol
	}
	return 0
}

func (m *WarrantData) GetAskVol() int64 {
	if m != nil && m.AskVol != nil {
		return *m.AskVol
	}
	return 0
}

func (m *WarrantData) GetVolume() int64 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *WarrantData) GetTurnover() float64 {
	if m != nil && m.Turnover != nil {
		return *m.Turnover
	}
	return 0
}

func (m *WarrantData) GetScore() float64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *WarrantData) GetPremium() float64 {
	if m != nil && m.Premium != nil {
		return *m.Premium
	}
	return 0
}

func (m *WarrantData) GetBreakEvenPoint() float64 {
	if m != nil && m.BreakEvenPoint != nil {
		return *m.BreakEvenPoint
	}
	return 0
}

func (m *WarrantData) GetLeverage() float64 {
	if m != nil && m.Leverage != nil {
		return *m.Leverage
	}
	return 0
}

func (m *WarrantData) GetIpop() float64 {
	if m != nil && m.Ipop != nil {
		return *m.Ipop
	}
	return 0
}

func (m *WarrantData) GetPriceRecoveryRatio() float64 {
	if m != nil && m.PriceRecoveryRatio != nil {
		return *m.PriceRecoveryRatio
	}
	return 0
}

func (m *WarrantData) GetConversionPrice() float64 {
	if m != nil && m.ConversionPrice != nil {
		return *m.ConversionPrice
	}
	return 0
}

func (m *WarrantData) GetStreetRate() float64 {
	if m != nil && m.StreetRate != nil {
		return *m.StreetRate
	}
	return 0
}

func (m *WarrantData) GetStreetVol() int64 {
	if m != nil && m.StreetVol != nil {
		return *m.StreetVol
	}
	return 0
}

func (m *WarrantData) GetAmplitude() float64 {
	if m != nil && m.Amplitude != nil {
		return *m.Amplitude
	}
	return 0
}

func (m *WarrantData) GetIssueSize() int64 {
	if m != nil && m.IssueSize != nil {
		return *m.IssueSize
	}
	return 0
}

func (m *WarrantData) GetHighPrice() float64 {
	if m != nil && m.HighPrice != nil {
		return *m.HighPrice
	}
	return 0
}

func (m *WarrantData) GetLowPrice() float64 {
	if m != nil && m.LowPrice != nil {
		return *m.LowPrice
	}
	return 0
}

func (m *WarrantData) GetImpliedVolatility() float64 {
	if m != nil && m.ImpliedVolatility != nil {
		return *m.ImpliedVolatility
	}
	return 0
}

func (m *WarrantData) GetDelta() float64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *WarrantData) GetEffectiveLeverage() float64 {
	if m != nil && m.EffectiveLeverage != nil {
		return *m.EffectiveLeverage
	}
	return 0
}

type S2C struct {
	LastPage             *bool          `protobuf:"varint,1,req,name=lastPage" json:"lastPage,omitempty"`
	AllCount             *int32         `protobuf:"varint,2,req,name=allCount" json:"allCount,omitempty"`
	WarrantDataList      []*WarrantData `protobuf:"bytes,3,rep,name=warrantDataList" json:"warrantDataList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a92bbe5a0ce691f, []int{2}
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

func (m *S2C) GetLastPage() bool {
	if m != nil && m.LastPage != nil {
		return *m.LastPage
	}
	return false
}

func (m *S2C) GetAllCount() int32 {
	if m != nil && m.AllCount != nil {
		return *m.AllCount
	}
	return 0
}

func (m *S2C) GetWarrantDataList() []*WarrantData {
	if m != nil {
		return m.WarrantDataList
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
	return fileDescriptor_3a92bbe5a0ce691f, []int{3}
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
	return fileDescriptor_3a92bbe5a0ce691f, []int{4}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetWarrant.C2S")
	proto.RegisterType((*WarrantData)(nil), "Qot_GetWarrant.WarrantData")
	proto.RegisterType((*S2C)(nil), "Qot_GetWarrant.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetWarrant.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetWarrant.Response")
}

func init() { proto.RegisterFile("Qot_GetWarrant.proto", fileDescriptor_3a92bbe5a0ce691f) }

var fileDescriptor_3a92bbe5a0ce691f = []byte{
	// 1154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x72, 0x13, 0x37,
	0x14, 0x9e, 0xb5, 0x63, 0x62, 0x2b, 0x21, 0x09, 0x22, 0xc0, 0xc1, 0x40, 0x58, 0x5c, 0x68, 0xb7,
	0xb4, 0xcd, 0x30, 0x1e, 0xae, 0x7a, 0xeb, 0xd2, 0xde, 0xc0, 0x4c, 0xba, 0x66, 0xe8, 0x65, 0x47,
	0x59, 0x8b, 0xa0, 0xc9, 0x7a, 0xe5, 0x4a, 0xb2, 0xd9, 0xf4, 0xb6, 0x7d, 0x83, 0x3e, 0x62, 0x5f,
	0xa4, 0x73, 0x8e, 0xf6, 0xdf, 0xa6, 0x77, 0xfb, 0xfd, 0xe8, 0xac, 0x56, 0x3a, 0xfb, 0x49, 0xec,
	0xf4, 0x57, 0xed, 0x7e, 0xff, 0x45, 0xba, 0xdf, 0x84, 0x31, 0x22, 0x73, 0xe7, 0x2b, 0xa3, 0x9d,
	0xe6, 0x47, 0x6d, 0x76, 0x7c, 0x38, 0xd3, 0xcb, 0xa5, 0xce, 0xbc, 0x3a, 0x3e, 0x41, 0xb5, 0xc9,
	0x4c, 0xfe, 0x19, 0xb1, 0xfe, 0x6c, 0x3a, 0xe7, 0xa7, 0x6c, 0x70, 0x29, 0xaf, 0x54, 0x06, 0x41,
	0xd8, 0x8b, 0x06, 0xb1, 0x07, 0xfc, 0x84, 0xf5, 0xb3, 0xf5, 0x12, 0x7a, 0xc4, 0xe1, 0x23, 0x7f,
	0xcc, 0x46, 0x56, 0x1b, 0xf7, 0xb3, 0x92, 0xe9, 0x02, 0xfa, 0xc4, 0xd7, 0x04, 0xbf, 0xcf, 0x6e,
	0x09, 0x9b, 0xc8, 0x6c, 0x01, 0x7b, 0x61, 0x2f, 0x1a, 0xc6, 0x05, 0xe2, 0x2f, 0xd9, 0x40, 0x7f,
	0xce, 0xa4, 0x81, 0x41, 0x18, 0x44, 0x07, 0xd3, 0xd3, 0xf3, 0xc6, 0x3c, 0xe6, 0x32, 0x59, 0x1b,
	0xe5, 0x6e, 0x62, 0x6f, 0xe1, 0x63, 0x36, 0x74, 0x37, 0x2b, 0xf9, 0x56, 0x59, 0x07, 0xb7, 0xc2,
	0x7e, 0x34, 0x88, 0x2b, 0xcc, 0xcf, 0x18, 0x53, 0xd6, 0xae, 0xa5, 0x21, 0x75, 0x9f, 0xd4, 0x06,
	0xc3, 0x23, 0x76, 0xbc, 0x14, 0x8e, 0xca, 0xbd, 0x57, 0x4b, 0xf9, 0x4e, 0x65, 0x30, 0x0c, 0x83,
	0x68, 0x14, 0x77, 0xe9, 0x2d, 0xa7, 0xc8, 0x61, 0xb4, 0xc3, 0x29, 0x72, 0xfc, 0x62, 0xb5, 0xd2,
	0x17, 0xd2, 0x28, 0xbd, 0x00, 0x16, 0x06, 0xf8, 0xc5, 0x15, 0x81, 0xea, 0xca, 0xa8, 0x44, 0xbe,
	0xbf, 0x59, 0x49, 0x38, 0xf0, 0x6a, 0x45, 0xe0, 0x7a, 0x58, 0x27, 0xdc, 0xda, 0xc2, 0x21, 0x49,
	0x05, 0xe2, 0x21, 0x3b, 0x48, 0xd6, 0xe6, 0x02, 0x7d, 0x38, 0xc7, 0xdb, 0x61, 0x10, 0x05, 0x71,
	0x93, 0x6a, 0x39, 0x44, 0x0e, 0x47, 0x1d, 0x87, 0xc8, 0xf9, 0xd7, 0xec, 0xc8, 0x3a, 0xa3, 0xae,
	0x65, 0x55, 0xe6, 0x98, 0x4c, 0x1d, 0xb6, 0xeb, 0x13, 0x39, 0x9c, 0x6c, 0xfb, 0xfc, 0x77, 0x5a,
	0x67, 0xa4, 0x74, 0x58, 0xea, 0x0e, 0x59, 0x6a, 0xa2, 0xa1, 0x8a, 0x1c, 0x78, 0x4b, 0x15, 0x39,
	0x7f, 0xce, 0x6e, 0x27, 0x3a, 0xdb, 0x48, 0x63, 0x95, 0xce, 0x70, 0xfc, 0x5d, 0x72, 0xb4, 0xc9,
	0x8e, 0x4b, 0xe4, 0x70, 0xba, 0xe5, 0x12, 0x39, 0xae, 0xd9, 0x46, 0xa7, 0x58, 0xe4, 0x5e, 0x18,
	0x44, 0x7b, 0x71, 0x81, 0x4a, 0x5e, 0xe4, 0x70, 0xbf, 0xe6, 0x45, 0x8e, 0x3d, 0xb1, 0x32, 0x72,
	0xa9, 0xd6, 0x4b, 0x1c, 0xf3, 0x80, 0x4a, 0x36, 0x98, 0xa6, 0x2e, 0x72, 0x80, 0xb6, 0x2e, 0x72,
	0xfe, 0x92, 0x9d, 0xa4, 0x72, 0x23, 0x8d, 0xb8, 0x92, 0xb1, 0x70, 0x4a, 0x63, 0x95, 0x87, 0xe4,
	0xda, 0xe2, 0xb7, 0xbd, 0x22, 0x87, 0xf1, 0x2e, 0xaf, 0xc8, 0xb1, 0x8f, 0x17, 0x32, 0x75, 0x02,
	0xeb, 0x3d, 0x22, 0x4f, 0x85, 0x6b, 0x4d, 0xe4, 0xf0, 0xb8, 0xa9, 0xf9, 0xef, 0x51, 0xcb, 0x55,
	0xaa, 0xe4, 0x02, 0x47, 0x3e, 0xf1, 0xf3, 0xad, 0x99, 0xa6, 0x2e, 0x72, 0x38, 0x6b, 0xeb, 0xfe,
	0x7b, 0x8c, 0x4c, 0xf4, 0x46, 0x9a, 0x9b, 0xaa, 0x33, 0x9e, 0xfa, 0x39, 0x76, 0xf9, 0x6d, 0xaf,
	0xc8, 0x21, 0xdc, 0xe5, 0x15, 0x39, 0x7f, 0xcd, 0xee, 0x51, 0x63, 0xc7, 0x85, 0x50, 0x2d, 0xd6,
	0x33, 0x1a, 0xb0, 0x5b, 0xfc, 0xc2, 0x28, 0x91, 0xc3, 0xe4, 0x8b, 0xa3, 0x44, 0x3e, 0xf9, 0x97,
	0xb1, 0x83, 0x22, 0xc1, 0x7e, 0x12, 0x4e, 0x60, 0x7e, 0x58, 0xa7, 0x93, 0x6b, 0x4a, 0xa7, 0x2f,
	0xe6, 0x07, 0x59, 0xea, 0xac, 0xe9, 0xfd, 0x9f, 0xd7, 0x67, 0x0d, 0x67, 0x7b, 0x98, 0x2d, 0x45,
	0x90, 0xd1, 0x33, 0xf6, 0x99, 0x4f, 0x14, 0xca, 0xb0, 0x41, 0x5c, 0x20, 0x3e, 0x61, 0x87, 0xcd,
	0x68, 0x80, 0x41, 0xd8, 0x8b, 0x46, 0x71, 0x8b, 0xe3, 0xdf, 0xb3, 0x3b, 0x4d, 0x6c, 0x9d, 0x58,
	0xae, 0xe0, 0x16, 0x7d, 0xe9, 0xb6, 0x80, 0x5d, 0x90, 0x2a, 0xeb, 0xa8, 0xda, 0x3e, 0x55, 0xab,
	0x30, 0xfe, 0x2b, 0xe5, 0xb3, 0xaf, 0x32, 0xf4, 0xff, 0x4a, 0x8b, 0x24, 0x97, 0xb0, 0xee, 0xbd,
	0x11, 0x0b, 0x49, 0x65, 0x46, 0x54, 0xa6, 0x4d, 0xf2, 0x73, 0xc6, 0x5b, 0x84, 0x2f, 0xc8, 0xa8,
	0xe0, 0x0e, 0x05, 0xab, 0xb6, 0x76, 0x9f, 0x72, 0x2d, 0x88, 0xdb, 0x24, 0x26, 0x68, 0xfd, 0xe3,
	0xd2, 0xce, 0xc1, 0x61, 0xd8, 0x8b, 0x82, 0xb8, 0x4b, 0x73, 0x60, 0xfb, 0xa9, 0x76, 0x73, 0xf5,
	0xa7, 0x84, 0xdb, 0xb4, 0xa4, 0x25, 0xc4, 0x94, 0x6b, 0xa4, 0x10, 0x1c, 0xd1, 0xf8, 0x26, 0x85,
	0xe9, 0x85, 0x33, 0x9c, 0xa5, 0xda, 0x16, 0xa6, 0x63, 0x32, 0x75, 0x58, 0xdc, 0xc9, 0x4c, 0x2c,
	0x25, 0x9c, 0xd0, 0x02, 0xd0, 0x33, 0xae, 0x6f, 0x19, 0x98, 0x70, 0x87, 0x46, 0x55, 0x18, 0xeb,
	0x52, 0xeb, 0xcd, 0x3e, 0x89, 0xec, 0x4a, 0x7e, 0x10, 0x29, 0x70, 0x5f, 0xb7, 0xcd, 0xe2, 0xdf,
	0x96, 0x10, 0x88, 0x85, 0x93, 0x70, 0x97, 0x3c, 0x0d, 0xa6, 0x91, 0xf0, 0xa7, 0xbe, 0x5b, 0x8a,
	0x84, 0x1f, 0xb3, 0xe1, 0xa5, 0x5a, 0xf8, 0x77, 0xdf, 0xf3, 0xef, 0x2e, 0x31, 0x6a, 0xc2, 0x5e,
	0x7b, 0xed, 0xbe, 0xd7, 0x4a, 0x8c, 0xf5, 0x2e, 0xd5, 0xe2, 0x83, 0x4e, 0xe1, 0x41, 0xd8, 0x8b,
	0xfa, 0x71, 0x81, 0xfc, 0xc9, 0x7a, 0x8d, 0x3c, 0x78, 0xde, 0xa3, 0x22, 0x15, 0xd7, 0x4b, 0x09,
	0x0f, 0x3d, 0xef, 0x11, 0x9d, 0xa2, 0x6b, 0x93, 0xe1, 0x7e, 0xc1, 0xd8, 0xbf, 0xa3, 0xc4, 0x78,
	0xd6, 0xdb, 0x44, 0x1b, 0x09, 0x8f, 0x48, 0xf0, 0x00, 0x77, 0xa9, 0x48, 0x45, 0x78, 0x4c, 0x7c,
	0x09, 0x71, 0xad, 0x2e, 0x8d, 0x14, 0xd7, 0x6f, 0x36, 0x32, 0xbb, 0xd0, 0x2a, 0x73, 0xf0, 0xc4,
	0xaf, 0x55, 0x9b, 0xa5, 0x7e, 0x2e, 0x52, 0x10, 0xce, 0xfc, 0x3b, 0x4b, 0x8c, 0xfb, 0xa3, 0x56,
	0x7a, 0x05, 0x4f, 0x89, 0xa7, 0x67, 0xec, 0xcb, 0xed, 0xdf, 0xbf, 0xc8, 0x9f, 0x1d, 0x4a, 0xbb,
	0xe3, 0xfc, 0xf2, 0x3d, 0xeb, 0x76, 0x9c, 0x5f, 0xc5, 0x33, 0xc6, 0xfc, 0xe1, 0x44, 0xbb, 0x36,
	0xf1, 0xbb, 0x56, 0x33, 0xf5, 0x69, 0x86, 0x0b, 0xfa, 0x15, 0x2d, 0x5c, 0x4d, 0xa0, 0x2a, 0x30,
	0x4f, 0xdd, 0x7a, 0x21, 0xe1, 0x39, 0x0d, 0xae, 0x09, 0xba, 0x0f, 0x60, 0x22, 0x50, 0x3f, 0xbf,
	0xf0, 0x63, 0x2b, 0x02, 0xd5, 0x4f, 0xea, 0xea, 0x93, 0x9f, 0xdd, 0x37, 0x7e, 0x6c, 0x45, 0xd0,
	0x0a, 0xe9, 0xcf, 0x5e, 0x8c, 0x8a, 0x15, 0x2a, 0x30, 0x66, 0x47, 0x91, 0xe2, 0x1f, 0x74, 0x2a,
	0x9c, 0x4a, 0x95, 0xbb, 0x81, 0x6f, 0x7d, 0x76, 0x6c, 0x09, 0xb8, 0x87, 0x74, 0x62, 0xc0, 0x4b,
	0x72, 0x78, 0x80, 0x35, 0xe4, 0xc7, 0x8f, 0x32, 0x71, 0x6a, 0x23, 0xdf, 0x96, 0x5b, 0xf1, 0x1d,
	0xbd, 0x68, 0x5b, 0x98, 0xfc, 0x1d, 0xb0, 0xfe, 0x7c, 0x3a, 0xa3, 0x59, 0x09, 0xeb, 0x2e, 0xd0,
	0x1c, 0xd0, 0xbd, 0xad, 0xc2, 0xd4, 0xab, 0x69, 0x3a, 0xd3, 0xeb, 0xcc, 0x15, 0xd7, 0xc0, 0x0a,
	0xf3, 0x37, 0xec, 0xf8, 0x73, 0x1d, 0xd2, 0x74, 0x25, 0xeb, 0x87, 0xfd, 0xe8, 0x60, 0xfa, 0xe8,
	0xbc, 0x73, 0x37, 0x6d, 0x64, 0x79, 0xdc, 0x1d, 0x33, 0x79, 0xc5, 0xf6, 0x63, 0xf9, 0xc7, 0x5a,
	0x5a, 0xc7, 0x5f, 0xb0, 0x7e, 0x32, 0xb5, 0x45, 0xca, 0xdf, 0xed, 0x56, 0x99, 0x4d, 0xe7, 0x31,
	0xea, 0x93, 0xbf, 0x02, 0x36, 0x8c, 0xa5, 0x5d, 0xe9, 0xcc, 0xe2, 0x5e, 0xef, 0x1b, 0xe9, 0xe8,
	0xfe, 0x45, 0x77, 0xd7, 0x1f, 0xf7, 0x7e, 0x78, 0xfd, 0xea, 0x55, 0x5c, 0x92, 0xf8, 0x87, 0x18,
	0xe9, 0xde, 0xd9, 0x2b, 0xe8, 0xd1, 0x05, 0xaf, 0x40, 0xd8, 0xef, 0xd2, 0x98, 0x99, 0x5e, 0x60,
	0xfc, 0xe3, 0xe5, 0xac, 0x84, 0x38, 0x0b, 0x3b, 0x4d, 0x60, 0x8f, 0xee, 0xaa, 0x5b, 0xb3, 0x98,
	0x4f, 0x67, 0x31, 0xea, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xa8, 0xbd, 0x50, 0x81, 0x0b,
	0x00, 0x00,
}
