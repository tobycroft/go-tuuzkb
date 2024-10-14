package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netReceiver"
	"main.go/netSender"
	"sync/atomic"
	"time"
)

type lastKey struct {
	Ctrl    atomic.Value
	Button0 atomic.Value
	Button1 atomic.Value
	Button2 atomic.Value
	Button3 atomic.Value
	Button4 atomic.Value
	Button5 atomic.Value
}

var LastPress = lastKey{}
var CurrentPress = lastKey{}

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		swap_key(&c)
		self.c = c
		fmt.Println("keybaordrecv", c)
		go self.kb_actvate()
		go self.kb_reboot()
		go self.kb_unbanall()
		//go self.kb_test()
		go self.key_main()
		go self.qe_main()
		go self.whel_main()
		//
		go self.kb_get_para()
		go self.kb_set_para()
		go self.kb_get_usbstring()
		go self.kb_set_usbstring()
		if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen) {
			go Kb_banSomeKeys()
			fmt.Println("ban_all")
		}
		self.SendKbGeneralDataRaw(c)

	}
	panic("键盘通道意外结束")
}

func swap_key(c *netSender.KeyboardData2) {
	if ctrl := CurrentPress.Ctrl.Swap(c.Ctrl); ctrl != nil {
		LastPress.Ctrl.Store(ctrl.(byte))
	} else {
		LastPress.Ctrl.Store(byte(0))
	}

	if btn0 := CurrentPress.Button0.Swap(c.Button[0]); btn0 != nil {
		LastPress.Button0.Store(btn0.(byte))
	} else {
		LastPress.Button0.Store(byte(0))
	}

	if btn1 := CurrentPress.Button1.Swap(c.Button[1]); btn1 != nil {
		LastPress.Button1.Store(btn1.(byte))
	} else {
		LastPress.Button1.Store(byte(0))
	}

	if btn2 := CurrentPress.Button2.Swap(c.Button[2]); btn2 != nil {
		LastPress.Button2.Store(btn2.(byte))
	} else {
		LastPress.Button2.Store(byte(0))
	}

	if btn3 := CurrentPress.Button3.Swap(c.Button[3]); btn3 != nil {
		LastPress.Button3.Store(btn3.(byte))
	} else {
		LastPress.Button3.Store(byte(0))
	}

	if btn4 := CurrentPress.Button4.Swap(c.Button[4]); btn4 != nil {
		LastPress.Button4.Store(btn4.(byte))
	} else {
		LastPress.Button4.Store(byte(0))
	}

	if btn5 := CurrentPress.Button5.Swap(c.Button[5]); btn5 != nil {
		LastPress.Button5.Store(btn5.(byte))
	} else {
		LastPress.Button5.Store(byte(0))
	}

}

func (self *Action) kb_actvate() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(0)
		fmt.Println("Reset")
		//go self.Km.KmNetLcdPicture_tempSet("Golang", "GolangGolang", "GolangGolangGolang", 1*time.Second)
	}
}

func Kb_banSomeKeys() {
	kb_add_masking(hid.CmdApplication, false)
	kb_add_masking(hid.CmdPrintScreen, false)
	kb_add_masking(hid.CmdPause, false)
	kb_add_masking(hid.CmdScrollLock, false)
	kb_add_masking(hid.RightCtrl, true)
	//kb_add_masking(hid.RightShift, true)
	kb_add_masking(hid.RightAlt, true)
}

func (self *Action) kb_unbanall() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		Mask.Button.Clear()
		Mask.Ctrl.Clear()
		fmt.Println("unbanall")
	}
}

func (self *Action) kb_get_para() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdScrollLock) {
		self.ClientTx.CmdGetParaCfg()
	}
}

func (self *Action) kb_set_para() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdScrollLock) {
		self.ClientTx.CmdGetParaCfg()
	}
}

func (self *Action) kb_get_usbstring() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdPrintScreen) {
		self.ClientTx.CmdGetUsbString(netSender.StrTypeManufacturer)
		self.ClientTx.CmdGetUsbString(netSender.StrTypeProduct)
		self.ClientTx.CmdGetUsbString(netSender.StrTypeSerial)
	}
}

func (self *Action) kb_set_usbstring() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift+hid.RightAlt, hid.CmdPrintScreen) {
		netReceiver.SetUsbString()
	}
}

func (self *Action) kb_reboot() {
	if self.checkKeyIsPressedByOrder(hid.LeftCtrl+hid.LeftShift, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		self.ClientTx.CmdSetDefaultCfg()
		time.Sleep(2 * time.Second)
		self.ClientTx.CmdReset()
	}
}

func (self *Action) kb_test() {
	if self.checkKeyIsPressedByOrder(0, hid.CmdScrollLock) {
		self.ClientTx.CmdGetParaCfg()
	}
}
