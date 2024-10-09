package ws

import (
	"fmt"
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
		fmt.Println(err)
		return
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn:    c.Conn,
		Type:    0,
		Message: []byte(""),
		Status:  true,
	}
	tjson, err := a.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	ld := &loginData{}
	err = sonic.Unmarshal(tjson, &ld)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ld)
	//jsonObj := sonic.Unmarshal(c.Message, &loginData{})
}
