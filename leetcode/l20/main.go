package l20

import "container/list"

func isValid(s string) bool {
	stack := list.New()

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack.PushBack(ch)
		} else {
			if stack.Len() == 0 {
				return false
			}

			pre := stack.Back().Value.(rune)
			stack.Remove(stack.Back())

			switch ch {
			case ')':
				if pre != '(' {
					return false
				}
			case ']':
				if pre != '[' {
					return false
				}
			case '}':
				if pre != '{' {
					return false
				}
			}
		}
	}

	return stack.Len() == 0
}
