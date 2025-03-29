package main

import "fmt"

func SomaAte(n int) int {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma
}

func main() {
	var n int
	fmt.Print("Insira um número: ")
	fmt.Scan(&n)

	resultado := SomaAte(n)
	fmt.Println("A soma de 1 até", n, "é:", resultado)
}
