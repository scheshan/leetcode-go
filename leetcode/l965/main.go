package l965

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root, root.Val)
}

func dfs(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}

	if root.Val != value {
		return false
	}

	return dfs(root.Left, value) && dfs(root.Right, value)
}
