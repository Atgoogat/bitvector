package bitvector

const (
	wordSize = 64
)

type Bitvector struct {
	data []uint64
	size int
}

// Create a new fixed size bitvector
func NewBitvector(size int) Bitvector {
	return Bitvector{
		size: size,
		data: make([]uint64, (size+wordSize-1)/wordSize),
	}
}

// Length (in bits)
func (bv Bitvector) Len() int {
	return bv.size
}

// Set bit at index
func (bv *Bitvector) Set(index int) {
	bv.data[bv.wordIndex(index)] |= 1 << bv.innerWordIndex(index)
}

// Clear bit at index
func (bv *Bitvector) Clear(index int) {
	bv.data[bv.wordIndex(index)] &^= (1 << bv.innerWordIndex(index))
}

// Get bit at index
func (bv Bitvector) Get(index int) bool {
	return (bv.data[bv.wordIndex(index)] & (1 << bv.innerWordIndex(index))) > 0
}

// Get word at index (word index)
func (bv Bitvector) GetWord(index int) uint64 {
	return bv.data[index]
}

// Length of underlying words
func (bv Bitvector) LenWords() int {
	return len(bv.data)
}

func (bv Bitvector) wordIndex(index int) int {
	return index / wordSize
}

func (bv Bitvector) innerWordIndex(index int) int {
	return index % wordSize
}

var _ ReadonlyBitvector = (*Bitvector)(nil)
