package trees

import "github.com/golang-collections/collections/stack"

// BinaryTree ...
type BinaryTree struct {
	Value      int
	LeftChild  *BinaryTree
	RightChild *BinaryTree
}

// PreOrder ...
func (tree *BinaryTree) PreOrder() []int {
	result := []int{tree.Value}
	if tree.LeftChild != nil {
		result = append(result, tree.LeftChild.PreOrder()...)
	}
	if tree.RightChild != nil {
		result = append(result, tree.RightChild.PreOrder()...)
	}
	return result
}

// IterativePreOrder ...
func (tree *BinaryTree) IterativePreOrder() []int {
	var _ = stack.New()
	return []int{}
}
