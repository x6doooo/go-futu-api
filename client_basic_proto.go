package go_futu_api

import (
	"github.com/x6doooo/go-futu-api/pb/InitConnect"
	"github.com/x6doooo/go-futu-api/pb/KeepAlive"
	"time"
)

func (me *Client) InitConnect() (resp *InitConnect.Response, err error) {
	req := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  &ClientID,
			ClientVer: &ClientVer,
		},
	}
	resp = &InitConnect.Response{}
	err = me.DoRequest(uint32(1001), req, resp)
	return
}

func (me *Client) KeepAlive() (resp *KeepAlive.Response, err error) {
	t := int64(time.Now().Second())
	req := &KeepAlive.Request{
		C2S: &KeepAlive.C2S{
			Time: &t,
		},
	}
	resp = &KeepAlive.Response{}
	err = me.DoRequest(uint32(1004), req, resp)
	return
}
