package action

import (
	"fmt"
	"main.go/define/hid"
)

func (self *Action) keyboard_runnable() {
	self.kb_add_masking(hid.CmdApplication, false)
	self.kb_add_masking(hid.CmdPrintScreen, false)
	self.kb_add_masking(hid.CmdPause, false)
	self.kb_add_masking(hid.CmdScrollLock, false)
	self.kb_add_masking(hid.RightCtrl, true)

	for c := range self.ClientRx.KeyboardRxChannel {
		self.c = c
		//fmt.Println("keybaordrecv", c)
		go self.kb_actvate()
		go self.kb_banSomeKeys()
		go self.kb_reboot()
		go self.kb_unbanall()
		//go self.kb_test()
		go self.key_main()
		go self.qe_main()
		go self.whel_main()
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
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen) {
		self.kb_add_masking(hid.CmdApplication, false)
		self.kb_add_masking(hid.CmdPrintScreen, false)
		self.kb_add_masking(hid.CmdPause, false)
		self.kb_add_masking(hid.CmdScrollLock, false)
		self.kb_add_masking(hid.RightCtrl, true)
		fmt.Println("ban_all")
	}
}

func (self *Action) kb_unbanall() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdPrintScreen) {
		Mask.Button.Clear()
		Mask.Ctrl.Clear()
		fmt.Println("unbanall")
	}
}

func (self *Action) kb_reboot() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		self.ClientTx.CmdReset()
	}
}

func (self *Action) kb_test() {
	if self.checkKeyIsPressedByOrder(0, hid.CmdScrollLock) {
		self.ClientTx.CmdSendMsRelWheel(-1)
	}
}
