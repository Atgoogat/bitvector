package bitvector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank_Empty(t *testing.T) {
	size := 4096
	bv := NewBitvector(size)

	rank := NewRank(bv)

	for i := 0; i <= size; i += 1 {
		assert.Equal(t, i, rank.Rank0(i))
		assert.Equal(t, 0, rank.Rank1(i))
	}
}

func TestRank_Sparse(t *testing.T) {
	size := 4096
	bv := NewBitvector(size)
	bv.Set(1)

	rank := NewRank(bv)

	assert.Equal(t, 0, rank.Rank0(0))
	assert.Equal(t, 1, rank.Rank0(1))
	for i := 2; i <= size; i += 1 {
		assert.Equal(t, i-1, rank.Rank0(i))
		assert.Equal(t, 1, rank.Rank1(i))
	}
}

func TestRank_Sparse_8193(t *testing.T) {
	size := 8193
	bv := NewBitvector(size)
	bv.Set(1)
	bv.Set(4095)
	bv.Set(8192)

	rank := NewRank(bv)

	assert.Equal(t, 4096-2, rank.Rank0(4096))
	assert.Equal(t, 8193-3, rank.Rank0(8193))
}

func TestRankOnce_Empty(t *testing.T) {
	size := 4096
	bv := NewBitvector(size)

	for i := 0; i <= size; i += 1 {
		assert.Equal(t, i, Rank0Once(bv, i))
		assert.Equal(t, 0, Rank1Once(bv, i))
	}
}

func TestRankOnce_Sparse(t *testing.T) {
	size := 4096
	bv := NewBitvector(size)
	bv.Set(1)

	assert.Equal(t, 0, Rank0Once(bv, 0))
	assert.Equal(t, 1, Rank0Once(bv, 1))
	for i := 2; i <= size; i += 1 {
		assert.Equal(t, i-1, Rank0Once(bv, i))
		assert.Equal(t, 1, Rank1Once(bv, i))
	}
}

func BenchmarkRank_16384(b *testing.B) {
	size := 2 << 14
	bv := NewBitvector(size)
	rank := NewRank(bv)
	b.ResetTimer()

	for i := 0; i < b.N; i += 1 {
		_ = rank.Rank0(2 << 13)
	}
}

func BenchmarkRankOnce_16384(b *testing.B) {
	size := 2 << 14
	bv := NewBitvector(size)

	b.ResetTimer()

	for i := 0; i < b.N; i += 1 {
		_ = Rank0Once(bv, 2<<13)
	}
}
