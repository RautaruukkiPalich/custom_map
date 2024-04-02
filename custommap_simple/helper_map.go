package custommap

import (
	"fmt"
	"hash/fnv"
)

func (m *customMap) incrementMapLength() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.length++
}

func (m *customMap) incrementBucketLength() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lenBuckets++
}

func (m *customMap) hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (m *customMap) createBucket(hashedKey uint32) *bucket {

	b := &bucket{topHash: hashedKey}

	defer m.incrementBucketLength()

	m.mu.Lock()
	defer m.mu.Unlock()
	

	m.buckets = append(m.buckets, b)
	return b
}

func (m *customMap) getBucket(hashedKey uint32) *bucket {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, bucket := range m.buckets {
		if bucket.topHash == hashedKey {
			return bucket
		}
	}
	return nil
}

func (m *customMap) getBucketByHash(hashedKey uint32) *bucket {
	b := m.getBucket(hashedKey)
	if b != nil {
		return b
	}
	return	m.createBucket(hashedKey)
}

func (m *customMap) setValueToBucket(key string, value any, b *bucket) {
	
	if b.getLength() >= bucketSize {
		//TODO: handle this
		m.evacuateBucket()
	}

	b.setKV(key, value)
}

func (m *customMap) getValueFromBucket(key string, b *bucket) (any, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return b.getValue(key)
}

func (m *customMap) evacuateBucket() {
	fmt.Println("evacuate bucket")
}
