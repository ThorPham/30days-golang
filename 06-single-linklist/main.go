package main

import "fmt"

type Node struct {
	Value    int
	nextNode *Node
}

func main() {
	node1 := Node{Value: 1}
	node2 := Node{Value: 2}
	node3 := Node{Value: 3}
	node1.nextNode = &node2
	node2.nextNode = &node3
	printNode(&node1)
	fmt.Println(lenLinkList(&node1))
	node4 := Node{Value: 4}
	appendLast(&node1, &node4)
	fmt.Println(lenLinkList(&node1))
	printNode(&node1)

}

func printNode(root *Node) {
	nodeWalk := root
	for nodeWalk.nextNode != nil {
		fmt.Println(nodeWalk.Value)
		nodeWalk = nodeWalk.nextNode
	}
	fmt.Println(nodeWalk.Value)

}
func lenLinkList(root *Node) int {
	nodeWalk := root
	count := 1
	for nodeWalk.nextNode != nil {
		nodeWalk = nodeWalk.nextNode
		count += 1
	}
	return count
}

// append node in last linklist

func appendLast(root *Node, newNode *Node) {
	nodeWalk := root
	for nodeWalk.nextNode != nil {
		nodeWalk = nodeWalk.nextNode
	}
	nodeWalk.nextNode = newNode
}

func Insert(root *Node, localion int, newNode *Node) error {
	nodeWalk := root
	count := 0
	for nodeWalk.nextNode != nil {
		if count == localion-1 {
			tmp := nodeWalk.nextNode
			newNode.nextNode = tmp
			nodeWalk.nextNode = newNode
			return nil
			nodeWalk = nodeWalk.nextNode
			count = count + 1

		}

	}
	return fmt.Errorf("Went past no of elements in list")
}
