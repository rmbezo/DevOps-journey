package coffeemaker

import "fmt"

type CoffeeMaker interface {
	Brew()
	CheckBeans() bool // Проверят есть ли зерна

}

type Espresso struct {
	BeansCount int
}

func (e *Espresso) Brew() {
	fmt.Println("☕ Наливаю крепкий Эспрессо...")
	e.BeansCount -= 10
}

func (e *Espresso) CheckBeans() bool { // beans - кол-во бобов, count - количество кружек эспрессо
	return e.BeansCount >= 10
}

type Cappucino struct {
	BeansCount int
}

func (c *Cappucino) Brew() {
	fmt.Println(" Взбиваю молоко...")
	fmt.Println(" Наливаем коффе...")
	fmt.Println("☕ Перемешиваем")
	c.BeansCount -= 20
}

func (c *Cappucino) CheckBeans() bool { // beans - кол-во бобов, count - количество кружек эспрессо
	return c.BeansCount >= 20
}

// Robot

type CoffeeRobot struct {
	Maker CoffeeMaker
}

func (r *CoffeeRobot) StartProcess() {
	fmt.Println("Робот: Проверяю наличие зерен...")

	// Вызываем метод интерфейса
	if r.Maker.CheckBeans() {
		r.Maker.Brew()
		fmt.Println("Робот: Готово! Приятного аппетита.")
		fmt.Println(" ")
	} else {
		fmt.Println("Робот: Ошибка! Мало зерен, сходи в магазин.")
		fmt.Println(" ")
	}
}
