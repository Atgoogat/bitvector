package main

import (
	"fmt"

	"github.com/Atgoogat/bitvector"
)

func main() {
	bv := bitvector.NewBitvector(10)

	for i := 0; i < 10; i += 2 {
		bv.Set(i) // set every second bit
	}

	rank := bitvector.NewRank(bv)

	for i := 0; i <= 10; i += 1 {
		fmt.Println(rank.Rank0(i))
	}
}

// Output:
// 0
// 0
// 1
// 1
// 2
// 2
// 3
// 3
// 4
// 4
// 5
