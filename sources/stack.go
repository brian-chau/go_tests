package main

type StackError struct {
	msg string // description of error
}

func (e *StackError) Error() string { return e.msg }

type stack struct {
	s []int
}

func NewStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	l := len(s.s)
	if l == 0 {
		return 0, &StackError{"empty stack"}
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Top() (int, error) {
	return s.Peek()
}

func (s *stack) Peek() (int, error) {
	l := len(s.s)
	if l == 0 {
		return 0, &StackError{"empty stack"}
	}

	return s.s[l-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) Size() int {
	return len(s.s)
}
