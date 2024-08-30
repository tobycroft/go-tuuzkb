package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

func (kb *Kb) CmdSetUsbString() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SET_USB_STRING
	kb.Ctx.Len = 2
	kb.data([]byte("2.4G Wireless Receiver")).sum()
	return kb
}
func (rx *ClientRx) CmdSetUsbStringRecv(buf []byte) string {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
	binary.Read(bs, binary.BigEndian, &crx)
	pa := [24]byte{}
	binary.Read(bs, binary.BigEndian, &pa)
	fmt.Println(crx)
	fmt.Println(string(pa[2:]))
	return string(pa[2:])
}
