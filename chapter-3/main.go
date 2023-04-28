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

// 04 - interface, interface kosong
type Mamalia interface {
	Menyusui()
}

func (h Human) Menyusui() {
	fmt.Println("human", h.Name, "menyusui")
}

type Animal struct {
	Name string
	Age int
}

func (a *Animal) Menyusui() {
	fmt.Println("animal", a.Name, "menyusui")
}


// 05 - penegasan tipe
func checkType(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println(x, "adalah int")
	case string:
		fmt.Println(x, "adalah string")
	default:
		fmt.Println("tipe", t, "tidak diketahui")
	}

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

	// 04 - interface / interface kosong
	var mamalia Mamalia
	myMom := Human{"Mommy", 40}
	mamalia = myMom
	mamalia.Menyusui()
	mamalia = &Animal{"Heli", 3}
	mamalia.Menyusui()

	// 05 - penegasan tipe
	var universal interface{} = "world"

	s := universal.(string)
	fmt.Println(s)

	s, canCasting := universal.(string)
	if canCasting {
		fmt.Println(s, "adalah string")
	}

	j, canCasting := universal.(int)
	if (!canCasting) {
		fmt.Println(universal, "bukan int")
		fmt.Println(j)
	}

		// switch dengan tipe
		checkType(10)
		checkType("hello")
		checkType(human)

}