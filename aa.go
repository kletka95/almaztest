package main
import "fmt"




//Даем два списка для перевода значений, один стринг интеджер, другой интеджер стринг
var romanArabic = map[string]int {
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

var arabicRoman = map[int]string {
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



func arabicToRoman(roman, string) (int) {
   for value := range(romanArabic){
    if roman == value{
      fmt.Println(value)}
    }

func main() {
  var roman string
  roman = "III"
  arabicToRoman(roman)
}
