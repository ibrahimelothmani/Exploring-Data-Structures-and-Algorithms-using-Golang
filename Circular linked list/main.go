package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
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

func main() {
	head := &Node{Value: 1}
	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtStart(head, 0)
	head = insertAtMiddle(head, 1)

	fmt.Println("Original List:")
	printList(head)

	fmt.Println("Reversed List:")
	reversedHead := reverseList(head)
	printList(reversedHead)

	fmt.Println("List after inserting 4 at middle:")
	head = insertAtMiddle(reversedHead, 4)
	printList(head)

	fmt.Println("List after inserting 5 at start:")
	head = insertAtStart(head, 5)
	printList(head)

	fmt.Println("List after inserting 6 at end:")
	head = insertAtEnd(head, 6)
	printList(head)
}