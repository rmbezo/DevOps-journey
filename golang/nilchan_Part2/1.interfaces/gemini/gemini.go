package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. ИНТЕРФЕЙС (Контракт остается прежним)
type Product interface {
	Information()
	Buy(num int)
	ChangePrice(targetPrice int)
}

// 2. БАЗОВАЯ СТРУКТУРА (Чтобы не копипастить!)
// Сюда мы выносим ВСЮ общую логику
type BaseProduct struct {
	Name  string
	Price int
	Info  string
}

func (bp *BaseProduct) Information() {
	fmt.Printf("\n🍅/🥦 Товар: %s\nИнфо: %s\nЦена за кг: %d руб.\n", bp.Name, bp.Info, bp.Price)
}

func (bp *BaseProduct) Buy(num int) {
	total := bp.Price * num
	fmt.Printf("🛒 Итого за %d кг: %d руб. Берем? (да/нет): ", num, total)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ans := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if ans == "да" || ans == "yes" || ans == "y" {
		fmt.Printf("✅ Успешно! Вы купили %s на сумму %d руб.\n", bp.Name, total)
	} else {
		fmt.Println("❌ Ну, на нет и суда нет.")
	}
}

// ИСПРАВЛЕННАЯ МЕХАНИКА ТОРГОВ
func (bp *BaseProduct) ChangePrice(target int) {
	fmt.Printf("🧐 Вы предлагаете %d руб. вместо %d руб...\n", target, bp.Price)

	minPrice := bp.Price / 2 // Не отдадим дешевле половины цены

	if target >= bp.Price {
		fmt.Println("🤑 О, ты щедрый! По рукам!")
		bp.Price = target
	} else if target < minPrice {
		fmt.Printf("😡 Ты что, издеваешься? Меньше %d не отдам! Уходи!\n", minPrice)
	} else {
		// Торгуемся: сойдемся посередине между ценой юзера и нашей
		dealPrice := (bp.Price + target) / 2
		fmt.Printf("🤝 Ладно, давай ни твое, ни мое: за %d руб. забирай!\n", dealPrice)
		bp.Price = dealPrice
	}
}

// 3. КОНКРЕТНЫЕ ОВОЩИ (Теперь они пустые, всё в базе!)
type Pomidor struct{ BaseProduct }
type Broccoli struct{ BaseProduct }

// 4. ГЛАВНЫЙ СИМУЛЯТОР (Стал чище)
func buyingSimulation(p Product) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		p.Information()
		fmt.Print("\nЧто делаем? (1 - Торговаться, 2 - Купить, 3 - Уйти): ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Твоя цена: ")
			scanner.Scan()
			price, _ := strconv.Atoi(scanner.Text())
			p.ChangePrice(price)
		case "2":
			fmt.Print("Сколько кг надо? ")
			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			p.Buy(num)
			return // Купили и вышли к следующему прилавку
		case "3":
			fmt.Println("🚶 Уходим к другому продавцу...")
			return
		default:
			fmt.Println("🤔 Не понял, говори четче.")
		}
	}
}

func main() {
	// Создаем объекты через адрес (&), чтобы менять цену
	p := &Pomidor{BaseProduct{Name: "Помидор", Price: 100, Info: "Красный, сочный"}}
	b := &Broccoli{BaseProduct{Name: "Брокколи", Price: 300, Info: "Зеленая и полезная"}}

	buyingSimulation(p)
	buyingSimulation(b)
}
