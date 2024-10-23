package main

import "fmt"
func main () {
	var a, b, angka int

	fmt.Print("input nilai a dan b : ")
	fmt.Scan(&a, &b)

	for i:=b; i<=a; i++{
		fmt.Print(i, " ")
		angka += i
	}
	fmt.Print("\nTotal = ", angka)

}