package action

import (
	"fmt"
	"main.go/common"
	"main.go/define/hid"
	"main.go/netSender"
	"time"
)

func (self *Action) keyboard_runnable() {
	for c := range self.ClientRx.KeyboardRxChannel {
		//self.ClientTx.CmdSendKbGeneralData(c)
		//fmt.Println("keybaordrecv", c.Ctrl, c)
		if self.maskingKeyBoard2(&c) > 0 {
			self.ClientTx.CmdSendKbGeneralData(c)
			self.kb_actvate(c)
			fmt.Println("keybaordrecv", c)
		}
	}
	panic("键盘通道意外结束")
}

func (self *Action) banKey(key byte) byte {
	if hid.CmdErrorRollOver == key {
		return 0x00
	}
	return key
}

func (self *Action) checkKeyIsPressed(c netSender.KeyboardData, Ctrl, Btn byte) bool {
	switch Btn {
	case c.Button0:
		return c.Ctrl == Ctrl

	case c.Button1:
		return c.Ctrl == Ctrl

	case c.Button2:
		return c.Ctrl == Ctrl

	case c.Button3:
		return c.Ctrl == Ctrl

	case c.Button4:
		return c.Ctrl == Ctrl

	case c.Button5:
		return c.Ctrl == Ctrl

	default:
		return false
	}
}

func (self *Action) maskingKeyBoard2(c *netSender.KeyboardData) int {
	num := 0
	if self.Ctrl != c.Ctrl {
		self.Ctrl = c.Ctrl
		num += 1
	}
	if self.Button0 != self.banKey(c.Button0) {
		self.Button0 = self.banKey(c.Button0)
		num += 1
	}
	if self.Button1 != self.banKey(c.Button1) {
		self.Button1 = self.banKey(c.Button1)
		num += 1
	}
	if self.Button2 != self.banKey(c.Button2) {
		self.Button2 = self.banKey(c.Button2)
		num += 1
	}
	if self.Button3 != self.banKey(c.Button3) {
		self.Button3 = self.banKey(c.Button3)
		num += 1
	}
	if self.Button4 != self.banKey(c.Button4) {
		self.Button4 = self.banKey(c.Button4)
		num += 1
	}
	if self.Button5 != self.banKey(c.Button5) {
		self.Button5 = self.banKey(c.Button5)
		num += 1
	}

	return num
}

func (self *Action) kb_actvate(c netSender.KeyboardData) {
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(0)
		//go self.Km.KmNetLcdPicture_tempSet("Golang", "GolangGolang", "GolangGolangGolang", 1*time.Second)
	} else {
		go self.key_main(c)
	}
}

func (self *Action) kb_add_masking(mask_key byte) {
	self.MaskKey.Store(mask_key, true)
}

func (self *Action) kb_banSomeKeys(c, d *function.KeyPressed) {
	if c.RightCtrl && c.RightAlt && c.ScrollLock {
		//self.km.KmNetSetVidPid(0x05ac, 0x0256)
		//self.km.KmNetSetVidPid(0x05ac, 0x0256)
		self.Km.KmNetMaskKeyboard(hid.CmdApplication)
		self.Km.KmNetMaskKeyboard(hid.CmdPrintScreen)
		self.Km.KmNetMaskKeyboard(hid.CmdPause)
		self.Km.KmNetMaskKeyboard(hid.CmdScrollLock)
		self.Km.KmNetMaskKeyboard(hid.CmdRightControl)
		self.Km.KmNetMaskKeyboard(1)
		go self.Km.KmNetLcdPicture_tempSet("Key", "Re ban", "Complete", 1*time.Second)
	} else if c.RightCtrl && c.RightAlt && c.PauseBreak {
		go self.test_key(c, d)
	}
}

func (self *Action) kb_banKey_direct() {
	for {
		time.Sleep(1 * time.Second)
		if self.Km.Fresh {
			self.Km.Fresh = false
			time.Sleep(2 * time.Second)
			self.Km.KmNetMaskKeyboard(hid.CmdApplication)
			self.Km.KmNetMaskKeyboard(hid.CmdPrintScreen)
			self.Km.KmNetMaskKeyboard(hid.CmdPause)
			self.Km.KmNetMaskKeyboard(hid.CmdScrollLock)
			self.Km.KmNetMaskKeyboard(hid.CmdRightControl)
			self.Km.KmNetMaskKeyboard(1)
			go self.Km.KmNetLcdPicture_tempSet("Key", "AutoBan", "Complete", 5*time.Second)
		}
	}
}

func (self *Action) kb_reboot(c, d *function.KeyPressed) {
	if c.RightCtrl && c.RightAlt && d.Application && d.PrintScreen {
		go self.Km.KmNetReboot()
		common.PrintRedis("manual reboot")
	}
}

func (self *Action) test_key(c, d *function.KeyPressed) {
	if c.RightCtrl && c.PauseBreak {
		//self.km.KmNetReboot()
		//self.km.KmNetMouseWheel(1)
		//time.Sleep(100 * time.Millisecond)
		//self.km.KmNetMouseWheel(-1)
		//self.km.KmNetMouseWheel(1)
		t := float64(250)
		start := time.Now().UnixMilli()
		for i := float64(0); i < t; i++ {
			self.Km.KmNetKeyDown(hid.CmdQ)
			self.Km.KmNetKeyUp(hid.CmdQ)
			self.Km.KmNetKeyDown(hid.CmdR)
			self.Km.KmNetKeyUp(hid.CmdR)
			//self.qe_auto()
		}
		end := float64(time.Now().UnixMilli() - start)
		de := end / (t * 4)
		qps := ((t * 4) / end) * 1000
		common.PrintRedis("时间使用:", end, time.Duration(end)*time.Millisecond, "总次数:", 4*t, "单次执行耗时ms:", de, "每秒QPS:", qps, "实际执行:", qps/2)
	}
}
