package l82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy

	var pre *ListNode
	duplicate := false

	for head != nil {
		if pre == nil {
			pre = head
		} else {
			if head.Val == pre.Val {
				duplicate = true
			} else {
				if !duplicate {
					tail.Next = pre
					tail = tail.Next
				}
				duplicate = false
				pre = head
			}
		}
		head = head.Next
	}

	if pre != nil && !duplicate {
		tail.Next = pre
		tail = tail.Next
	}

	tail.Next = nil

	return dummy.Next
}
