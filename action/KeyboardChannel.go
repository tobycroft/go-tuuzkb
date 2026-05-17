package action

import (
	"fmt"
	"sync/atomic"
	"time"

	"main.go/define/hid"
	"main.go/netReceiver"
	"main.go/netSender"
)

type lastKey struct {
	Ctrl   atomic.Uint32
	Button [6]atomic.Uint32
}

var LastPress = &lastKey{}
var CurrentPress = &lastKey{}
var OnchangePress = &lastKey{}

func (self *Action) keyboard_runnable() {
	self.ready()
	for c := range netReceiver.Crx.KeyboardRxChannel {
		swap_key(c)
		//fmt.Println("keybaordrecv", c, OnchangePress.Ctrl.Load(), OnchangePress.Button)
		self.kb_reset()
		self.kb_reboot()
		//self.kb_unbanall()
		//self.kb_test()
		self.key_main()
		self.qe_main()
		self.whel_main()
		//self.kb_get_para()
		//self.kb_set_para()
		//self.kb_get_usbstring()
		//self.kb_set_usbstring()
		//self.kb_bansomeKeys()
		self.SendKbGeneralDataRaw()

	}
	panic("键盘通道意外结束")
}

func swap_key(c *netSender.KeyboardData2) {
	if byte(CurrentPress.Ctrl.Load()) == c.Ctrl {
		OnchangePress.Ctrl.Store(uint32(0))
	} else {
		OnchangePress.Ctrl.Store(uint32(c.Ctrl))
	}
	LastPress.Ctrl.Store(CurrentPress.Ctrl.Load())
	CurrentPress.Ctrl.Store(uint32(c.Ctrl))
	for i := 0; i < 6; i++ {
		if c.Button[i] == byte(CurrentPress.Button[i].Load()) {
			OnchangePress.Button[i].Store(uint32(0))
		} else {
			if i < 5 && c.Button[i] == byte(CurrentPress.Button[i+1].Load()) {
				OnchangePress.Button[i].Store(uint32(0))
			} else {
				OnchangePress.Button[i].Store(uint32(c.Button[i]))
			}
		}
	}
	for i := 0; i < 6; i++ {
		LastPress.Button[i].Store(CurrentPress.Button[i].Load())
		CurrentPress.Button[i].Store(uint32(c.Button[i]))
	}
}

func (self *Action) ready() {
	CurrentPress.Ctrl.Store(uint32(0))
	LastPress.Ctrl.Store(uint32(0))
	OnchangePress.Ctrl.Store(uint32(0))

	Endpoint_delay.Store(0)
	Endpoint_BeforeDelay.Store(50)
	Endpoint_BeforeDelay_Random.Store(20)

	for i := range OnchangePress.Button {
		CurrentPress.Button[i].Store(uint32(0))
		LastPress.Button[i].Store(uint32(0))
		OnchangePress.Button[i].Store(uint32(0))
	}
}

func (self *Action) kb_reset() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(50)
		Endpoint_BeforeDelay_Random.Store(20)
		go fmt.Println("Reset")
		go Lcd_refresh()
		//go self.Km.KmNetLcdPicture_tempSet("Golang", "GolangGolang", "GolangGolangGolang", 1*time.Second)
	}
}

func (self *Action) kb_bansomeKeys() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen) {
		go Kb_banSomeKeys()
		go fmt.Println("ban_all")
	}
}

func (self *Action) kb_unbanall() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		Mask.ButtonMu.Lock()
		Mask.Button = make(map[byte]bool)
		Mask.ButtonMu.Unlock()
		Mask.CtrlMu.Lock()
		Mask.Ctrl = make(map[byte]bool)
		Mask.CtrlMu.Unlock()
		go fmt.Println("unbanall")
		go Lcd_refresh()
	}
}

func (self *Action) kb_get_para() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdScrollLock) {
		netSender.Ctx.CmdGetParaCfg()
	}
}

func (self *Action) kb_set_para() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdScrollLock) {
		netSender.Ctx.CmdGetParaCfg()
	}
}

func (self *Action) kb_get_usbstring() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdPrintScreen) {
		netSender.Ctx.CmdGetUsbString(netSender.StrTypeManufacturer)
		netSender.Ctx.CmdGetUsbString(netSender.StrTypeProduct)
		netSender.Ctx.CmdGetUsbString(netSender.StrTypeSerial)
	}
}

func (self *Action) kb_set_usbstring() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift+hid.RightAlt, hid.CmdPrintScreen) {
		netReceiver.SetUsbString()
	}
}

func (self *Action) kb_reboot() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightShift, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		//netSender.Ctx.CmdSetDefaultCfg()
		time.Sleep(2 * time.Second)
		netSender.Ctx.CmdReset()
	}
}

func (self *Action) kb_test() {
	if self.checkKeyIsPressedByOrder(0, hid.CmdScrollLock) {
		netSender.Ctx.CmdGetParaCfg()
	}
}
