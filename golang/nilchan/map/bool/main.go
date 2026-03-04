package main

import "fmt"

func main() {
	// Notebook - Is having 16 Gb RAM?
	notebook := make(map[string]bool, 8)

	notebook["Macbook Air M4 16/512"] = true
	notebook["Macbook Air M5 16/512"] = true
	notebook["Macbook Air M4 8/256"] = false
	notebook["Macbook Air M2 8/256"] = false
	notebook["Macbook Pro M3 16/512"] = true
	notebook["Macbook Pro M1 16/512"] = true

	// fmt.Println(notebook["Macbook Air M5 16/512"])

	//Rober

	Criminal := make(map[string]bool, 8)

	Criminal["Gosha"] = true
	Criminal["Jora"] = false
	Criminal["Dasha"] = true
	Criminal["Perplexity"] = false
	Criminal["Gemini"] = false
	Criminal["Gosha Vorovskoy"] = true

	var name string
	fmt.Println("FBI BASE")
	fmt.Print("Type name to know is he was arested or no: ")
	fmt.Scan(&name)
	c, ok := Criminal[name]
	if !ok {
		fmt.Println("Not founded in base")
		return
	}

	fmt.Println("This guy is founded in base!")
	if c {
		fmt.Println("He was arrested before")
	} else {
		fmt.Println("He is clear")
	}

}
