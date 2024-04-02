package custommap

type MapInterface interface {
	Set(string, any)        // set value into map
	Get(string) (any, bool) // get value from map
	Len() int               // length of map keys
}

func NewMap() MapInterface {
	return &customMap{
		length: 0,
		salt:   0,

		lenBuckets: 0,
		buckets:    []*bucket{},
		oldBuckets: []*bucket{},
	}
}

// without memory alignment
type customMap struct {
	length uint
	salt   uint32

	lenBuckets uint8
	buckets    []*bucket
	oldBuckets []*bucket
}

func (m *customMap) Len() int {
	return int(m.length)
}

func (m *customMap) Set(key string, value any) {
	// create hashKey for the key
	hashedKey := m.hashKey(key)

	// looking for bucket
	bucketIdx := m.getIdxOrCreateBucketFromHash(hashedKey)

	// if key is not unique save val to bucket
	_, ok := m.getValueFromBucketByBucketIdx(key, bucketIdx)

	// if key is unique -> increase counter
	if !ok {
		m.incrementMapLength()
	}

	// save key and value to bucket
	m.setValueToBucketByIdx(key, value, bucketIdx)
}

func (m *customMap) Get(key string) (any, bool) {
	// create hashKey for the key
	hashedKey := m.hashKey(key)

	// looking for bucket
	bucketIdx := m.getIdxOrCreateBucketFromHash(hashedKey)

	return m.getValueFromBucketByBucketIdx(key, bucketIdx)
}
