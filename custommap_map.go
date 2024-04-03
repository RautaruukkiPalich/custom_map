package custommap

import (
	"hash/fnv"
	"sync/atomic"
)

const (
	loadFactor = 6
)

func (m *customMap) incrementMapLength() {
	atomic.AddInt32(&m.length, 1)
}

func (m *customMap) incrementBucketLength() {
	atomic.AddUint32(&m.lenBuckets, 1)
}

func (m *customMap) getBucketLength() uint32 {
	return 1 << atomic.LoadUint32(&m.lenBuckets)
}

func (m *customMap) hashKey(key *string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(*key))
	h.Write([]byte{m.salt})
	return h.Sum32()
}

func (m *customMap) getBucket(id uint32) *bucket{
	return m.buckets.list[id]
}

func (m *customMap) getBucketByHash(hashedKey uint32) *bucket {

	b := m.getBucket(hashedKey % m.getBucketLength())

	if b.getLength() > loadFactor {
		m.scaleBucketCapacity()
		b = m.getBucket(hashedKey % m.getBucketLength())
	}

	return b
}

func (m *customMap) setValueToBucket(key *string, value *any, b *bucket) {
	b.setKV(key, value)
}

func (m *customMap) getValueFromBucket(key *string, b *bucket) (any, bool) {
	return b.getValue(key)
}

func (m *customMap) scaleBucketCapacity() {

	m.oldBuckets = m.buckets
	m.incrementBucketLength()
	m.buckets = createBuckets(int(m.getBucketLength()))
	m.migrateData()
}

func (m *customMap) migrateData() {
	// перезаписать данные из старых бакетов
	for _, b := range m.oldBuckets.list {
		for i := 0; i < int(b.getLength()); i++ {
			hashed := m.hashKey(&b.keys[i])
			bct := m.getBucketByHash(hashed)
			bct.setKV(&b.keys[i], &b.values[i])
		}
	}

	// удаляем старые бакеты
	m.oldBuckets = nil
}

