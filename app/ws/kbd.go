package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
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

	case "cfg3k":
		netSender.Ctx.CmdSetParaCfg(netSender.BaudRate300k, 0x05ac, 0x0256)
		break

	case "cfg115k":
		netSender.Ctx.CmdSetParaCfg(netSender.BaudRate115200, 0x05ac, 0x0256)
		break

	case "cfg9k":
		netSender.Ctx.CmdSetParaCfg(netSender.BaudRate9600, 0x05ac, 0x0256)
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
		fmt.Println("Reset")
		break

	default:
		break
	}
	Net.WsConns.Range(func(key, value interface{}) bool {
		Net.WsServer_WriteChannel <- Net.WsData{
			Conn:    value.(*websocket.Conn),
			Type:    websocket.TextMessage,
			Message: []byte("update"),
		}
		return true
	})
}
