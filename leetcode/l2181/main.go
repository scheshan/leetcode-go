package l2181

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	var last *ListNode

	for head != nil {
		if head.Val == 0 {
			if last != nil {
				tail.Next = last
				tail = last
				last = nil
			}
		} else {
			if last == nil {
				last = &ListNode{Val: head.Val}
			} else {
				last.Val += head.Val
			}
		}
		head = head.Next
	}

	return dummy.Next
}
