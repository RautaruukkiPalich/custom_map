package custommap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMap(t *testing.T) {
	m := NewMap()
	assert.NotNil(t, m)
}

func TestSetElems(t *testing.T) {
	m := NewMap()
	m.Set("1", 1)
	m.Set("2", 2)
	assert.Equal(t, 2, m.Len())
}

func TestLenElems(t *testing.T) {
	m := NewMap()
	m.Set("1", 1)
	m.Set("2", 2)
	assert.Equal(t, 2, m.Len())

	m.Set("2", 4)
	assert.Equal(t, 2, m.Len())
}

func TestGetElems(t *testing.T) {
	m := NewMap()
	m.Set("1", 1)
	m.Set("2", []string{"1", "2", "3"})

	val, ok := m.Get("1")
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, val)

	val, ok = m.Get("2")
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"1", "2", "3"}, val)

	val, ok = m.Get("3")
	assert.Equal(t, false, ok)
	assert.Equal(t, nil, val)
}