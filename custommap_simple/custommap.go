package custommap_simple

import (
	"sync"
)

type MapInterface interface {
	Set(string, any)        // set value into map
	Get(string) (any, bool) // get value from map
	Len() int               // length of map keys
}

func NewMap() MapInterface {
	return &customMap{
		length: 0,
		// salt:   0,

		lenBuckets: 0,
		buckets:    make([]*bucket, 0, 10),
		// oldBuckets: []*bucket{},
	}
}

// without memory alignment
type customMap struct {
	length int

	lenBuckets uint8
	buckets    []*bucket

	mu sync.RWMutex
}

func (m *customMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.length
}

func (m *customMap) Set(key string, value any) {

	// create hashKey for the key
	hashedKey := m.hashKey(key)

	// looking for bucket
	bucket := m.getBucketByHash(hashedKey)

	// if key is not unique save val to bucket
	_, ok := m.getValueFromBucket(key, bucket)

	// if key is unique -> increase counter
	if !ok {
		m.incrementMapLength()
	}

	// save key and value to bucket
	m.setValueToBucket(key, value, bucket)
}

func (m *customMap) Get(key string) (any, bool) {
	// create hashKey for the key
	hashedKey := m.hashKey(key)

	// looking for bucket
	bucket := m.getBucketByHash(hashedKey)

	return m.getValueFromBucket(key, bucket)
}
