package main

import (
	"fmt"
)

func main() {
	var arr [6]int
	for i := range arr {
		fmt.Printf("Insira o dígito %d do RA: ", i+1)
		fmt.Scan(&arr[i])

		fmt.Println(arr[i])
	}
	fmt.Println("R.A.: ", arr)
}
