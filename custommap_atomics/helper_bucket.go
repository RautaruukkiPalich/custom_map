package custommap

import (
	"sync"
	"sync/atomic"
)

const (
	bucketSize = 8
)

type buckets struct {
	list []*bucket
}

func createBuckets(count int) *buckets {

	buckets := &buckets{
		list: make([]*bucket, count),
	}
	for i := 0; i < count; i++ {
		idx := i
		buckets.list[idx] = &bucket{topHash: uint32(idx)}
	}

	return buckets
}

type bucket struct {
	topHash uint32
	length  uint32
	keys 	[bucketSize]string
	values	[bucketSize]any
	mu      sync.RWMutex
}

func (b *bucket) getLength() uint32 {
	return atomic.LoadUint32(&b.length)
}

func (b *bucket) increaseLength() {
	atomic.AddUint32(&b.length, 1)
}

func (b *bucket) setKV(key string, value any) {

	b.mu.Lock()
	defer b.mu.Unlock()

	idx, isExist := b.getIndex(key)
	if !isExist {
		b.keys[idx] = key
		b.increaseLength()
	}
	b.values[idx] = value
}

func (b *bucket) getValue(key string) (any, bool) {
	idx, isExist := b.getIndex(key)
	if !isExist {
		return nil, false
	}
	return b.values[idx], true
}

func (b *bucket) getIndex(key string) (uint32, bool) {
	for i, bKey := range b.keys {
		if bKey == key {
			return uint32(i), true
		}
	}
	return b.getLength(), false
}
