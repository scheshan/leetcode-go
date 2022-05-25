package o24

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		t := head.Next
		head.Next = res
		res = head
		head = t
	}

	return res
}
