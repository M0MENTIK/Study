package main

import (
	"fmt"
	"sync"
)

// var n int = 0
// var n atomic.Int64
var slice []int

var mtx sync.Mutex

func increace(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		//n.Add(1)
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go increace(wg)
	go increace(wg)
	go increace(wg)
	go increace(wg)
	go increace(wg)
	wg.Wait()

	mtx.Lock()
	fmt.Println(len(slice))
	mtx.Unlock()
	//fmt.Println(n.Load())
}
