package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	pre := "1"
	for i := 2; i <= n; i++ {
		cur := strings.Builder{}

		for ind := 0; ind < len(pre); {
			ch := pre[ind]
			right := ind + 1
			for right < len(pre) && pre[right] == pre[right-1] {
				right++
			}
			cnt := right - ind

			cur.WriteString(strconv.Itoa(cnt))
			cur.WriteByte(ch)

			ind = right
		}

		pre = cur.String()
	}

	return pre
}

func main() {
	countAndSay(4)
}
