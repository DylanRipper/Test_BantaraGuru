package main

import "fmt"

func main() {
	hasil := 1
	var n int
	fmt.Print("Input Angka : ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		hasil = hasil * i
	}
	fmt.Print("hasil: ", hasil)
}
