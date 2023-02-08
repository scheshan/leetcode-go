package p60

func dicesProbability(n int) []float64 {
	arr1 := make([]float64, 6)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = 1.0 / 6.0
	}

	var arr2 []float64
	for i := 2; i <= n; i++ {
		arr2 = make([]float64, i*5+1)
		for j := 0; j < len(arr1); j++ {
			for k := 0; k < 6; k++ {
				arr2[j+k] += arr1[j] / 6.0
			}
		}
		arr1 = arr2
	}
	return arr1
}
