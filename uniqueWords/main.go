package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type data struct {
		k string
		v int
	}

	var s []data

	for k, v := range words {
		s = append(s, data{k, v})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].v > s[j].v
	})

	for _, s := range s {
		fmt.Println(s.k, "appears", s.v, "times")
	}

}
