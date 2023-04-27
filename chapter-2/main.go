package main

import (
	"fmt"
)

func belajarDefer() {
	// output -> dimas adalah programmer baru di sts
	fmt.Print("dimas ")				// 1
	defer fmt.Println("di sts")		// 5
	defer fmt.Print("baru ")		// 4
	defer fmt.Print("programmer ")	// 3
	fmt.Print("adalah ")			// 2
}

func main() {
	// 01 - defer
	belajarDefer()

	// 02 - pointer
	// pointer = penyimpan alamat ( *T )
	var alamatRumahDimas *string

	// & = "address-of" operator
	rumahDimas := "Grobogan"
	alamatRumahDimas = &rumahDimas

	// *pointerVariable = mendapatkan 
	// nilai dari alamat yang 
	// di simpan pointerVariable
	fmt.Println("Rumah dimas ada di", *alamatRumahDimas)

	// dimas pindah rumah
	*alamatRumahDimas = "Semarang"
	fmt.Println("Dimas sekarang pindah ke", *alamatRumahDimas)
	fmt.Println(rumahDimas)
}