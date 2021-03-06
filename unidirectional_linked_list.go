package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node) //o primeiro nó chamado cabeça nó

var tail *Node = new(Node) //o ultimo nó chamado nó de cauda

func initial() {
	head.data = "San Francisco"
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Freemont"
	tail.next = nil
	nodeBerkeley.next = tail
}

func add(data string) {
	var newNode *Node = &Node{data: data, next: nil}
	tail.next = newNode
	tail = newNode

}

func output(node *Node) {
	var p = node

	for {
		if p == nil {
			break
		}
		fmt.Printf("%s-> ", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")
}

func main() {
	initial()

	fmt.Printf("Append a new node name: Walnut to the end: \n")
	add("Walnut")

	output(head)
}
