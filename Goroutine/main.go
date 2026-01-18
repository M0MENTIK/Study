package main

import (
	"fmt"
	"time"
)

// Function descripted a trip to the mine
// Input
// 1. Number a trip to the mine
// Return
// 1. main coal
func mine(transferPoint chan int, n int) {
	fmt.Println("A trip to the mine number:", n, "start")
	time.Sleep(1 * time.Second)
	fmt.Println("A trip to the mine number:", n, "finish")

	transferPoint <- 10
	fmt.Println("A trip number", n, "give coal")

}

func main() {
	coal := 0

	transferPoint := make(chan int, 2)

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Mine", coal, "coal!")
	fmt.Println("Spend time:", time.Since(initTime))

	// go foo()

	// go func() {
	// 	for {
	// 		fmt.Println("Anon")
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()

	// time.Sleep(1 * time.Second)
}

// func foo() {
// 	for {
// 		fmt.Println("Hello")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }
