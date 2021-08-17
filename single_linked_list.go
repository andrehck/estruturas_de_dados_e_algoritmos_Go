package single

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node) //o primeiro nó chamad cabeça nó

func initial() {
	head.data = "San Francisco"
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", next: nil}
	nodeOakland.next = nodeBerkeley

	var nodeRiodeJaneiro *Node = &Node{data: "Rio de Janeiro", next: nil}
	nodeBerkeley.next = nodeRiodeJaneiro
	//o ultimo nó chamado nó de cauda

	var tail *Node = &Node{data: "Fremont", next: nil}
	nodeRiodeJaneiro.next = tail

}

func output(node *Node) {
	var p = node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")
}

func main() {
	initial()
	output(head)
}
