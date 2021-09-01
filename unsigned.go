package varint

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"
)

var Unsigned VarintCodec = unsignedCodec{}

type unsignedCodec struct{}

// EncodedLen returns the number of bytes necessary to encode the given integer.
// The function will panic if the value is less than zero.
func (unsignedCodec) EncodedLen(i *big.Int) int {
	s := i.Sign()
	if i.Sign() < 0 {
		panic(errors.New("unsigned varints must be non-negative"))
	} else if s == 0 {
		return 1
	}

	bitLength := i.BitLen()
	q, r := bitLength/7, bitLength%7
	if r > 0 {
		return q + 1
	} else {
		return q
	}
}

func (codec unsignedCodec) EncodeToBytes(i *big.Int) []byte {
	l := codec.EncodedLen(i)
	buf := bytes.NewBuffer(make([]byte, 0, l))
	n, err := codec.Write(buf, i)
	if err != nil {
		panic(err)
	} else if n != l {
		panic(fmt.Errorf("bad length: expected %d, got %d", l, n))
	} else {
		return buf.Bytes()
	}
}

func (codec unsignedCodec) DecodeBytes(data []byte) (*big.Int, error) {
	i := big.NewInt(0)
	buf := bytes.NewBuffer(make([]byte, 0, len(data)))
	if _, err := buf.Write(data); err != nil {
		panic(err)
	} else if _, err := codec.Read(buf, i); err != nil {
		return nil, err
	} else {
		return i, nil
	}
}

func (unsignedCodec) Write(w io.ByteWriter, i *big.Int) (n int, err error) {
	if i.Sign() < 0 {
		panic(errors.New("unsigned varints must be non-negative"))
	}

	// allocate a new big.Int so that we can mutate it
	i = big.NewInt(0).Set(i)

	for j := big.NewInt(0); limit.Cmp(i) < 0; n++ {
		j = j.And(i, limit)
		j = j.SetBit(j, 7, 1)
		err = w.WriteByte(byte(j.Uint64()))
		if err != nil {
			return
		}

		i = i.Rsh(i, 7)
	}

	err = w.WriteByte(byte(i.Uint64()))
	if err != nil {
		return
	}

	n++
	return
}

func (unsignedCodec) Read(r io.ByteReader, i *big.Int) (n int, err error) {
	i = i.SetUint64(0)
	delta := big.NewInt(0)
	for b := byte(0); ; n++ {
		b, err = r.ReadByte()
		if err != nil {
			return
		}

		delta = delta.SetUint64(uint64(b))
		delta = delta.SetBit(delta, 7, 0)

		if n > 0 {
			delta = delta.Lsh(delta, uint(n)*7)
		}

		i = i.Add(i, delta)

		if b < 0x80 {
			n++
			break
		}
	}

	return
}
