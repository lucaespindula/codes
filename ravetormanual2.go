package main

import (
	"fmt"
)

func main() {
	var arr [7]int
	arr[1] = 1
	arr[2] = 2
	arr[3] = 4
	arr[4] = 4
	arr[5] = 3
	arr[6] = 1
	fmt.Println(arr[1], arr[2], arr[3], arr[4], arr[5], arr[6])
}
