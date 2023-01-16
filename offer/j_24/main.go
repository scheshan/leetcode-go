package j_24

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		t := head.Next
		head.Next = res
		res = head
		head = t
	}

	return res
}
