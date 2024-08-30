package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

type Usbstr struct {
	HidStingType byte
	HidLen       byte
	Data         string
}

func (kb *Kb) CmdGetUsbString() *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_GET_USB_STRING
	kb.data([]byte{0x1}).sum()
	return kb
}
func (rx *ClientRx) CmdGetUsbStringRecv(buf []byte) Usbstr {
	bs := bytes.NewReader(buf)
	crx := ClientRx{}
	binary.Read(bs, binary.BigEndian, &crx)
	us := Usbstr{}
	binary.Read(bs, binary.BigEndian, &us)
	fmt.Println(crx)
	fmt.Println(us)
	return us
}
