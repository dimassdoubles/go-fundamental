package main

import (
	"encoding/json"
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

// 04 - struct tag / meta
type Employee struct {
	EmployeeName string `json:"employee_name"`
	Age  int    `json:"age"`
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

	// 04 - struct tag / meta
	// struct to json
	dimaStruct := Employee{"Dimas Saputro", 21}
	dimasJson, err := json.Marshal(dimaStruct)
	if err != nil {
		fmt.Println("Gagal")
	} else {
		fmt.Println(string(dimasJson))
	}

		// json to struct
		// json harus menggunakan backticks `
		anggaJson := `{"employee_name": "Angga", "age": 18}`
		var anggaStruct Employee
		err = json.Unmarshal([]byte(anggaJson), &anggaStruct)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(anggaStruct)
		}
}