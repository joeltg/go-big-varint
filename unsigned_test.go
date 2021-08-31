package varint

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
	"math/rand"
	"testing"
)

func TestEncodeUnsignedZero(t *testing.T) {
	i := big.NewInt(0)
	data := make([]byte, 10)
	l := Unsigned.Encode(data, i)
	if l != 1 {
		log.Println(data, i, l)
		t.Fatal("bad uint64 encoding length")
	}

	if bytes.Equal(data, make([]byte, 10)) {
		return
	} else {
		log.Println(data, i, l)
		t.Fatal("bad uint64 encoding result")
	}
}

func TestDecodeUnsignedZero(t *testing.T) {
	data := make([]byte, 10)
	i, l := Unsigned.Decode(data)
	if l != 1 {
		log.Println(data, i, l)
		t.Fatal("bad uint64 decoding length")
	}
	if i.Sign() != 0 {
		log.Println(data, i, l)
		t.Fatal("bad uint64 decoding result")
	}
}

// TestEncodeUint64 encodes a random uint64 using both binary.PutUvarint and Encode
// and checks that the encoded results are equal using bytes.Equal.
// The randomness here is not real randomness; rand.Uint64() yields
// the same sequence of values every time the test suite is run.
func TestEncodeUint64(t *testing.T) {
	a := rand.Uint64()
	log.Println("using uint64", a)
	i := big.NewInt(0)
	for v := a; v > 0; v = v >> 1 {
		i.SetUint64(v)
		d1 := make([]byte, binary.MaxVarintLen64)
		d2 := make([]byte, binary.MaxVarintLen64)
		l1 := Unsigned.Encode(d1, i)
		l2 := binary.PutUvarint(d2, i.Uint64())
		if l1 != l2 {
			log.Println(v, d1, l1, d2, l2)
			t.Fatal("bad uint64 encoding length")
		}
		if bytes.Equal(d1, d2) {
			continue
		} else {
			log.Println(v, d1, l1, d2, l2)
			t.Fatal("bad uint64 encoding result")
		}
	}
}

// TestDecodeUint64 encodes a random uint64 using binary.PutUvarint and
// decodes it using Decode, and checks that the same value is recovered.
// Again the randomness here is not real randomness.
func TestDecodeUint64(t *testing.T) {
	a := rand.Uint64()
	log.Println("using uint64", a)
	for v := a; v > 0; v = v >> 1 {
		data := make([]byte, binary.MaxVarintLen64)
		length := binary.PutUvarint(data, v)
		i, l := Unsigned.Decode(data)
		if l != length {
			log.Println(v, data, length, i, l)
			t.Fatal("bad uint64 decoding length")
		} else if j := i.Uint64(); j != v {
			log.Println(v, data, length, i, l)
			t.Fatal("bad uint64 decoding result")
		}
	}
}

// TestEncodeDecodeUint128 encodes a random *big.Int in the 128-bit range
// and checks that decoding the result recovers the original value.
// Again the randomness here is not real randomness.
func TestEncodeDecodeUint128(t *testing.T) {
	a, b := rand.Uint64(), rand.Uint64()
	log.Println("using uint64s", a, b)
	i := big.NewInt(0).SetUint64(a)
	i = i.Lsh(i, 64)
	i = i.Add(i, big.NewInt(0).SetUint64(b))
	for i.BitLen() > 0 {
		length := Unsigned.EncodingLength(i)
		data := make([]byte, length)
		if Unsigned.Encode(data, i) != length {
			log.Println(i, length, data)
			t.Fatal("bad uint128 encoding length")
		} else if j, l := Unsigned.Decode(data); l != length {
			t.Fatal("bad uint128 decoding length")
		} else if j.Cmp(i) != 0 {
			t.Fatal("bad uint128 decoding result")
		}
		i = i.Rsh(i, 1)
	}
}
