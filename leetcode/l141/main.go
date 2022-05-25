package l141

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil {
		if slow == fast {
			return true
		}

		if fast != nil {
			fast = fast.Next
		}
		if fast != nil {
			fast = fast.Next
		}

		if slow != nil {
			slow = slow.Next
		}
	}

	return false
}
