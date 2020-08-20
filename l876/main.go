package l876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	low, fast := head, head.Next

	for fast != nil {
		low = low.Next

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return low
}
