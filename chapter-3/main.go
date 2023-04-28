package main

import (
	"fmt"
	"time"
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

// 06 - stringer
func (h Human) String() string {
	return fmt.Sprint("Aku adalah human dengan nama ", h.Name, " dan berumur ", h.Age, " tahun")
}

// 07 - error
type MyError struct {
	Message string
}

func (err MyError) Error() string {
	return err.Message
}

func getError() error {
	return MyError{"Ini error yang saya buat sendiri"}
}

// 08 - go routine
func masakMie() {
	fmt.Println("Mulai memasak mie")
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Mie on process")
	}
	fmt.Println("Mie siap dihidangkan")
}

func masakRendang() {
	fmt.Println("Mulai masak rendang")
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Rendang on process")
	}
	fmt.Println("Rendang siap dihidangkan")
}

// 09 - channel
func pesanMie(dapur chan string) {
	fmt.Println("Mulai memasak mie")
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Mie on process")
	}
	dapur <- "Mie siap dihidangkan"
}

func pesanRendang(dapur chan string) {
	fmt.Println("Mulai masak rendang")
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Rendang on process")
	}
	dapur <- "Rendang siap dihidangkan"
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

	// 06 - stringer
	fmt.Println(human)

	// 07 - Error
	if err:= getError(); err!= nil {
		fmt.Println(err)
	}

	// 08 - go routine
	go masakRendang()
	masakMie()


	// 09 - channel
	dapur := make(chan string)
	go pesanRendang(dapur)
	go pesanMie(dapur)

	fmt.Println(<-dapur)
	fmt.Println(<-dapur)

		// channel dengan buffer
		dapur2 := make(chan string, 2)
		pesanRendang(dapur2)
		pesanMie(dapur2)
		pesanRendang(dapur2) // error karena dapur 2 hanya bisa melayani 2 pesanan

		fmt.Println(<-dapur2)
		fmt.Println(<-dapur2)
		fmt.Println(<-dapur2)
}