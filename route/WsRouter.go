package route

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/websocket"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		fmt.Println(c.Conn.RemoteAddr(), string(c.Message), c.Status)
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

		default:
			Net.WsServer_WriteChannel <- c
			break
		}
	}
}
