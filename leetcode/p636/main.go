package p636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	deque := make([][]int, 0, n)
	res := make([]int, n)

	for _, log := range logs {
		arr := strings.Split(log, ":")
		id, _ := strconv.Atoi(arr[0])
		typ := arr[1]
		ts, _ := strconv.Atoi(arr[2])

		switch typ {
		case "start":
			if len(deque) > 0 {
				res[deque[len(deque)-1][0]] += ts - deque[len(deque)-1][1]
				deque[len(deque)-1][1] = ts
			}
			deque = append(deque, []int{id, ts})
		default:
			res[deque[len(deque)-1][0]] += ts - deque[len(deque)-1][1] + 1
			deque = deque[:len(deque)-1]
			if len(deque) > 0 {
				deque[len(deque)-1][1] = ts + 1
			}
		}
	}
	return res
}
