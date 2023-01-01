package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	listItem := &ListItem{Value: v, Prev: nil, Next: nil}
	if l.head == nil {
		l.tail = listItem
	} else {
		l.head.Prev = listItem
	}
	listItem.Next = l.head
	l.head = listItem
	l.len++
	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := &ListItem{Next: nil, Prev: nil, Value: v}
	if l.head == nil {
		l.head = listItem
	} else {
		l.tail.Next = listItem
	}
	listItem.Prev = l.tail
	l.tail = listItem
	l.len++
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Next == nil:
		i.Prev.Next = nil
	case i.Prev == nil:
		i.Next.Prev = nil
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	return new(list)
}
