package l1022

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, n int) int {
	if root == nil {
		return 0
	}

	n = n<<1 + root.Val

	if root.Left == nil && root.Right == nil {
		return n
	} else if root.Left == nil {
		return dfs(root.Right, n)
	} else if root.Right == nil {
		return dfs(root.Left, n)
	} else {
		return dfs(root.Left, n) + dfs(root.Right, n)
	}
}
