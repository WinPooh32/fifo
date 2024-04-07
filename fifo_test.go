package fifo

import (
	"fmt"
	"math/rand"
	"testing"
)

func mustFillCache[K comparable, V any](t *testing.T, capacity int, keys []K, values []V) *Cache[K, V] {
	t.Helper()
	if len(keys) != len(values) {
		t.Fatal("keys and values lengths must be equal")
	}

	cache := New[K, V](capacity)

	for i, k := range keys {
		cache.Set(k, values[i])
	}

	return cache
}

func ExampleCache() {
	const capacity = 3

	keys := []string{"a", "b", "c", "d", "e", "f"}
	data := []int{0, 1, 2, 3, 4, 5}

	cache := New[string, int](capacity)

	for i, key := range keys {
		cache.Set(key, data[i])
	}

	for _, key := range keys {
		fmt.Println(cache.Get(key))
	}

	// Output:
	// 0 false
	// 0 false
	// 0 false
	// 3 true
	// 4 true
	// 5 true
}

func TestCache_Set_One(t *testing.T) {
	tests := []struct {
		name  string
		key   int
		value int
		want  int
	}{
		{
			name:  "set one value",
			key:   367,
			value: 999,
			want:  999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := New[int, int](1)

			cache.Set(tt.key, tt.value)

			got, ok := cache.Get(tt.key)
			if !ok {
				t.Fatalf("value at the key %v must be presented at the cache", tt.key)
			}

			if got != tt.want {
				t.Fatalf("got value %v at the key %v, but want %v", got, tt.key, tt.want)
			}
		})
	}
}

func TestCache_Set_Series(t *testing.T) {
	tests := []struct {
		name  string
		key   int
		value int
		want  int
	}{
		{
			name:  "set one value",
			key:   367,
			value: 999,
			want:  999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := New[int, int](1)

			cache.Set(tt.key, tt.value)

			got, ok := cache.Get(tt.key)
			if !ok {
				t.Fatalf("value at the key %v must be presented at the cache", tt.key)
			}

			if got != tt.want {
				t.Fatalf("got value %v at the key %v, but want %v", got, tt.key, tt.want)
			}
		})
	}
}

func TestCache_Get_Many(t *testing.T) {
	tests := []struct {
		name         string
		capacity     int
		cachedKeys   []int
		cachedValues []int
		getKeys      []int
		wantValues   []int
		wantOks      []bool
	}{
		{
			name:         "values fits to the cache capacity and all values are found",
			capacity:     3,
			cachedKeys:   []int{1, 2, 3},
			cachedValues: []int{10, 20, 30},
			getKeys:      []int{1, 2, 3},
			wantValues:   []int{10, 20, 30},
			wantOks:      []bool{true, true, true},
		},
		{
			name:         "values does not fit to the cache capacity and only last values are found",
			capacity:     5,
			cachedKeys:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			cachedValues: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
			getKeys:      []int{1, 5, 7},
			wantValues:   []int{0, 50, 70},
			wantOks:      []bool{false, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mustFillCache(t, tt.capacity, tt.cachedKeys, tt.cachedValues)

			n := 0
			for i, k := range tt.getKeys {
				got, ok := cache.Get(k)
				if tt.wantOks[i] {
					if !ok {
						t.Fatalf("value at the key %v must be presented at the cache", k)
					}
				} else {
					if ok {
						t.Fatalf("value at the key %v must not be found at the cache", k)
					}
					n++
					continue
				}

				if want := tt.wantValues[i]; got != want {
					t.Fatalf("got value %v at the key %v, but want %v", got, k, want)
				}

				n++
			}

			if n != len(tt.wantValues) {
				t.Fatalf("not all keys were found: want %v, but found %v", len(tt.wantValues), n)
			}
		})
	}
}

func TestCache_Get_Many_After_Reset(t *testing.T) {
	tests := []struct {
		name         string
		capacity     int
		cachedKeys   []int
		cachedValues []int
		getKeys      []int
		wantValues   []int
		wantOks      []bool
	}{
		{
			name:         "values fits to the cache capacity and all values are found",
			capacity:     3,
			cachedKeys:   []int{1, 2, 3},
			cachedValues: []int{10, 20, 30},
			getKeys:      []int{1, 2, 3},
			wantValues:   []int{10, 20, 30},
			wantOks:      []bool{true, true, true},
		},
		{
			name:         "values does not fit to the cache capacity and only last values are found",
			capacity:     5,
			cachedKeys:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			cachedValues: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
			getKeys:      []int{1, 5, 7},
			wantValues:   []int{0, 50, 70},
			wantOks:      []bool{false, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mustFillCache(t, tt.capacity, tt.cachedKeys, tt.cachedValues)

			cache.Reset()

			cache = mustFillCache(t, tt.capacity, tt.cachedKeys, tt.cachedValues)

			n := 0
			for i, k := range tt.getKeys {
				got, ok := cache.Get(k)
				if tt.wantOks[i] {
					if !ok {
						t.Fatalf("value at the key %v must be presented at the cache", k)
					}
				} else {
					if ok {
						t.Fatalf("value at the key %v must not be found at the cache", k)
					}
					n++
					continue
				}

				if want := tt.wantValues[i]; got != want {
					t.Fatalf("got value %v at the key %v, but want %v", got, k, want)
				}

				n++
			}

			if n != len(tt.wantValues) {
				t.Fatalf("not all keys were found: want %v, but found %v", len(tt.wantValues), n)
			}
		})
	}
}

func TestCache_Reset(t *testing.T) {
	tests := []struct {
		name         string
		capacity     int
		cachedKeys   []int
		cachedValues []int
	}{
		{
			name:         "all values are not found after reset",
			capacity:     3,
			cachedKeys:   []int{1, 2, 3},
			cachedValues: []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := mustFillCache(t, tt.capacity, tt.cachedKeys, tt.cachedValues)

			cache.Reset()

			n := 0
			for _, k := range tt.cachedKeys {
				_, ok := cache.Get(k)
				if ok {
					t.Fatalf("value at the key %v must not be found at the cache", k)
				}
				n++
			}

			if n != len(tt.cachedKeys) {
				t.Fatalf("not all keys were tested: want %v, but found %v", len(tt.cachedKeys), n)
			}
		})
	}
}

var v int

func BenchmarkCache_Get_1024(b *testing.B) {
	cache := New[int, int](1024)

	for i := 0; i < 1024*4; i++ {
		cache.Set(i, rand.Int())
	}

	for i := 0; i < b.N; i++ {
		if value, ok := cache.Get(i); ok {
			v = value
		}
	}
}

func BenchmarkCache_Set_1024(b *testing.B) {
	cache := New[int, int](1024)

	for i := 0; i < b.N; i++ {
		cache.Set(i, i)
	}
}
