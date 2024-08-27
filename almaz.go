package main

import (
        "fmt"
       
)




func main() {
    romeToArab := map[string]int {
                'I': 1,
                'V': 5,
                'X': 10,
                'L': 50,
                'C': 100,
                'D': 500,
                'M': 1000,
            
    }
    arabToRome := map[int]string{

            1: 'I',
            5: 'V',
            10: 'X',
            50: 'L',
            100: 'C'
    }
   res := romeToArab["I"] + romeToArab["II"]
   fmr.Println(romeToArab[res])
    var first int
    var sec int
    var oper string
    var c bool
    fmt.Scan(&first, &oper, &sec)
        switch c{
            case first < 10:
                fmt.Println("Stop")
            case sec < 10:
                fmt.Println("Stop")
             default: switch oper{
            case "+":
                fmt.Println(first+sec)
            
            case "*":
                fmt.Println(first*sec)    
                
            case "/":
                fmt.Println(first/sec)    
                
            
            case "-":
                
                fmt.Println(first-sec)    
                }    
                }
        

     
}
