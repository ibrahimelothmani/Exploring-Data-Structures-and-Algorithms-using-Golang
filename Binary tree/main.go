package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

func printTree(node *Node, level int) {
	if node == nil {
		return
	}
	printTree(node.right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}
	fmt.Println(node.data)
	printTree(node.left, level+1)
}

func findData(node *Node, data string) *Node {
	if node == nil {
		return nil
	}
	if node.data == data {
		return node
	}
	leftResult := findData(node.left, data)
	if leftResult != nil {
		return leftResult
	}
	return findData(node.right, data)
}

func main() {
	A := Node{data: "A"}
	B := Node{data: "B"}
	C := Node{data: "C"}
	D := Node{data: "D"}
	E := Node{data: "E"}
	F := Node{data: "F"}
	G := Node{data: "G"}

	A.left = &B
	A.right = &C
	B.left = &D
	B.right = &E
	C.left = &F
	C.right = &G

	fmt.Println("Binary Tree Structure:")
	printTree(&A, 0)
	dataToFind := "E"
	foundNode := findData(&A, dataToFind)
	if foundNode != nil {
		fmt.Printf("Found node: %s\n", foundNode.data)
	} else {
		fmt.Println("Node not found")
	}
}