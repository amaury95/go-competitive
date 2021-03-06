package sorting

// QuickSort https://medium.com/@rgalus/sorting-algorithms-quick-sort-implementation-in-go-9ebfd91fe95f
func QuickSort(arr []int) []int {
	result := append([]int{}, arr...)
	sort(result, 0, len(arr)-1)
	return result
}

func sort(arr []int, start, end int) {
	if end-start < 1 {
		return
	}

	pivot := arr[end]
	split := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[split] = arr[split], arr[i]
			split++
		}
	}

	arr[end] = arr[split]
	arr[split] = pivot

	sort(arr, start, split-1)
	sort(arr, split+1, end)
}
