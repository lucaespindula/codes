package main

import (
	"fmt"
)

// Stack representa uma pilha de inteiros
type Stack struct {
	items []int
}

// Push adiciona um elemento ao topo da pilha
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o elemento do topo da pilha
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Retorna falso se a pilha estiver vazia
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Remove o último elemento
	return item, true
}

// Top retorna o elemento do topo sem remover
func (s *Stack) Top() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty verifica se a pilha está vazia
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)

	top, exists := s.Top()
	if exists {
		fmt.Println("Elemento no topo:", top)
	} else {
		fmt.Println("A pilha está vazia")
	}

	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println("Removido:", val)
	}
}
