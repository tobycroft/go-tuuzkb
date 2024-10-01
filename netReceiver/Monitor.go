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
