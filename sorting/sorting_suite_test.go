package sorting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSorting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sorting Suite")
}

var Array = []int{4, 2, 234, 523123, 4, 3564, 367, 345, 213, 41, 234, 23, 5632, 456}
var Sorted = []int{2, 4, 4, 23, 41, 213, 234, 234, 345, 367, 456, 3564, 5632, 523123}
