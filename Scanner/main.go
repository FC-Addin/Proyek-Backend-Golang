package main

import "fmt"

func main(){
	var suhu float64
	fmt.Println("Masukan suhu")
	fmt.Scanf("%f", &suhu)
	// printf untuk mengganti formatingnya
	fmt.Printf("Suhu yang diinput adalah: %f\n", suhu)

	var uang uint64
	fmt.Println("Masukan uang")
	fmt.Scanf("%d", &uang)
	fmt.Printf("Uang yang diinput adalah: %d\n", uang)
}
