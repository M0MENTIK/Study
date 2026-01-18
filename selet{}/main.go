package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "friend 1",
				Text:   "Hi",
			}

			time.Sleep(10000 * time.Millisecond)

		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "friend 2",
				Text:   "How are you?",
			}

			time.Sleep(100 * time.Millisecond)

		}
	}()

	for {
		select {
		case msg1 := <-messageChan1:
			fmt.Println("I get message from: ", msg1.Author, "Text:", msg1.Text)
		case msg2 := <-messageChan2:
			fmt.Println("I get message from: ", msg2.Author, "Text:", msg2.Text)
		}

	}

}
