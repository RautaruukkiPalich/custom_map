package custommap_simple

// import (
// 	"math/rand"
// 	"strconv"
// 	"sync"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateMap(t *testing.T) {
// 	m := NewMap()
// 	assert.NotNil(t, m)
// }

// func TestSetElems(t *testing.T) {
// 	m := NewMap()
// 	m.Set("1", 1)
// 	m.Set("2", 2)
// 	assert.Equal(t, 2, m.Len())
// }

// func TestLenElems(t *testing.T) {
// 	m := NewMap()
// 	m.Set("1", 1)
// 	m.Set("2", 2)
// 	assert.Equal(t, 2, m.Len())

// 	m.Set("2", 4)
// 	assert.Equal(t, 2, m.Len())
// }

// func TestGetElems(t *testing.T) {
// 	m := NewMap()
// 	m.Set("1", 1)
// 	m.Set("2", []string{"1", "2", "3"})

// 	val, ok := m.Get("1")
// 	assert.Equal(t, true, ok)
// 	assert.Equal(t, 1, val)

// 	val, ok = m.Get("2")
// 	assert.Equal(t, true, ok)
// 	assert.Equal(t, []string{"1", "2", "3"}, val)

// 	val, ok = m.Get("3")
// 	assert.Equal(t, false, ok)
// 	assert.Equal(t, nil, val)
// }

// func TestSetElemWithRaces(t *testing.T) {
// 	m := NewMap()
// 	iter := 10000
// 	wg := &sync.WaitGroup{}
// 	wg.Add(iter)
// 	for i := 0; i < iter; i++ {
// 		go func() {
// 			m.Set("1", 2)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	assert.Equal(t, 1, m.Len())
// }


// func TestSetAndGet10kElemsWithRaces(t *testing.T) {
// 	m := NewMap()
// 	iter := 10_000
// 	randoms := 100

// 	wg := &sync.WaitGroup{}
// 	wg.Add(iter*2)
// 	for i := 0; i < iter; i++ {
// 		go func(i int) {
// 			m.Set(strconv.Itoa(rand.Intn(randoms)), i)
// 			wg.Done()
// 		}(i)
// 		go func(i int) {
// 			m.Get(strconv.Itoa(rand.Intn(randoms)))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()

// 	assert.Equal(t, randoms, m.Len())
// }


// func TestSet10kElemsCheckValues(t *testing.T) {
// 	m := NewMap()
// 	iter := 10_000
// 	wg := &sync.WaitGroup{}
// 	wg.Add(iter)
// 	for i := 0; i < iter; i++ {
// 		go func(i int) {
// 			m.Set(strconv.Itoa(i), i)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// 	assert.Equal(t, iter, m.Len())

// 	for i := 0; i < iter; i++ {
// 		res, ok := m.Get(strconv.Itoa(i))
// 		assert.Equal(t, i, res)
// 		assert.Equal(t, true, ok)
// 	}

// 	for i := iter; i < iter+iter; i++ {
// 		res, ok := m.Get(strconv.Itoa(i))
// 		assert.Equal(t, nil, res)
// 		assert.Equal(t, false, ok)
// 	}
// }
