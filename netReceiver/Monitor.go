package netReceiver

import "sync"

func (self *Km) MonitorKeyboard() {
	for client_tx := range self.keyboardReport {
		self.KeyState.waitGroup = sync.WaitGroup{}
		self.KeyState.waitGroup.Add(12)
		keyPressed := &KeyPressed{KeyPressDebug: self.KeyState.KeyBoardDebug.MessagePress}
		keyStayPressed := &KeyPressed{KeyPressDebug: self.KeyState.KeyBoardDebug.MessagePressAndHold}
		go self.KeyState.keyboard_function(client_tx.Buttons, keyPressed)
		go self.KeyState.keyboard_function(client_tx.Buttons, keyStayPressed)
		go self.keyboard_state1(client_tx.Data[0], keyPressed, keyStayPressed)
		go self.keyboard_state2(client_tx.Data[1], keyPressed, keyStayPressed)
		go self.keyboard_state3(client_tx.Data[2], keyPressed, keyStayPressed)
		go self.keyboard_state4(client_tx.Data[3], keyPressed, keyStayPressed)
		go self.keyboard_state5(client_tx.Data[4], keyPressed, keyStayPressed)
		go self.keyboard_state6(client_tx.Data[5], keyPressed, keyStayPressed)
		go self.keyboard_state7(client_tx.Data[6], keyPressed, keyStayPressed)
		go self.keyboard_state8(client_tx.Data[7], keyPressed, keyStayPressed)
		go self.keyboard_state9(client_tx.Data[8], keyPressed, keyStayPressed)
		go self.keyboard_state10(client_tx.Data[9], keyPressed, keyStayPressed)
		self.KeyState.waitGroup.Wait()
		self.KeyChannel <- KeyAll{keyPressed, keyStayPressed}
	}
}
