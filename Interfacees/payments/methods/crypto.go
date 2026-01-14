package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd float64) int {
	fmt.Println("Payment crypto")
	fmt.Println("Size payment: ", usd, "USDT")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Cancel crypto opeation ID:", id)
}
