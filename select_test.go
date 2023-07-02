package bitvector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect_Empty(t *testing.T) {
	bv := NewBitvector(128)

	assert.Equal(t, 0, Select0Once(bv, 0))
	assert.Equal(t, 0, Select0Once(bv, 1))
	assert.Equal(t, 1, Select0Once(bv, 2))
	assert.Equal(t, 127, Select0Once(bv, 128))

	assert.Equal(t, 128, Select1Once(bv, 1))
	assert.Equal(t, 128, Select1Once(bv, 128))
}

func TestSelect_Sparse(t *testing.T) {
	bv := NewBitvector(128)
	bv.Set(1)
	bv.Set(64)

	assert.Equal(t, 0, Select0Once(bv, 1))
	assert.Equal(t, 2, Select0Once(bv, 2))
	assert.Equal(t, 71, Select0Once(bv, 70))
	assert.Equal(t, 127, Select0Once(bv, 128-2))
	assert.Equal(t, 128, Select0Once(bv, 128))

	assert.Equal(t, 1, Select1Once(bv, 1))
	assert.Equal(t, 64, Select1Once(bv, 2))
	assert.Equal(t, 128, Select1Once(bv, 3))
}

func BenchmarkSelectOnce_16384(b *testing.B) {
	size := 2 << 14
	bv := NewBitvector(size)

	b.ResetTimer()

	var index int
	for i := 0; i < b.N; i += 1 {
		index = Select0Once(bv, 2<<13)
	}
	b.StopTimer()

	assert.Equal(b, (2<<13)-1, index)
}

func TestInverse(t *testing.T) {
  bv := NewBitvector(64)
  bv.Set(12)
  bv.Set(14)
  bv.Set(1)
  bv.Set(0)

  s := Select0Once(bv, 6)
  r := Rank0Once(bv, s + 1)

  assert.Equal(t, 6, r)
}
