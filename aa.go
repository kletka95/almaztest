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
 
  return 0, errors.New("Введенного римского числа не существует.")
  
}


func arabicToRoman(arabic int) (string, error) {
 if arabic <= 0 {
  
  return " ",  errors.New("римские числа должны быть больше нуля")
  
 }
 if arabic > 100 {
 
  return " ", errors.New("Значение не может быть больше ста")

 }
  return arabicRoman[arabic], nil
}

func calulate( a, b int, oper string) (int, error){
 switch oper {
 case: "+":
        return a + b, nil
 case: "-":
        return a - b, nil
 case: "*":
        return a * b, nil
 case: "/":
        if b == 0{
          return 0, errors.New("Деление на 0 невозможно")
          } else{
             return a / b, nil
                    }
default:
        0, errors.New("Низвестный оператор")

        }


func main() {
 
calculate(100, 100, "-")
}
