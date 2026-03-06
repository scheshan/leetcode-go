package p141

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
