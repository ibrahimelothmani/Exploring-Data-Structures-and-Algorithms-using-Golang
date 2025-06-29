package main

import "fmt"

type Node struct {
 Value    int
 Next     *Node
 Previous *Node
 }

 func printList(head *Node){
	current := head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
 }
func reverseList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		current.Previous = next // Set Previous pointer for doubly linked list
		prev = current
		current = next
	}
	return prev
}

func insertAtMiddle(head *Node, value int) *Node {
	newNode := &Node{Value: value}
	if head == nil {
		return newNode
	}

	slow, fast := head, head
	var prev *Node

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = newNode
		newNode.Previous = prev
	}
	newNode.Next = slow
	if slow != nil {
		slow.Previous = newNode
	}

	return head
}

func insertAtStart(head *Node, value int) *Node {
	newNode := &Node{Value: value}
	if head == nil {
		return newNode
	}
	newNode.Next = head
	head.Previous = newNode
	return newNode
}
func insertAtEnd(head *Node, value int) *Node {
	newNode := &Node{Value: value}
	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	newNode.Previous = current
	return head
}

func deleteNode(node *Node) {
	if node == nil || node.Next == nil {
		return // Cannot delete the last node or nil node
	}
	node.Value = node.Next.Value
	node.Next = node.Next.Next
	if node.Next != nil {
		node.Next.Previous = node
	}
}

func main() {
	head := &Node{Value: 1}
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtStart(head, 0)
	head = insertAtMiddle(head, 1)

	fmt.Println("Original List:")
	printList(head)
	reversedHead := reverseList(head)
	fmt.Println("\nReversed List:")
	printList(reversedHead)
	fmt.Println("\nInserting 4 at the end:")
	reversedHead = insertAtEnd(reversedHead, 4)
	printList(reversedHead)
	fmt.Println("\nInserting 5 at the start:")
	reversedHead = insertAtStart(reversedHead, 5)
	printList(reversedHead)
	fmt.Println("\nInserting 6 at the middle:")
	reversedHead = insertAtMiddle(reversedHead, 6)
	printList(reversedHead)
	fmt.Println("\nDeleting a node:")
	deleteNode(reversedHead.Next) // Deleting the second node
	printList(reversedHead)
	fmt.Println("\nEnd of Doubly Linked List Operations")
}