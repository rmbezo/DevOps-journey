package main

import "fmt"

func main() {
	weather := map[int]int{
		10: +10,
		11: +9,
		12: +4,
		13: +6,
		14: +5,
		15: +9,
	}
	// var day int
	// fmt.Print("Day for forecast prediction: ")
	// fmt.Scan(&day)
	// c, ok := weather[day]
	// if ok {
	// 	fmt.Println(c)
	// } else if day > 31 || day <= 0 {
	// 	fmt.Println("Not valid day.")
	// } else {
	// 	fmt.Println("Doesn't have data for this day yet.")
	// }
	fmt.Println("before: ")
	fmt.Println(weather)
	AddData(weather)
	fmt.Println("after: ")
	fmt.Println(weather)
}

func AddData(m map[int]int) {
	var date, temperature int
	fmt.Print("Type the date: ")
	fmt.Scan(&date)
	fmt.Print("Type prediction: ")
	fmt.Scan(&temperature)

	fmt.Println("Updating statistics....")
	m[date] = temperature
	fmt.Println("Statistic was updated")
}
