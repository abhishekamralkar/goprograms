package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dectobase "github.com/abhishekamralkar/dectobase"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a decimal number: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid input for number")
	}

	b, _ := reader.ReadString('\n')
	b = strings.TrimSpace(b)

	base, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Invalid input for base")
	}
	fmt.Println(dectobase.ConvertToBase(number, base))

}
