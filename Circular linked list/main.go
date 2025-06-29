package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func printList(head *Node) {
	if head == nil {
		fmt.Println()
		return
	}
	current := head
	for {
		fmt.Print(current.Value, " ")
		current = current.Next
		if current == head {
			break
		}
	}
	fmt.Println()
}

func reverseList(head *Node) *Node {
	if head == nil {
		return nil
	}
	current := head
	var prev *Node
	for {
		next := current.Next
		current.Next = current.Previous
		current.Previous = next
		prev = current
		current = next
		if current == head {
			break
		}
	}
	return prev
}

func insertAtStart(head *Node, value int) *Node {
	newNode := &Node{Value: value}
	if head == nil {
		newNode.Next = newNode
		newNode.Previous = newNode
		return newNode
	}
	last := head.Previous
	newNode.Next = head
	newNode.Previous = last
	head.Previous = newNode
	last.Next = newNode
	return newNode
}

func insertAtEnd(head *Node, value int) *Node {
	newNode := &Node{Value: value}
	if head == nil {
		newNode.Next = newNode
		newNode.Previous = newNode
		return newNode
	}
	last := head.Previous
	last.Next = newNode
	newNode.Previous = last
	newNode.Next = head
	head.Previous = newNode
	return head
}

func insertAtMiddle(head *Node, value int) *Node {
	if head == nil {
		newNode := &Node{Value: value}
		newNode.Next = newNode
		newNode.Previous = newNode
		return newNode
	}
	slow, fast := head, head
	for fast.Next != head && fast.Next.Next != head {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newNode := &Node{Value: value}
	next := slow.Next
	slow.Next = newNode
	newNode.Previous = slow
	newNode.Next = next
	next.Previous = newNode
	return head
}

func main() {
	head := &Node{Value: 1}
	head.Next = head
	head.Previous = head

	head = insertAtEnd(head, 2)
	head = insertAtEnd(head, 3)
	head = insertAtStart(head, 0)
	head = insertAtMiddle(head, 1)

	fmt.Println("Original List:")
	printList(head)

	fmt.Println("Reversed List:")
	head = reverseList(head)
	printList(head)

	fmt.Println("List after inserting 4 at middle:")
	head = insertAtMiddle(head, 4)
	printList(head)

	fmt.Println("List after inserting 5 at start:")
	head = insertAtStart(head, 5)
	printList(head)

	fmt.Println("List after inserting 6 at end:")
	head = insertAtEnd(head, 6)
	printList(head)

}
