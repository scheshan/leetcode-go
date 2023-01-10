package o60

import "container/heap"

type Heap struct {
	arr [][]int
}

func (t *Heap) Len() int {
	return len(t.arr)
}

func (t *Heap) Less(i, j int) bool {
	return t.arr[i][1] < t.arr[j][1]
}

func (t *Heap) Swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Heap) Push(x interface{}) {
	t.arr = append(t.arr, x.([]int))
}

func (t *Heap) Pop() interface{} {
	arr := t.arr[len(t.arr)-1]
	t.arr = t.arr[:len(t.arr)-1]
	return arr
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = m[num] + 1
	}

	h := &Heap{
		arr: make([][]int, 0, k),
	}

	for num, v := range m {
		if h.Len() < k {
			heap.Push(h, []int{num, v})
		} else {
			if v > h.arr[0][1] {
				heap.Pop(h)
				heap.Push(h, []int{num, v})
			}
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([]int)[0]
	}

	return res
}
