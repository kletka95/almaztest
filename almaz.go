package main

import "fmt"

func main() {
	fmt.Print("Введите любое число")
	var first int
	var second int
	fmt.Scanln(&first, &second)
	
	fmt.Print(first)
}
