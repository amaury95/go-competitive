package basic

import (
	"strconv"
)

// ConvertString2Int converts the given string to integer
func ConvertString2Int(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}

// ConvertString2Float converts the given string to float
func ConvertString2Float(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

// ConvertString2Boolean converts the given string to boolean
func ConvertString2Boolean(value string) bool {
	result, _ := strconv.ParseBool(value)
	return result
}

// ConvertNumber2String converts two numbers to string and contatenates them 
func ConvertNumber2String(value int) string {
	return strconv.Itoa(value)
}