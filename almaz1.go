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
func main(){
 var k string = "V"
 if val, ok := m[k]; ok{
	 fmt.Println(val)
}
