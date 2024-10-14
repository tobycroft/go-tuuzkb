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
	Ctrl   atomic.Value
	Button [6]atomic.Value
}

var LastPress = &lastKey{}
var CurrentPress = &lastKey{}
var OnchangePress = &lastKey{}

func (self *Action) keyboard_runnable() {
	self.ready()
	for c := range self.ClientRx.KeyboardRxChannel {
		swap_key(&c)
		//fmt.Println("keybaordrecv", c, OnchangePress.Ctrl.Load(), OnchangePress.Button)
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
	if CurrentPress.Ctrl.Load().(byte) == c.Ctrl {
		OnchangePress.Ctrl.Store(byte(0))
	} else {
		OnchangePress.Ctrl.Store(c.Ctrl)
	}
	LastPress.Ctrl.Store(CurrentPress.Ctrl.Load().(byte))
	CurrentPress.Ctrl.Store(c.Ctrl)
	for i := 0; i < 6; i++ {
		if c.Button[i] == CurrentPress.Button[i].Load() {
			OnchangePress.Button[i].Store(byte(0))
		} else {
			if i < 6 && c.Button[i] == CurrentPress.Button[i+1].Load() {
				OnchangePress.Button[i].Store(byte(0))
			} else {
				OnchangePress.Button[i].Store(c.Button[i])
			}
		}
	}
	for i := 0; i < 6; i++ {
		LastPress.Button[i].Store(CurrentPress.Button[i].Load().(byte))
		CurrentPress.Button[i].Store(c.Button[i])
	}
}

func (self *Action) ready() {
	CurrentPress.Ctrl.Store(byte(0))
	LastPress.Ctrl.Store(byte(0))
	OnchangePress.Ctrl.Store(byte(0))
	for i := range OnchangePress.Button {
		CurrentPress.Button[i].Store(byte(0))
		LastPress.Button[i].Store(byte(0))
		OnchangePress.Button[i].Store(byte(0))
	}
}

func (self *Action) kb_actvate() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl, hid.CmdScrollLock) {
		Endpoint_delay.Store(0)
		Endpoint_BeforeDelay.Store(0)
		fmt.Println("Reset")
		go Lcd_refresh()
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
	go Lcd_refresh()
}

func (self *Action) kb_unbanall() {
	if self.checkKeyIsPressedByOrder(hid.RightCtrl+hid.RightAlt, hid.CmdPrintScreen, hid.CmdScrollLock, hid.CmdPause) {
		Mask.Button.Clear()
		Mask.Ctrl.Clear()
		fmt.Println("unbanall")
		go Lcd_refresh()
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
