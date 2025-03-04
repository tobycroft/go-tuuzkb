package netReceiver

import (
	"encoding/hex"
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
	"sync"
	"sync/atomic"
	"time"
)

var Crx = &ClientRx{}

type ClientRx struct {
	keyboardMain chan *netSender.KeyboardData2
	mouseMain    chan *netSender.MouseData

	KeyboardRxChannel chan *netSender.KeyboardData2
	MouseRxChannel    chan *netSender.MouseData
}

var originCtrl = &atomic.Value{}
var originButton = &[6]*atomic.Value{}

var OriginalButton = &sync.Map{}
var OriginCtrl = &sync.Map{}

func (self *ClientRx) Ready() {
	self.keyboardMain = make(chan *netSender.KeyboardData2, 1)
	self.mouseMain = make(chan *netSender.MouseData, 1)

	self.MouseRxChannel = make(chan *netSender.MouseData, 1)
	self.KeyboardRxChannel = make(chan *netSender.KeyboardData2, 1)

	originCtrl.Store(byte(hid.CmdNone))
	for i := range originButton {
		originButton[i] = &atomic.Value{}
		originButton[i].Store(byte(hid.CmdNone))
	}

	go self.RouterKeyboard()
	go self.RouterMouse()
}

type keyframe struct {
	DataLength byte   // 长度
	Ident      byte   // 标识
	KeyData    []byte // 键值
	Index      byte   // 序列号
	Checksum   byte   // 校验
}

func (self *ClientRx) MessageRouter(Data []byte) {
	if len(Data) < 1 {
		return
	}
	switch Data[0] {

	case 0x00:
		if len(Data) < 2 {
			fmt.Println("Router9239-data-err：", Data)
			break
		}
		self.Router9239(Data[1:])
		break

	case 0x99:
		//fmt.Println("ping")
		break

	case 0x81:
		fmt.Println("链接")
		break

	case 0x82:
		//go fmt.Println("状态请求:", Data[0], Data)
		break

	case 0x80:
		//go fmt.Println("键值改变帧", Data[1:])
		break

	case 0x86:
		fmt.Println("设备断开")
		break

	case 0x88:
		////fmt.Println("kff", len(Data), Data[1], Data[2], 2+int(Data[1]))
		//frame := keyframe{
		//	DataLength: Data[1],
		//	Ident:      Data[2],
		//	KeyData:    Data[3 : int(Data[1])-2], // 根据 DataLength 解析数据
		//	Index:      Data[int(Data[1])-1],
		//	Checksum:   Data[int(Data[1])],
		//}
		////switch frame.Ident & hid.Bit0 {
		////case 0x00:
		////	fmt.Println("上键值数据结构：")
		////	break
		////case 0x01:
		////	fmt.Println("下键值数据结构：")
		////	break
		////}
		//if frame.Ident&hid.Bit0 == 0 {
		//
		//}
		//fmt.Println("fma1:", frame.Ident&hid.Bit5, frame.Ident&hid.Bit4, frame.Ident&hid.Bit5&hid.Bit4, "fma2:", frame.Ident&hid.Bit2, frame.Ident&hid.Bit1, "port:", frame.Ident&hid.Bit0)
		////kbreport := netSender.KeyboardData2{}
		////buf := bytes.NewReader(Data[1:])
		////err := binary.Read(buf, binary.BigEndian, &kbreport)
		////if err != nil {
		////	fmt.Println(len(Data), Data)
		////	fmt.Println(hex.EncodeToString(Data))
		////	panic(err.Error())
		////}
		//fmt.Println("键值数据帧：", Data[1:])
		//fmt.Println("键值数据结构：", frame)
		////self.keyboardMain <- kbreport
		break

	case 0x01:
		//fmt.Println("kb-recv", hex.EncodeToString(Data))
		if len(Data) < 9 {
			fmt.Println("键盘数据帧：", Data)
			break
		}
		kbreport := &netSender.KeyboardData2{
			Ctrl:   Data[1],
			Resv:   Data[2],
			Button: [6]byte{Data[3], Data[4], Data[5], Data[6], Data[7], Data[8]},
		}
		//buf := bytes.NewReader(Data[1:])
		//err := binary.Read(buf, binary.BigEndian, kbreport)
		//if err != nil {
		//	fmt.Println(len(Data), Data)
		//	fmt.Println(hex.EncodeToString(Data))
		//	panic(err.Error())
		//}
		//go fmt.Println(kbreport)
		self.keyboardMain <- kbreport
		//go fmt.Println("键盘数据帧：", Data[1:9])
		break

	case 0x02:
		//go fmt.Println("鼠标数据帧2：", Data[1:5])
		if len(Data) < 5 {
			fmt.Println("键盘数据帧：", Data)
			break
		}
		mouseReport := &netSender.MouseData{
			Resv:       0x01,
			ButtonBits: Data[1],
			X:          Data[2],
			Y:          Data[3],
			Wheel:      Data[4],
		}
		self.mouseMain <- mouseReport
		//netSender.Ctx.CmdSendMsRelData(*mouseReport)
		break

	case 0x04:
		if len(Data) < 9 {
			fmt.Println("鼠标数据帧错误-data-err：", Data)
			break
		}
		fmt.Println("鼠标数据帧4：", Data[1:8])
		break

	default:
		fmt.Println("main_unreco:", hex.EncodeToString(Data))

	}
}

func (self *ClientRx) Router9239(Data []byte) {
	if len(Data) < 1 {
		return
	}
	switch Data[0] {

	case 0x99:
		//fmt.Println("ping")
		break

	case 0x81:
		//fmt.Println("9239:PowerUp:", hex.EncodeToString(Data[2:]))
		if len(Data) < 5 {
			fmt.Println("控制器识别错误-data-err：", Data)
			break
		}
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
		//fmt.Println("键盘数据帧：", hex.EncodeToString(Data[0:]))
		if len(Data) < 4 {
			fmt.Println("CmdGetParaCfgRecv-data-err：", Data)
			break
		}
		netSender.CmdGetParaCfgRecv(Data[2:])
		break

	case 0x85:
		//fmt.Println("KBMS-操作成功")
		break

	case 0x8a:
		if len(Data) < 3 {
			fmt.Println("键盘产商字符串描述符-data-err：", Data)
			break
		}
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
		if len(Data) < 2 {
			fmt.Println("键盘字符串设定-data-err：", Data)
			break
		}
		if Data[1] == 0x01 {
			fmt.Println("键盘字符串设定成功")
		}
		break

	case 0xca, 0xc0, 0xc1, 0xc2, 0xc3, 0xc4:
		fmt.Println("9239错误：", hex.EncodeToString(Data))
		break

	case 0x8f:
		fmt.Println("设备重启完成")
		go func() {
			time.Sleep(3 * time.Second)
			SetUsbString()
			netSender.Ctx.CmdGetParaCfg()
		}()

		break

	case 0x89:
		fmt.Println("设备CFG设定成功")
		break

	default:
		fmt.Println("9239_unreco:", hex.EncodeToString(Data))

	}
}

func SetUsbString() {
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeManufacturer, "2.4G MonkaKeyboard")
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeProduct, "2.4G MonkaReciever")
	netSender.Ctx.CmdSetUsbString(netSender.StrTypeSerial, "001202208")
}
