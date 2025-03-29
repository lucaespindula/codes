package main

import "fmt"

func ClassificarNota(nota int) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "Excelente"
	case nota >= 7 && nota <= 8:
		return "Bom"
	case nota >= 5 && nota <= 6:
		return "Regular"
	case nota >= 3 && nota <= 4:
		return "Ruim"
	case nota >= 0 && nota <= 2:
		return "Muito Ruim"
	default:
		return "Nota inválida"
	}
}

func main() {
	var nota int
	fmt.Print("Insira a nota (0 a 10): ")
	fmt.Scan(&nota)

	resultado := ClassificarNota(nota)
	fmt.Println("Classificação:", resultado)
}
