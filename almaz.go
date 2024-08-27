package main

import (
        "fmt"
       
)




func main() {
    var first int
    var sec int
    var oper string
    var c bool
    fmt.Scan(&first, &oper, &sec)
        switch c{
            case first > 10:
                fmt.Println("ture")
            case sec > 10:
                fmt.Println("true too")

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
