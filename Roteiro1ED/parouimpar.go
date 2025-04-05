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
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)
	resultado := ParOuImpar(numero)
	fmt.Println("O número é:", resultado)
}
