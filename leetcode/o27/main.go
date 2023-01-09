package o27

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
