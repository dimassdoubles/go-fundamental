package main

import (
	"fmt"
	"math/cmplx"
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

	// deklarasi variable singkat
	// country := "indonesia" -> error

// 05 - tipe dasar, konversi, inferensi tipe
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	sq      complex128 = cmplx.Sqrt(-5 + 12i)
)


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

		// deklarasi variable singkat
		fruit := "jambu"
		fmt.Println(fruit)
	
	// 05 - tipe dasar, konversi, inferensi tipe
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", sq, sq)

		// nilai kosong (nilai default tipe data)
		var i2 int 		// 0
		var f float64 	// 0
		var b bool		// false
		var s string	// ""
		fmt.Printf("%v %v %v %q\n", i2, f, b, s)

		// konversi tipe
		// var i3 int = 42
		// var f2 float64 = float64(i)
		// var u uint = uint(f)
		i3 := 42
		f2 := float64(i)
		u := uint(f)
		fmt.Println(i3, f2, u);

		// inferensi tipe
		m := 42           	// int
		f3 := 3.142        	// float64
		g := 0.867 + 0.5i 	// complex128
		fmt.Printf("m bertipe %T\n", m)
		fmt.Printf("f3 bertipe %T\n", f3)
		fmt.Printf("g bertipe %T\n", g)

	// 06 - Kontanta
	const shape = "Lingkaran"
	const pi = 3.14
	// const  := "test" -> error
	fmt.Println("pi untuk", shape, "adalah", pi)
}

