// Package ds implements a simple library of data structures.
package ds

import (
	"bytes"
	"fmt"
	"strconv"
)

// Node - an element in the linked list. Contains its data and a pointer to the next element.
type Node struct {
	Data int
	Next *Node
}

// LinkedList - a simple singlely-linked list.
type LinkedList struct {
	Head *Node
}

// func CreateLinkedList() *LinkedList {
// 	return &LinkedList{ Head: nil }
// }

// func CreateLinkedList(Data int) *LinkedList {
// 	return &LinkedList{ Head: &Node{ Data: Data, Next: nil }}
// }

// CreateLinkedList will create a linked list with an element for each parameter passed in.
func CreateLinkedList(data ...int) *LinkedList {
	var ll *LinkedList

	for i := 0; i < len(data); i++ {
		ll.Add(data[i])
	}

	return ll
}

// Add an element to the end of the linked list.
func (ll *LinkedList) Add(data int) {
	newNode := &Node{Data: data, Next: nil}

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

// Print the linked list in the form, "head -> next -> ... -> tail"
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

// Add two linked lists together - not implemented yet
func Add(a, b LinkedList) LinkedList {
	return LinkedList{&Node{0, nil}}
}
