package netReceiver

import (
	"encoding/hex"
	"fmt"
	"net"
	"sync"
)

type netReciever struct {
	MonitorPort    uint32
	connMonitor    *net.UDPConn
	KeyState       keyboardState
	keyboardReport chan StandardKeyboardReport

	IP             string
	Port           int
	UUID           string
	SendHoldMode   bool //开启后模拟单线程模式
	Queue          bool //开启后模拟单线程模式
	SeparateQueue  bool //开启后模拟单线程模式
	DebugClient    bool
	DebugDelay     bool
	KeyChannel     chan KeyAll
	MouseChannel   chan any
	KeyboardData   KeyboardData
	CurrentPressed CurrentPressed
	Fresh          bool //是否是第一次启动
}

type CurrentPressed struct {
	sync.Map
}

func (self *netReciever) Ready() {

}

func TtlRouter(Data []byte) {
	if len(Data) < 1 {
		return
	}
	switch Data[0] {

	case 0x81:
		fmt.Println("链接")
		break

	case 0x82:
		fmt.Println(Data[0], Data)
		break

	case 0x80:
		//fmt.Println("键值改变帧", Data[1:6])
		break

	case 0x86:
		//fmt.Println("设备断开")
		break

	case 0x88:
		fmt.Println("键值数据帧：", Data[1:])
		break

	case 0x01:
		fmt.Println("键盘数据帧：", Data[1:9])

	case 0x02:
		fmt.Println("鼠标数据帧2：", Data[1:5])

	case 0x04:
		fmt.Println("鼠标数据帧4：", Data[1:8])

	default:
		fmt.Println(hex.EncodeToString(Data[:1]), hex.EncodeToString(Data[1:]))

	}
}
