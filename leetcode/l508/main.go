package l508

import "github.com/scheshan/leetcode/common"

type TreeNode = common.TreeNode

func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)
	dfs(root, m)

	max := 0
	res := make([]int, 0)
	for k, n := range m {
		if n > max {
			max = n
			res = res[:0]
			res = append(res, k)
		} else if n == max {
			res = append(res, k)
		}
	}

	return res
}

func dfs(node *TreeNode, m map[int]int) int {
	if node == nil {
		return 0
	}

	n1 := dfs(node.Left, m)
	n2 := dfs(node.Right, m)

	n := n1 + n2 + node.Val
	m[n] = m[n] + 1
	return n
}
