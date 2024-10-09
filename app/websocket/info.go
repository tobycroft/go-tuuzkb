package websocket

import Net "github.com/tobycroft/TuuzNet"

func Info(c *Net.WsData) {
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn:    c.Conn,
		Type:    0,
		Message: []byte(""),
		Status:  true,
	}
}
