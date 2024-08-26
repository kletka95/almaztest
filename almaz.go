package main

import "fmt"



func main() {
    var first int
    var sec int
    var operator string
    fmt.Scan(&first, &operator, &sec)
    
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
