package j_40

import "container/heap"

type Heap []int

func (t *Heap) Len() int {
	return len(*t)
}

func (t *Heap) Less(i, j int) bool {
	return (*t)[i] > (*t)[j]
}

func (t *Heap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Heap) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *Heap) Pop() interface{} {
	n := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return n
}

func (t *Heap) Peek() int {
	return (*t)[0]
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	items := make([]int, 0, k)
	h := Heap(items)
	hp := &h

	for _, num := range arr {
		if hp.Len() < k {
			heap.Push(hp, num)
		} else if num < h.Peek() {
			heap.Pop(hp)
			heap.Push(hp, num)
		}
	}

	return *hp
}
