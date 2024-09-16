package main

import "fmt"

func main() {

	var number1, number2 float64
	var operator string

	for {
		fmt.Print("Input Operator (-, +, *, /): ")
		fmt.Scanln(&operator)

		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			break
		} else {
			fmt.Println("Komputer tidak mengerti input user, tolong masukkan operator yang benar!")
		}
	}

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&number2)

	switch operator {

	case "+":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Printf("%f %s %f = %f", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("Invalid Operator")

	}
}
