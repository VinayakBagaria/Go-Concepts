package lrucache

import (
	"errors"
	"fmt"

	doublylinkedlist "container/list"
)

type cacheKey int
type cacheValue interface{}

type CacheObject struct {
	key   cacheKey
	value cacheValue
}

type LRUCache struct {
	capacity int
	list     *doublylinkedlist.List
	elements map[cacheKey]*doublylinkedlist.Element
}

func New(capacity int) (*LRUCache, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("capacity cannot be 0")
	}

	return &LRUCache{
		capacity: capacity,
		list:     doublylinkedlist.New(),
		elements: make(map[cacheKey]*doublylinkedlist.Element, capacity),
	}, nil
}

func (cache *LRUCache) Get(key cacheKey) (cacheValue, error) {
	elem, ok := cache.elements[key]
	if !ok {
		return nil, errors.New("cache key was not found")
	}

	cache.list.MoveToFront(elem)
	value := elem.Value.(CacheObject).value
	return value, nil
}

func (cache *LRUCache) Put(key cacheKey, value cacheValue) {
	elem, ok := cache.elements[key]
	cacheObj := CacheObject{
		key:   key,
		value: value,
	}

	if ok {
		elem.Value = cacheObj
		cache.list.MoveToFront(elem)
	} else {
		if cache.list.Len() == cache.capacity {
			lastNode := cache.list.Back()
			key := lastNode.Value.(CacheObject).key
			delete(cache.elements, key)
			cache.list.Remove(lastNode)
		}

		element := cache.list.PushFront(cacheObj)
		cache.elements[key] = element
	}
}

func (cache *LRUCache) Purge() {
	cache.list = doublylinkedlist.New()
	cache.elements = make(map[cacheKey]*doublylinkedlist.Element, cache.capacity)
}

func (cache *LRUCache) Print() {
	currentElement := cache.list.Front()
	if currentElement == nil {
		fmt.Println("nothing to show")
		return
	}

	for currentElement != nil {
		cacheObject := currentElement.Value.(CacheObject)
		fmt.Printf("Key: %d, Value: %+v\n", cacheObject.key, cacheObject.value)
		currentElement = currentElement.Next()
	}

	fmt.Println()
}

func DoWork() {
	lru, _ := New(5)
	lru.Put(1, "hello")
	lru.Put(2, 20)
	lru.Put(3, 3.4)
	lru.Print()

	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println()

	lru.Put(1, true)
	lru.Print()

	lru.Purge()
	lru.Print()
	fmt.Println()

	lru.Put(1, "back")
	lru.Put(2, 2)
	lru.Put(3, true)
	lru.Print()
}
