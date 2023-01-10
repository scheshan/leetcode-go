package l2236

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
