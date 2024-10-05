package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *Action) keyboard_runnable() {
	self.kb_add_masking(hid.CmdApplication, false)
	self.kb_add_masking(hid.CmdPrintScreen, false)
	self.kb_add_masking(hid.CmdPause, false)
	self.kb_add_masking(hid.CmdScrollLock, false)
	self.kb_add_masking(hid.RightCtrl, true)

	for c := range self.ClientRx.KeyboardRxChannel {
		fmt.Println("keybaordrecv", c)
		go self.kb_actvate(c)
		go self.kb_banSomeKeys(c)
		go self.kb_reboot(c)
		go self.kb_unbanall(c)
		self.SendKbGeneralDataRaw(c)

	}
	panic("键盘通道意外结束")
}

func (self *Action) kb_actvate(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(0)
		fmt.Println("aaa")
		//go self.Km.KmNetLcdPicture_tempSet("Golang", "GolangGolang", "GolangGolangGolang", 1*time.Second)
	} else {
		//go self.key_main(c)
	}
}

func (self *Action) kb_banSomeKeys(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen) {
		self.kb_add_masking(hid.CmdApplication, false)
		self.kb_add_masking(hid.CmdPrintScreen, false)
		self.kb_add_masking(hid.CmdPause, false)
		self.kb_add_masking(hid.CmdScrollLock, false)
		self.kb_add_masking(hid.RightCtrl, true)
	}
}

func (self *Action) kb_unbanall(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdPrintScreen) {
		self.Mask.Button.Clear()
		self.Mask.Ctrl.Clear()
		fmt.Println("unbanall")
	}
}

func (self *Action) kb_reboot(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen, hid.CmdPause, hid.CmdScrollLock) {
		self.ClientTx.CmdReset()
	}
}

func (self *Action) kb_test(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdA) {
		fmt.Println("testa")
	}
}
