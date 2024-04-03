package custommap

import (
	"sync"
	"sync/atomic"
)

type MapInterface interface {
	Set(string, any)        // set value into map
	Get(string) (any, bool) // get value from map
	Len() int               // length of map keys
}

func NewMap() MapInterface {

	log2BucketLength := uint32(2) // number of buckets = 4

	return &customMap{
		//header
		length: 0,
		salt:   0,
		lenBuckets: log2BucketLength, //ln2 length of buckets

		buckets:    createBuckets(1 << log2BucketLength),
		oldBuckets: nil,
	}
}

// without memory alignment
type customMap struct {
	length int32
	lenBuckets uint32
	salt uint8

	buckets    *buckets
	oldBuckets *buckets

	lock sync.Mutex
}

func (m *customMap) Len() int {
	return int(atomic.LoadInt32(&m.length))
}

func (m *customMap) Set(key string, value any) {
	
	hashedKey := m.hashKey(key)

	m.lock.Lock()
	defer m.lock.Unlock()
	
	b := m.getBucketByHash(hashedKey)
	_, ok := m.getValueFromBucket(key, b)
	if !ok {
		m.incrementMapLength()
	}

	m.setValueToBucket(key, value, b)
}

func (m *customMap) Get(key string) (any, bool) {
	
	hashedKey := m.hashKey(key)

	m.lock.Lock()
	defer m.lock.Unlock()

	bucket := m.getBucketByHash(hashedKey)

	return m.getValueFromBucket(key, bucket)
}
