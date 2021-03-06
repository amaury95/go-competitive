package trees

// Tree write...
type Tree struct {
	value    int
	children TreeSlice
}

// TreeSlice write...
type TreeSlice []Tree

// Node write..
func Node(value int) Tree {
	return Tree{value: value, children: []Tree{}}
}

// Value write...
func (t *Tree) Value() int {
	return t.value
}

// Children write...
func (t *Tree) Children() []Tree {
	return t.children
}

// AddChild wirte...
func (t *Tree) AddChild(c ...Tree) {
	t.children = append(t.children, c...)
}

// IsBinary write...
func (t *Tree) IsBinary() bool {
	if len(t.children) > 2 {
		return false
	}

	for _, n := range t.children {
		if !n.IsBinary() {
			return false
		}
	}

	return true
}

// PreOrder write...
func (t *Tree) PreOrder() []int {
	result := []int{t.value}
	for _, e := range t.children {
		result = append(result, e.PreOrder()...)
	}
	return result
}
