package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd float64) int {
	fmt.Println("Payment card")
	fmt.Println("Size payment: ", usd, "USD")

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Cancel card opeation ID:", id)
}
