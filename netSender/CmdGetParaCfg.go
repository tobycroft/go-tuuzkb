package netSender

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

func (self *ClientTx) CmdGetParaCfg() *ClientTx {
	self.head()
	self.sendData.Cmd = cmd.CMD_GET_PARA_CFG
	self.data([]byte{}).send()
	return self
}
func (self *ClientTx) CmdGetParaCfgRecv(buf []byte) Para {
	bs := bytes.NewReader(buf)
	crx := sendData{}
	binary.Read(bs, binary.BigEndian, &crx)
	pa := Para{}
	binary.Read(bs, binary.BigEndian, &pa)
	fmt.Println(crx)
	fmt.Println(pa)
	return pa
}

const (
	GetModeKeyMouse = 0x80
	GetModeKey      = 0x81
	GetModeMouse    = 0x82
	GetModeHidRaw   = 0x83
)

const (
	GetCfgNorm       = 0x80
	GetCfgASCII      = 0x81
	GetCfgPassthough = 0x82
)
