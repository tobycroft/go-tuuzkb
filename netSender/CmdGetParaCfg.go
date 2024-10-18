package netSender

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/gorilla/websocket"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/common"
	"main.go/define/cmd"
	"math/bits"
)

func (self *ClientTx) CmdGetParaCfg() *ClientTx {
	new(SendFrame).Head(cmd.CMD_GET_PARA_CFG).Data([]byte{}).Send()
	return self
}

func CmdGetParaCfgRecv(buf []byte) Para {
	bs := bytes.NewReader(buf)
	//crx := sendData{}
	//binary.Read(bs, binary.BigEndian, &crx)
	pa := Para{}
	binary.Read(bs, binary.BigEndian, &pa)
	go common.PrintRedis("获取参数配置", pa)
	//fmt.Println(pa)
	switch pa.Mode {
	case 0x00, 0x80:
		fmt.Println("工作模式：键盘鼠标")
		KbMode.Store(0x00)
		break

	case 0x01, 0x81:
		fmt.Println("工作模式：键盘")
		KbMode.Store(0x01)
		break

	case 0x02, 0x82:
		fmt.Println("工作模式：鼠标")
		KbMode.Store(0x02)
		break

	case 0x03, 0x83:
		fmt.Println("工作模式：HID Raw")
		KbMode.Store(0x03)
		break
	}

	switch pa.Cfg {
	case 0x00, 0x80:
		fmt.Println("配置：协议传输")
		KbCfg.Store(0x00)
		break

	case 0x01, 0x81:
		fmt.Println("配置：ASCII")
		KbCfg.Store(0x01)
		break

	case 0x02, 0x82:
		fmt.Println("配置：Passthough")
		KbCfg.Store(0x02)
		break
	}

	fmt.Println("通信地址:", pa.ComAddress)
	fmt.Println("波特率:", pa.BaudRate)
	BaudRate.Store(pa.BaudRate)
	fmt.Println("通信包间隔:", pa.SepDelay)
	SepDelay.Store(uint32(pa.SepDelay))
	fmt.Println("PID:", hex.EncodeToString([]byte{byte(pa.Pid), byte(pa.Pid >> 8)}), "VID:", hex.EncodeToString([]byte{byte(pa.Vid), byte(pa.Vid >> 8)}))
	Pid.Store(uint32(bits.ReverseBytes16(pa.Pid)))
	Vid.Store(uint32(bits.ReverseBytes16(pa.Vid)))

	fmt.Println("USB字符串:", pa.UsbStringSign)

	go func() {
		Net.WsConns.Range(func(key, value interface{}) bool {
			Net.WsServer_WriteChannel <- Net.WsData{
				Conn:    value.(*websocket.Conn),
				Type:    websocket.TextMessage,
				Message: []byte("update"),
			}
			return true
		})
	}()

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
