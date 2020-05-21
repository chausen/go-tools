package ds

import (
	"bytes"
	"strconv"
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func CreateLinkedList(Data int) *LinkedList {
	return &LinkedList{ Head: &Node{ Data: Data, Next: nil }}
}

func (ll *LinkedList) Add(Data int) {
	newNode := &Node{ Data: Data, Next: nil }

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		currentNode := ll.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

func (ll *LinkedList) Print() {
	var b bytes.Buffer
	n := ll.Head

	if n == nil {
		fmt.Println("Empty")
	}

	b.WriteString(strconv.Itoa(n.Data))
	n = n.Next

	for n != nil {
		b.WriteString(" -> ")
		b.WriteString(strconv.Itoa(n.Data))
		n = n.Next
	}

	fmt.Println(b.String())
}

func Add(a, b LinkedList) LinkedList {
	return LinkedList{&Node{0, nil}}
}