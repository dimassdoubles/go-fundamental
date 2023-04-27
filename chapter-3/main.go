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

// 02 - method
type Human struct {
	Name string
	Age int
}

type Plant struct {
	Name string
	Age int
}

func (h Human) makan() {
	fmt.Println(h.Name, "telah makan")
}

// 03 - pointer receiver
func (h *Human) ulangTahun() {
	h.Age += 1
}

func main() {
	// 01 - fungsi closure
	fib := fibonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	// 02 - method
	human := Human{"dimas", 21}
	human.makan()

	// plant := Plant{"mawar", 2}
	// plant.makan() -> error

	// 03 - pointer receiver
	fmt.Println("umur before:", human.Age)
	human.ulangTahun()
	fmt.Println("umur after:", human.Age)

}