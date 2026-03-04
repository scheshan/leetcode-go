package p150

import "strconv"

func evalRPN(tokens []string) int {
	deque := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			n1 := deque[len(deque)-2]
			n2 := deque[len(deque)-1]
			deque = deque[:len(deque)-2]
			deque = append(deque, n1+n2)
		case "-":
			n1 := deque[len(deque)-2]
			n2 := deque[len(deque)-1]
			deque = deque[:len(deque)-2]
			deque = append(deque, n1-n2)
		case "*":
			n1 := deque[len(deque)-2]
			n2 := deque[len(deque)-1]
			deque = deque[:len(deque)-2]
			deque = append(deque, n1*n2)
		case "/":
			n1 := deque[len(deque)-2]
			n2 := deque[len(deque)-1]
			deque = deque[:len(deque)-2]
			deque = append(deque, n1/n2)
		default:
			n, _ := strconv.Atoi(token)
			deque = append(deque, n)
		}
	}

	return deque[0]
}
