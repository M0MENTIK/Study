package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd float64) int {
	fmt.Println("Payment PayPaL:")
	fmt.Println("Size payment: ", usd, "usd")

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Cancel PaypPal opeation ID:", id)
}
