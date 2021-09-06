package main

import (
	"bufio"
	"os"
)

var (
	nums   []int = []int{1, 2, 3, 4, 5}
	inputs []string
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

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
	// findLinks 递归实现
	//ch5.Input(inputs)

	// findLinks2 函数值实现 前序、后序遍历
	//ch5.Input2(inputs)

	// extract 匿名函数+函数值实现
	//ch5.Extract(inputs[0])

	/*courselist := ch5.TopologicalSort(prereqs)
	for _, course := range courselist {
		fmt.Println(course)
	}*/

}
