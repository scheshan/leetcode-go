package j_34

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func pathSum(root *TreeNode, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	dfs(root, target, &path, &res)
	return res
}

func dfs(root *TreeNode, target int, path *[]int, res *[][]int) {
	if root == nil {
		return
	}

	*path = append(*path, root.Val)
	target -= root.Val

	if root.Left == nil && root.Right == nil {
		if target == 0 {
			arr := make([]int, len(*path))
			copy(arr, *path)
			*res = append(*res, arr)
		}
	} else {
		dfs(root.Left, target, path, res)
		dfs(root.Right, target, path, res)
	}

	*path = (*path)[:len(*path)-1]
}
