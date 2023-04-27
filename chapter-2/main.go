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

	// 05 - array, slice, looping dengan range
	var buahBuahan [3]string
	buahBuahan[0] = "melon"
	fmt.Println(buahBuahan[0])

	pegawaiSts := [5]string{"dimas", "agus", "angga", "brian", "ozi"}
	fmt.Println(pegawaiSts[0])
		
		// slice
		primes := [6]int{2, 3, 5, 7, 11, 13}

		var s []int = primes[1:4] // [3 5 7]
		fmt.Println(s)

		// mengubah elemen slice 
		// akan mengubah elemen arraynya
		s[0] = 100
		fmt.Println(primes) // [2 100 5 7 11 13]

		// inisialisasi slice
		sliceMakanan := []string{"martabak", "bebek goreng", "soto"}
		fmt.Println(sliceMakanan)

		// nilai default slice
		arrayMinuman := [3]string{"sirup", "kopi", "teh"}
		sliceMinuman1 := arrayMinuman[:]  // default [0:3], 3 adalah panjang arrayMinuman
		sliceMinuman2 := arrayMinuman[:3]
		sliceMinuman3 := arrayMinuman[0:]
		// sliceMinuman1 - 3 sama
		fmt.Println(sliceMinuman1, sliceMinuman2, sliceMinuman3)

		// panjang dan kapasitas
		sliceA := []int{1, 2, 3, 4}
		// kapasitas adalah panjang maksimal yang bisa dimiliki slice
		fmt.Println("panjang sliceA", len(sliceA), ",kapasitas sliceA", cap(sliceA))
		sliceA = sliceA[: 0]
		fmt.Println("panjang sliceA", len(sliceA), ",kapasitas sliceA", cap(sliceA))
		sliceA = sliceA[:2]
		fmt.Println("panjang sliceA", len(sliceA), ",kapasitas sliceA", cap(sliceA))

		// slice kosong
		var emptySlice []int
		fmt.Println("panjang emptySlice", len(emptySlice), ",kapasitas emptySlice", cap(emptySlice))
		if emptySlice == nil {
			fmt.Println("nil!")
		}

		// membuat slice dengan make
		mySlice := make([]int, 5)
		fmt.Println(len(mySlice))

		// nested slice
		board := [][]string{
					{"_", "_", "_"},
					{"_", "_", "_"},
					{"_", "_", "_"},
				}
		board[0][1] = "X"
		board[1][1] = "O"
		board[2][1] = "X"
		fmt.Println(board)

		// menambahkan elemen ke slice
		var sliceKu []int
		sliceKu = append(sliceKu, 1)
		fmt.Println(sliceKu)
		sliceKu = append(sliceKu, 2, 3, 4)
		fmt.Println(sliceKu)

		// looping dengan for
		sliceX := []string{"a", "b", "c"}
		for i, v := range(sliceX) {
			fmt.Println(i, v)
		}
			// jika perlu indeks saja
			for i, _ := range(sliceX) {
				fmt.Println(i)
			}

			for i := range(sliceX) {
				fmt.Println(i)
			}

			// jika perlu value saja
			for _, v := range(sliceX) {
				fmt.Println(v)
			}
	// 06 - map
	mapku := make(map[string]string)
	mapku["makanan"] = "bakso"
	fmt.Println(mapku["makanan"])
			
			// inisialisasi map
			map2 := map[string]string {
				"key1": "value1",
				"key2": "value2",
			}
			fmt.Println(map2)

			// operasi map
			// mengisi map
			map2["key3"] = "value3"
			// mengambil elemen map
			elemenke3 := map2["key3"]
			fmt.Println(elemenke3)
			// menghapus elemen map
			delete(map2, "key3")
			fmt.Println(map2)

			// menguji apakah key ada di dalam map
			value, exist := map2["key3"]
			if (exist) {
				fmt.Println(value)
			} else {
				fmt.Println("key tidak terdaftar")
			}

	// 07 - iterate map
	mymap := map[string]int {
		"f": 1,
		"e": 2,
		"d": 3,
		"b": 4,
		"c": 5,
		"a": 6,
	}
	for key, value := range mymap {
		fmt.Println(key, value)
	}			
}