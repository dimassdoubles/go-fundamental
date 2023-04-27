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

// 03 - struct
type WebApp struct {
	backend string
	frontend string
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

	// 03 - struct
	fmt.Println(WebApp{"Golang", "Vue"})

		// mengakses bagian struct
		iClinicWeb := WebApp{"Laravel", "Angular Js"}
		fmt.Println("iClinic Web menggunakan backend", iClinicWeb.backend, ", dan frontend", iClinicWeb.frontend)

		// pointer struct
		pointeriClinic := &iClinicWeb
		// syntax (*pointeriClinic).backend sama dengan => pointeriClinic.backend
		fmt.Println((*pointeriClinic).backend)
		fmt.Println(pointeriClinic.backend) // prefer this
		pointeriClinic.backend = "Golang"
		fmt.Println(pointeriClinic.backend)

		// inisialisasi struct
		j := WebApp{"Laravel", "React"}		// j bertipe WebApp
		k := WebApp{backend: "Firebase"}	// frontend = ""
		l := WebApp{}						// backend = "", frontend = ""
		pointerWeb := &WebApp{}				// pointerWeb bertipe *WebApp
		fmt.Println(j, k, l, pointerWeb)

}