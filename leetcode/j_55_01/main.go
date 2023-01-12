package j_55_01

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
