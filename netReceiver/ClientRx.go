package netReceiver

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"main.go/netSender"
	"net"
	"sync"
)

var Crx = &ClientRx{}

type ClientRx struct {
	keyboardMain chan netSender.KeyboardData2
	mouseMain    chan any

	KeyboardRxChannel chan netSender.KeyboardData2
	MouseRxChannel    chan any

	keys           netSender.KeyboardData2
	OriginalButton sync.Map
	OriginCtrl     sync.Map
}

func (self *ClientRx) Ready() {
	self.keyboardMain = make(chan netSender.KeyboardData2)
	self.mouseMain = make(chan any)

	self.MouseRxChannel = make(chan any)
	self.KeyboardRxChannel = make(chan netSender.KeyboardData2)

	go self.RouterKeyboard()
}

func (self *ClientRx) MessageRouter(Data []byte, Addr net.Addr, PackConn net.PacketConn) {
	if len(Data) < 1 {
		return
	}
	switch Data[0] {

	case 0x00:
		go self.Router9239(Data[1:], Addr, PackConn)
		break

	case 0x81:
		go fmt.Println("链接")
		break

	case 0x82:
		go fmt.Println(Data[0], Data)
		break

	case 0x80:
		//go fmt.Println("键值改变帧", Data[1:6])
		break

	case 0x86:
		//go fmt.Println("设备断开")
		break

	case 0x88:
		go fmt.Println("键值数据帧：", Data[1:])
		break

	case 0x01:
		kbreport := netSender.KeyboardData2{}
		buf := bytes.NewReader(Data[1:9])
		err := binary.Read(buf, binary.NativeEndian, &kbreport)
		if err != nil {
			panic(err.Error())
		}
		//go fmt.Println(kbreport)
		self.keyboardMain <- kbreport
		//go fmt.Println("键盘数据帧：", Data[1:9])

	case 0x02:
		go fmt.Println("鼠标数据帧2：", Data[1:5])

	case 0x04:
		go fmt.Println("鼠标数据帧4：", Data[1:8])

	default:
		go fmt.Println("unreco:", Addr, Data[0], hex.EncodeToString(Data[:1]), hex.EncodeToString(Data[1:]))

	}
}

func (self *ClientRx) Router9239(Data []byte, Addr net.Addr, PackConn net.PacketConn) {
	switch Data[0] {

	case 0x82:
		//fmt.Println("CMD_SEND_KB_GENERAL_DATA键盘执行结果:", hex.EncodeToString(Data[1:]))
		break

	default:
		go fmt.Println("rcv_unreco:", hex.EncodeToString(Data[:0]), hex.EncodeToString(Data[1:]))

	}
}
