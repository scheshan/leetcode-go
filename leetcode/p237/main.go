package p237

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
