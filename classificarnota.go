package main

import "fmt"

func ClassificarNota(nota int) string {
	switch {
	case nota >= 9:
		return "Excelente"
	case nota >= 7:
		return "Bom"
	case nota >= 5:
		return "Regular"
	case nota >= 3:
		return "Ruim"
	case nota >= 0:
		return "Muito Ruim"
	default:
		return "Nota inválida"
	}
}

func main() {
	var nota int
	fmt.Print("Insira a nota (de 0 a 10): ")
	fmt.Scan(&nota)

	resultado := ClassificarNota(nota)
	fmt.Println("Classificação:", resultado)
}
