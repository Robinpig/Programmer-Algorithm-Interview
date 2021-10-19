package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
)

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	head2 := &ListNode{}
	CreateListNode(head, 8)
	CreateListNode(head2, 4)
	sumHead := addTwoNumbers(head, head2)

	PrintListNode("head : ", head)
	PrintListNode("head2 : ", head2)
	PrintListNode("sum: ", sumHead)
}

func CreateListNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max*2; i++ {
		cur.Next = &ListNode{}
		cur.Next.Val = rand.Intn(max)
		cur = cur.Next
	}
}

func PrintListNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
