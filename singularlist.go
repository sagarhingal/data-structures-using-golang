package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	value interface{}
	next  *Node
}

// createnode : This function starts the list
// by creating the head node and then calling
// the other child function which creates the rest
// of the nodes.
func createnode(s string, mylist *LinkedList) {

	// check if it's the first node
	// if Yes - create the head node
	// if No - traverse the list and
	// add at the end.
	if mylist.head == nil {
		mylist.head = &Node{s, nil}
		mylist.head.next = nil
		return
	} else {
		traverselist(mylist.head, s)
	}

}

// traverselist : This function goes through all the
// nodes and adds a new node at the end.
func traverselist(root *Node, s string) {
	if root.next == nil {
		// reached last node
		root.next = &Node{s, nil}
	} else {
		traverselist(root.next, s)
	}
}

// makemylist : A parent function responsible for
// initiating the node creation process.
func makemylist(myarray []string, mylist *LinkedList) {

	for _, value := range myarray {
		createnode(value, mylist)
	}

}

// printnodes : This function prints all the nodes
// in the list by traversing in a recursive manner.
func printnodes(root *Node) {

	if root.next == nil {
		// Tail condition
		fmt.Printf("|%s|X|\n\n", root.value)
	} else {
		fmt.Printf("|%s|->", root.value)
		// Repeat the process for next node
		printnodes(root.next)
	}

}

func main() {

	// Initialise the list
	mylist := LinkedList{}

	// A character list for example
	a := []string{"a", "b", "c", "d", "e", "f", "h"}
	makemylist(a, &mylist)
	fmt.Printf("\n|H|%s|->", mylist.head.value)
	printnodes(mylist.head.next)
}
