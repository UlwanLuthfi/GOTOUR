package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new element to the end of the list.
func (l *List[T]) Push(val T) {
	last := l
	for last.next != nil {
		last = last.next
	}
	last.next = &List[T]{val: val}
}

// Len returns the length of the list.
func (l *List[T]) Len() int {
	n := 0
	for curr := l; curr != nil; curr = curr.next {
		n++
	}
	return n
}

// ForEach calls the given function for each element in the list.
func (l *List[T]) ForEach(f func(T)) {
	for curr := l; curr != nil; curr = curr.next {
		f(curr.val)
	}
}

func main() {
	l := new(List[string])
	l.Push("a")
	l.Push("b")
	fmt.Println(l.Len())

	l.ForEach(func(s string) {
		fmt.Println(s)
	})
}
