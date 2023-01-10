package l451

import (
	"container/heap"
	"strings"
)

type Heap struct {
	arr [][]int
}

func (t *Heap) Len() int {
	return len(t.arr)
}

func (t *Heap) Less(i, j int) bool {
	return t.arr[i][1] > t.arr[j][1]
}

func (t *Heap) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Heap) Push(x interface{}) {
	t.arr = append(t.arr, x.([]int))
}

func (t *Heap) Pop() interface{} {
	res := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return res
}

func frequencySort(s string) string {
	m := make(map[rune]int)

	for _, ch := range s {
		m[ch] = m[ch] + 1
	}

	h := &Heap{arr: make([][]int, 0, len(m))}
	for ch, cnt := range m {
		heap.Push(h, []int{int(ch), cnt})
	}

	builder := strings.Builder{}
	for builder.Len() < len(s) {
		p := heap.Pop(h).([]int)

		for i := 0; i < p[1]; i++ {
			builder.WriteRune(rune(p[0]))
		}
	}

	return builder.String()
}
