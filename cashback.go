package main

import "fmt"
func main () {
	var digit_pertama, kode, digit_terakhir int
	var keluaran bool

	fmt.Print("masukkan 4 digit kode : ")
	fmt.Scan(&kode)

	digit_pertama  = (kode / 1000)
	digit_terakhir =  (kode % 10)

	keluaran = digit_pertama == digit_terakhir
	fmt.Print(keluaran)

}