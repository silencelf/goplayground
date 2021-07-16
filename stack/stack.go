package stack

import (
	"errors"
	"sync"
)

type Stack struct {
	lock  sync.Mutex
	value []int
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.value = append(s.value, v)
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.value)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}
	res := s.value[l-1]
	s.value = s.value[:l-1]

	return res, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.value) == 0 {
		return 0, errors.New("Empty Stack")
	}

	return s.value[len(s.value)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.value) == 0
}
