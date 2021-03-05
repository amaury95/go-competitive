package basic

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitIntegers splits the given input into an array of integers
func SplitIntegers(value string) []int {
	trimmed := strings.TrimSpace(value)
	split := strings.Split(trimmed, " ")

	result := make([]int, len(split))

	for i, v := range split {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

// Round4 rounds a given number 4 decimal points
func Round4(value float64) string {
	result := fmt.Sprintf("%.4f", value)
	return result
}