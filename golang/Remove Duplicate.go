package main

import (
	_ "fmt"
	. "github.com/isdamir/gotype"
	"math/rand"
)

func RemoveDup(head *LNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	m := map[interface{}]interface{}{}

	cur := head.Next
	prev := head.Next
	for cur != nil {
		_, ok := m[cur.Data]
		if ok {
			prev.Next = cur.Next
		} else {
			m[cur.Data] = cur.Data
			prev = cur
		}
		cur = cur.Next
	}

}

/**
remove duplicate Node

1. use HashSet
2. sort list before remove

How about a sorted/un-sorted list?

*/
func main2() {

	head := &LNode{}
	CreateNodeT(head, 10)
	PrintNode("Before remove: ", head)
	RemoveDup(head)
	PrintNode("After remove: ", head)
}

func CreateNodeT(node *LNode, max int) {
	cur := node
	for i := 1; i < max*2; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = rand.Intn(max)
		cur = cur.Next
	}
}
