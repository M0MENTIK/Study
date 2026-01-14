package payments

type PaymentMethod interface {
	Pay(usd float64) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo map[int]PaymentInfo

	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принимает:
// 1. Метод проводимой оплаты
// 2. Сумма оплаты
// Возвращает:
// 1. Id операции
func (p PaymentModule) Pay(descroption string, usd float64) int {
	// 1. Proved payment
	// 2. give id payment

	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: descroption,
		Usd:         usd,
		Cancelled:   false,
	}

	// 3. save info about operation
	// - description operation
	// - How much spend
	// - This operation does cansel?
	p.paymentsInfo[id] = info

	// 4. return id payment
	return id

}

// Метод Cancel
// Принимает:
// 1. Id operation
// Возвращает:
// nothing
func (p PaymentModule) Cancel(id int) {

	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}
	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentsInfo[id] = info

}

// Метод Info
// Принимает:
// 1. ID operation
// Возвращает:
// 1. Info about for operation
func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод AllInfo()
// Принимает:
// Nothing
// Возвращает:
// 1. Info about all operations
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	return tempMap
}
