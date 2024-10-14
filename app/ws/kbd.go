package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
	"main.go/netReceiver"
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
		netReceiver.SetUsbString()
		time.Sleep(500 * time.Millisecond)
		netSender.Ctx.CmdReset()
		break

	case "05ac":
		netSender.Pid.Store(uint32(0x05ac))
		netSender.Vid.Store(uint32(0x0256))
		netSender.BaudRate.Store(netSender.BaudRate115200)
		netSender.Ctx.CmdSetParaCfg()
		break

	case "cfg3k":
		netSender.BaudRate.Store(netSender.BaudRate300k)
		netSender.Ctx.CmdSetParaCfg()
		break

	case "cfg115k":
		netSender.BaudRate.Store(netSender.BaudRate115200)
		netSender.Ctx.CmdSetParaCfg()
		break

	case "cfg9k":
		netSender.BaudRate.Store(netSender.BaudRate9600)
		netSender.Ctx.CmdSetParaCfg()
		break

	case "cfgget":
		netSender.Ctx.CmdGetParaCfg()
		break

	case "setusb":
		go netReceiver.SetUsbString()
		break

	case "setting_reset":
		action.Endpoint_delay.Store(0)
		action.Endpoint_BeforeDelay.Store(0)
		action.Mode.Store(0)
		action.Endpoint_dynamic_mode.Store(0)
		fmt.Println("设置重置")
		go action.Lcd_refresh()
		break

	case "bankey":
		action.Kb_banSomeKeys()
		break

	case "unbanall":
		action.Mask.Button.Clear()
		action.Mask.Ctrl.Clear()
		go action.Lcd_refresh()
		break

	default:
		break
	}

}
