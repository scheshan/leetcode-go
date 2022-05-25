package m02_01

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	visit := make([]bool, 20001)

	dummy := new(ListNode)
	tail := dummy

	for head != nil {
		if !visit[head.Val] {
			visit[head.Val] = true
			tail.Next = head
			tail = tail.Next
		}
		head = head.Next
	}
	tail.Next = nil

	return dummy.Next
}
