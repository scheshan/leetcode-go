package l2

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	flag := false
	for l1 != nil || l2 != nil {
		n := 0
		if l1 != nil {
			n += l1.Val
		}
		if l2 != nil {
			n += l2.Val
		}

		if flag {
			flag = false
			n++
		}
		if n >= 10 {
			flag = true
			n -= 10
		}

		node := &ListNode{Val: n}
		tail.Next = node
		tail = node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if flag {
		tail.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
