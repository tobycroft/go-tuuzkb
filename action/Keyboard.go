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
		out := self.kb_gen_output(c)
		if out.Resv != self.lastPress {
			self.lastPress = out.Resv
			self.ClientTx.CmdSendKbGeneralDataRaw(out)
			fmt.Println("keybaordsnd", out)
		}

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
		self.kb_add_masking(hid.LeftWindows, true)
		//fmt.Println("bankey")
		self.Mask.Button.Range(func(key, value interface{}) bool {
			fmt.Println("banbutton", key, value)
			return true
		})
		self.Mask.Ctrl.Range(func(key, value interface{}) bool {
			fmt.Println("banctrl", key, value)
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

func (self *Action) kb_add_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		self.Mask.Ctrl.Store(key, true)
	} else {
		self.Mask.Button.Store(key, true)
	}
}

func (self *Action) kb_remove_masking(key byte, is_ctrl bool) {
	if is_ctrl {
		self.Mask.Ctrl.Delete(key)
	} else {
		self.Mask.Button.Delete(key)
	}
}

func (self *Action) kb_chec_mask(key byte, is_ctrl bool) bool {
	if is_ctrl {
		_, ok := self.Mask.Ctrl.Load(key)
		return ok
	} else {
		_, ok := self.Mask.Button.Load(key)
		return ok
	}
}
