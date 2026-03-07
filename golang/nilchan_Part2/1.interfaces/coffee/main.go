package main

import (
	coffeemaker "coffee/coffee-maker"
	"fmt"
)

func main() {
	randBeans := 50 // Для теста

	// Cappucino
	capp := &coffeemaker.Cappucino{BeansCount: randBeans}
	fmt.Println("Beans before: ", capp.BeansCount)
	fmt.Println(" ")
	robot := coffeemaker.CoffeeRobot{
		Maker: capp,
	}
	robot.StartProcess()

	fmt.Println("Beans after: ", capp.BeansCount)
	fmt.Println(" ")

	// ESPRESSO
	esp := &coffeemaker.Espresso{BeansCount: randBeans}

	fmt.Println("Beans before: ", esp.BeansCount)
	fmt.Println(" ")

	robot.Maker = esp
	robot.StartProcess()
	robot.StartProcess()
	robot.StartProcess()
	robot.StartProcess()

	fmt.Println("Beans before: ", esp.BeansCount)
	fmt.Println(" ")
}
