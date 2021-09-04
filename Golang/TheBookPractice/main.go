package main

import (
	"bufio"
	"os"

	"go-thebook-practice/ch5"
)

var (
	nums   []int = []int{1, 2, 3, 4, 5}
	inputs []string
)

type test struct {
	name string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		inputs = append(inputs, scanner.Text())
	}
	//ch5.Input(inputs)
	ch5.Input2(inputs)
}
