package bitvector

import "math/bits"

// Return the index where the ith-zero is.
// When there is not ith-zero bv.Len() will be returned.
// Complexity: O(n)
func Select0Once(bv ReadonlyBitvector, ith int) (index int) {
	remaining := ith
	for wordIndex := 0; wordIndex < bv.LenWords(); wordIndex += 1 {
		zeros := 64 - bits.OnesCount64(bv.GetWord(wordIndex))
		if remaining <= zeros {
			break
		}
		remaining -= zeros
		index += wordSize
	}

	for index < bv.Len() && remaining > 0 {
		if !bv.Get(index) {
			remaining -= 1
		}
		index += 1
	}
	if index > 0 {
		index -= 1
	}

	if remaining > 0 {
		return bv.Len()
	}
	return
}

// Return the index where the ith-one is.
// When there is not ith-one bv.Len() will be returned.
// Complexity: O(n)
func Select1Once(bv ReadonlyBitvector, ith int) (index int) {
	remaining := ith
	for wordIndex := 0; wordIndex < bv.LenWords(); wordIndex += 1 {
		ones := bits.OnesCount64(bv.GetWord(wordIndex))
		if remaining <= ones {
			break
		}
		remaining -= ones
		index += wordSize
	}

	for index < bv.Len() && remaining > 0 {
		if bv.Get(index) {
			remaining -= 1
		}
		index += 1
	}
	if index > 0 {
		index -= 1
	}

	if remaining > 0 {
		return bv.Len()
	}
	return
}
