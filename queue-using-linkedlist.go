package main

import "fmt"

type Queue struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
}

// enqueue : This function adds the element to the
// last position of the queue using the tail marker
// in the linked list.
func enqueue(a int, myqueue *Queue) {

	// check for the first node in the queue
	if myqueue.head == nil {
		myqueue.head = &Node{a, nil}
		myqueue.tail = myqueue.head
	} else {
		newtail := &Node{a, nil}
		myqueue.tail.next = newtail
		myqueue.tail = newtail
	}
}

// dequeue : This function removes the number of elements
// from the given list based on the given count using recursion.
// e.g if 2 is provided, then 2 elements will be removed
// from the list.
func dequeue(count int, myqueue *Queue) {

	if count > 0 {
		newhead := myqueue.head.next
		nextnode := myqueue.head.next.next
		myqueue.head = newhead
		myqueue.head.next = nextnode
		count -= 1
		dequeue(count, myqueue)
	}

}

// createqueue : This method acts as a base to run
// the enqueue function "n" times with respect to
// the size of the data array.
func createqueue(mylist []int, root *Queue) {

	for _, value := range mylist {
		enqueue(value, root)
	}
}

// printqueue : This method prints all the elements
// of the queue by traversing through it using
// recursive technique.
func printqueue(root *Node) {

	if root.next == nil {
		// last element
		fmt.Printf("|%d|\n", root.value)
	} else {
		fmt.Printf("|%d|-", root.value)
		printqueue(root.next)
	}
}

func main() {

	myqueue := Queue{}
	a := []int{1, 2, 3, 4, 5}
	createqueue(a, &myqueue)
	// the original queue
	fmt.Printf("Original list before dequeue -> ")
	printqueue(myqueue.head)

	// remove required amount of elements
	dequeue(2, &myqueue)
	// print the new queue
	fmt.Printf("Original list after dequeue --> ")
	printqueue(myqueue.head)

}
