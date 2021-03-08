package basic

import "errors"

// Stack ...
type Stack struct {
	values []int
	length int
}

// Values ...
func (s *Stack) Values() []int {
	return s.values[:s.length]
}

// Len ...
func (s *Stack) Len() int {
	return s.length
}

// Peek ...
func (s *Stack) Peek() (int, error) {
	if s.length == 0 {
		return 0, errors.New("empty stack")
	}

	return s.values[s.length-1], nil
}

// Push ...
func (s *Stack) Push(value int) {
	if len(s.values) == s.length {
		s.values = append(s.values, make([]int, 32)...)
	}
	s.values[s.length] = value
	s.length++
}

// Pop ...
func (s *Stack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("stack is empty")
	}

	s.length--
	return s.values[s.length], nil
}
