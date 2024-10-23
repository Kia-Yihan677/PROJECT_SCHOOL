package main

import "fmt"
func main () {
	var a, b, c float64

	fmt.Println("masukkan nilai a, b, c : ")
	fmt.Scan(&a, &b, &c)

	if a+b>c && a+c>b && b+c>a {
		fmt.Println(a, b, "dan", c, "segitiga? true")
	} else {
		fmt.Println(a, b, "dan", c, "segitiga? false")
	}







}