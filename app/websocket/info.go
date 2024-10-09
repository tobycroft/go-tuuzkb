package websocket

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
)

func Info(c *Net.WsData) {
	data := map[string]any{
		"Endpoint_delay":        action.Endpoint_delay.Load(),
		"Endpoint_BeforeDelay":  action.Endpoint_BeforeDelay.Load(),
		"Endpoint_dynamic_mode": action.Endpoint_dynamic_mode.Load(),
		"vhits":                 action.VHits.Load(),
		"vlast":                 action.VLast.Load(),

		"LCD1": action.LcdLine1,
		"LCD2": action.LcdLine2,
		"LCD3": action.LcdLine3,
	}
	bt, err := sonic.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn:    c.Conn,
		Type:    0,
		Message: bt,
		Status:  true,
	}
}
