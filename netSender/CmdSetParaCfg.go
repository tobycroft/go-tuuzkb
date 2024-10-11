package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
	"math/bits"
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
	Blank2               uint64
	Blank3               uint32
}

const BaudRate300k = 0x493e0
const BaudRate115200 = 0x1c200
const BaudRate9600 = 0x2580

//0x05ac 0x0256

func (self *ClientTx) CmdSetParaCfg(BaudRate uint32, Pid, Vid uint16) *ClientTx {
	pa := Para{
		Mode:                 SetModeKeyMouse,
		Cfg:                  SetCfgNorm,
		ComAddress:           0x00,
		BaudRate:             BaudRate,
		Blank1:               0x0000,
		SepDelay:             0x1,
		Pid:                  bits.ReverseBytes16(Pid),
		Vid:                  bits.ReverseBytes16(Vid),
		KeyboardDelay:        0x00,
		KeyboardReleaseDelay: 0x01,
		EnterSignAuto:        0x00,
		EnterSign:            936748722493063168,
		KeyboardFilter:       0x0000000000000000,
		UsbStringSign:        0x01,
		FastUploadSign:       0x00,
		Blank2:               0x00000000,
		Blank3:               0x00000000,
	}
	bb := bytes.Buffer{}
	err := binary.Write(&bb, binary.BigEndian, pa)
	if err != nil {
		panic(fmt.Sprintln("binary编译失败", err))
	}
	//fmt.Println(bb.Len())
	self.head(cmd.CMD_SET_PARA_CFG).data(pa).send()
	return self
}
func (self *ClientTx) CmdSetParaCfgRecv(buf []byte) [50]byte {
	bs := bytes.NewReader(buf)
	crx := sendData{}
	binary.Read(bs, binary.BigEndian, &crx)
	dats := [50]byte{}
	binary.Read(bs, binary.BigEndian, &dats)
	fmt.Println(crx)
	fmt.Println(dats)
	return dats
}

const (
	SetModeKeyMouse = 0x00
	SetModeKey      = 0x01
	SetModeMouse    = 0x02
	SetModeHidRaw   = 0x03
	ss              = 1
)

const (
	SetCfgNorm       = 0x00
	SetCfgASCII      = 0x01
	SetCfgPassthough = 0x02
)
