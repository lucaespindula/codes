package main

import "fmt"

// Estrutura do nó
type Node struct {
	data string
	next *Node
}

// Estrutura da lista encadeada
type LinkedList struct {
	head *Node
}

// Adiciona um elemento no final da lista
func (l *LinkedList) Append(data string) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Remove o primeiro elemento da lista
func (l *LinkedList) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

// Exibe os elementos da lista
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
}

func main() {
	list := LinkedList{}
	list.Append("Fazer bet")
	list.Append("ir no arq")
	list.Append("comprar ingresso pro show do embaixador")
	fmt.Println("Lista Encadeada Simples:")
	list.Print()
	list.RemoveFirst()
	fmt.Println("Após remover o primeiro:")
	list.Print()
}
