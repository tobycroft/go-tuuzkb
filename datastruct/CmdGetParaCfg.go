package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

func (kb *Kb) CmdGetParaCfg() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_PARA_CFG
	kb.data([]byte{}).sum()
	return kb
}
func (rx *ClientRx) CmdGetParaCfgRecv(buf []byte) [50]byte {
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
