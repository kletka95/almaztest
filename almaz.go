package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println(a...:"Введите что-то")
	text, _ :=  reader.ReadString(delim: '\n')
	text = strings.TrimSpace(text)
	toNumber, _ := strconv.Atoi(text)
	fmt.Println(toNumber+8)		
	
}
