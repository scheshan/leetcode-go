package j2_025

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)

	dummy := &ListNode{}
	tail := dummy
	flag := false
	for l1 != nil || l2 != nil || flag {
		n := 0
		if flag {
			n = 1
			flag = false
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
			flag = true
		}

		node := &ListNode{Val: n}
		tail.Next = node
		tail = node
	}

	return reverse(dummy.Next)
}

func reverse(node *ListNode) *ListNode {
	var res *ListNode

	for node != nil {
		next := node.Next
		node.Next = res
		res = node
		node = next
	}

	return res
}
