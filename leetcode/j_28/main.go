package j_28

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return travel(root.Left, root.Right)
}

func travel(left *TreeNode, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		return travel(left.Left, right.Right) && travel(left.Right, right.Left)
	} else if left == nil && right == nil {
		return true
	} else {
		return false
	}
}
