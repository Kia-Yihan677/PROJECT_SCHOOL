package main

import "fmt"
func main () {
	var n, i, j int
	var star string

	fmt.Println("jumlah baris yang diinginkan : ")
	fmt.Scan(&n)

	for i = 1; i<= n; i++ {
		star = " "
		for j = 1; j <= i; j++ {
			star = star + "*"
		}
		fmt.Println(star)
	}

}
	