package mapwithexpiration

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	value      interface{}
	lastAccess int64
}

type TtlMap struct {
	m  map[string]*Item
	mu sync.Mutex
}

func New(size int, maxTtl int) *TtlMap {
	// map with given size
	m := &TtlMap{m: make(map[string]*Item, size)}

	// Go-routine to clean up old items
	go func() {
		for now := range time.Tick(time.Second) {
			m.mu.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTtl) {
					delete(m.m, k)
				}
			}
			m.mu.Unlock()
		}
	}()

	return m
}

func (m *TtlMap) Put(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	it, ok := m.m[key]
	if !ok {
		it = &Item{
			value: value,
		}
	}

	it.value = value
	it.lastAccess = time.Now().Unix()
	m.m[key] = it
}

func (m *TtlMap) Get(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if it, ok := m.m[key]; ok {
		it.lastAccess = time.Now().Unix()
		return it.value, true
	}

	return nil, false
}

func (m *TtlMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, key)
}

func DoWork() {
	m := New(100, 4)

	m.Put("key1", "string1")
	v, _ := m.Get("key1")

	fmt.Println(v)

	var wg sync.WaitGroup
	wg.Add(1)

	time.AfterFunc(time.Second*5, func() {
		m.Put("key2", "string2")
		fmt.Println(m)
		wg.Done()
	})

	wg.Wait()

}
