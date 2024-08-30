package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

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
