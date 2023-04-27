package main

import (
	"fmt"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	// 01 - hello world
	fmt.Println("Hello World!")

	// 02 - package, import
	fmt.Println("Bilangan kesukaan saya adalah", rand.Intn(10))

	// 03 - function
	fmt.Println(add(25, 31))

}

