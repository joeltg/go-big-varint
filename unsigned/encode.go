package unsigned

import (
	"math/big"
)

// Encode a big integer into a byte slice, returning the number of bytes written.
// The function will panic if the value is less than zero.
func Encode(data []byte, i *big.Int) int {
	if i.Sign() < 0 {
		panic(ErrLessThanZero)
	}

	// allocate a new big.Int so that we can mutate it
	i = big.NewInt(0).Set(i)

	offset := 0
	j := big.NewInt(0)
	for ; limit.Cmp(i) < 0; offset++ {
		j = j.And(i, pad)
		j = j.Or(j, msb)
		data[offset] = byte(j.Uint64())
		i = i.Div(i, msb)
	}

	c := big.NewInt(0)
	for ; c.And(i, notRest).Sign() != 0; offset++ {
		j = j.And(i, pad)
		j = j.Or(j, msb)
		data[offset] = byte(j.Uint64())
		i = i.Rsh(i, 7)
	}

	data[offset] = byte(i.Uint64())
	return offset + 1
}

// EncodingLength returns the number of bytes necessary to encode the given integer.
// The function will panic if the value is less than zero.
func EncodingLength(i *big.Int) int {
	if i.Sign() < 0 {
		panic(ErrLessThanZero)
	}

	bitLength := i.BitLen()
	q, r := bitLength/7, bitLength%7
	if r > 0 {
		return q + 1
	} else {
		return q
	}
}
