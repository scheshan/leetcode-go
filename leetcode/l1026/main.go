package l1026

import (
	"github.com/scheshan/leetcode/common"
)

type TreeNode = common.TreeNode

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, minNum int, maxNum int) int {
	if root == nil {
		return 0
	}

	res := max(abs(root.Val-minNum), abs(root.Val-maxNum))

	if root.Val < minNum {
		minNum = root.Val
	}
	if root.Val > maxNum {
		maxNum = root.Val
	}
	res = max(res, dfs(root.Left, minNum, maxNum))
	res = max(res, dfs(root.Right, minNum, maxNum))

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
