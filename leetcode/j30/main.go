package j30

type MinStack struct {
	arr [][]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		arr: make([][]int, 0),
	}
	return stack
}

func (this *MinStack) Push(x int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, []int{x, x})
	} else {
		min := this.arr[len(this.arr)-1][1]
		if x < min {
			min = x
		}
		this.arr = append(this.arr, []int{x, min})
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1][0]
}

func (this *MinStack) Min() int {
	return this.arr[len(this.arr)-1][1]
}
