package main

import "fmt"

func main() {
	fmt.Print("Введите любое число")
	var first string
	fmt.Scanln(&first)
	fmt.Print(first)
}
