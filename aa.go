package main
import (
        "fmt"
        "errors"
        )




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
  2:  "II",
  3: "III",
  4:  "IV",
  5: "V",
  6:  "VI",
  7:  "VII",
  8: "VIII",
  9:  "IX",
  10: "X",
  }



func romanToArabic(roman string) (int, error) {
 if value, exists := romanArabic[roman]; exists {
  return value, nil
 }
  err := errors.New("Введенного римского числа не существует.")
  return 0, err
  fmt.Print(err)
}


func arabicToRoman(arabic int) (string, error) {
 if arabic <= 0 {
  err := errors.New("римские числа должны быть больше нуля")
  return " ",  err
  fmt.Print(err)
 }
 if arabic > 100 {
  err := errors.New("Значение не может быть больше ста")
  return " ", err
  fmt.Print(err)
 }
  return arabicRoman[arabic], nil
}


func main() {
  var roman string
  roman = "C"
  romanToArabic(roman)
  arabicToRoman(1122)
  arabicToRoman(5)
  arabicToRoman(-1)
 if error != nil {
  fmt.Println("NOO")
 }

}
