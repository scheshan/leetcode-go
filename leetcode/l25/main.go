package l25

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for head != nil {
		var node *ListNode
		nodeTail := head

		i := 0
		for ; i < k && head != nil; i++ {
			next := head.Next
			head.Next = node
			node = head
			head = next
		}

		if i == k {
			tail.Next = node
			tail = nodeTail
		} else {
			var node2 *ListNode
			for node != nil {
				next := node.Next
				node.Next = node2
				node2 = node
				node = next
			}
			tail.Next = node2
		}
	}

	return dummy.Next
}
