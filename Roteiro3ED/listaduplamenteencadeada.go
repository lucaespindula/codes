package main

import "fmt"

// Estrutura do nó
type Node struct {
	data int
	next *Node
}

// Estrutura da lista encadeada
type LinkedList struct {
	head *Node
}

// Adiciona um elemento no final da lista
func (l *LinkedList) Append(data int) {
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
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("Lista Encadeada Simples:")
	list.Print()
	list.RemoveFirst()
	fmt.Println("Após remover o primeiro:")
	list.Print()
}
