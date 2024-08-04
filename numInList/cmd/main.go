package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	findnumber "github.com/abhishekamralkar/numInList"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a single number: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	number, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid inout, please enter a valid number")
		return
	}

	fmt.Println("Enter a list of numbers separated by spaces: ")
	rawList, _ := reader.ReadString('\n')
	rawList = strings.TrimSpace(rawList)

	strList := strings.Split(rawList, " ")

	var list []int
	for _, strNum := range strList {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Invalid input, please enter valid integers")
			return
		}
		list = append(list, num)

	}

	if findnumber.FindNumber(list, number) {
		fmt.Printf("number %d exist in %v\n", number, list)
	} else {
		fmt.Printf("number %d does not exist in %v\n", number, list)
	}
}
