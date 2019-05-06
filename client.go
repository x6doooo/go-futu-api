package go_futu_api

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/x6doooo/go-futu-api/pb/InitConnect"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
)

var (
	ClientVer = int32(100)
	ClientID  = "go-futu-api-100"
)

type Client struct {
	packChanMap sync.Map
	socket      *SocketClient
}

func NewClient(addr string) *Client {
	return &Client{
		socket: NewSocketClient(addr),
	}
}

func (me *Client) Run() (err error) {
	err = me.socket.Connect()
	if err != nil {
		return
	}
	go func() {
		log.Println("futu api client run, addr:", me.socket.addr)
		me.WatchMessage()
	}()
	return
}

func (me *Client) SyncDo(requestPack *FutuPack) (responsePack *FutuPack, err error) {
	ch := make(chan *FutuPack)
	defer close(ch)
	err = requestPack.Send(me.socket.conn)
	sid := requestPack.GetSerialNoStr()
	me.packChanMap.Store(sid, ch)
	//log.Println("send msg", sid)
	responsePack = <-ch
	return
}

//func (me *Client) DoRequest(protoId uint32, request proto.Message) (*FutuPack, error) {
//	pack := &FutuPack{}
//	pack.SetProtoID(protoId)
//	body, err := proto.Marshal(request)
//	if err != nil {
//		return nil, err
//	}
//	pack.SetBody(body)
//	return me.SyncDo(pack)
//}

func (me *Client) DoRequest(protoId uint32, request proto.Message, response proto.Message) error {
	pack := &FutuPack{}
	pack.SetProtoID(protoId)
	body, err := proto.Marshal(request)
	if err != nil {
		return err
	}
	pack.SetBody(body)
	respPack, err := me.SyncDo(pack)
	if err != nil {
		return err
	}
	b := respPack.GetBody()
	return proto.Unmarshal(b, response)
}

func (me *Client) WatchMessage() {
	scanner := bufio.NewScanner(me.socket.conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})
	//buf := make([]byte, 64 * 1024)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	//scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		scannedPack := new(FutuPack)
		err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			return
		}

		resp := &InitConnect.Response{}
		err = proto.Unmarshal(scannedPack.GetBody(), resp)

		sid := scannedPack.GetSerialNoStr()
		chInterface, ok := me.packChanMap.Load(sid)
		if ok {
			ch := chInterface.(chan *FutuPack)
			ch <- scannedPack
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("无效数据包", err)
	}
}
