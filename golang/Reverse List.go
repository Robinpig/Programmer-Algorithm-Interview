package main

import (
	_ "fmt"
	. "github.com/isdamir/gotype"
)

func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	cur := node.Next.Next
	var next *LNode
	node.Next.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur // reverse
		cur = next
	}

}

func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	node.Next = RecursiveReverseChild(firstNode)
}

// Dummy node: head
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *LNode
	var pre *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre // reverse
		pre = next
		next = cur
	}
	node.Next = pre
}

func main() {
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("before reverse: ", head)
	Reverse(head)
	PrintNode("after reverse: ", head)
}
