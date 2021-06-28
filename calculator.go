package main

import (
	"fmt"
	"strings"
)


func calc() {
	fmt.Println("What operation do you want to do: ")
	fmt.Println("Type 'add' or 'subtract' or 'divide' or 'multiply'")
	var Option string
	fmt.Scanln(&Option)
	var Input = strings.ToLower(Option)
	if Input == "add" || Input == "a" {
		fmt.Println("Enter first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Enter second number:")
		var num2 int
		fmt.Scanln(&num2)
		var sum int = num1 + num2
		fmt.Println("The sum is", sum)
	} else if Input == "subtract" || Input == "s" {
		fmt.Println("Enter first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Enter second number:")
		var num2 int
		fmt.Scanln(&num2)
		var diff int = num1 - num2
		fmt.Println("The difference is", diff)
	} else if Input == "multiply" || Input == "m" {
		fmt.Println("Enter first number:")
		var num1 int
		fmt.Scanln(&num1)
		fmt.Println("Enter second number:")
		var num2 int
		fmt.Scanln(&num2)
		var prod int = num1 * num2
		fmt.Println("The product is", prod)
	} else if Input == "divide" || Input == "d" {
		fmt.Println("Enter first number:")
		var num1 float64
		fmt.Scanln(&num1)
		fmt.Println("Enter second number:")
		var num2 float64
		fmt.Scanln(&num2)
		var quot float64 = num1 / num2
		fmt.Println("The quotient is", quot)
	}
}