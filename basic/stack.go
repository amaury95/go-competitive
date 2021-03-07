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

// Length ...
func (s *Stack) Length() int {
	return s.length
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
