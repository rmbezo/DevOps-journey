package payments

type PaymentMethod interface {
	Pay(usd int, desc string) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
	}
}

// Принимает описание оплаты
// Возвращает ID проведенной операции
func (p *PaymentModule) Pay(desc string, usd int) int {

}

// Принимает ID
// Отменяет операцию
func (p PaymentModule) Cancel() {}

// Принимает ID
// Отдает ID + Description
func (p PaymentModule) Info() {}

// Ничего не принимает
// Возвращает всед
func (p PaymentModule) AllInfo() {}
