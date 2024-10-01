package netReceiver

import (
	"fmt"
)

func (self *Reciever) MonitorKeyboard() {
	for report := range self.KeyboardReport {
		fmt.Println(report, report.Button, report.Ctrl)
		//self.KeyState.waitGroup.Add(8)
		//keyPressed := &KeyPressed{KeyPressDebug: self.KeyState.KeyBoardDebug.MessagePress}
		//keyStayPressed := &KeyPressed{KeyPressDebug: self.KeyState.KeyBoardDebug.MessagePressAndHold}
		//go self.KeyState.keyboard_function(client_tx.Buttons, keyPressed)
		//go self.KeyState.keyboard_function(client_tx.Buttons, keyStayPressed)
		//go self.keyboard_state1(client_tx.Data[0], keyPressed, keyStayPressed)
		//go self.keyboard_state2(client_tx.Data[1], keyPressed, keyStayPressed)
		//go self.keyboard_state3(client_tx.Data[2], keyPressed, keyStayPressed)
		//go self.keyboard_state4(client_tx.Data[3], keyPressed, keyStayPressed)
		//go self.keyboard_state5(client_tx.Data[4], keyPressed, keyStayPressed)
		//go self.keyboard_state6(client_tx.Data[5], keyPressed, keyStayPressed)
		//self.KeyState.waitGroup.Wait()
		//self.KeyChannel <- KeyAll{keyPressed, keyStayPressed}
	}
}

//
//func (self *Reciever) keyboard_state1(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state1.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下1:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state1.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state1.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开1:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state1.Button, false, keypress)
//		self.KeyState.state1.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持1:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *Reciever) keyboard_state2(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state2.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下2:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state2.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state2.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开2:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state2.Button, false, keypress)
//		self.KeyState.state2.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持2:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *Reciever) keyboard_state3(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state3.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下3:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state3.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state3.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开3:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state3.Button, false, keypress)
//		self.KeyState.state3.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持3:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *Reciever) keyboard_state4(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state4.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下4:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state4.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state4.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开4:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state4.Button, false, keypress)
//		self.KeyState.state4.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持4:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *Reciever) keyboard_state5(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state5.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下5:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state5.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state5.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开5:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state5.Button, false, keypress)
//		self.KeyState.state5.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持5:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *Reciever) keyboard_state6(key uint8, keypress *KeyPressed, keyStayPressed *KeyPressed) {
//	if self.KeyState.state6.Button == hid.CmdNone && key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序按下6:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		self.KeyState.keyboard_keys(key, true, keypress)
//		self.KeyState.state6.Button = key
//		keypress.KeyDown, keyStayPressed.KeyDown = key, key
//	} else if key == hid.CmdNone && self.KeyState.state6.Button != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "顺序松开6:")
//		}
//		//self.KeyState.keyboard_keys(self.KeyState.state6.Button, false, keypress)
//		self.KeyState.state6.Button = key
//		keypress.KeyUp, keyStayPressed.KeyUp = key, key
//	} else if key != hid.CmdNone {
//		if self.KeyState.KeyBoardDebug.MessagePress {
//			common.PrintRedis("Keyboard", "按键保持6:")
//		}
//		self.KeyState.keyboard_keys(key, true, keyStayPressed)
//		keypress.KeyCurrent, keyStayPressed.KeyCurrent = key, key
//	}
//	self.KeyState.waitGroup.Done()
//}
//
//func (self *keyboardState) keyboard_keys(key uint8, state bool, keypress *KeyPressed) {
//	keypress.KeyCurrent = key
//	switch key {
//	//F1-F12
//	case hid.CmdF1:
//		keypress.F1 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F1")
//		}
//		break
//	case hid.CmdF2:
//		keypress.F1 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F2")
//		}
//		break
//	case hid.CmdF3:
//		keypress.F3 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F3")
//		}
//		break
//	case hid.CmdF4:
//		keypress.F4 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F4")
//		}
//		break
//	case hid.CmdF5:
//		keypress.F5 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F5")
//		}
//		break
//	case hid.CmdF6:
//		keypress.F6 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F6")
//		}
//		break
//	case hid.CmdF7:
//		keypress.F7 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F7")
//		}
//		break
//	case hid.CmdF8:
//		keypress.F8 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F8")
//		}
//		break
//	case hid.CmdF9:
//		keypress.F9 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F9")
//		}
//		break
//	case hid.CmdF10:
//		keypress.F10 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F10")
//		}
//		break
//	case hid.CmdF11:
//		keypress.F11 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F11")
//		}
//		break
//	case hid.CmdF12:
//		keypress.F12 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F12")
//		}
//		break
//
//	//special keys
//	case hid.CmdGraveAccentAndTilde:
//		keypress.Backquote = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "~")
//		}
//		break
//	case hid.CmdMinusUnderscore:
//		keypress.Minus = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "-")
//		}
//		break
//	case hid.CmdEqualPlus:
//		keypress.Equals = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "=")
//		}
//		break
//	case hid.CmdOBracketAndOBrace:
//		keypress.BraceL = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "{[")
//		}
//		break
//	case hid.CmdCBRacketAndCBrace:
//		keypress.BraceR = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "}]")
//		}
//		break
//	case hid.CmdBackslashVerticalBar:
//		keypress.Backslash = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "\\|")
//		}
//		break
//	case hid.CmdSemicolonColon:
//		keypress.Semicolon = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", ";:")
//		}
//		break
//	case hid.CmdSingleAndDoubleQuote:
//		keypress.Quote = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "'\"")
//		}
//		break
//	case hid.CmdCommaAndLess:
//		keypress.Comma = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "<,")
//		}
//		break
//	case hid.CmdDotGreater:
//		keypress.Dot = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", ".>")
//		}
//		break
//	case hid.CmdSlashQuestion:
//		keypress.QuestionSlash = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "?/")
//		}
//		break
//
//	//0-9
//	case hid.Cmd0CParenthesis:
//		keypress.Num0 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "0")
//		}
//		break
//	case hid.Cmd1ExclamationMark:
//		keypress.Num1 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "1")
//		}
//		break
//	case hid.Cmd2At:
//		keypress.Num2 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "2")
//		}
//		break
//	case hid.Cmd3NumberSign:
//		keypress.Num3 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "3")
//		}
//		break
//	case hid.Cmd4Dollar:
//		keypress.Num4 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "4")
//		}
//		break
//	case hid.Cmd5Percent:
//		keypress.Num5 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "5")
//		}
//		break
//	case hid.Cmd6Caret:
//		keypress.Num6 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "6")
//		}
//		break
//	case hid.Cmd7Ampersand:
//		keypress.Num7 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "7")
//		}
//		break
//	case hid.Cmd8Asterisk:
//		keypress.Num8 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "8")
//		}
//		break
//	case hid.Cmd9OParenthesis:
//		keypress.Num9 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "9")
//		}
//		break
//
//	//A-Z
//	case hid.CmdA:
//		keypress.A = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "A")
//		}
//		break
//	case hid.CmdB:
//		keypress.B = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "B")
//		}
//		break
//	case hid.CmdC:
//		keypress.C = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "C")
//		}
//		break
//	case hid.CmdD:
//		keypress.D = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "D")
//		}
//		break
//	case hid.CmdE:
//		keypress.E = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "E")
//		}
//		break
//	case hid.CmdF:
//		keypress.F = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "F")
//		}
//		break
//	case hid.CmdG:
//		keypress.G = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "G")
//		}
//		break
//	case hid.CmdH:
//		keypress.H = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "H")
//		}
//		break
//	case hid.CmdI:
//		keypress.I = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "I")
//		}
//		break
//	case hid.CmdJ:
//		keypress.J = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "J")
//		}
//		break
//	case hid.CmdK:
//		keypress.K = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "K")
//		}
//		break
//	case hid.CmdL:
//		keypress.L = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "L")
//		}
//		break
//	case hid.CmdM:
//		keypress.M = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "M")
//		}
//		break
//	case hid.CmdN:
//		keypress.N = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "N")
//		}
//		break
//	case hid.CmdO:
//		keypress.O = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "O")
//		}
//		break
//	case hid.CmdP:
//		keypress.P = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "P")
//		}
//		break
//	case hid.CmdQ:
//		keypress.Q = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Q")
//		}
//		break
//	case hid.CmdR:
//		keypress.R = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "R")
//		}
//		break
//	case hid.CmdS:
//		keypress.S = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "S")
//		}
//		break
//	case hid.CmdT:
//		keypress.T = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "T")
//		}
//		break
//	case hid.CmdU:
//		keypress.U = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "U")
//		}
//		break
//	case hid.CmdV:
//		keypress.V = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "V")
//		}
//		break
//	case hid.CmdW:
//		keypress.W = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "W")
//		}
//		break
//	case hid.CmdX:
//		keypress.X = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "X")
//		}
//		break
//	case hid.CmdY:
//		keypress.Y = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Y")
//		}
//		break
//	case hid.CmdZ:
//		keypress.Z = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Z")
//		}
//		break
//
//	//function key
//	case hid.CmdPrintScreen:
//		keypress.PrintScreen = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "PrintScreen")
//		}
//		break
//	case hid.CmdScrollLock:
//		keypress.ScrollLock = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "ScrollLock")
//		}
//		break
//	case hid.CmdPause:
//		keypress.PauseBreak = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "PauseBreak")
//		}
//		break
//	case hid.CmdInsert:
//		keypress.Insert = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Insert")
//		}
//		break
//	case hid.CmdHome:
//		keypress.Home = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Home")
//		}
//		break
//	case hid.CmdPageUp:
//		keypress.PageUp = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "PageUp")
//		}
//		break
//	case hid.CmdEnd1:
//		keypress.End = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "End")
//		}
//		break
//	case hid.CmdPageDown:
//		keypress.PageDown = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "PageDown")
//		}
//		break
//	case hid.CmdCapsLock:
//		keypress.CapsLock = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "CapsLock")
//		}
//		break
//	case hid.CmdTab:
//		keypress.Tab = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Tab")
//		}
//		break
//	case hid.CmdBackspace:
//		keypress.Backspace = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Backspace")
//		}
//		break
//	case hid.CmdEnter:
//		keypress.Enter = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Enter")
//		}
//		break
//	case hid.CmdSpacebar:
//		keypress.Space = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Space")
//		}
//		break
//	case hid.CmdEscape:
//		keypress.Esc = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Esc")
//		}
//		break
//	case hid.CmdDelete:
//		keypress.Delete = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Delete")
//		}
//		break
//
//	//arrow keys
//	case hid.CmdUpArrow:
//		keypress.ArrowUp = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "ArrowUp")
//		}
//		break
//	case hid.CmdDownArrow:
//		keypress.ArrowDown = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "ArrowDown")
//		}
//		break
//	case hid.CmdLeftArrow:
//		keypress.ArrowLeft = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "ArrowLeft")
//		}
//		break
//	case hid.CmdRightArrow:
//		keypress.ArrowRight = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "ArrowRight")
//		}
//		break
//
//	//keypad
//	case hid.CmdKeypad0Insert:
//		keypress.NumPad0 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad0")
//		}
//		break
//	case hid.CmdKeypad1End:
//		keypress.NumPad1 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad1")
//		}
//		break
//	case hid.CmdKeypad2DownArrow:
//		keypress.NumPad2 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad2")
//		}
//		break
//	case hid.CmdKeypad3PageDown:
//		keypress.NumPad3 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad3")
//		}
//		break
//	case hid.CmdKeypad4LeftArrow:
//		keypress.NumPad4 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad4")
//		}
//		break
//	case hid.CmdKeypad5:
//		keypress.NumPad5 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad5")
//		}
//		break
//	case hid.CmdKeypad6RightArrow:
//		keypress.NumPad6 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad6")
//		}
//		break
//	case hid.CmdKeypad7Home:
//		keypress.NumPad7 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad7")
//		}
//		break
//	case hid.CmdKeypad8UpArrow:
//		keypress.NumPad8 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad8")
//		}
//		break
//	case hid.CmdKeypad9PageUp:
//		keypress.NumPad9 = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPad9")
//		}
//		break
//	case hid.CmdKeypadMinus:
//		keypress.NumPadMinus = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPadMinus")
//		}
//		break
//	case hid.CmdKeypadPlus:
//		keypress.NumPadPlus = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "CmdKeypadPlus")
//		}
//		break
//	case hid.CmdKeypadAsterisk:
//		keypress.NumPadMultiply = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPadMultiply")
//		}
//		break
//	case hid.CmdKeypadNumLockAndClear:
//		keypress.NumLock = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumLock")
//		}
//		break
//	case hid.CmdKeypadSlash:
//		keypress.NumPadDivide = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPadDivide")
//		}
//		break
//	case hid.CmdKeypadEnter:
//		keypress.NumPadEnter = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPadEnter")
//		}
//		break
//	case hid.CmdKeypadDecimalSeparatorDelete:
//		keypress.NumPadDot = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "NumPadDot")
//		}
//		break
//
//	case hid.CmdApplication:
//		keypress.Application = state
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "Application")
//		}
//
//	default:
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "未能识别出的按键16进制：", key)
//		}
//		break
//	}
//}
//
//func (self *keyboardState) keyboard_function_reset(keypress *KeyPressed) {
//	keypress.LeftCtrl = false
//	keypress.LeftShift = false
//	keypress.LeftAlt = false
//	keypress.RightCtrl = false
//	keypress.RightShift = false
//	keypress.RightAlt = false
//	keypress.LeftWindows = false
//	keypress.RightWindows = false
//}
//
//func (self *keyboardState) keyboard_function(key uint8, keypress *KeyPressed) {
//	if key == hid.CmdNone && self.currentFunctionKey != hid.CmdNone {
//		self.currentFunctionKey = hid.CmdNone
//		self.keyboard_function_reset(keypress)
//		if keypress.KeyPressDebug {
//			common.PrintRedis("Keyboard", "松开所有功能键")
//		}
//	} else {
//		self.currentFunctionKey = key
//		self.keyboard_function_reset(keypress)
//		switch key {
//		case hid.CmdNone:
//			break
//		case hid.LeftCtrl:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了左 Ctrl 键")
//			} // 设置左 Ctrl 键按下的状态为 true
//			keypress.LeftCtrl = true
//			break
//		case hid.LeftShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了左 Shift 键")
//			} // 设置左 Shift 键按下的状态为 true
//			keypress.LeftShift = true
//			break
//		case hid.LeftCtrl + hid.LeftShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl 和左 Shift 键")
//			} // 设置左 Ctrl 和左 Shift 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.LeftShift = true
//			break
//		case hid.LeftCtrl + hid.LeftAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl 和左 Alt 键")
//			} // 设置左 Ctrl 和左 Alt 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.LeftAlt = true
//			break
//		case hid.LeftCtrl + hid.LeftAlt + hid.LeftShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl、左 Alt 和左 Shift 键")
//			} // 设置左 Ctrl、左 Alt 和左 Shift 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.LeftAlt = true
//			keypress.LeftShift = true
//			break
//		case hid.RightCtrl:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了右 Ctrl 键")
//			} // 设置右 Ctrl 键按下的状态为 true
//			keypress.RightCtrl = true
//			break
//		case hid.LeftAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了左 Alt 键")
//			} // 设置左 Alt 键按下的状态为 true
//			keypress.LeftAlt = true
//			break
//		case hid.LeftAlt + hid.LeftShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Alt 和左 Shift 键")
//			} // 设置左 Alt 和左 Shift 键同时按下的状态为 true
//			keypress.LeftAlt = true
//			keypress.LeftShift = true
//			break
//		case hid.RightCtrl + hid.RightShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了右 Ctrl 和右 Shift 键")
//			} // 设置右 Ctrl 和右 Shift 键同时按下的状态为 true
//			keypress.RightCtrl = true
//			keypress.RightShift = true
//			break
//		case hid.LeftCtrl + hid.RightShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl 和右 Shift 键")
//			} // 设置左 Ctrl 和右 Shift 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.RightShift = true
//			break
//		case hid.LeftCtrl + hid.RightAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl 和右 Alt 键")
//			} // 设置左 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.RightAlt = true
//			break
//		case hid.RightAlt + hid.RightShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了右 Alt 和右 Shift 键")
//			} // 设置右 Alt 和右 Shift 键同时按下的状态为 true
//			keypress.RightAlt = true
//			keypress.RightShift = true
//			break
//		case hid.LeftCtrl + hid.LeftAlt + hid.RightCtrl + hid.RightAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了左 Ctrl、左 Alt、右 Ctrl 和右 Alt 键")
//			} // 设置左 Ctrl、左 Alt、右 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.LeftCtrl = true
//			keypress.LeftAlt = true
//			keypress.RightCtrl = true
//			keypress.RightAlt = true
//			break
//		case hid.RightAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了右 Alt 键")
//			} // 设置右 Alt 键按下的状态为 true
//			keypress.RightAlt = true
//			break
//		case hid.RightShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "按下了右 Shift 键")
//			} // 设置右 Shift 键按下的状态为 true
//			keypress.RightShift = true
//			break
//		case hid.RightCtrl + hid.RightAlt:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了右 Ctrl 和右 Alt 键")
//			} // 设置右 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.RightCtrl = true
//			keypress.RightAlt = true
//			break
//		case hid.RightCtrl + hid.RightAlt + hid.RightShift:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "同时按下了右 Ctrl 和右 Alt 键 和右Shift")
//			} // 设置右 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.RightCtrl = true
//			keypress.RightAlt = true
//			keypress.RightShift = true
//			break
//
//		case hid.LeftWindows:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "左Windows")
//			} // 设置右 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.LeftWindows = true
//			break
//
//		case hid.RightWindows:
//			if keypress.KeyPressDebug {
//				common.PrintRedis("Keyboard", "右Windows")
//			} // 设置右 Ctrl 和右 Alt 键同时按下的状态为 true
//			keypress.RightWindows = true
//			break
//
//		default:
//			common.PrintRedis("Keyboard", "按下了未知的键", key)
//			break
//		}
//	}
//	self.waitGroup.Done()
//}
