package main 

import "fmt"


func main() {
     var num1, num2 float64
     var operator string 

     fmt.Println("Type first number: ")
     fmt.Scan(&num1)
     fmt.Println("Type operator: ")
     fmt.Scan(&operator)
     fmt.Println("Type second number: ")
     fmt.Scan(&num2)

     var result float64
     switch operator {
	     case "+":
		     result = num1 + num2
	     case "-":
		     result = num1 - num2
	     case ":":
		     if num2 != 0{
		    	 result = num1 / num2 
		     } else {
		    		fmt.Println("Error, not possible to split on 0")
			    	fmt.Scanln()
						return
		      }	
	     case "*":
		     result = num1 * num2
			 default:
				 fmt.Println("Not right operator, retype the operator.")
				 fmt.Println("Press enter to close the program: ")
				 fmt.Scanln()
				 return
}
    fmt.Println("The result is: ", num1, operator, num2, "=", result)
    fmt.Println("Press Enter to end the program: ")
    fmt.Scanln()
}












