package problem30

import "errors"

type Stack struct {
	List []int
}

func (s *Stack) Push(a int) {
	s.List = append(s.List, a)
}
func (s *Stack) Pop() (int, error) {
	length := len(s.List)
	if length == 0 {
		return -1, errors.New("栈为空")
	}
	out := s.List[length-1]
	s.List = s.List[:length-1]
	return out, nil
}
func (s *Stack) Top() (int, error) {
	length := len(s.List)
	if length == 0 {
		return -1, errors.New("栈为空")
	}
	out := s.List[length-1]
	return out, nil
}

type StackWithMin struct {
	StackData Stack
	StackMin  Stack
}

func (s *StackWithMin) Push(a int) {
	s.StackData.Push(a)
	tempMin, _ := s.StackMin.Top()
	if len(s.StackMin.List) == 0 || a < tempMin {
		s.StackMin.Push(a)
	} else {
		s.StackMin.Push(tempMin)
	}
}

func (s *StackWithMin) Pop() error {
	if len(s.StackData.List) <= 0 || len(s.StackMin.List) <= 0 {
		return errors.New("栈为空")
	} else {
		s.StackData.Pop()
		s.StackMin.Pop()
	}
	return nil
}
func (s *StackWithMin) Min() (int, error) {
	if len(s.StackData.List) <= 0 || len(s.StackMin.List) <= 0 {
		return -1, errors.New("栈为空")
	} else {
		return s.StackMin.Top()
	}
}
