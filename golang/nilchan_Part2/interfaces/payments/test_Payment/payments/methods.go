package payments

import (
	"fmt"
	"math/rand/v2"
)

// Stripe — конкретная реализация
type Stripe struct{}

func (s Stripe) Pay(usd int, desc string) int {
	fmt.Printf("[Stripe] Оплата: %s на сумму $%d\n", desc, usd)

	min := 1555555
	max := 9999999
	randomID := min + rand.IntN(max-min+1)
	return randomID // Возвращаем фейковый ID транзакции
}

func (s Stripe) Cancel(id int) {
	fmt.Printf("[Stripe] Транзакция %d отменена\n", id)
}

type Crypto struct{}

func (c Crypto) Pay(usd int, desc string) int {
	fmt.Printf("[Crypto] Перевод %d USD в BTC за '%s'\n", usd, desc)
	min := 2555555
	max := 9999999
	randomID := min + rand.IntN(max-min+1)
	return randomID // Свой ID для крипты
}

func (c Crypto) Cancel(id int) {
	fmt.Printf("[Crypto] Ошибка: в блокчейне нельзя отменить транзакцию %d!\n", id)
}
