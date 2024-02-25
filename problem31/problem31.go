package problem31

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

func (s *Stack) Empty() bool {
	return len(s.List) == 0
}

func IsPopOrder(push, pop []int) bool {
	result := false
	s := Stack{}
	length := len(pop)
	i := 0
	j := 0
	for ; i < length; i++ {
		top, _ := s.Top()
		for ; s.Empty() || top != pop[i]; j++ {
			if j == length {
				break
			}
			s.Push(push[j])
			top, _ = s.Top()
		}
		if top != pop[i] {
			break
		}
		s.Pop()
	}
	if s.Empty() && i == length {
		result = true
	}
	return result
}
