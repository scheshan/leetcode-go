package l142

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for slow != nil && fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
