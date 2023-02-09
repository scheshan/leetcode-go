package p445

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)

	dummy := &ListNode{}
	tail := dummy
	promotion := false

	for l1 != nil || l2 != nil || promotion {
		n := 0
		if promotion {
			n = 1
			promotion = false
		}
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		if n >= 10 {
			n -= 10
			promotion = true
		}
		tail.Next = &ListNode{Val: n}
		tail = tail.Next
	}
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}
