package main

import (
        "fmt"
        "strings"
)




func main() {
    var first int
    var sec int
    var oper string
    fmt.Scan(&first, &oper, &sec)
        switch first{
             case first > 10:
                fmt.Println("Не больше 10")
                }
        switch sec{
            case sec > 10:
                fmt.Println("Не больше 10")
                }
        switch oper{
            case "+":
                fmt.Println(first+sec)
            
            case "*":
                fmt.Println(first*sec)    
                
            case "/":
                fmt.Println(first/sec)    
                
            
            case "-":
                fmt.Println(first-sec)    
                }    
    fmt.Println(&oper)
     
}
