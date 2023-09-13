package adt

type Set struct {
	size int
	items [6]string
}

func NewSet() *Set {
	return &Set{0, [6]string{}}
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) Add(item string) {
	s.size++
	s.items[s.size] = item
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Contains(item string) bool {
	return s.items[s.size] == item
}