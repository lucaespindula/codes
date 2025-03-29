package main

import (
	"fmt"
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	p := Produto{
		Nome:       "Café Pelé",
		Preco:      50,
		Quantidade: 10,
	}

	fmt.Println(p.Nome)
	fmt.Printf("%f R$\n", p.Preco)
	fmt.Printf("Há %d unidades no estoque", p.Quantidade)
}
