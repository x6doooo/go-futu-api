package go_futu_api

import (
	"github.com/x6doooo/go-futu-api/pb/Qot_Common"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetBasicQot"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetPlateSecurity"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetPlateSet"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetSecuritySnapshot"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetStaticInfo"
	"github.com/x6doooo/go-futu-api/pb/Qot_GetTradeDate"
)

func (me *Client) GetBasicQot(list []*Qot_Common.Security) (resp *Qot_GetBasicQot.Response, err error) {
	req := &Qot_GetBasicQot.Request{
		C2S: &Qot_GetBasicQot.C2S{
			SecurityList: list,
		},
	}
	resp = &Qot_GetBasicQot.Response{}
	err = me.DoRequest(uint32(3004), req, resp)
	return
}


func (me *Client) GetTradeDate(
	market Qot_Common.QotMarket,
	beginTime string,
	endTime string,
) (resp *Qot_GetTradeDate.Response, err error) {
	marketInt32 := int32(market)
	req := &Qot_GetTradeDate.Request{
		C2S: &Qot_GetTradeDate.C2S{
			Market: &marketInt32,
			BeginTime: &beginTime,
			EndTime: &endTime,
		},
	}
	resp = &Qot_GetTradeDate.Response{}
	err = me.DoRequest(uint32(3200), req, resp)
	return
}


func (me *Client) GetSecuritySnapshot(list []*Qot_Common.Security) (resp *Qot_GetSecuritySnapshot.Response, err error) {
	req := &Qot_GetSecuritySnapshot.Request{
		C2S: &Qot_GetSecuritySnapshot.C2S{
			SecurityList: list,
		},
	}
	resp = &Qot_GetSecuritySnapshot.Response{}
	err = me.DoRequest(uint32(3203), req, resp)
	return
}

func (me *Client) GetPlateSet(
	market Qot_Common.QotMarket,
	plateSetType Qot_Common.PlateSetType,
) (resp *Qot_GetPlateSet.Response, err error) {
	marketParam := int32(market)
	plateSetTypeParam := int32(plateSetType)
	req := &Qot_GetPlateSet.Request{
		C2S: &Qot_GetPlateSet.C2S{
			Market:       &marketParam,
			PlateSetType: &plateSetTypeParam,
		},
	}
	resp = &Qot_GetPlateSet.Response{}
	err = me.DoRequest(uint32(3204), req, resp)
	return
}

func (me *Client) GetPlateSecurity(plate *Qot_Common.Security) (resp *Qot_GetPlateSecurity.Response, err error) {
	req := &Qot_GetPlateSecurity.Request{
		C2S: &Qot_GetPlateSecurity.C2S{
			Plate: plate,
		},
	}
	resp = &Qot_GetPlateSecurity.Response{}
	err = me.DoRequest(uint32(3205), req, resp)
	return
}

func (me *Client) GetStaticInfo(
	market *Qot_Common.QotMarket,
	secType *Qot_Common.SecurityType,
	list []*Qot_Common.Security,
) (resp *Qot_GetStaticInfo.Response, err error) {
	c2s := &Qot_GetStaticInfo.C2S{}
	if market != nil {
		marketInt32 := int32(*market)
		c2s.Market = &marketInt32
	}
	if secType != nil {
		secTypeInt32 := int32(*secType)
		c2s.SecType = &secTypeInt32
	}
	if len(list) != 0 {
		c2s.SecurityList = list
	}
	req := &Qot_GetStaticInfo.Request{
		C2S: c2s,
	}
	resp = &Qot_GetStaticInfo.Response{}
	err = me.DoRequest(uint32(3202), req, resp)
	return
}