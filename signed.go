package varint

import (
	"math/big"
)

var Signed VarintCodec = signedCodec{}

type signedCodec struct{}

// EncodingLength returns the number of bytes necessary to encode the given big integer
func (signedCodec) EncodingLength(i *big.Int) int {
	j := big.NewInt(0)
	j = j.Mul(i, two)
	if i.Sign() < 0 {
		j = j.Not(j)
	}
	return Unsigned.EncodingLength(j)
}

// Encode a big integer into a byte slice, returning the number of bytes written
func (signedCodec) Encode(data []byte, i *big.Int) int {
	j := big.NewInt(0)
	j = j.Mul(i, two)
	if i.Sign() < 0 {
		j = j.Not(j)
	}
	return Unsigned.Encode(data, j)
}

// Decode a big integer from a byte slice, returning the result and the number of bytes read
func (signedCodec) Decode(data []byte) (*big.Int, int) {
	i, l := Unsigned.Decode(data)
	if i.Bit(0) == 1 {
		i = i.Not(i)
	}
	i = i.Div(i, two)
	return i, l
}
