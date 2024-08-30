package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

func (kb *Kb) CmdGetUsbString() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
	kb.data([]byte{0x0}).sum()
	return kb
}
func (rx *ClientRx) CmdGetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
	binary.Read(bs, binary.BigEndian, &crx)
	pa := [24]byte{}
	binary.Read(bs, binary.BigEndian, &pa)
	fmt.Println(crx)
	fmt.Println(string(pa[2:]))
	return string(pa[2:])
}
