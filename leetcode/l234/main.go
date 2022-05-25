package l234

import (
	"container/list"
	"github.com/scheshan/leetcode/common"
)

type ListNode = common.ListNode

func isPalindrome(head *ListNode) bool {
	l := list.New()

	for head != nil {
		l.PushBack(head.Val)
		head = head.Next
	}

	for l.Len() > 1 {
		v1 := l.Remove(l.Front()).(int)
		v2 := l.Remove(l.Back()).(int)

		if v1 != v2 {
			return false
		}
	}

	return true
}
