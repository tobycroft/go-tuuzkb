package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/netSender"
	"time"
)

func Kbd(c *Net.WsData) {
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
	switch Type {
	case "reset":
		time.Sleep(1 * time.Second)
		netSender.Ctx.CmdReset()
		break

	default:
		break
	}
}
