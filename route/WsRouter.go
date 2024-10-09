package route

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/ws"
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
			ws.Login(&c)
			break

		case "info":
			ws.Info(&c)
			break

		default:
			fmt.Println(c.Conn.RemoteAddr().String())
			break
		}
	}
}
