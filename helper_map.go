package custommap

import "hash/fnv"

const (
	
)

func (m *customMap) incrementMapLength() {
	m.length++
}

func (m *customMap) incrementBucketLength() {
	m.lenBuckets++
}

func (m *customMap) hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (m *customMap) getIdxOrCreateBucketFromHash(hashedKey uint32) int {
	topHash := int(hashedKey)

	for bucketIdx, bucket := range m.buckets {
		if bucket.topHash == topHash {
			return bucketIdx
		}
	}

	bucketIdx := m.lenBuckets 
	m.buckets = append(m.buckets, &bucket{topHash: topHash})

	m.incrementBucketLength()

	return int(bucketIdx)
}

func (m *customMap) setValueToBucketByIdx(key string, value any, bucketIdx int) {
	m.buckets[bucketIdx].increaseLength()

	current_pos := m.buckets[bucketIdx].length

	if current_pos > bucketSize {
		//TODO: handle this
		m.evacuateBucket()
	}
	m.buckets[bucketIdx].keys[current_pos] = key
	m.buckets[bucketIdx].values[current_pos] = value
}

func (m *customMap) getValueFromBucketByBucketIdx(key string, bucketIdx int) (any, bool) {
	for i, bucketKey := range m.buckets[bucketIdx].keys {
		if key == bucketKey {
			return m.buckets[bucketIdx].values[i], true
		}
	}
	return nil, false
}

func (m *customMap) evacuateBucket() {}
