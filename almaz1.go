package main

import "fmt"

 var m = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
 }
 
 var m2 =  map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	}

func romToArab(rom string) (int) {
	if val, ok := m[rom]; ok{
	    return val
	}
	return 0
}
func ArabicToRom(ara int) string{
     if val, ok := m2[ara]; ok{
	    return val
	     }
	return "Числа нет"
}

func calculate (a, b int, oper string)(int, string){
  
	switch oper {
	case  "+":
	    return a+b, ""
	case  "-":
	    return a-b, ""
	case  "*":
	    return a*b, ""
	case  "/":
	    return a+b, ""
		
	default:
	    return 0, "ss"
		}
}

func main() {
  var a = romToArab("VI")
  fmt.Println(a)
  var b = ArabicToRom(512)
  fmt.Println(b)
  var c = calculate(5, 6, "+")
  fmt.Println(c)
}
