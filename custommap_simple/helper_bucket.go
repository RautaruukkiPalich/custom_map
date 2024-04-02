package custommap

import (
	"sync"
)

const (
	bucketSize = 8
)

type bucket struct {
	topHash uint32
	length  uint
	keys    [bucketSize]string
	values  [bucketSize]any
	mu      sync.RWMutex
}

func (b *bucket) increaseLength() {
	b.length++
}

func (b *bucket) getLength() uint {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.length
}

func (b *bucket) setKV(key string, value any) {

	b.mu.Lock()
	defer b.mu.Unlock()

	for idx, k := range b.keys {
		if k == key {
			i := idx
			b.values[i] = value
			return 
		}
	}

	b.keys[b.length] = key
	b.values[b.length] = value

	b.increaseLength()
}

func (b *bucket) getValue(key string) (any, bool){
	b.mu.RLock()
	defer b.mu.RUnlock()

	for i, bucketKey := range b.keys {
		if bucketKey == key {
			return b.values[i], true
		}
	}
	return nil, false
}

