package l129

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return num*10 + root.Val
	}

	return dfs(root.Left, num*10+root.Val) + dfs(root.Right, num*10+root.Val)
}
