package main

import "fmt"
func main() {
	var jam, menit, satuan_detik_pada_bumi, satuan_detik_pada_mars int

	fmt.Println("input dalam satuan detik : ")
	fmt.Scan(&satuan_detik_pada_bumi)

	jam = satuan_detik_pada_bumi / 4500
	menit = (satuan_detik_pada_bumi % 4500) / 75
	satuan_detik_pada_mars = satuan_detik_pada_bumi % 75

	fmt.Println("maka hasil konversinya adalah", jam, "jam", menit, "menit", "dan", satuan_detik_pada_mars, "detik di mars")
	





}
	 