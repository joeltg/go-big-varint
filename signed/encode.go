package signed

import (
	"math/big"

	"github.com/joeltg/go-big-varint/unsigned"
)

// Encode a big integer into a byte slice, returning the number of bytes written
func Encode(data []byte, i *big.Int) int {
	j := big.NewInt(0)
	j = j.Mul(i, two)
	if i.Sign() < 0 {
		j = j.Not(j)
	}
	return unsigned.Encode(data, j)
}

// EncodingLength returns the number of bytes necessary to encode the given big integer
func EncodingLength(i *big.Int) int {
	j := big.NewInt(0)
	j = j.Mul(i, two)
	if i.Sign() < 0 {
		j = j.Not(j)
	}
	return unsigned.EncodingLength(j)
}
