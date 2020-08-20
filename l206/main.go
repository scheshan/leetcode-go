package l206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	var tmp *ListNode

	for head != nil {
		tmp = head.Next
		head.Next = res
		res = head
		head = tmp
	}

	return res
}
