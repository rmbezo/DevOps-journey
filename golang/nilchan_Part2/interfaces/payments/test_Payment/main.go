package main

import (
	"fmt"
	"v1/payments"
)

func main() {
	// --- СЦЕНАРИЙ 1: Работаем через Stripe ---
	fmt.Println("=== РАБОТАЕМ СО STRIPE ===")
	stripe := payments.Stripe{}
	module := payments.NewPaymentModule(stripe)

	// Делаем пару покупок
	id1 := module.Pay("Подписка на Netflix", 15)
	id2 := module.Pay("Новый монитор", 300)

	fmt.Printf("Созданы транзакции: %d и %d\n", id1, id2)

	// Проверяем информацию по одной покупке
	fmt.Println("Инфо по id1:", module.Info(id1))

	// Смотрим весь журнал (AllInfo)
	fmt.Println("Весь журнал модуля:", module.AllInfo())

	// --- СЦЕНАРИЙ 2: Отмена заказа ---
	fmt.Println("\n=== ОТМЕНА ЗАКАЗА ===")
	module.Cancel(id1)
	fmt.Println("Журнал после отмены id1:", module.AllInfo())
	fmt.Println("Пытаемся узнать инфо об отмененном id1:", module.Info(id1))

	// --- СЦЕНАРИЙ 3: Подменяем мастера на лету (Крипта) ---
	fmt.Println("\n=== ПЕРЕКЛЮЧАЕМСЯ НА КРИПТУ ===")
	// Мы создаем НОВЫЙ модуль с криптой, либо можно было бы
	// добавить метод SetMethod в сам модуль. Но создадим новый для чистоты.
	crypto := payments.Crypto{}
	cryptoModule := payments.NewPaymentModule(crypto)

	id3 := cryptoModule.Pay("Покупка видеокарты", 1000)

	fmt.Println("Инфо по крипто-платежу:", cryptoModule.Info(id3))

	// Попробуем отменить крипту (помним, что наш Иваныч-Crypto это не любит)
	cryptoModule.Cancel(id3)
}
