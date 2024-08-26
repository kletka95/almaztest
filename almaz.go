package main

import "fmt"

func main() {
	fmt.Print("Введите любое число")
	var first int
	fmt.Scanln(&first)
	fmt.Print(first)
}
