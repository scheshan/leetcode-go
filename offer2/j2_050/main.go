package main

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func pathSum(root *TreeNode, targetSum int) int {
	m := make(map[int]int)
	m[0] = 1
	return dfs(root, targetSum, 0, m)
}

func dfs(node *TreeNode, target int, cur int, m map[int]int) int {
	if node == nil {
		return 0
	}

	res := 0
	cur += node.Val
	delta := cur - target
	res += m[delta]

	m[cur]++

	res += dfs(node.Left, target, cur, m)
	res += dfs(node.Right, target, cur, m)

	m[cur]--

	return res
}

func main() {
	pathSum(&TreeNode{Val: 1}, 1)
}
