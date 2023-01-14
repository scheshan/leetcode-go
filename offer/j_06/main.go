package j_06

import "github.com/scheshan/leetcode/common"

type ListNode = common.ListNode

func reversePrint(head *ListNode) []int {
	res := make([]int, 10000, 10000)
	count := 0

	for head != nil {
		res[count] = head.Val
		count++
		head = head.Next
	}

	l, r := 0, count-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res[0:count]
}
