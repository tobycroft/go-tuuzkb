package ws

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
)

func SemiConfig(c *Net.WsData) {
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
	case "Mode":
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
