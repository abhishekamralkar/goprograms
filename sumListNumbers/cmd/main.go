package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	sumlistnumbers "github.com/abhishekamralkar/sumListNumbers"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a list of numbers separated by comma: ")
	rawList, _ := reader.ReadString('\n')
	rawList = strings.TrimSpace(rawList)

	strList := strings.Split(rawList, " ")

	var list []int
	for _, n := range strList {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Not a valid integers")
		}

		list = append(list, num)
	}

	fmt.Printf("Total for %v is: %d\n", list, sumlistnumbers.SumList(list))

}
