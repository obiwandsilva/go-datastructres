package main

import (
	"fmt"
	"go-datastructures/linkedqueue"
	"strings"
)

func algorithm(name string, handler func()) {
	header := fmt.Sprintf("########## %s ##########", name)

	fmt.Println(header)
	handler()

	var fotter strings.Builder

	for i := 0; i < len(header); i++ {
		fotter.WriteString("#")
	}

	fmt.Println(fotter.String())
}

func workWithLinkedQueue() {
	queue := linkedqueue.NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	queue.Enqueue("E")
	queue.Enqueue("F")
	queue.Enqueue("G")
	queue.Display()
	queue.Reverse()
	queue.Display()
	queue.Dequeue()
	queue.Dequeue()
	queue.Display()
	queue.Reverse()
	queue.Display()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	// Last queue item
	if val, err := queue.Dequeue(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	// Empty queue
	if val, err := queue.Dequeue(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	queue.Display()
}

func main() {
	algorithm("Linked Queue", workWithLinkedQueue)
}
