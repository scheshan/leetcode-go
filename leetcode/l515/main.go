package l515

import (
	"github.com/scheshan/leetcode/common"
	"math"
)

type TreeNode = common.TreeNode

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		max := math.MinInt

		for i := 0; i < size; i++ {
			node := queue[0]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		res = append(res, max)
	}
	return res
}
