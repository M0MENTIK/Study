package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	number int,
	power int,
) {

	defer wg.Done()
	for {
		fmt.Printf("I'm miner №%d. Start mining coal\n", number)

		select {
		case <-ctx.Done():
			fmt.Printf("I'm miner №%d. My working day is finish\n", number)
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("I'm miner №%d. Mine coal\n", number)
		}

		select {
		case <-ctx.Done():
			fmt.Printf("I'm miner №%d. My working day is finish\n", number)
			return
		case transferPoint <- power:
			fmt.Printf("I'm miner №%d. Transferred %d coal\n", number, power)
		}

		//

		//

		// select {
		// case <-ctx.Done():
		// 	fmt.Printf("I'm miner №%d. My working day is finish\n", number)
		// 	return
		// default:
		// 	fmt.Printf("I'm miner №%d. Start mining coal\n", number)
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Printf("I'm miner №%d. Mine coal\n", number)

		// 	transferPoint <- power
		// 	fmt.Printf("I'm miner №%d. Transferred %d coal\n", number, power)
		// }

	}

}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
