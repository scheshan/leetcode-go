package j2_021

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head
	var pre *ListNode

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		pre = slow
		fast = fast.Next
		slow = slow.Next
	}

	if pre == nil {
		return head.Next
	} else {
		pre.Next = pre.Next.Next
	}
	return head
}
