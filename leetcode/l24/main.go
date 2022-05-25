package l24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	var odd *ListNode
	var next *ListNode
	i := 0
	for head != nil {
		next = head.Next

		if i == 0 {
			odd = head
			i = 1
		} else {
			head.Next = odd
			cur.Next = head
			cur = odd
			i = 0
		}

		head = next
	}
	cur.Next = odd
	if odd != nil {
		odd.Next = nil
	}

	return dummy.Next
}
