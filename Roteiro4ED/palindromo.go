package main

import (
	"fmt"
)

// Definição da pilha
type Stack struct {
	items []rune
}

// Push adiciona um caractere ao topo da pilha
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o caractere do topo da pilha
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// Função para verificar se uma palavra é um palíndromo usando pilha
func isPalindrome(word string) bool {
	var stack Stack

	for _, char := range word {
		stack.Push(char)
	}

	reversed := ""
	for range word {
		char, _ := stack.Pop()
		reversed += string(char)
	}

	return word == reversed
}

func main() {
	// Testes
	words := []string{"radar", "arara", "golang", "level", "hello"}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}
