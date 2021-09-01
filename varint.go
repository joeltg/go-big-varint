package varint

import (
	"io"
	"math/big"
)

type VarintCodec interface {
	EncodedLen(i *big.Int) int
	EncodeToBytes(i *big.Int) []byte
	DecodeBytes(data []byte) (*big.Int, error)
	Write(w io.ByteWriter, i *big.Int) (int, error)
	Read(r io.ByteReader, i *big.Int) (int, error)
}
