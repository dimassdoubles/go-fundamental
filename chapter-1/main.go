package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// naked return, return dengan nama
func getParent(child string) (father, mother string) {
	father = child + " father"
	mother = child + " mother"

	return
}

func main() {
	// 01 - hello world
	fmt.Println("Hello World!")

	// 02 - package, import
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))

	// 03 - function
	fmt.Println(add(25, 31))

		// fungsi dengan banyak return
		fmt.Println(swap("hello", "world"))
		
		// naked return, return dengan nama
		fmt.Println(getParent("dimas"))
}

