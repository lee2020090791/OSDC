package internal

type Element[T any] struct { // element in a linked list
	prev, next *Element[T]
	list       *List[T]

	Value T
}

func (e *Element[T]) Next() *Element[T] {
	if e.list == nil || e.next == &e.list.root {
		return nil
	}
	return e.next
}

func (e *Element[T]) Prev() *Element[T] {
	if e.list == nil || e.prev == &e.list.root {
		return nil
	}
	return e.prev
}

// list implement a generic linked list
// this contains the minimum functionally required for an lru cache
type List[T any] struct {
	root Element[T]
	len  int
}

func NewList[T any]() *List[T] {
	l := &List[T]{}
	l.Init()
	return l
}

func (l *List[T]) Init() {
	l.root = Element[T]{}
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *List[T]) move(e, at *Element[T]) *Element[T] {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

func (l *List[T]) Remove(e *Element[T]) T {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.list.len--
	e.next = nil
	e.prev = nil
	e.list = nil
	return e.Value
}

func (l *List[T]) PushFront(value T) *Element[T] {
	return l.insert(&Element[T]{Value: value}, &l.root)
}

func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}
