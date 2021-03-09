package l1047

import "container/list"

func removeDuplicates(S string) string {
	stack := list.New()

	for _, ch := range S {
		if stack.Len() > 0 && stack.Back().Value.(rune) == ch {
			stack.Remove(stack.Back())
			continue
		}

		stack.PushBack(ch)
	}

	res := make([]rune, stack.Len())
	i := 0
	for stack.Len() > 0 {
		res[i] = stack.Remove(stack.Front()).(rune)
		i++
	}

	return string(res)
}
