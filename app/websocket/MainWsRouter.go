package websocket

import (
	"fmt"
	Net "github.com/tobycroft/TuuzNet"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		fmt.Println(c)
	}
}
