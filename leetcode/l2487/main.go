package l2487

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func removeNodes(head *ListNode) *ListNode {
	node, _ := dfs(head)
	return node
}

func dfs(node *ListNode) (*ListNode, int) {
	if node == nil {
		return nil, 0
	}

	next, n := dfs(node.Next)
	if next == nil {
		return node, node.Val
	}

	if node.Val < n {
		return next, n
	} else {
		node.Next = next
		return node, node.Val
	}
}
