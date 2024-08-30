package datastruct

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"main.go/define/cmd"
)

// 设置字符串描述符配置
func (kb *Kb) CmdSetUsbString(HidStingType byte, str string) *Kb {
	kb.head()
	kb.Ctx.Cmd = cmd.CMD_SET_USB_STRING
	//kb.Ctx.Len = 2
	usbstr := Usbstr{
		HidStingType: HidStingType,
		HidLen:       0,
	}
	usbstr.HidLen = byte(len(str))
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, usbstr)
	buf.WriteString(str)
	fmt.Println(string(buf.Bytes()))
	kb.data(buf.Bytes()).send()
	return kb
}
