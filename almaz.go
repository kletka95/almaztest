package main

import (
 "errors"
 "fmt"
 "strconv"
 "strings"
)

var romanToArabicMap = map[string]int{
 "I":  1,
 "II": 2,
 "III": 3,
 "IV": 4,
 "V":  5,
 "VI": 6,
 "VII": 7,
 "VIII": 8,
 "IX": 9,
 "X": 10,
}

var arabicToRomanMap = []string{
 "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX",
 "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
}

func romanToArabic(roman string) (int, error) {
 if value, exists := romanToArabicMap[roman]; exists {
  return value, nil
 }
 return 0, errors.New("некорректное римское число")
}

func arabicToRoman(arabic int) (string, error) {
 if arabic <= 0 {
  return "", errors.New("римские числа должны быть больше нуля")
 }
 if arabic > 100 {
  return "", errors.New("значение результата не может быть представлено римскими цифрами")
 }
 return arabicToRomanMap[arabic], nil
}

func calculate(a, b int, operator string) (int, error) {
 switch operator {
 case "+":
  return a + b, nil
 case "-":
  return a - b, nil
 case "*":
  return a * b, nil
 case "/":
  if b == 0 {
   return 0, errors.New("деление на ноль невозможно")
  }
  return a / b, nil
 default:
  return 0, errors.New("неизвестная операция")
 }
}

func parseAndCalculate(input string) (string, error) {
 operators := []string{"+", "-", "*", "/"}
 var operator string

 for _, op := range operators {
  if strings.Contains(input, op) {
   operator = op
   break
  }
 }

 if operator == "" {
  return "", errors.New("оператор не найден")
 }

 parts := strings.Split(input, operator)
 if len(parts) != 2 {
  return "", errors.New("некорректный формат ввода")
 }

 operand1, operand2 := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

 // Проверка на римские или арабские числа
 isRoman := false
 if _, exists := romanToArabicMap[operand1]; exists {
  isRoman = true
 }
 if _, err := strconv.Atoi(operand1); err == nil && !isRoman {
  isRoman = false
 } else if err != nil && !isRoman {
  return "", errors.New("неверный формат числа")
 }

 if isRoman {
  a, err := romanToArabic(operand1)
  if err != nil {
   return "", err
  }
  b, err := romanToArabic(operand2)
  if err != nil {
   return "", err
  }

  result, err := calculate(a, b, operator)
  if err != nil {
   return "", err
  }

  if result <= 0 {
   return "", errors.New("результат операции римскими числами должен быть больше нуля")
  }

  romanResult, err := arabicToRoman(result)
  if err != nil {
   return "", err
  }

  return romanResult, nil
 } else {
  a, err := strconv.Atoi(operand1)
  if err != nil {
   return "", err
  }
  b, err := strconv.Atoi(operand2)
  if err != nil {
   return "", err
  }

  result, err := calculate(a, b, operator)
  if err != nil {
   return "", err
  }

  return strconv.Itoa(result), nil
 }
}

func main() {
 var input string
 fmt.Print("Введите выражение: ")
 fmt.Scanln(&input)

 result, err := parseAndCalculate(input)
 if err != nil {
  panic(err)
 }

 fmt.Println("Результат:", result)
}
