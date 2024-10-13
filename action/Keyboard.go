package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netReceiver"
	"main.go/netSender"
	"time"
)

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		self.c = c
		//fmt.Println("keybaordrecv", c)
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
			go self.kb_banSomeKeys()
			fmt.Println("ban_all")
		}
		self.SendKbGeneralDataRaw()

	}
	panic("键盘通道意外结束")
}

func (self *Action) kb_actvate() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(0)
		fmt.Println("Reset")
		//go self.Km.KmNetLcdPicture_tempSet("Golang", "GolangGolang", "GolangGolangGolang", 1*time.Second)
	}
}

func (self *Action) kb_banSomeKeys() {
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
