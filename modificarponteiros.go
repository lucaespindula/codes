package main

import "fmt"

func main() {
	var b float64 = 1
	var p2 *float64 = &b
	*p2 = 77
	fmt.Println("Valor de b após a modificação:", b)
}
