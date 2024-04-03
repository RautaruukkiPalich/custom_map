package custommap_simple

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
	extra  *bucket
	mu      sync.RWMutex
}

func (b *bucket) getLength() uint {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.length
}

func (b *bucket) getBucketAndIndexByKey(key string) (*bucket, int, bool) {
	
	b.mu.RLock()
	defer b.mu.RUnlock()

	for idx, k := range b.keys {
		if k == key {
			return b, idx, true
		}
	}
	if b.extra != nil {
		return b.extra.getBucketAndIndexByKey(key)
	}
	if b.length == bucketSize {
		return b.extendBucket(), 0, false
	}
	return b, int(b.getLength()), false
}

func (b *bucket) setKV(key string, value any, idx int, isExist bool) {

	b.mu.Lock()
	defer b.mu.Unlock()

	b.keys[b.length] = key
	b.values[b.length] = value

	if !isExist {
		b.length++
	}
}

func (b *bucket) getValue(key string) (any, bool){
	b.mu.RLock()
	defer b.mu.RUnlock()

	for i, bucketKey := range b.keys {
		if bucketKey == key {
			return b.values[i], true
		}
	}
	if b.extra != nil {
		return b.extra.getValue(key)
	}
	return nil, false
}

func (b *bucket) extendBucket() *bucket {
	b.extra = &bucket{}
	return b.extra
}