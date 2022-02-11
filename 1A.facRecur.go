package main

import "fmt"

func fac(n int) int {
	if n == 1 {
		return 1
	}
	return n * fac(n-1)
}

func main() {
	var n int
	fmt.Print("input Angka : ")
	fmt.Scan(&n)
	fmt.Println("hasil : ", fac(n))
}
