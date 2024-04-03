package custommap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkSetAndGetElemsWithRaces(b *testing.B) {
	m := NewMap()
	iter := b.N //100_000
	randoms := 1000

	wg := &sync.WaitGroup{}
	wg.Add(iter*2)
	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Set(strconv.Itoa(rand.Intn(randoms)), i)
			wg.Done()
		}(i)
		go func(i int) {
			m.Get(strconv.Itoa(rand.Intn(randoms)))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkSyncMapSetAndGetElemsWithRaces(b *testing.B) {
	m := sync.Map{}
	iter := b.N //100_000
	randoms := 1000

	wg := &sync.WaitGroup{}
	wg.Add(iter*2)
	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Store(strconv.Itoa(rand.Intn(randoms)), i)
			wg.Done()
		}(i)
		go func(i int) {
			m.Load(strconv.Itoa(rand.Intn(randoms)))
			wg.Done()
		}(i)
	}
	wg.Wait()
}


func BenchmarkSet(b *testing.B) {
	m := NewMap()
	iter := b.N
	wg := &sync.WaitGroup{}
	wg.Add(iter)

	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkSyncMapSet(b *testing.B) {
	m := sync.Map{}
	iter := b.N
	wg := &sync.WaitGroup{}
	wg.Add(iter)

	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Store(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}


func BenchmarkSetElemsCheckValues(b *testing.B) {
	m := NewMap()
	iter := b.N
	wg := &sync.WaitGroup{}
	wg.Add(iter)

	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Set(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < iter; i++ {
		m.Get(strconv.Itoa(i))
	}

	for i := iter; i < iter+iter; i++ {
		m.Get(strconv.Itoa(i))
	}
}

func BenchmarkSyncMapSetElemsCheckValues(b *testing.B) {
	m := sync.Map{}
	iter := b.N
	wg := &sync.WaitGroup{}
	wg.Add(iter)

	for i := 0; i < iter; i++ {
		go func(i int) {
			m.Store(strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < iter; i++ {
		m.Load(strconv.Itoa(i))
	}

	for i := iter; i < iter+iter; i++ {
		m.Load(strconv.Itoa(i))
	}
}