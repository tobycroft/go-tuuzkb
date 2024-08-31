package main

import (
	"fmt"
	"log"

	gowinkey "github.com/pizixi/KeyBoardListen"
)

func main() {

	events := gowinkey.Listen()
	for e := range events {
		if e.State != 0 {
			// log.Println("触发:", e.VirtualKey)
			var pressKeys string
			for key := range e.PressedKeys {
				// fmt.Printf("Key: %s\n", key)
				pressKeys += fmt.Sprintf("[%v] ", key)
			}
			if e.PressedKeys.ContainsAll(gowinkey.VK_LMENU, gowinkey.VK_Z) {
				log.Println("快捷键alt+z被触发:", pressKeys)
				// ...
			}
		}
	}

}
