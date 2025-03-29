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
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	resultado := SomaAte(numero)
	fmt.Println("A soma dos números até o número", numero, "é", resultado)
}
