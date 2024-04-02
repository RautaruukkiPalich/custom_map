package custommap

const (
	bucketSize = 8
)

type bucket struct {
	topHash  int
	length   uint
	keys     [bucketSize]string
	values   [bucketSize]any
}

func (b *bucket) increaseLength() {
	b.length++
}
