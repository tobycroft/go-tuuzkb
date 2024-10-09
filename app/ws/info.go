package ws

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
		"VHits":                 action.VHits.Load(),
		"VLast":                 action.VLast.Load(),

		"LCD1": action.LcdLine1,
		"LCD2": action.LcdLine2,
		"Mode": action.Mode.Load(),
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
