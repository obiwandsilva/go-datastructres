package linkedqueue

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	value string
	next  *node
}

// Queue ...
type Queue struct {
	length int
	head   *node
	tail   *node
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{0, nil, nil}
}

// Enqueue ...
func (q *Queue) Enqueue(value string) string {
	defer q.incLength()
	node := &node{value, nil}

	if q.isEmpty() {
		q.head = node
		q.tail = node

		return value
	}

	if q.length == 1 {
		q.head.next = node
		q.tail = node

		return value
	}

	q.tail.next = node
	q.tail = node

	return value
}

// Dequeue ...
func (q *Queue) Dequeue() (string, error) {
	if q.isEmpty() {
		return "", errors.New("Empty queue")
	}

	defer q.decLength()
	value := q.head.value

	if q.length == 1 {
		q.head = nil
		q.tail = nil

		return value, nil
	}

	q.head = q.head.next

	return value, nil
}

// Display ...
func (q *Queue) Display() {
	if q.isEmpty() {
		fmt.Println("()")
		return
	}

	var builder strings.Builder
	head := q.head

	builder.WriteString(fmt.Sprintf("(%s)", head.value))

	for head = head.next; head != nil; head = head.next {
		builder.WriteString(fmt.Sprintf("->(%s)", head.value))
	}

	fmt.Println(builder.String())
}

// Reverse ...
func (q *Queue) Reverse() {
	if q.isEmpty() || q.length == 1 {
		return
	}

	newTail := reverse(q.head)
	q.head = q.tail
	q.tail = newTail
}

func (q *Queue) incLength() {
	q.length++
}

func (q *Queue) decLength() {
	q.length--
}

func (q *Queue) isEmpty() bool {
	if q.length > 0 {
		return false
	}

	return true
}

func reverse(head *node) *node {
	if head.next == nil {
		return head
	}

	newHead := reverse(head.next)
	newHead.next = head
	head.next = nil

	return head
}
