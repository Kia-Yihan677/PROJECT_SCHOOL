package main

import "fmt"
func main(){
    var tahun_berapa int
	var tahun_apa bool
	fmt.Println("Tahun = ")
	fmt.Scan(&tahun_berapa)

	tahun_apa = tahun_berapa % 4 == 0 && tahun_berapa % 5 == 0 && tahun_berapa % 100 != 0
	fmt.Print(tahun_apa)

}