package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
	"main.go/define/hid"
	"math/bits"
	"sync/atomic"
)

type Para struct {
	Mode                 byte
	Cfg                  byte
	ComAddress           byte   //字节窗口通信地址
	BaudRate             uint32 //4 个字节芯片串口通信波特率，高字节在前，默认为 0x00002580，即波特率为 9600bps
	Blank1               uint16
	SepDelay             uint16
	Pid                  uint16
	Vid                  uint16
	KeyboardDelay        uint16
	KeyboardReleaseDelay uint16
	EnterSignAuto        byte
	EnterSign            uint64 //8个字节
	KeyboardFilter       uint64
	UsbStringSign        byte
	FastUploadSign       byte
	Blank2               [12]byte
}

const BaudRate300k = uint32(0x493e0)
const BaudRate115200 = uint32(0x1c200)
const BaudRate9600 = uint32(0x2580)

var KbMode = atomic.Uint32{}
var KbCfg = atomic.Uint32{}

// 0x05ac 0x0256
var SepDelay = atomic.Uint32{}
var BaudRate = atomic.Uint32{}

var Pid = atomic.Uint32{}
var Vid = atomic.Uint32{}

func (self *ClientTx) CmdSetParaCfg() *ClientTx {
	if SepDelay.Load() < uint32(0x01) || SepDelay.Load() > uint32(0x100) {
		SepDelay.Store(uint32(0x03))
	}
	if Pid.Load() == uint32(0x0) {
		Pid.Store(uint32(0x05ac))
	}
	if Vid.Load() == uint32(0x0) {
		Vid.Store(uint32(0x0256))
	}
	switch BaudRate.Load() {
	case BaudRate300k, BaudRate115200, BaudRate9600:
		break
	default:
		BaudRate.Store(BaudRate115200)
		break
	}
	pa := Para{
		Mode:                 byte(KbMode.Load()),
		Cfg:                  byte(KbCfg.Load()),
		ComAddress:           0x00,
		BaudRate:             BaudRate.Load(),
		Blank1:               0x0800,
		SepDelay:             uint16(SepDelay.Load()),
		Pid:                  bits.ReverseBytes16(uint16(Pid.Load())),
		Vid:                  bits.ReverseBytes16(uint16(Vid.Load())),
		KeyboardDelay:        0x0000,
		KeyboardReleaseDelay: 0x0001,
		EnterSignAuto:        0x00,
		EnterSign:            0x0d00000000000000,
		KeyboardFilter:       0x0000000000000000,
		UsbStringSign:        hid.Bit0 + hid.Bit1 + hid.Bit2 + hid.Bit7,
		FastUploadSign:       0x00,
		Blank2:               [12]byte{},
	}
	bb := bytes.Buffer{}
	err := binary.Write(&bb, binary.BigEndian, pa)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	//fmt.Println(bb.Len())
	SApi.Head(cmd.CMD_SET_PARA_CFG).Data(pa).Send()
	return self
}
func CmdSetParaCfgRecv(buf []byte) [50]byte {
	bs := bytes.NewReader(buf)
	dats := [50]byte{}
	binary.Read(bs, binary.BigEndian, &dats)
	fmt.Println(dats)
	return dats
}

const (
	SetModeKeyMouse = 0x00
	SetModeKey      = 0x01
	SetModeMouse    = 0x02
	SetModeHidRaw   = 0x03
)

const (
	SetCfgNorm       = 0x00
	SetCfgASCII      = 0x01
	SetCfgPassthough = 0x02
)
