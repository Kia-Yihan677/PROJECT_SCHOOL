package main

import "fmt"
func main (){
	var n int
	var g, y, v, y0 float64
	g = 9.8
	y0 = 0

	fmt.Println("lama pengamatan dan kecepatan benda saat dilempar : ")
	fmt.Scan(&n, &v)

	for t := 0; t <= n; t++ {
		y = y0 + v*float64(t) + 0.5*g*float64(t)*float64(t)
		fmt.Printf("%.5f\n", y)
	}
		
	} 






