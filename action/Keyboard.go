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
		self.ClientTx.CmdSendKbGeneralData(c)
		go self.kb_actvate(c)
		go self.kb_banSomeKeys(c)
		go self.kb_reboot(c)
		fmt.Println("keybaordrecv", c)
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
	if self.checkKeyIsPressed(c, hid.RightCtrl+hid.RightAlt, hid.CmdApplication, hid.CmdPrintScreen) {
		self.kb_reboot()
		return
	}
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

func (self *Action) kb_reboot() {
	self.ClientTx.CmdReset()
}

//func (self *Action) test_key(c, d *function.KeyPressed) {
//	if c.RightCtrl && c.PauseBreak {
//		//self.km.KmNetReboot()
//		//self.km.KmNetMouseWheel(1)
//		//time.Sleep(100 * time.Millisecond)
//		//self.km.KmNetMouseWheel(-1)
//		//self.km.KmNetMouseWheel(1)
//		t := float64(250)
//		start := time.Now().UnixMilli()
//		for i := float64(0); i < t; i++ {
//			self.Km.KmNetKeyDown(hid.CmdQ)
//			self.Km.KmNetKeyUp(hid.CmdQ)
//			self.Km.KmNetKeyDown(hid.CmdR)
//			self.Km.KmNetKeyUp(hid.CmdR)
//			//self.qe_auto()
//		}
//		end := float64(time.Now().UnixMilli() - start)
//		de := end / (t * 4)
//		qps := ((t * 4) / end) * 1000
//		common.PrintRedis("时间使用:", end, time.Duration(end)*time.Millisecond, "总次数:", 4*t, "单次执行耗时ms:", de, "每秒QPS:", qps, "实际执行:", qps/2)
//	}
//}

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
