package main

import "fmt"
func main () {
	var n, a, b, i int

	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		a = i
		b = n + 1 - i
		fmt.Println(a, "Telkom university", b)
	}

}