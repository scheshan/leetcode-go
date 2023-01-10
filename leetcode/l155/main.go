package l155

type MinStack struct {
	arr [][]int
}

func Constructor() MinStack {
	stack := MinStack{
		arr: make([][]int, 0),
	}
	return stack
}

func (this *MinStack) Push(val int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, []int{val, val})
	} else {
		min := this.arr[len(this.arr)-1][1]
		if val < min {
			min = val
		}
		this.arr = append(this.arr, []int{val, min})
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1][1]
}
