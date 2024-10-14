package ws

import (
	"encoding/hex"
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/action"
	"main.go/netSender"
)

func Info(c *Net.WsData) {
	action.Key_set_lcd()
	maskctrl, maskbutton := []string{}, []string{}
	action.Mask.Ctrl.Range(func(key, value interface{}) bool {
		maskctrl = append(maskctrl, hex.EncodeToString([]byte{key.(byte)}))
		return true
	})
	action.Mask.Button.Range(func(key, value interface{}) bool {
		maskbutton = append(maskbutton, hex.EncodeToString([]byte{key.(byte)}))
		return true
	})
	data := map[string]any{
		"Endpoint_delay":        action.Endpoint_delay.Load(),
		"Endpoint_BeforeDelay":  action.Endpoint_BeforeDelay.Load(),
		"Endpoint_dynamic_mode": action.Endpoint_dynamic_mode.Load(),
		"VHits":                 action.VHits.Load(),
		"VLast":                 action.VLast.Load(),

		"LCD1": action.LcdLine1,
		"LCD2": action.LcdLine2,
		"Mode": action.Mode.Load(),

		"MaskCtrl":   maskctrl,
		"MaskButton": maskbutton,

		"sep":  netSender.SepDelay.Load(),
		"baud": netSender.BaudRate.Load(),
		"pid":  hex.EncodeToString([]byte{byte(netSender.Pid.Load() >> 8), byte(netSender.Pid.Load())}),
		"vid":  hex.EncodeToString([]byte{byte(netSender.Vid.Load() >> 8), byte(netSender.Vid.Load())}),
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
