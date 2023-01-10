package main

import (
	"fmt"
	"github.com/scheshan/leetcode/common"
)

type ListNode = common.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	l1 := &ListNode{}
	t1 := l1

	for i := 1; i < left; i++ {
		t1.Next = head
		t1 = head
		head = head.Next
	}

	l2 := head
	t2 := head
	head = head.Next

	for i := left + 1; i <= right; i++ {
		next := head.Next
		head.Next = l2
		l2 = head
		head = next
	}

	t1.Next = l2
	t1 = t2
	t1.Next = nil

	for head != nil {
		t1.Next = head
		t1 = head
		head = head.Next
	}

	return l1.Next
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 5}}
	node := reverseBetween(head, 1, 2)

	fmt.Print(node.Val)
}
