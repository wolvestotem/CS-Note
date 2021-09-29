// ex5.10 sorts courses topologically based on hard-coded dependencies.
package main

import (
	"fmt"
	"time"
)

// prereqs maps computer science courses to their prerequisites.
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

func main() {
	// for i, course := range topoSort(prereqs) {
	// 	fmt.Printf("%d:\t%s\n", i+1, course)
	// }
	fmt.Printf("time: %v", int32(time.Now().AddDate(0, 0, -1).Unix()))
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)

	visitAll = func(items []string) {
		for _, v := range items {
			if !seen[v] {
				seen[v] = true
				visitAll(m[v])
				order = append(order, v)
			}
		}
	}

	for k := range m {
		visitAll([]string{k})
	}
	return order
}
