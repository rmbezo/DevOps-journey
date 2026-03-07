package payments

import "fmt"

type PaymentMethod interface {
	Pay(usd int, desc string) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
	history       map[int]string // Добавили хранилище: ID -> Описание
}

func NewPaymentModule(method PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: method,
		history:       make(map[int]string),
	}
}

func (p *PaymentModule) Pay(desc string, usd int) int {
	id := p.paymentMethod.Pay(usd, desc)
	p.history[id] = desc // Запоминаем операцию
	return id
}

func (p *PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)
	delete(p.history, id) // Удаляем из истории при отмене
}

func (p *PaymentModule) Info(id int) string {
	desc, ok := p.history[id]
	if !ok {
		return "Операция не найдена"
	}
	return fmt.Sprintf("ID: %d | Описание: %s", id, desc)
}

func (p *PaymentModule) AllInfo() map[int]string {
	return p.history
}
