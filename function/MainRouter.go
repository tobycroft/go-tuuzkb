package function

import (
	"main.go/common"
	"main.go/define/cmd"
	"time"
)

type RecvRouter struct {
	Keyboard chan byte
}

func (self RecvRouter) MainRouter() {
	go self.mainRouter()
}

func (self RecvRouter) mainRouter() {
	for client_tx := range self.recv {
		Mode := client_tx.Head.Cmd
		switch Mode {

		case cmd.CmdReboot:
			common.PrintRedis("盒子重启", client_tx.Head.Indexpts)

			time.Sleep(5 * time.Second)
			if self.conn != nil {
				self.conn.Close()
			}
			if self.connMonitor != nil {
				self.connMonitor.Close()
			}
			if self.debug.connDebug != nil {
				self.debug.connDebug.Close()
			}
			time.Sleep(1 * time.Second)
			_ = self.kmReconnect()
			self.kmNetMonitor()
			break

		case cmd.CmdConnect:
			common.PrintRedis("连接成功！", client_tx.Head.Indexpts)
			self.kmNetMonitor()
			self.Fresh = true
			break

		case cmd.CmdDebug:
			common.PrintRedis("Debug:", client_tx.Head.Indexpts, client_tx)
			self.kmNetDebugServer()
			break

		case cmd.CmdMonitor:
			self.kmNetMonitorServer()
			if self.DebugClient {
				common.PrintRedis("Monitor推送准备完毕，等待数据！", client_tx.Head.Indexpts)
			}
			break

		case cmd.CmdSetVidPid:
			common.PrintRedis("VIDPID设定完成", client_tx.Head.Rand)
			self.KmNetReboot()
			break

		case cmd.CmdUnmaskAll:
			if self.DebugClient {
				common.PrintRedis("清除所有按键屏蔽", client_tx.Head.Rand)
			}
			break

		case cmd.CmdMaskMouse:
			if self.DebugClient {
				common.PrintRedis("按键屏蔽成功", client_tx.Head.Rand)
			}
			break

		case cmd.CmdBezierMove:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdBezierMove", client_tx.CmdMouse.X, client_tx.CmdMouse.Y)
			}
			break

		case cmd.CmdKeyboardAll:
			if self.SendHoldMode {
				RecvPts.Swap(client_tx.Head.Indexpts)
				self.recv_keyboard_operation <- client_tx.Head.Indexpts
			}
			if self.KeyState.KeyBoardDebug.MessageRecv {
				common.PrintRedis("CmdKeyboardAll", client_tx.Head.Indexpts, client_tx.CmdKeyboard)
			}
			break

		case cmd.CmdMouseAutoMove:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseAutoMove", client_tx.CmdMouse.X, client_tx.CmdMouse.Y)
			}
			break

		case cmd.CmdMouseWheel:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseWheel", client_tx.CmdMouse.Wheel)
			}
			break

		case cmd.CmdMouseLeft:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseLeft", client_tx.CmdMouse.Button)
			}
			break

		case cmd.CmdMouseRight:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseRight", client_tx.CmdMouse.Button)
			}
			break

		case cmd.CmdMouseMiddle:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseMiddle", client_tx.CmdMouse.Button)
			}
			break

		case cmd.CmdMouseMove:
			if self.KeyState.MouseDebug.MessageRecv {
				common.PrintRedis("CmdMouseMove", client_tx.CmdMouse.X, client_tx.CmdMouse.Y)
			}
			break

		case cmd.CmdShowPic:
			if self.DebugClient {
				common.PrintRedis("图片设定")
			}
			break

		default:
			common.PrintRedis("未定义事件:", client_tx.Head.Cmd, client_tx)
			break

		}
	}
}
