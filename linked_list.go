package main

import "errors"

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data: data}
}

type linked_list[T any] struct {
	head *Node[T]
	size int
}

func NewLinckedList[T any]() linked_list[T] {
	return linked_list[T]{
		head: nil,
		size: 0,
	}
}

func (l *linked_list[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *linked_list[T]) AddToHead(value T) {
	newNode := NewNode[T](value)
	newNode.next = l.head
	l.head = newNode
	l.size++
}

func (l *linked_list[T]) AddToTail(value T) {
	newNode := NewNode[T](value)
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size++
}

func (l *linked_list[T]) Remove(value T) error {
	if l.IsEmpty() {
		return errors.New("List is Empty")
	}
	current := l.head
	for current.next != nil {
		
	}
}
