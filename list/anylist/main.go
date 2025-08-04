package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(val T) *List[T] {
	if l == nil {

		return &List[T]{
			val:  val,
			next: nil,
		}
	}

	current := l
	for current.next != nil {
		current = current.next
	}

	current.next = &List[T]{val: val, next: nil}

	return l
}

func (l *List[T]) Add(val T) *List[T] {

	return &List[T]{val: val, next: l}
}

func (l *List[T]) Length() int {
	count := 0

	current := l
	for current.next != nil {
		count++
		current = current.next
	}

	return count
}

func (l *List[T]) String() string {
	var sb strings.Builder

	current := l

	for current.next != nil {
		sb.WriteString(fmt.Sprintf("%v ", current.val))

		current = current.next
	}

	return sb.String()
}

func main() {
	list := &List[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)
	list.Append(9)
	list.Append(10)
	list.Add(0)
	fmt.Println(list)

	fmt.Println(list.Length())
	fmt.Println(list.String())
}
