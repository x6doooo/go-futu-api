package go_futu_api

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"sync/atomic"
)

var serialNo uint32 = 0 //uint32(rand.Int31())

type FutuPack struct {
	szHeaderFlag  [2]uint8  // u8_t szHeaderFlag[2];
	nProtoID      uint32    // u32_t nProtoID;
	nProtoFmtType uint8     // u8_t nProtoFmtType;
	nProtoVer     uint8     // u8_t nProtoVer;
	nSerialNo     uint32    // u32_t nSerialNo;
	nBodyLen      uint32    // u32_t nBodyLen;
	arrBodySHA1   [20]uint8 // u8_t arrBodySHA1[20];
	arrReserved   [8]uint8  // u8_t arrReserved[8];
	body          []byte    // []byte add;
}

// SetProtoID set nProtoID
func (p *FutuPack) SetProtoID(nProtoID uint32) {
	p.nProtoID = nProtoID
}

// SetBody set body
func (p *FutuPack) SetBody(body []byte) {
	p.body = body
	p.nBodyLen = uint32(len(body))

	sha := sha1.New()
	sha.Write(p.body)
	arrBodySHA1 := sha.Sum(nil)
	copy(p.arrBodySHA1[:], arrBodySHA1)
}

func (p *FutuPack) GetSerialNoStr() string {
	return fmt.Sprintf("p_%d", p.nSerialNo)
}

// Pack pack
func (p *FutuPack) Send(writer io.Writer) error {
	var err error

	// szHeaderFlag
	p.szHeaderFlag[0] = uint8('F')
	p.szHeaderFlag[1] = uint8('T')

	// nProtoFmtType
	p.nProtoFmtType = uint8(0)

	// nProtoVer
	p.nProtoVer = 0

	// nSerialNo
	//serialNo += 1
	atomic.AddUint32(&serialNo, 1)
	p.nSerialNo = atomic.LoadUint32(&serialNo)

	// arrBodySHA1
	// fmt.Println("debug: ", p.arrBodySHA1)
	// os.Exit(1)

	// arrReserved
	var arrReservedTmp [20]uint8
	copy(p.arrReserved[:], arrReservedTmp[:20])

	err = binary.Write(writer, binary.LittleEndian, &p.szHeaderFlag)
	err = binary.Write(writer, binary.LittleEndian, &p.nProtoID)
	err = binary.Write(writer, binary.LittleEndian, &p.nProtoFmtType)
	err = binary.Write(writer, binary.LittleEndian, &p.nProtoVer)
	err = binary.Write(writer, binary.LittleEndian, &p.nSerialNo)
	err = binary.Write(writer, binary.LittleEndian, &p.nBodyLen)
	err = binary.Write(writer, binary.LittleEndian, &p.arrBodySHA1)
	err = binary.Write(writer, binary.LittleEndian, &p.arrReserved)

	err = binary.Write(writer, binary.LittleEndian, &p.body)

	return err
}

// Unpack unpack
func (p *FutuPack) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.LittleEndian, &p.szHeaderFlag)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoID)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoFmtType)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoVer)
	err = binary.Read(reader, binary.LittleEndian, &p.nSerialNo)
	err = binary.Read(reader, binary.LittleEndian, &p.nBodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.arrBodySHA1)
	err = binary.Read(reader, binary.LittleEndian, &p.arrReserved)

	p.body = make([]byte, p.nBodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.body)

	return err
}

// to string
func (p *FutuPack) String() string {
	return fmt.Sprintf("nBodyLen: %d body: %s",
		p.nBodyLen,
		p.body,
	)
}

// GetBody get body data
func (p *FutuPack) GetBody() []byte {
	return p.body
}



