package main

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
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
		if head.Val != right.Val {
			return false
		}
		head = head.Next
		right = right.Next
	}

	return true
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	isPalindrome(head)
}
