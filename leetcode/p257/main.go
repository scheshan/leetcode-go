package p257

import (
	"github.com/scheshan/leetcode/common"
	"strconv"
	"strings"
)

type TreeNode = common.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	path := make([]int, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			b := strings.Builder{}
			for _, num := range path {
				if b.Len() > 0 {
					b.WriteString("->")
				}
				b.WriteString(strconv.Itoa(num))
			}
			res = append(res, b.String())
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}
