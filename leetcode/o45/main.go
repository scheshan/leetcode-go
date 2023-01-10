package o45

import (
	"container/list"
	"github.com/scheshan/leetcode/common"
)

type TreeNode = common.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	res := 0

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		res = queue.Front().Value.(*TreeNode).Val

		q := list.New()

		for queue.Len() > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
		queue = q
	}

	return res
}
