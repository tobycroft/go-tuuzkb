package netReceiver

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"main.go/define/hid"
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
		self.Router9239(Data[1:], Addr, PackConn)
		break

	case 0x99:
		//fmt.Println("ping")
		break

	case 0xab:
		self.Router9239(Data[2:], Addr, PackConn)
		break

	case 0x57:
		self.Router9239(Data[2:], Addr, PackConn)
		break

	case 0x81:
		go fmt.Println("链接")
		break

	case 0x82:
		//go fmt.Println("状态请求:", Data[0], Data)
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
		err := binary.Read(buf, binary.BigEndian, &kbreport)
		if err != nil {
			panic(err.Error())
		}
		//go fmt.Println(kbreport)
		self.keyboardMain <- kbreport
		//go fmt.Println("键盘数据帧：", Data[1:9])
		break

	case 0x02:
		go fmt.Println("鼠标数据帧2：", Data[1:5])
		break

	case 0x04:
		go fmt.Println("鼠标数据帧4：", Data[1:8])
		break

	default:
		go fmt.Println("main_unreco:", Addr, Data[0], hex.EncodeToString(Data[:2]), hex.EncodeToString(Data))

	}
}

func (self *ClientRx) Router9239(Data []byte, Addr net.Addr, PackConn net.PacketConn) {
	switch Data[0] {

	case 0x00:
		break

	case 0x99:
		//fmt.Println("ping")
		break

	case 0x81:
		//fmt.Println("9239:PowerUp:", hex.EncodeToString(Data[2:]))
		fmt.Print("9239:Version:1.", Data[2]-0x30)
		if Data[3] == 0x00 {
			fmt.Print(":控制器识别失败")
		} else {
			fmt.Print(":控制器识别成功")
		}
		fmt.Println(":Lockers:", "Numlock:", Data[4]&hid.Bit0, "Capslock:", Data[4]&hid.Bit1, "Scrolllock:", Data[4]&hid.Bit2)
		break

	case 0x82:
		//fmt.Println("CMD_SEND_KB_GENERAL_DATA键盘执行结果:", hex.EncodeToString(Data[1:]))
		break

	case 0x88:
		//fmt.Println("键盘数据帧：", hex.EncodeToString(Data[2:]))
		go netSender.CmdGetParaCfgRecv(Data[2:])
		break

	case 0x85:
		//fmt.Println("KBMS-操作成功")
		break

	case 0x8a:
		switch Data[2] {
		//，0x00 表示厂商字符串描述符；0x01 表示产品字符串描述符；
		//0x02 表示序列号字符串描述符
		case 0x00:
			fmt.Println("键盘产商字符串描述符：", netSender.CmdGetUsbStringRecv(Data))
			break

		case 0x01:
			fmt.Println("键盘产品字符串描述符：", netSender.CmdGetUsbStringRecv(Data))
			break

		case 0x02:
			fmt.Println("键盘序列号字符串描述符：", netSender.CmdGetUsbStringRecv(Data))
			break
		}
		//fmt.Println(hex.EncodeToString(Data), netSender.CmdGetUsbStringRecv(Data))
		break

	case 0x8b:
		if Data[1] == 0x01 {
			go fmt.Println("键盘字符串设定成功")
		}
		break

	case 0xca:
		go fmt.Println("错误：", hex.EncodeToString(Data[2:2]))
		break

	case 0x8f:
		fmt.Println("设备重启完成")
		break

	default:
		go fmt.Println("9239_unreco:", hex.EncodeToString(Data))

	}
}
