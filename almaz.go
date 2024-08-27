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
             case  > 10:
                fmt.Println("Не больше 10")
                }
        switch sec{
            case  > 10:
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
