package main

import "fmt"
func main () {
	var uang, bunga float64

	fmt.Print("modal yang dimiliki : ")
	fmt.Scan(&uang)

	bunga = 0.05


	for tahun:=1; tahun <= 10; tahun++ {
		uang += uang * bunga
		fmt.Printf("Tahun ke-%d : Rp%2f/n", tahun, uang)
	}
}