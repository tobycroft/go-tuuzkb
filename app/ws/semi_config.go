package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
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

	default:
		fmt.Println(c.Conn.RemoteAddr().String(), Type)
		break
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn:    c.Conn,
		Type:    websocket.TextMessage,
		Message: []byte("update"),
	}
}
