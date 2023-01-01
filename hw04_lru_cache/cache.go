package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mx       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mx.Lock()
	defer l.mx.Unlock()
	if l.capacity <= l.queue.Len() {
		l.queue.PushFront(l.items[key])
		l.queue.Remove(l.queue.Back())
		return false
	}
	if l.items[key] != nil {
		l.items[key].Value = value
		l.queue.MoveToFront(l.items[key])
		return true
	}
	l.queue.PushFront(value)
	l.items[key] = l.queue.Front()
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mx.Lock()
	defer l.mx.Unlock()
	if _, ok := l.items[key]; !ok {
		return nil, false
	}
	l.queue.MoveToFront(l.items[key])
	return l.items[key].Value, true
}

func (l *lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
