package main

import (
	"fmt"
)

func main() {
	var arr [6]int
	for i := 0; i < 6; i++ {
		fmt.Printf("Insira o dígito %d do RA: ", i+1)
		fmt.Scan(&arr[i])

		fmt.Println(arr[i])
	}
	fmt.Println("R.A.: ", arr)
}
