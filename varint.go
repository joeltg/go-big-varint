package varint

import (
	"math/big"
)

type VarintCodec interface {
	EncodingLength(i *big.Int) int
	Encode(data []byte, i *big.Int) int
	Decode(data []byte) (*big.Int, int)
}
