package main

import "fmt"
func main () { 
	var n1, n2, n3, n4, n5, digit, total, total2, digit1, digit2 int

	fmt.Println("masukkan 5 digit angka : ")
    fmt.Scan(&digit)
	

	n1 = digit / 10000
	n2 = (digit % 10000) / 1000
	n3 = (digit % 1000) / 100
	n4 = (digit % 100) / 10
	n5 = digit % 10

	total = n1 + n2 + n3 + n4 + n5

	digit1 = total / 10
	digit2 = total % 10

	total2 = digit1 + digit2

	fmt.Println(total2)


}