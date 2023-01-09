package m04_03

import (
	"github.com/scheshan/leetcode/common"
)

type TreeNode = common.TreeNode
type ListNode = common.ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	res := make([]*ListNode, 0)
	if tree == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)

	for len(queue) > 0 {
		size := len(queue)
		dummy := &ListNode{}
		tail := dummy
		for i := 0; i < size; i++ {
			node := queue[0]

			tail.Next = &ListNode{Val: node.Val}
			tail = tail.Next

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

		res = append(res, dummy.Next)
	}

	return res
}
