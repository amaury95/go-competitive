package basic

// Push appends an element to the end of the given array
func Push(s []string, data string) []string {
	return append(s, data)
}

// Pop removes an element from the last index of the array
func Pop(s []string) []string {
	return s[:len(s)-1]
}

// RemoveAt removes the ith element of the array
func RemoveAt(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
