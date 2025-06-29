package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func printList(head *Node) {
    current := head
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}

func reverseList(head *Node) *Node {
    var prev *Node
    current := head
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}

func findMiddle(head *Node) *Node {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func listLength(head *Node) int {
    length := 0
    current := head
    for current != nil {
        length++
        current = current.Next
    }
    return length
}

func nthFromEnd(head *Node, n int) *Node {
    fast, slow := head, head
    for i := 0; i < n; i++ {
        if fast == nil {
            return nil
        }
        fast = fast.Next
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}

func deleteNode(node *Node) {
	if node == nil || node.Next == nil {
		return // Cannot delete the last node or nil node
	}
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    aa.Next = &bb
    bb.Next = &cc

    fmt.Println("Singly Linked List:")
    printList(&aa)
    fmt.Println("End of Singly Linked List")

    fmt.Println("Reversing the Singly Linked List:")
    reversedHead := reverseList(&aa)
    printList(reversedHead)
    fmt.Println("End of Reversed Singly Linked List")

    fmt.Println("Finding the middle of the Singly Linked List:")
    middle := findMiddle(reversedHead)
    fmt.Println("Middle element:", middle.Value)
    fmt.Println("End of Finding Middle Element")

    fmt.Println("Finding the length of the Singly Linked List:")
    length := listLength(reversedHead)
    fmt.Println("Length of Singly Linked List:", length)
    fmt.Println("End of Length Calculation")

    fmt.Println("Finding the nth node from the end of the Singly Linked List:")
    n := 2 // Change this value to find the nth node from the end
    nth := nthFromEnd(reversedHead, n)
    if nth != nil {
        fmt.Println("Nth node from the end is:", nth.Value)
    } else {
        fmt.Println("List is shorter than n")
    }
    fmt.Println("End of Finding Nth Node")
	fmt.Println("Deleting a node in the Singly Linked List:")
	deleteNode(reversedHead)
	printList(reversedHead)
	fmt.Println("End of Deleting a Node")
}