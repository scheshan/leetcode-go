package o06

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0, 10000)
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
