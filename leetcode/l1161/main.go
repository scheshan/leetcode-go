package l1161

import (
	"github.com/scheshan/leetcode/common"
	"math"
)

type TreeNode = common.TreeNode

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxNumber := math.MinInt
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	row := 0
	for len(queue) > 0 {
		row++
		size := len(queue)
		num := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			num += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}

		if num > maxNumber {
			maxNumber = num
			res = row
		}
	}

	return res
}
