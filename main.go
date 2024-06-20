package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&op)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		fmt.Printf("%f + %f = %f\n", a, b, a+b)
	case "-":
		fmt.Printf("%f - %f = %f\n", a, b, a-b)
	case "*":
		fmt.Printf("%f * %f = %f\n", a, b, a*b)
	case "/":
		if b != 0 {
			fmt.Printf("%f / %f = %f\n", a, b, a/b)
		} else {
			fmt.Println("Division by zero!")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
