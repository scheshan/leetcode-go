package main

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	} else {
		if root.Val == 2 {
			return evaluateTree(root.Left) || evaluateTree(root.Right)
		} else {
			return evaluateTree(root.Left) && evaluateTree(root.Right)
		}
	}
}
