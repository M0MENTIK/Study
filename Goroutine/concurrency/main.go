package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	initTime := time.Now()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("----------Work day for miners finishhed----------")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("----------Work day for postmans finishhed----------")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 3)
	mailTransferPoint := postman.PostmanPool(postmanContext, 3)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, m)
			mtx.Unlock()
		}

	}()

	wg.Wait()

	// isCoalClosed := false
	// isMailClosed := false

	// for !isCoalClosed || !isMailClosed {
	// 	select {
	// 	case c, ok := <-coalTransferPoint:
	// 		if !ok {
	// 			isCoalClosed = true
	// 			continue
	// 		}
	// 		coal += c
	// 	case m, ok := <-mailTransferPoint:
	// 		if !ok {
	// 			isMailClosed = true
	// 			continue
	// 		}
	// 		mails = append(mails, m)
	// 	}
	// }

	fmt.Println("Sum mine coal:", coal.Load())

	mtx.Lock()
	fmt.Println("Sum count get letters:", len(mails))
	mtx.Unlock()
	fmt.Println("Spendtime:", time.Since(initTime))
}
