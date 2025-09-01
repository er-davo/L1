package main

import "sync"

type SyncMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func New[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	val, ok := m.m[key]
	m.mu.RUnlock()
	return val, ok
}

func (m *SyncMap[K, V]) Set(key K, val V) {
	m.mu.Lock()
	m.m[key] = val
	m.mu.Unlock()
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

func (m *SyncMap[K, V]) Len() int {
	if m == nil {
		return 0
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.m)
}

func main() {
	m := New[int, int]()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			m.Set(i, i)
			if m.Len() == 123 {
				m.Delete(i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			if i%3 == 0 {
				m.Get(i)
			}
		}
	}()

	wg.Wait()
}
