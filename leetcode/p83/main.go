package p83

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	pre := -101

	for head != nil {
		next := head.Next
		if head.Val != pre {
			tail.Next = head
			tail = tail.Next
		}
		pre = head.Val
		head = next
	}

	tail.Next = nil
	return dummy.Next
}
