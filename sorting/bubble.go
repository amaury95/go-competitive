package sorting

// BubbleSort sorts the array
func BubbleSort(arr []int) []int {
	var sorted bool = false

	for !sorted {
		sorted = true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
			}
		}
	}

	return arr
}
