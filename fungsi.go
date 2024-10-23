package main

import "fmt"
func main(){
     var n1, n2, fx float64
	 
	fmt.Print("masukkan nilai x : ")
	fmt.Scan(&n1)
	 
	fmt.Print("masukkan nilai y : ")
	fmt.Scan(&n2)
	
	fx=5/(2*n2 + 7) + n1*n1 - 3*n2
	fmt.Println("Hasil perhitungan persamaan f(n1,n2) : ", fx)

	fx=5/(2*n1 + 7) + n2*n2 - 3*n1
	fmt.Println("Hasil perhitungan persamaan f(n2,n1) : ", fx)

}
	
	 