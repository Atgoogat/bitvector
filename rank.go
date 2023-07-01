package bitvector

import "math/bits"

const (
	blockSize      = 512
	superblockSize = 4096

	wordsPerBlock       = blockSize / wordSize
	blocksPerSuperblock = superblockSize / blockSize
)

// ReadonlyInterface for Bitvector
type ReadonlyBitvector interface {
	Get(index int) bool
	GetWord(index int) uint64
	Len() int
	LenWords() int
}

// Datastructure to allow for fast rank queries
type Rank struct {
	// underlying bitvector
	bv ReadonlyBitvector
	// zeros from beginning of containing super block
	blocks []int16
	// zeros from beginning of bitvector
	superblocks []int
}

// Build rank datastructure for this bitvector
//
// Caution: Any changes to the underlying bitvector lead to undefined behaviour regarding the rank queries
// Complexity: O(n) (n == bv.Len())
func NewRank(bv ReadonlyBitvector) Rank {
	blocks, superblocks := buildDataStructure(bv)
	return Rank{
		bv:          bv,
		blocks:      blocks,
		superblocks: superblocks,
	}
}

// Returns how many zeros there are before index.
// For index = 0 => 0
// Complexity: O(1)
func (r Rank) Rank0(index int) (rank int) {
	containingBlock := index / blockSize
	containingSuperblock := index / superblockSize

	if containingSuperblock > 0 {
		rank += r.superblocks[containingSuperblock-1]
	}
	if containingBlock%blocksPerSuperblock > 0 {
		rank += int(r.blocks[containingBlock-1])
	}

	wordIndex := containingBlock * (blockSize / wordSize)
	for ; wordIndex < index/wordSize; wordIndex += 1 {
		rank += 64 - bits.OnesCount64(r.bv.GetWord(wordIndex))
	}

	index -= 1
	for (index+wordSize)%wordSize != wordSize-1 {
		if !r.bv.Get(index) {
			rank += 1
		}
		index -= 1
	}

	return
}

// Return how many ones there are before index.
// For index = 0 => 0
// Complexity: O(1)
func (r Rank) Rank1(index int) (rank int) {
	return index - r.Rank0(index)
}

// Return how many zeros there are before index.
// For index = 0 => 0
// Complexity: O(n) (n = index)
// Consider using Rank.Rank0 if you do multiple requests
func Rank0Once(bv ReadonlyBitvector, index int) int {
	rank := 0
	for i := 0; i < index/wordSize; i += 1 {
		rank += 64 - bits.OnesCount64(bv.GetWord(i))
	}
	index -= 1
	for (index+wordSize)%wordSize != wordSize-1 {
		if !bv.Get(index) {
			rank += 1
		}
		index -= 1
	}
	return rank
}

// Return how many ones there are before index.
// For index = 0 => 0
// Complexity: O(n) (n = index)
// Consider using Rank.Rank1 if you do multiple requests
func Rank1Once(bv ReadonlyBitvector, index int) int {
	return index - Rank0Once(bv, index)
}

func buildDataStructure(bv ReadonlyBitvector) (blocks []int16, superblocks []int) {
	superblocks = make([]int, bv.Len()/superblockSize)
	blocks = make([]int16, bv.Len()/blockSize)

	superblockIndex := 0
	wordIndex := 0
	totalSum := 0
	sum := 0
	for blockIndex := 0; blockIndex < len(blocks); blockIndex += 1 {
		for i := 0; i < wordsPerBlock; i += 1 {
			sum += 64 - bits.OnesCount64(bv.GetWord(wordIndex))
			wordIndex += 1
		}
		blocks[blockIndex] = int16(sum)
		if blockIndex%blocksPerSuperblock == blocksPerSuperblock-1 {
			totalSum += sum
			sum = 0
			superblocks[superblockIndex] = totalSum
			superblockIndex += 1
		}
	}
	return
}
