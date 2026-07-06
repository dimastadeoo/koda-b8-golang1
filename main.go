package main

import "fmt"

func luasLingkaran(jari int) float32{
	var phi float32
	if (jari % 7 == 0){
		phi = 22/7
	}else {
		phi = 3.14
	}
	result := phi * float32(jari) * float32(jari)
	return result
}

func kelLingkaran(jari int) float32{
	var phi float32
	if (jari % 7 == 0){
		phi = 22/7
	}else{
		phi = 3.14
	}
	result := 2 * phi * float32(jari)
	return result
}

func main(){
	r := 30
	fmt.Printf("Lingkaran dengan Jari - Jari %d: \n", r)
	fmt.Printf("Luas Lingkaran %f \n", luasLingkaran(r))
	fmt.Printf("Keliling Lingkaran %f \n", kelLingkaran(r))
}