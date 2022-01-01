package datastructres

import "errors"

type Slice struct {
	list []int
}

type ISlice interface {
	Init([]int)
	GetIndexByValue(int)
	GetValueByIndex(int)
}

func (s *Slice) GetIndexByValue(value int) int {
	for i, v := range s.list {
		if v == value {
			return i
		}
	}
	return -1
}
func (s *Slice) GetValueByIndex(index int) (int, error) {
	if index < 0 {
		return 0, errors.New("index not valid")
	}
	for i, v := range s.list {
		if i == index {
			return v, nil
		}
	}
	return 0, errors.New("index out of range")
}
func (s *Slice) Init(array []int) {
	s.list = array
}
