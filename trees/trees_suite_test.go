package trees_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trees Suite")
}
