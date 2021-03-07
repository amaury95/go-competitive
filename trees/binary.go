package trees

import "github.com/golang-collections/collections/stack"

// BinaryTree ...
type BinaryTree struct {
	Value      int
	LeftChild  *BinaryTree
	RightChild *BinaryTree
}

// RecursivePreOrder ...
func (tree *BinaryTree) RecursivePreOrder() []int {
	result := []int{tree.Value}
	if tree.LeftChild != nil {
		result = append(result, tree.LeftChild.RecursivePreOrder()...)
	}
	if tree.RightChild != nil {
		result = append(result, tree.RightChild.RecursivePreOrder()...)
	}
	return result
}

// IterativePreOrder ...
func (tree *BinaryTree) IterativePreOrder() []int {
	var result = []int{}

	var nodes = stack.New()
	nodes.Push(tree)

	for nodes.Len() > 0 {
		node := nodes.Pop().(*BinaryTree)
		result = append(result, node.Value)
		if node.RightChild != nil {
			nodes.Push(node.RightChild)
		}
		if node.LeftChild != nil {
			nodes.Push(node.LeftChild)
		}
	}

	return result
}

// RecursiveInOrder ...
func (tree *BinaryTree) RecursiveInOrder() []int {
	result := []int{}
	if tree.LeftChild != nil {
		result = append(result, tree.LeftChild.RecursiveInOrder()...)
	}

	result = append(result, tree.Value)

	if tree.RightChild != nil {
		result = append(result, tree.RightChild.RecursiveInOrder()...)
	}
	return result
}

// IterativeInOrder ...
func (tree *BinaryTree) IterativeInOrder() []int {
	var result = []int{}

	var current *BinaryTree = tree
	var nodes = stack.New()

	for {
		if current == nil && nodes.Len() == 0 {
			break
		}

		if current != nil {
			nodes.Push(current)
			current = current.LeftChild
		} else {
			node := nodes.Pop().(*BinaryTree)
			result = append(result, node.Value)
			current = node.RightChild
		}
	}

	return result
}
