package varint

import (
	"errors"
	"math/big"
)

var Unsigned VarintCodec = unsignedCodec{}

type unsignedCodec struct{}

// EncodingLength returns the number of bytes necessary to encode the given integer.
// The function will panic if the value is less than zero.
func (unsignedCodec) EncodingLength(i *big.Int) int {
	if i.Sign() < 0 {
		panic(errors.New("unsigned varints must be non-negative"))
	}

	bitLength := i.BitLen()
	q, r := bitLength/7, bitLength%7
	if r > 0 {
		return q + 1
	} else {
		return q
	}
}

// Encode a big integer into a byte slice, returning the number of bytes written.
// The function will panic if the value is less than zero.
func (codec unsignedCodec) Encode(data []byte, i *big.Int) int {
	if i.Sign() < 0 {
		panic(errors.New("unsigned varints must be non-negative"))
	} else if len(data) < codec.EncodingLength(i) {
		panic(errors.New("the provided byte slice is too small to encode the provided integer"))
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

// Decode a big integer from a byte slice, returning the result and the number of bytes read.
// The result is guaranteed to be non-negative.
func (unsignedCodec) Decode(data []byte) (*big.Int, int) {
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
