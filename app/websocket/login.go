package websocket

import (
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
)

type loginData struct {
	Username string
	Password string
}

func Login(c *Net.WsData) {
	a, err := sonic.Get(c.Message, "data")
	if err != nil {
		return
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn:    c.Conn,
		Type:    0,
		Message: []byte(""),
		Status:  true,
	}
	tjson, err := a.MarshalJSON()
	sonic.Unmarshal(tjson, &loginData{})
	//jsonObj := sonic.Unmarshal(c.Message, &loginData{})
}
