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

// 04 - variable
var c bool
var python string

	// variable dengan inisialisasi
	var a, b int = 1, 2


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

	// 04 - variable
	var i int
	fmt.Println(i, c, python)

		// variable dengan inisialisasi
		var x, y, z = true, false, "no!"
		fmt.Println(a,b, x, y, z)
}

