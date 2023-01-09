package l2265

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func averageOfSubtree(root *TreeNode) int {
	_, _, find := travel(root)
	return find
}

func travel(root *TreeNode) (total int, count int, find int) {
	if root == nil {
		return 0, 0, 0
	}

	t1, c1, f1 := travel(root.Left)
	t2, c2, f2 := travel(root.Right)

	total = t1 + t2 + root.Val
	count = c1 + c2 + 1
	find = f1 + f2

	if root.Val == total/count {
		find++
	}

	return
}
