package main

import (
	"sort"
)

type Set interface {
	Contains(int) bool
	Add(int)
	Remove(int)
}

type LinearSlice []int

func (l LinearSlice) Contains(value int) bool {
	for _, element := range l {
		if element == value {
			return true
		}
	}
	return false
}

func (l *LinearSlice) Add(value int) {
	list := *l
	for list.Contains(value) {
		return
	}
	list = append(list, value)
	*l = list
}

func (l *LinearSlice) Remove(value int) {
	list := *l
	for idx, element := range list {
		if element == value {
			list = append(list[:idx], list[idx+1:]...)
			*l = list
			break
		}
	}
}

type SortedInts []int

func (l SortedInts) binarySearch(value int) (int, bool) {
	idx := sort.Search(len(l), func(i int) bool { return l[i] >= value })
	exists := idx != len(l) && l[idx] == value
	return idx, exists
}
func (l SortedInts) Contains(value int) bool {
	_, ok := l.binarySearch(value)
	return ok
}
func (l *SortedInts) Add(value int) {
	list := *l
	idx, ok := list.binarySearch(value)
	if ok {
		return
	}
	list = append(list, 0)
	copy(list[idx+1:], list[idx:])
	list[idx] = value
	*l = list
}
func (l *SortedInts) Remove(value int) {
	list := *l
	idx, ok := list.binarySearch(value)
	if !ok {
		return
	}
	list = append(list[:idx], list[idx+1:]...)
	*l = list
}

type IntHashSet map[int]struct{}

var void = struct{}{}

func (s IntHashSet) Contains(value int) bool {
	_, ok := s[value]
	return ok
}
func (s IntHashSet) Add(value int) {
	s[value] = void
}
func (s IntHashSet) Remove(value int) {
	delete(s, value)
}
