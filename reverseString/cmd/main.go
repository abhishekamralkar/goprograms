package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abhishekamralkar/reverseString"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	fmt.Println(reverseString.ReverseIt(str))
}
