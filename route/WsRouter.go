package route

import (
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/websocket"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		nd, err := sonic.Get(c.Message, "route")
		if err != nil {
			continue
		}
		r, err := nd.String()
		if err != nil {
			continue
		}
		switch r {
		case "login":
			websocket.Login(&c)
			break

		case "info":
			websocket.Info(&c)
			break

		default:
			Net.WsServer_WriteChannel <- c
			break
		}
	}
}
