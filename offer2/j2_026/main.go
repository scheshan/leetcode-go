package main

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func reorderList(head *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next
		fast = fast.Next
		slow = slow.Next
	}

	var right *ListNode
	node := slow.Next
	slow.Next = nil

	for node != nil {
		next := node.Next
		node.Next = right
		right = node
		node = next
	}

	for head != nil && right != nil {
		n1 := head.Next
		n2 := right.Next

		head.Next = right
		right.Next = n1

		head = n1
		right = n2
	}
}
