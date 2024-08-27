package main
import (
	"fmt"
	
)
var romanMap = []struct {
	decVal int
	symbol string
}{
    {1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}
func decimalToRomanRecursive(num int) string {
	if num == 0 {
    	return ""
	}
	for _, pair := range romanMap {
    	if num >= pair.decVal {
            return pair.symbol + decimalToRomanRecursive(num-pair.decVal)
        }
    }
	return ""
}

func decimalTooRomanRecursive(num2 int) string {
	if num2 == 0 {
    	return ""
	}
	for _, pair := range romanMap {
    	if num2 >= pair.decVal {
            return pair.symbol + decimalToRomanRecursive(num-pair.decVal)
        }
    }
	return ""
}


func main() {
   
    var oper string
    fmt.Scan(&num, &oper, &num2)
        switch first{
             case first > 10:
                fmt.Println("Не больше 10")
                }
        switch sec{
                case sec > 10
                        fmt.Println("Не больше 10")
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
