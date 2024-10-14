package route

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/ws"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		nd, err := sonic.Get(c.Message, "route")
		if err != nil {
			fmt.Println("err", string(c.Message))
			continue
		}
		r, err := nd.String()
		if err != nil {
			continue
		}
		switch r {
		case "login":
			go ws.Login(&c)
			break

		case "info":
			go ws.Info(&c)
			break

		case "semi-config":
			go ws.SemiConfig(&c)
			break

		case "kbd":
			go ws.Kbd(&c)
			break

		case "ping":
			Net.WsServer_WriteChannel <- Net.WsData{
				Conn:    c.Conn,
				Type:    websocket.TextMessage,
				Message: []byte("pong"),
			}
			break

		default:
			fmt.Println(c.Conn.RemoteAddr().String())
			break
		}
	}
}
