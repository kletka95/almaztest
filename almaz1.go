package main
import (
 "errors"
 "fmt"
 "strings"
)

// Функция для перевода римских цифр в арабские
func romanToArabic(roman string) (int, error) {
 romanValues := map[byte]int{
  'I': 1,
  'V': 5,
  'X': 10,
  'L': 50,
  'C': 100,
  'D': 500,
  'M': 1000,
 }

 total := 0
 prevValue := 0

 for i := len(roman) - 1; i >= 0; i-- {
  value := romanValues[roman[i]]
  if value < prevValue {
   total -= value
  } else {
   total += value
  }
  prevValue = value
 }

 return total, nil
}

// Функция для перевода арабских цифр в римские
func arabicToRoman(arabic int) (string, error) {
 if arabic <= 0 || arabic > 3999 {
  return "", errors.New("value out of range")
 }

 romanNumerals := []struct {
  Value int
  Symbol string
 }{
  {1000, "M"},
  {900, "CM"},
  {500, "D"},
  {400, "CD"},
  {100, "C"},
  {90, "XC"},
  {50, "L"},
  {40, "XL"},
  {10, "X"},
  {9, "IX"},
  {5, "V"},
  {4, "IV"},
  {1, "I"},
 }

 var result strings.Builder
 for _, numeral := range romanNumerals {
  for arabic >= numeral.Value {
   result.WriteString(numeral.Symbol)
   arabic -= numeral.Value
  }
 }

 return result.String(), nil
}

// Функция для обработки выражений и выполнения операций
func calculate(input string) (string, error) {
 var a, b int
 var operator string

 // Попробуем обработать как арабские цифры
 _, err := fmt.Sscanf(input, "%d %s %d", &a, &operator, &b)
 if err != nil {
  // Если не удалось, попробуем как римские
  var romanA, romanB string
  _, err = fmt.Sscanf(input, "%s %s %s", &romanA, &operator, &romanB)
  if err != nil {
   return "", errors.New("invalid input")
  }

  a, err = romanToArabic(romanA)
  if err != nil {
   return "", err
  }
  b, err = romanToArabic(romanB)
  if err != nil {
   return "", err
  }
 }

 var result int
 switch operator {
 case "+":
  result = a + b
 case "-":
  result = a - b
 case "*":
  result = a * b
 case "/":
  if b == 0 {
   return "", errors.New("division by zero")
  }
  result = a / b
 default:
  return "", errors.New("unsupported operation")
 }

 if strings.ContainsAny(input, "IVXLCDM") {
  // Если входные данные содержали римские цифры, возвращаем результат в римских
  return arabicToRoman(result)
 }

 return fmt.Sprintf("%d", result), nil
}

func main() {
 input := ""
 fmt.Println("Ввести что-то")
 fmt.Scan(&input)
 result, err := calculate(input)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }
 fmt.Println(input, result)
}
