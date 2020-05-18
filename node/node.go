package node

import "fmt"

type node struct {
	Val string
	Next *node	
}

// constructor method
func New(val string) node {
	n := node {}
	n.Val = val
	n.Next = nil 
	return n
}

// Adds elements at the end of the list
func (n *node) Append(s string) {
	x := New(s)
	y := &x
	temp := n
	for temp.Next != nil {
		temp = temp.Shift()
	}
	temp.Next = y
}

// Reads the linked list elements from head to tail
func (n *node) Read() {  
	temp := n
	for temp.Next != nil {
		fmt.Println(temp.Val)
		temp = temp.Shift()
	}
	fmt.Println(temp.Val)
}

// move from a node to the next
func (n *node) Shift() (*node) {
	return n.Next
}

// Inserts new node at specified position
func (n *node) Insert(position int, word string) {
	temp := n
	for i := 1; i < position-1; i++ {
		temp = temp.Shift()
	}
	newNode := New(word)
	nodeBefore := temp
	nodeAfter := nodeBefore.Next
	nodeBefore.Next = &newNode
	newNode.Next = nodeAfter
}

// Prepend node at beginning of linked list
func (n *node) Prepend(word string) {
	temp := n
	firstNode := New(word)
	firstNode.Next = temp 
	n = &firstNode
}