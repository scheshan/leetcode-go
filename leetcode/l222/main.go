package l222

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
