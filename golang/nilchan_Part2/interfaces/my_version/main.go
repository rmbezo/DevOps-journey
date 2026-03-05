package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Interface
type Product interface {
	Information()
	Buy(num int)
	ChangePrice(price int)
}

func buyingSimulation(product Product) {

	for {
		fmt.Printf("Покупаем:")
		fmt.Println(" ")
		product.Information()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println(" ")
		fmt.Println("Вы хотите попробывать торгануться?")
		fmt.Print("Да/Нет: ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		if len(input) > 1 {
			fmt.Println("discarding process.")
			return
		}
		cmd := strings.ToLower(input[0])
		switch cmd {
		case "yes", "да":
			fmt.Print("Введите цену: ")
			var price int
			fmt.Scan(&price)
			product.ChangePrice(price)
		case "no", "нет":
			fmt.Print("Введите количество: ")
			scanner.Scan()
			countStr := scanner.Text()
			count, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Println("Error!")
				continue
			}
			product.Buy(count)
			continue
		default:
			fmt.Println("Не понял вас.")
			fmt.Println(" ")
			continue
		}
		fmt.Print("Введите количество: ")
		scanner.Scan()
		countStr := scanner.Text()
		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("Error!")
			continue
		}
		product.Buy(count)
		return
	}

}

// Pomidor logics
type Pomidor struct {
	Price int
	Info  string
}

func (p *Pomidor) Information() {
	fmt.Println(p.Info)
	fmt.Printf("Цена за киллограмм: %v", p.Price)
}

func (p *Pomidor) Buy(num int) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		price := p.Price * num
		fmt.Printf("Стоимость будет: %v\n", price)
		fmt.Print("Вы готовы к покупке?(Да/Нет или Yes/No):")
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		if len(input) > 1 {
			fmt.Println("discarding process.")
			return
		}
		cmd := strings.ToLower(input[0])
		switch cmd {
		case "yes", "да", "y", "д", "ye":
			fmt.Printf("Купили %v кг помидоров.\n", num)
			fmt.Println(" ")
			return
		case "no", "нет", "n", "не", "н":
			fmt.Println("Досвидания!")
			return
		default:
			fmt.Println("Некорректный ввод. Повторите попытку.")
		}
	}
}

func (p *Pomidor) ChangePrice(price int) {
	if price < p.Price || p.Price/2 < price {
		fmt.Println("Ну ничего себе. Вы захотели скидочку?")
		if price+5 < p.Price && price > p.Price/2 {
			torg := price + 5
			fmt.Printf("Ну ладно тебе как родному за %v отдам\n", torg)
			p.Price = torg
		} else {
			fmt.Println("Не брат давай без скидки, бери либо уходи")
			return
		}
	} else {
		fmt.Println("Не брат, много просишь. Бери либо уходи.")
		return
	}
}

//
//
//	 Broccoli logics
//
//

type Broccoli struct {
	Price int
	Info  string
}

func (b *Broccoli) Information() {
	fmt.Println(b.Info)
	fmt.Printf("Цена за киллограмм: %v", b.Price)
}

func (b *Broccoli) Buy(num int) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		price := b.Price * num
		fmt.Printf("Стоимость будет: %v\n", price)
		fmt.Print("Вы готовы к покупке?(Да/Нет или Yes/No):")
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		if len(input) > 1 {
			fmt.Println("discarding process.")
			return
		}
		cmd := strings.ToLower(input[0])
		switch cmd {
		case "yes", "да", "y", "д", "ye":
			fmt.Printf("Купили %v кг брокколи.\n", num)
			fmt.Println(" ")
			return
		case "no", "нет", "n", "не", "н":
			fmt.Println("Досвидания!")
			return
		default:
			fmt.Println("Некорректный ввод. Повторите попытку.")
		}
	}
}

func (b *Broccoli) ChangePrice(price int) {
	if price < b.Price || b.Price/2 < price {
		fmt.Println("Ну ничего себе. Вы захотели скидочку?")
		if price+15 < b.Price && price > b.Price/2 {
			torg := price + 15
			fmt.Printf("Ну ладно тебе как родному за %v отдам\n", torg)
			b.Price = torg
		} else {
			fmt.Println("Не брат давай без скидки, бери либо уходи")
			return
		}
	} else {
		fmt.Println("Не брат, много просишь. Бери либо уходи.")
		return
	}
}

//
// Main function
//

func main() {
	fmt.Println()

	pomidor := &Pomidor{
		Price: 50,
		Info:  "Помидор, это красная, вкусная и полезная ягода (признан овощем)",
	}
	broccoli := &Broccoli{
		Price: 250,
		Info:  "Брокколи, это основа любого здорового питания, зеленый и похож на маленькое дерево",
	}

	fmt.Println("Покупаем помидор. ")
	fmt.Println(" ")
	buyingSimulation(pomidor)
	fmt.Println("Покупаем брокколи. ")
	fmt.Println(" ")
	buyingSimulation(broccoli)
}
