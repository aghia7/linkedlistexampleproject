package linkedlist

import (
	"errors"

	"askar.khaimuldin/example/list"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrEmptyList       = errors.New("list is empty")
)

type myNode[T any] struct {
	data T
	next *myNode[T]
	prev *myNode[T]
}

func newMyNode[T any](data T) *myNode[T] {
	return &myNode[T]{
		data: data,
	}
}

type MyLinkedList[T any] struct {
	head   *myNode[T]
	tail   *myNode[T]
	length int
}

func New[T any]() *MyLinkedList[T] {
	return &MyLinkedList[T]{}
}

func (l *MyLinkedList[T]) AddFront(data T) {
	newNode := newMyNode(data)
	if l.head != nil {
		newNode.next = l.head
		l.head.prev = newNode
	} else {
		l.tail = newNode
	}

	l.head = newNode
	l.length++
}

func (l *MyLinkedList[T]) AddLast(data T) {
	if l.head == nil {
		l.AddFront(data)

		return
	}

	newNode := newMyNode(data)
	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
	l.length++
}

func (l *MyLinkedList[T]) Add(index int, data T) error {
	if index < 0 || index > l.length {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		l.AddFront(data)

		return nil
	}

	if index == l.length {
		l.AddLast(data)

		return nil
	}

	current := l.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	newNode := newMyNode(data)
	newNode.prev = current
	newNode.next = current.next
	newNode.next.prev = newNode
	current.next = newNode
	l.length++

	return nil
}

func (l *MyLinkedList[T]) RemoveFront() error {
	if l.head == nil {
		return ErrEmptyList
	}

	l.head = l.head.next
	l.head.prev = nil
	l.length--

	if l.length <= 1 {
		l.tail = l.head
	}

	return nil
}

func (l *MyLinkedList[T]) RemoveLast() error {
	if l.head == nil {
		return ErrEmptyList
	}

	if l.length == 1 {
		return l.RemoveFront()
	}

	l.tail.prev.next = nil
	l.length--

	return nil
}

func (l *MyLinkedList[T]) Remove(index int) error {
	if index < 0 || index >= l.length {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		return l.RemoveFront()
	}

	if index == l.length-1 {
		return l.RemoveLast()
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.prev.next = current.next
	current.next.prev = current.prev
	l.length--

	return nil

}

func (l *MyLinkedList[T]) Get(index int) (T, error) {
	var t T
	if index < 0 || index >= l.length {
		return t, ErrIndexOutOfRange
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data, nil
}

func (l *MyLinkedList[T]) Length() int {
	return l.length
}

func (l *MyLinkedList[T]) ToSlice() []T {
	values := make([]T, 0, l.length)
	current := l.head
	for current != nil {
		values = append(values, current.data)
		current = current.next
	}

	return values
}

func (l *MyLinkedList[T]) Iterator() list.Iterator[T] {
	return newMyLinkedListIterator(l)
}

type myLinkedListIterator[T any] struct {
	current *myNode[T]
}

func newMyLinkedListIterator[T any](list *MyLinkedList[T]) *myLinkedListIterator[T] {
	return &myLinkedListIterator[T]{
		current: list.head,
	}
}

func (i *myLinkedListIterator[T]) HasNext() bool {
	return i.current != nil
}

func (i *myLinkedListIterator[T]) Next() (T, error) {
	var data T
	if i.current == nil {
		return data, ErrIndexOutOfRange
	}

	data = i.current.data
	i.current = i.current.next

	return data, nil
}
