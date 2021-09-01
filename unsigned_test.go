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
	data := Unsigned.EncodeToBytes(i)
	if len(data) != 1 {
		log.Println(data, i)
		t.Fatal("bad uint64 encoding length")
	} else if data[0] != 0 {
		log.Println(data, i)
		t.Fatal("bad uint64 encoding result")
	}
}

func TestDecodeUnsignedZero(t *testing.T) {
	data := make([]byte, 10)
	i, err := Unsigned.DecodeBytes(data)
	if err != nil {
		t.Fatal(err)
	} else if i.Sign() != 0 {
		log.Println(data, i)
		t.Fatal("bad uint64 decoding result")
	}
}

// TestEncodeUint64 encodes a random uint64 using both binary.PutUvarint and Encode
// and checks that the encoded results are equal using bytes.Equal.
// The randomness here is not real randomness; rand.Uint64() yields
// the same sequence of values every time the test suite is run.
func TestEncodeUint64(t *testing.T) {
	a := rand.Uint64()
	i := big.NewInt(0)
	for v := a; v > 0; v = v >> 1 {
		log.Println("encode uint64", v)
		i.SetUint64(v)
		d1 := Unsigned.EncodeToBytes(i)
		d2 := make([]byte, binary.MaxVarintLen64)
		l := binary.PutUvarint(d2, v)
		if len(d1) != l {
			log.Println(v, d1, d2, l)
			t.Fatal("bad uint64 encoding length")
		}

		if bytes.Equal(d1, d2[:l]) {
			continue
		} else {
			log.Println(v, d1, d2, l)
			t.Fatal("bad uint64 encoding result")
		}
	}
}

// TestDecodeUint64 encodes a random uint64 using binary.PutUvarint and
// decodes it using Decode, and checks that the same value is recovered.
// Again, the randomness here is not real randomness.
func TestDecodeUint64(t *testing.T) {
	a := rand.Uint64()
	for v := a; v > 0; v = v >> 1 {
		log.Println("decode uint64", v)
		data := make([]byte, binary.MaxVarintLen64)
		length := binary.PutUvarint(data, v)
		i, err := Unsigned.DecodeBytes(data)
		if err != nil {
			t.Fatal(err)
		} else if i.Uint64() != v {
			log.Println(v, data, length, i)
			t.Fatal("bad uint64 decoding result")
		}
	}
}

// TestEncodeDecodeUint128 encodes a random *big.Int in the 128-bit range
// and checks that decoding the result recovers the original value.
// Again, the randomness here is not real randomness.
func TestEncodeDecodeUint128(t *testing.T) {
	a, b := rand.Uint64(), rand.Uint64()

	i := big.NewInt(0).SetUint64(a)
	i = i.Lsh(i, 64)
	i = i.Add(i, big.NewInt(0).SetUint64(b))

	for i.BitLen() > 64 {
		log.Println("encode/decode uint128", i)
		length := Unsigned.EncodedLen(i)
		data := Unsigned.EncodeToBytes(i)
		if len(data) != length {
			log.Println(i, length, data)
			t.Fatal("bad uint128 encoding length")
		} else if j, err := Unsigned.DecodeBytes(data); err != nil {
			t.Fatal(err)
		} else if j.Cmp(i) != 0 {
			t.Fatal("bad uint128 decoding result")
		}
		i = i.Rsh(i, 1)
	}
}
