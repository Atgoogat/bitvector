# Bitvector library written in go

This library contains a simple bitvector implementation and efficient rank queries.

## Complexity

|            | Time (Query) | Time (Construction) | Space |
| ---------- | ------------ | ------------------- | ----- |
| Bitvector  | O(1)         | O(n)                | o(n)  |
| Rank.Rank0 | O(1)         | O(n)                | o(n)  |
| Rank.Rank1 | O(1)         | O(n)                | o(n)  |
| Rank0Once  | O(n)         | -                   | -     |
| Rank1Once  | O(n)         | -                   | -     |

## Example

```
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
```
