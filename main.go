package main

import "fmt"

func AboutMe() {
		fmt.Println("Enter your name: ")
		var name string
		fmt.Scanln(&name)
		fmt.Println("Enter your age: ")
		var age int64
		fmt.Scanln(&age)
	fmt.Printf("Hello my name is %v and my age is %v\n", name, age)
}

func main() {
	calc()
}