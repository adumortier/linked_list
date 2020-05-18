package main

import (
	"github.com/adumortier/linked_list/node"
)

func main() {
	list := node.New("pizza")
	list.Append("potatoe")
	list.Append("fries")
	list.Append("many things")
	
	list.Insert(1, "prout")
	
	list.Prepend("first")
	list.Read()	
}