package main

import "fmt"

func luasLingkaran(jari int) float32{
	var phi float32
	if (jari % 7 == 0){
		phi = 22.0 / 7.0
	}else {
		phi = 3.14
	}
	result := phi * float32(jari) * float32(jari)
	return result
}

func kelLingkaran(jari int) float32{
	var phi float32
	if (jari % 7 == 0){
		phi = 22.0 / 7.0
	}else{
		phi = 3.14
	}
	result := 2 * phi * float32(jari)
	return result
}

func main(){
	var r int
	fmt.Print("Masukkan Jari - Jari Lingkaran: ")
	fmt.Scan(&r)
	fmt.Println("----------------------------------------")
	fmt.Printf("Luas Lingkaran %.2f \n", luasLingkaran(r))
	fmt.Printf("Keliling Lingkaran %.2f \n", kelLingkaran(r))
	fmt.Println("----------------------------------------")
}