package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
	"main.go/netSender"
	"time"
)

func Kbd(c *Net.WsData) {
	a, err := sonic.Get(c.Message, "type")
	if err != nil {
		fmt.Println(err)
		return
	}
	Type, err := a.String()
	if err != nil {
		fmt.Println(err)
		return
	}
	switch Type {
	case "reset":
		action.SetUsbString()
		time.Sleep(1 * time.Second)
		go netSender.Ctx.CmdReset()
		break

	case "cfg3k":
		go netSender.Ctx.CmdSetParaCfg(netSender.BaudRate300k, 0x05ac, 0x0256)
		break

	case "cfg115k":
		go netSender.Ctx.CmdSetParaCfg(netSender.BaudRate115200, 0x05ac, 0x0256)
		break

	case "cfg9k":
		go netSender.Ctx.CmdSetParaCfg(netSender.BaudRate9600, 0x05ac, 0x0256)
		break

	case "cfgget":
		go netSender.Ctx.CmdGetParaCfg()
		break

	case "setusb":
		go action.SetUsbString()

	default:
		break
	}
}
