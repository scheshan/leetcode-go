package l203

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for head != nil {
		if head.Val != val {
			tail.Next = head
			tail = head
		}
		head = head.Next
	}

	tail.Next = nil
	return dummy.Next
}
