package problem09

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

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) AppendTail(a int) {
	q.stack1.Push(a)
}
func (q *Queue) DeleteHead() (int, error) {
	if len(q.stack2.List) == 0 {
		for len(q.stack1.List) != 0 {
			in, _ := q.stack1.Pop()
			q.stack2.Push(in)
		}
	}
	if len(q.stack2.List) == 0 {
		return -1, errors.New("队列为空")
	}
	return q.stack2.Pop()
}
