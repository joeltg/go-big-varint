package varint

import (
	"io"
	"math/big"
)

var Signed VarintCodec = signedCodec{}

type signedCodec struct{}

// signedToUnsigned allocates a new *big.Int
func signedToUnsigned(i *big.Int) *big.Int {
	j := big.NewInt(0)
	j = j.Mul(i, two)
	if i.Sign() < 0 {
		j = j.Not(j)
	}
	return j
}

// unsignedToSigned DOES NOT allocate a new *big.Int
func unsignedToSigned(i *big.Int) *big.Int {
	if i.Bit(0) == 1 {
		i = i.Not(i)
	}

	return i.Div(i, two)
}

// EncodedLen returns the number of bytes necessary to encode the given big integer
func (signedCodec) EncodedLen(i *big.Int) int {
	return Unsigned.EncodedLen(signedToUnsigned(i))
}

// EncodeToBytes encodes a big integer into a byte slice
func (signedCodec) EncodeToBytes(i *big.Int) []byte {
	return Unsigned.EncodeToBytes(signedToUnsigned(i))
}

// DecodeBytes decodes a big integer from a byte slice
func (signedCodec) DecodeBytes(data []byte) (i *big.Int, err error) {
	i, err = Unsigned.DecodeBytes(data)
	_ = unsignedToSigned(i)
	return
}

// Write encodes a *big.Int to the Encoder's io.ByteWriter
func (signedCodec) Write(w io.ByteWriter, i *big.Int) (int, error) {
	return Unsigned.Write(w, signedToUnsigned(i))
}

// Read decodes a signed varint from the Decoder's io.ByteReader
// and stores the result in the provided *big.Int.
func (signedCodec) Read(r io.ByteReader, i *big.Int) (int, error) {
	n, err := Unsigned.Read(r, i)
	_ = unsignedToSigned(i)
	return n, err
}
