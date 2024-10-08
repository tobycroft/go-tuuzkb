package netSender

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"main.go/define/cmd"
)

func (self *ClientTx) CmdGetParaCfg() *ClientTx {
	self.head(cmd.CMD_GET_PARA_CFG).data([]byte{}).send()
	return self
}
func CmdGetParaCfgRecv(buf []byte) Para {
	bs := bytes.NewReader(buf)
	//crx := sendData{}
	//binary.Read(bs, binary.BigEndian, &crx)
	pa := Para{}
	binary.Read(bs, binary.BigEndian, &pa)
	switch pa.Mode {
	case 0x00:
		fmt.Println("工作模式：键盘鼠标")
		break

	case 0x01:
		fmt.Println("工作模式：键盘")
		break

	case 0x02:
		fmt.Println("工作模式：鼠标")
		break

	case 0x03:
		fmt.Println("工作模式：HID Raw")
		break
	}

	switch pa.Cfg {
	case 0x00:
		fmt.Println("配置：协议传输")
		break

	case 0x01:
		fmt.Println("配置：ASCII")
		break

	case 0x02:
		fmt.Println("配置：Passthough")
		break
	}

	fmt.Println("通信地址:", pa.ComAddress)
	fmt.Println("波特率:", pa.BaudRate)
	fmt.Println("通信包间隔:", pa.SepDelay)
	fmt.Println("PID:", hex.EncodeToString([]byte{byte(pa.Pid), byte(pa.Pid >> 8)}), "VID:", hex.EncodeToString([]byte{byte(pa.Vid), byte(pa.Vid >> 8)}))
	fmt.Println("USB字符串:", pa.UsbStringSign)
	return pa
}

const (
	GetModeKeyMouse = 0x80
	GetModeKey      = 0x81
	GetModeMouse    = 0x82
	GetModeHidRaw   = 0x83
)

const (
	GetCfgNorm       = 0x80
	GetCfgASCII      = 0x81
	GetCfgPassthough = 0x82
)
