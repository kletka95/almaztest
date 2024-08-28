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
   "I",
   "II",
   "III",
   "IV",
   "V",
  "VI",
  "VII",
   "VIII",
   "IX",
   "X",
  }



func romanToArabic(roman string) (int) {
 if value, exists := romanArabic[roman]; exists {
  return value
 }
  return fmt.Println("Введенного римского числа не существует.")
}


func arabicToRoman(arabic int) string {
 if arabic <= 0 {
  return fmt.Println("римские числа должны быть больше нуля")
 }
 if arabic > 100 {
  return fmt.Println("Значение не может быть больше ста")
 }
  return arabic
}


func main() {
  var roman string
  roman = "C"
  romanToArabic(roman)
  arabicToRoman(11)
  arabicToRoman(5)
}
