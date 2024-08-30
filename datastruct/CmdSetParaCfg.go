package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

type Para struct {
	Mode                 byte
	Cfg                  byte
	ComAddress           byte //字节窗口通信地址
	BaudRate             byte //4 个字节芯片串口通信波特率，高字节在前，默认为 0x00002580，即波特率为 9600bps
	Blank1               uint16
	SepDelay             uint16
	PidVid               uint32
	KeyboardDelay        byte
	KeyboardReleaseDelay byte
	EnterSignAuto        byte
	EnterSign            uint64 //8个字节
	KeyboardFilter       uint64
	FastUploadSign       byte
}

func (kb *Kb) CmdSetParaCfg() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SET_PARA_CFG
	kb.data([]byte{}).sum()
	return kb
}
func (rx *ClientRx) CmdSetParaCfgRecv(buf []byte) [50]byte {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
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
)

const (
	SetCfgNorm       = 0x00
	SetCfgASCII      = 0x01
	SetCfgPassthough = 0x02
)
