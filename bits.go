package immutable

import "math/bits"

func bitindex(bmap, bitpos uint) int {
	return bits.OnesCount(bmap & (bitpos - 1))
}
