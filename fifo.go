package fifo

import "github.com/WinPooh32/ring"

type Cache[K comparable, V any] struct {
	r ring.Ring[K]
	m map[K]V
}

func New[K comparable, V any](cap int) *Cache[K, V] {
	return &Cache[K, V]{
		r: ring.Make[K](cap),
		m: make(map[K]V, cap),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	if k, pop := c.r.Push(key); pop {
		delete(c.m, k)
	}
	c.m[key] = value
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	value, ok = c.m[key]
	return
}

func (c *Cache[K, V]) Reset() {
	for k := range c.m {
		delete(c.m, k)
	}
	c.r.Reset()
}
