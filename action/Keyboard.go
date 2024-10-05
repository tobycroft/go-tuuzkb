package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netSender"
)

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		//self.ClientTx.CmdSendKbGeneralData(c)
		//fmt.Println("keybaordrecv", c.Ctrl, c)
		go self.kb_actvate(c)
		go self.kb_banSomeKeys(c)
		go self.kb_reboot(c)
		fmt.Println("keybaordrecv", c)
		self.ClientTx.CmdSendKbGeneralData(c)

	}
	panic("键盘通道意外结束")
}

func (self *Action) kb_washing(c int16) {
	self.MaskKey.Store(c, true)
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
		self.kb_add_masking(hid.CmdApplication)
		self.kb_add_masking(hid.CmdPrintScreen)
		self.kb_add_masking(hid.CmdPause)
		self.kb_add_masking(hid.CmdScrollLock)
		self.kb_add_masking(hid.CmdRightControl)
		//fmt.Println("bankey")
		self.MaskKey.Range(func(key, value interface{}) bool {
			fmt.Println("bankey", key, value)
			return true
		})
	}
}

func (self *Action) kb_reboot(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdPrintScreen) {
		self.ClientTx.CmdReset()
	}
}

func (self *Action) kb_test(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdA) {
		fmt.Println("testa")
	}
}

func (self *Action) kb_add_masking(mask_key byte) {
	self.MaskKey.Store(mask_key, true)
}

func (self *Action) kb_remove_masking(mask_key byte) {
	self.MaskKey.Delete(mask_key)
}

func (self *Action) kb_chec_mask(mask_key byte) bool {
	_, ok := self.MaskKey.Load(mask_key)
	return ok
}
