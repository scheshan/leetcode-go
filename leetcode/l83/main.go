package l83

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val: -200,
	}
	tail := dummy

	for head != nil {
		if head.Val != tail.Val {
			tail.Next = head
			tail = head
		}

		n := head.Next
		head.Next = nil
		head = n
	}

	return dummy.Next
}
