package main

import (
	"fmt"
)

func main() {
	// 01 - defer
	// output -> dimas adalah programmer baru di sts
	fmt.Print("dimas ")				// 1
	defer fmt.Println("di sts")		// 5
	defer fmt.Print("baru ")		// 4
	defer fmt.Print("programmer ")	// 3
	fmt.Print("adalah ")			// 2

}