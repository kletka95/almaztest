
package main

import (
    "fmt"
)

func main() {
    a, b, operator := getInput()

    var result float64

    switch operator {
    case "+":
        result = add(a, b)
    case "-":
        result = subtract(a, b)
    case "*":
        result = multiply(a, b)
    case "/":
        result = divide(a, b)
    default:
        fmt.Println("Неизвестный оператор")
        return
    }

    fmt.Printf("Результат: %.2f\n", result)
}
