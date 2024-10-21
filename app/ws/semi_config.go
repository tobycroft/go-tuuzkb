package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
	"main.go/netSender"
)

func SemiConfig(c *Net.WsData) {
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
	d, err := sonic.Get(c.Message, "data")
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := d.Map()
	if err != nil {
		fmt.Println(err)
		return
	}
	switch Type {
	case "Mode":
		action.Mode.Store(Calc.Any2Int64(data["Mode"]))
		action.SwitchMode()
		break

	case "Endpoint_dynamic_mode":
		action.Endpoint_dynamic_mode.Store(Calc.Any2Int64(data["Endpoint_dynamic_mode"]))
		action.SwitchDynamicMode()
		break

	case "Endpoint_delay":
		action.Endpoint_delay.Store(Calc.Any2Int64(data["Endpoint_delay"]))
		break

	case "Endpoint_BeforeDelay":
		action.Endpoint_BeforeDelay.Store(Calc.Any2Int64(data["Endpoint_BeforeDelay"]))
		break

	case "Endpoint_BeforeDelay_Random":
		action.Endpoint_BeforeDelay_Random.Store(Calc.Any2Int64(data["Endpoint_BeforeDelay_Random"]))
		break

	case "sep":
		if Calc.Any2Int64(data["sep"]) < 1 {
			data["sep"] = 1
		}
		netSender.SepDelay.Store(uint32(Calc.Any2Int64(data["sep"])))
		break

	case "kbmode":
		netSender.KbMode.Store(uint32(Calc.Any2Int64(data["kbmode"])))
		netSender.Ctx.CmdSetParaCfg()
		break

	case "kbcfg":
		netSender.KbCfg.Store(uint32(Calc.Any2Int64(data["kbcfg"])))
		netSender.Ctx.CmdSetParaCfg()
		break

	default:
		fmt.Println(c.Conn.RemoteAddr().String(), Type)
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
