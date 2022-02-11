package main

import (
	"fmt"
	"strconv"
)

func strip(n int) string {
	var before int = -1
	var output string
	inputStr := strconv.Itoa(n)
	for i := 0; i < len(inputStr); i++ {
		save := inputStr[i] - 48
		if save%2 == 0 && before%2 == 0 {
			output = output + "-"
		}
		output = output + string(inputStr[i])
		before = int(save)
	}
	return output
}

// func Genap(input int) {
// 	ubah := strconv.Itoa(input)
// 	for i := 0; i < len(ubah); i++ {
// 		str := string(ubah[i])
// 		if string(ubah[i]) == "2" || string(ubah[i]) == "4" ||
// 			string(ubah[i]) == "6" || string(ubah[i]) == "8" {
// 			str += "-"
// 		}
// 		fmt.Print(str)
// 	}
// }
func main() {
	fmt.Println(strip(24686428))
}
