package ds

import (
	"bytes"
	"strconv"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func CreateLinkedList(data int) *LinkedList {
	return &LinkedList{ head: &Node{ data: data, next: nil }}
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{ data: data, next: nil }

	if ll.head == nil {
		ll.head = newNode
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (ll *LinkedList) Print() {
	var b bytes.Buffer
	n := ll.head

	if n == nil {
		fmt.Println("Empty")
	}

	b.WriteString(strconv.Itoa(n.data))
	n = n.next

	for n != nil {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(n.data))
		n = n.next
	}

	fmt.Println(b.String())
}

func Add(a, b LinkedList) LinkedList {
	return LinkedList{&Node{0, nil}}
}