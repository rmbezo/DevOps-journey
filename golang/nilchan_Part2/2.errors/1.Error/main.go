package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance < usd {
		return errors.New("Недостаточно средств")
	}

	user.Balance -= usd
	return nil

	// Сетевой запрос
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-20 > 0 {
		speed := rand.IntN(150)
		fmt.Println("Нажимаю на газ.")
		c.Armor -= 20
		return speed, nil
	}
	return 0, errors.New("Машина в критическом состояние, требуется обслуживание.\n")
}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника!", p)
		}
	}()
	// PANIC - problem in code
	a := 0
	b := 1 / a
	fmt.Println("b: ", b)

	car := Car{
		Armor: 100,
	}

	for {
		kmch, err := car.Gas()
		if err != nil {
			fmt.Println(err.Error())
			break
		} else {
			fmt.Printf("Скорость: %v, Состояние: %v \n", kmch, car.Armor)
			fmt.Println(" ")
		}
	}

	//
	// Practice before
	//
	// for i := 0; i < 11; i++ {
	// 	fmt.Println("Моя скорость: ", car.Gas())
	// 	fmt.Println("Моя прочность: ", car.Armor)
	// }

	// user := &User{
	// 	Name:    "Oleg",
	// 	Balance: 50,
	// }

	// randomPay := rand.IntN(100)
	// err := Pay(user, randomPay)
	// if err != nil {
	// 	fmt.Println("Ошибка!!!! Оплата не была произведена, причина: ", err.Error())
	// } else {
	// 	fmt.Println("Успешно!!!! Оплата была произведена")
	// }

	// fmt.Println("Before buying stuff:")
	// pp.Println(user)
	// fmt.Println(" ")
	// Pay(user, 6000000000)

	// fmt.Println("After buying stuff:")
	// pp.Println(user)
	// fmt.Println(" ")
}
