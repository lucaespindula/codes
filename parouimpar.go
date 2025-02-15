package main

import "fmt"

func ParOuImpar(n int) string {
	if n%2 == 0 {
		return "Par"
	}
	return "Ímpar"
}

func main() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	resultado := ParOuImpar(numero)
	fmt.Println("O número", numero, "é", resultado)
}
