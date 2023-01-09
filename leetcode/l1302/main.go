package l1302

import (
	"container/list"
	"github.com/scheshan/leetcode/common"
)

type TreeNode = common.TreeNode

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	level := 0
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(0)

	for queue.Len() > 0 {
		value := queue.Remove(queue.Front())
		switch value.(type) {
		case int:
			res = level
			level = 0
			if queue.Len() > 0 {
				queue.PushBack(0)
			}
		case *TreeNode:
			node := value.(*TreeNode)
			level += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res
}
