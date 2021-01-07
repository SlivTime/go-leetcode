package main

import "fmt"

type Node struct { //We create Node type here to build our linked list of those
	val  int
	next *Node
}

type MyLinkedList struct { //Our Singly Linked List has head (which is reference to our "anchor" Node in memory) and the size
	head *Node
	size int
}

func toString(l *MyLinkedList) string {
	result := ""
	if l != nil {
		result += "["
		cur := l.head
		for {
			if cur == nil {
				break
			}
			result += fmt.Sprintf("%d", cur.val)
			cur = cur.next
			if cur != nil {
				result += ","
			}
		}
		result += "]"
	}
	return result
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: new(Node), //Since we use new() function to create our "anchor" Node, it will fill up struct with Empty Values (val: 0, next: nil)
		size: 0,         // But we want to define the size explicitly
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	curr := this.head
	for ; index > 0; index-- {
		curr = curr.next
		if curr == nil {
			return -1
		}
	}
	return curr.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newHead := Node{val, this.head}
	if this != nil {
		this.head = &newHead
	}
	if this.size == 0 {
		newHead.next = nil
	}
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tailNode := Node{val, nil}
	curr := this.head

	if this.size == 0 {
		this.AddAtHead(val)
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &tailNode
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	curr := this.head
	prev := &Node{-1, curr}
	if index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	for ; index > 0; index-- {
		prev = curr
		curr = curr.next
	}

	newNode := Node{val, curr}
	prev.next = &newNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	curr := this.head

	if index >= this.size {
		return
	} else if this.size == 1 {
		this.head.val = 0
		this.size--
		return
	} else if index == 0 {
		this.head = curr.next
	}

	for ; index > 1; index-- {
		curr = curr.next
	}
	curr.next = curr.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
