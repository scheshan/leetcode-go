package j2_024

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		next := head.Next
		head.Next = res
		res = head

		head = next
	}

	return res
}
