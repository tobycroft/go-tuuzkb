package netSender

import (
	"main.go/define/cmd"
)

// 发送 USB 键盘媒体功能
func (self *ClientTx) CmdSendKbMediaData(keybytes []byte) {
	SApi.Head(cmd.CMD_SEND_KB_MEDIA_DATA).Data(keybytes).Send()
}
