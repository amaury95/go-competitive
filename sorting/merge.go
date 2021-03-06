package sorting

// MergeSort returns a sorted array by using mergesort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	return merge(MergeSort(arr[:middle]), MergeSort(arr[middle:]))
}

func merge(left, right []int) []int {
	l, r := 0, 0

	result := make([]int, len(left)+len(right))

	for i := 0; i < len(result); i++ {
		if l == len(left) {
			result[i] = right[r]
			r++
		} else if r == len(right) {
			result[i] = left[l]
			l++
		} else if left[l] <= right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
	}
	return result
}
