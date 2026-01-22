package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println("I'm postman. I sent letter:", text, i, "times")
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman("News", wg)

	wg.Add(1)
	go postman("Play magazine", wg)

	wg.Add(1)
	go postman("Autochronic", wg)

	wg.Wait()

	fmt.Println("main finish")

}
