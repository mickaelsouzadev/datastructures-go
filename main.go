package main

import (
	queue "datastructures/queue"
	stack "datastructures/stack"
	"fmt"
)

func main() {
	stack := stack.InitStack()
	stack.Push("Krakatoa")
	stack.Push("Vesuvius")
	stack.Push("Kilauea")
	fmt.Println("Pilha 1: ", stack.Peek())

	stack.Pop()
	stack.Pop()
	fmt.Println("Pilha 2", stack.Peek())

	queue := queue.InitQueue()
	queue.Enqueue("Krakatoa")
	queue.Enqueue("Vesuvius")
	queue.Enqueue("Kilauea")
	fmt.Println("Fila 1: ", queue.Peek())

	queue.Dequeue()
	fmt.Println("Fila 2: ", queue.Peek())
	queue.Dequeue()
	fmt.Println("Fila 3: ", queue.Peek())
	queue.Dequeue()
	fmt.Println("Fila 4: ", queue.Peek())
}
