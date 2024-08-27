package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите что-то")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	
	
}
