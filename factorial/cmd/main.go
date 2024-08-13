package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhishekamralkar/factorial"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a number:")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid input for number")
	}

	fmt.Println(factorial.Fac(number))
}
