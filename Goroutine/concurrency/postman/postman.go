package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	number int,
	mail string,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("I'm postman №%d. My working day is finish\n", number)
			return
		default:
			fmt.Printf("I'm postman №%d. Take letter\n", number)
			time.Sleep(1 * time.Second)
			fmt.Printf("I'm postman №%d. Delivered the letter to the post office: %s\n", number, mail)

			transferPoint <- mail
			fmt.Printf("I'm postman №%d. Delivered the letter: %s\n", number, mail)
		}

	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}
	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Information",
		2: "News from parrents",
		3: "Tax information",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}
	return mail

}
