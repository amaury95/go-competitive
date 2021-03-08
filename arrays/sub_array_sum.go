package arrays

// IterativeSumSubArray ...
func IterativeSumSubArray(arr []int, sum int) []int {
	var index, accum int

	for i := 0; i < len(arr); i++ {
		if accum+arr[i] > sum {
			for accum+arr[i] > sum && index < i {
				accum -= arr[index]
				index++
			}
		}
		accum += arr[i]
		if accum == sum {
			return arr[index : i+1]
		}
	}

	return []int{}
}
