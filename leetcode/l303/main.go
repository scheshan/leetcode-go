package l303

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i, n := range nums {
		prefix[i+1] += prefix[i] + n
	}

	res := NumArray{
		prefix: prefix,
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefix[j+1] - this.prefix[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
