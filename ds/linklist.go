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

// func CreateLinkedList() *LinkedList {
// 	return &LinkedList{ Head: nil }
// }

// func CreateLinkedList(Data int) *LinkedList {
// 	return &LinkedList{ Head: &Node{ Data: Data, Next: nil }}
// }

func CreateLinkedList(data ...int) *LinkedList {	
	ll := &LinkedList{ Head: &Node{ Data: data[0], Next: nil }}

	for i := 1; i < len(data); i++ {
		ll.Add(data[i])
	}
	
	return ll
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{ Data: data, Next: nil }

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