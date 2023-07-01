package bitvector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBit(t *testing.T) {
	bv := NewBitvector(64)

	assert.False(t, bv.Get(0))
	bv.Set(0)
	assert.True(t, bv.Get(0))
}

func TestSetBit_3(t *testing.T) {
	bv := NewBitvector(64)

	assert.False(t, bv.Get(2))
	bv.Set(2)
	assert.True(t, bv.Get(2))
}

func TestSetBitInNextWord(t *testing.T) {
	bv := NewBitvector(65)
	bv.Set(64)
	assert.True(t, bv.Get(64))
}

func TestClearBit(t *testing.T) {
	bv := NewBitvector(64)
	bv.Set(3)
	assert.True(t, bv.Get(3))
	bv.Clear(3)
	assert.False(t, bv.Get(3))
}

func TestClearAlreadyClearedBit(t *testing.T) {
	bv := NewBitvector(64)
	assert.False(t, bv.Get(3))
	bv.Clear(3)
	assert.False(t, bv.Get(3))
}
