package action

import (
	"fmt"
	"main.go/define/hid"
	"main.go/netReceiver"
	"main.go/netSender"
	"sync"
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

// 任务类型
type taskType int

const (
	taskKeyMain taskType = iota
	taskQeMain
	taskWhelMain
	taskKbReset
	taskKbReboot
)

// 工作池配置
const (
	workerCount = 4
	taskQueueSize = 32
)

var (
	taskQueue    chan func()
	workerPoolOnce sync.Once
)

// 初始化工作池
func initWorkerPool() {
	workerPoolOnce.Do(func() {
		taskQueue = make(chan func(), taskQueueSize)
		for i := 0; i < workerCount; i++ {
			go worker()
		}
	})
}

// 工作协程
func worker() {
	for task := range taskQueue {
		task()
	}
}

// 提交任务到工作池
func submitTask(task func()) {
	select {
	case taskQueue <- task:
	default:
		// 如果队列已满，直接在当前 goroutine 执行（防止阻塞）
		task()
	}
}

func (self *Action) keyboard_runnable() {
	initWorkerPool()
	self.ready()
	for c := range netReceiver.Crx.KeyboardRxChannel {
		swap_key(c)
		//fmt.Println("keybaordrecv", c, OnchangePress.Ctrl.Load(), OnchangePress.Button)
		submitTask(func() { self.kb_reset() })
		submitTask(func() { self.kb_reboot() })
		submitTask(func() { self.key_main() })
		submitTask(func() { self.qe_main() })
		submitTask(func() { self.whel_main() })
		self.SendKbGeneralDataRaw()

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
			if i < 5 && c.Button[i] == CurrentPress.Button[i+1].Load() {
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

	Endpoint_delay.Store(0)
	Endpoint_BeforeDelay.Store(50)
	Endpoint_BeforeDelay_Random.Store(20)

	for i := range OnchangePress.Button {
		CurrentPress.Button[i].Store(byte(0))
		LastPress.Button[i].Store(byte(0))
		OnchangePress.Button[i].Store(byte(0))
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
		Mask.Button.Clear()
		Mask.Ctrl.Clear()
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