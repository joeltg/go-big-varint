package unsigned

import (
	"errors"
	"math/big"
)

// Decode a big integer from a byte slice, returning the result and the number of bytes read.
// The result is guaranteed to be non-negative.
func Decode(data []byte) (*big.Int, int) {
	offset := 0
	i, shift := big.NewInt(0), uint(0)
	delta, c := big.NewInt(0), big.NewInt(0)
	for {
		if offset < len(data) {
			b := data[offset]
			delta = delta.SetUint64(uint64(b))
			delta = delta.And(delta, rest)
			if shift < 28 {
				delta = delta.Lsh(delta, shift)
			} else {
				c = c.SetUint64(1)
				c = c.Lsh(c, shift)
				delta = delta.Mul(delta, c)
			}
			i = i.Add(i, delta)
			shift += 7
			offset++
			if b < 0x80 {
				break
			} else {
				continue
			}
		} else {
			panic(errors.New("out of range"))
		}
	}

	return i, offset
}
