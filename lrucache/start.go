package lrucache

import (
	"errors"
	"fmt"
	"go-concepts/doublylinkedlist"
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
	elements map[cacheKey]*doublylinkedlist.Node
}

func New(capacity int) (*LRUCache, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("capacity cannot be 0")
	}

	return &LRUCache{
		capacity: capacity,
		list:     &doublylinkedlist.List{},
		elements: make(map[cacheKey]*doublylinkedlist.Node, capacity),
	}, nil
}

func (cache *LRUCache) Get(key cacheKey) (cacheValue, error) {
	elem, ok := cache.elements[key]
	if !ok {
		return nil, errors.New("cache key was not found")
	}

	value := elem.Value.(*CacheObject).value
	cache.list.MoveToFront(elem)
	return value, nil
}

func (cache *LRUCache) Put(key cacheKey, value cacheValue) {
	elem, ok := cache.elements[key]
	cacheObj := &CacheObject{
		key:   key,
		value: value,
	}

	if ok {
		elem.Value = cacheObj
		cache.list.MoveToFront(elem)
	} else {
		if cache.list.Length() == cache.capacity {
			lastNode := cache.list.Back()
			key := lastNode.Value.(*CacheObject).key
			delete(cache.elements, key)
			cache.list.Remove(lastNode)
		}

		node := &doublylinkedlist.Node{Value: cacheObj}
		cache.list.PushFront(node)
		cache.elements[key] = node
	}
}

func (cache *LRUCache) Purge() {
	cache.list = &doublylinkedlist.List{}
	cache.elements = make(map[cacheKey]*doublylinkedlist.Node, cache.capacity)
}

func (cache *LRUCache) Print() {
	currentElement := cache.list.Front()
	if currentElement == nil {
		fmt.Println("nothing to show")
		return
	}

	for currentElement != nil {
		cacheObject := currentElement.Value.(*CacheObject)
		fmt.Printf("Key: %d, Value: %+v\n", cacheObject.key, cacheObject.value)
		currentElement = currentElement.Next
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

	lru.Put(1, true)
	lru.Print()

	lru.Purge()
	lru.Print()

	lru.Put(1, "back")
	lru.Put(2, 2)
	lru.Put(3, true)
	lru.Print()
}
