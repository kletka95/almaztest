package main

import (
        "fmt")


// this is a comment

func main() {
    var first int
    var sec int
    var oper string
    fmt.Scan(&first, &oper, &sec)
        switch oper {
                case "+":
                        fmt.Println(first+sec)
                case "*":
                        fmt.Println(first*sec)
                }
     
}
