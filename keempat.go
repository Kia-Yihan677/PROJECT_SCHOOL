package main

import "fmt"
func main(){
	var x, N int

	fmt.Println("masukkan 2 bilangan bulat positif : ")
	fmt.Scan(&x, &N)

	if x % N == 0 {
		fmt.Println("Bilangan x dan N = true")
	} else {
		fmt.Println("Bilangan x dan N = false")
	}

}
