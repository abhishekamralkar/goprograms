package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abhishekamralkar/fizzbuzz"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid number")
		return
	}

	fizzbuzz.FizzBuzz(number)

}
