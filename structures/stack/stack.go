package stack

import "fmt"

type Stack struct {
	values []int
}

func New() *Stack {
	return &Stack{values: []int{}}
}

func (s *Stack) Append(v int) {
	s.values = append(s.values, v)
}

func (s *Stack) Pop() int {
	n := len(s.values) - 1
	v := s.values[n]
	s.values = s.values[:n]
	return v
}

func (s *Stack) PrintTop() {
	fmt.Println(s.values[len(s.values)-1])
}

func Test() {
	stack := New()

	for i := 0; i < 10; i++ {
		stack.Append(i)
		stack.PrintTop()
	}

	for i := 0; i < 10; i++ {
		stack.PrintTop()
		fmt.Println(stack.Pop())
	}
}
