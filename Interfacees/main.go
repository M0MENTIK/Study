package main

import (
	"system_payment/payments"
	"system_payment/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("BURGER", 5)
	idPhone := paymentModule.Pay("PHONE", 500)
	idGame := paymentModule.Pay("Game", 20)

	paymentModule.Cancel(idPhone)

	pp.Println("All payments: ", paymentModule.AllInfo())
	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info:", gameInfo)
}
