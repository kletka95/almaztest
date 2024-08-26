package main

import "fmt"

func main() {
	fmt.Print("Введите любое число")
	var first int
	var second int
	var thrid int
	var ss = fmt.Scanln(&first, &thrid, &second)
	 
	fmt.Print(ss)
}
