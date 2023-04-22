package lrucache

import (
	"container/list"
	"errors"
	"fmt"
)

type cacheKey int
type cacheValue interface{}

type CacheObject struct {
	key   cacheKey
	value cacheValue
}

type LruCache struct {
	capacity int
	list     *list.List
	elements map[cacheKey]*list.Element
}

func New(capacity int) (*LruCache, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("capacity cannot be 0")
	}

	return &LruCache{
		capacity: capacity,
		list:     new(list.List),
		elements: make(map[cacheKey]*list.Element, capacity),
	}, nil
}

func (cache *LruCache) Get(key cacheKey) (cacheValue, error) {
	elem, ok := cache.elements[key]
	if !ok {
		return nil, errors.New("cache key was not found")
	}

	value := elem.Value.(*CacheObject).value
	cache.list.MoveToFront(elem)
	return value, nil
}

func (cache *LruCache) Put(key cacheKey, value cacheValue) {
	elem, ok := cache.elements[key]
	if ok {
		elem.Value = &CacheObject{
			key:   key,
			value: value,
		}
		cache.list.MoveToFront(elem)
	} else {
		if cache.list.Len() == cache.capacity {
			key := cache.list.Back().Value.(*CacheObject).key
			delete(cache.elements, key)
			cache.list.Remove(cache.list.Back())
		}

		cacheObj := &CacheObject{
			key:   key,
			value: value,
		}
		pointer := cache.list.PushFront(cacheObj)
		cache.elements[key] = pointer
	}
}

func (cache *LruCache) Purge() {
	cache.list = new(list.List)
	cache.elements = make(map[cacheKey]*list.Element, cache.capacity)
}

func (cache *LruCache) Print() {
	currentElem := cache.list.Front()
	if currentElem == nil {
		fmt.Println("nothing to show")
		return
	}

	for {
		if currentElem.Next() == nil {
			fmt.Printf("Key: %d, Value: %+v\n", currentElem.Value.(*CacheObject).key, currentElem.Value.(*CacheObject).value)
			break
		}

		fmt.Printf("Key: %d, Value: %+v\n", currentElem.Value.(*CacheObject).key, currentElem.Value.(*CacheObject).value)
		currentElem = currentElem.Next()
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
	lru.Print()

	lru.Put(1, true)
	lru.Print()

	lru.Purge()
	lru.Print()

	lru.Put(1, "back")
	lru.Put(2, 2)
	lru.Put(3, true)
	lru.Print()
}
