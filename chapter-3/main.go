package main

import (
	"fmt"
)

// 01 - fungsi closure
func fibonaci() func() int {
	next := 1
	prev := 0

	return func() int {
		result := prev
		prev, next = next, prev + next
		return result
	}
}

func main() {
	// 01 - fungsi closure
	fib := fibonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}