package varint

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
	"math/rand"
	"testing"
)

func TestEncodeSignedZero(t *testing.T) {
	i := big.NewInt(0)
	data := Signed.EncodeToBytes(i)
	if len(data) != 1 {
		log.Println(data, i)
		t.Fatal("bad int64 encoding length")
	} else if data[0] != 0 {
		log.Println(data, i)
		t.Fatal("bad int64 encoding result")
	}
}

func TestDecodeSignedZero(t *testing.T) {
	data := make([]byte, 10)
	i, err := Signed.DecodeBytes(data)
	if err != nil {
		t.Fatal(err)
	} else if i.Sign() != 0 {
		log.Println(data, i)
		t.Fatal("bad int64 decoding result")
	}
}

// TestEncodeInt64 encodes a random int64 using both binary.PutVarint and Encode
// and checks that the encoded results are equal using bytes.Equal.
// The randomness here is not real randomness; rand.Int63() yields
// the same sequence of values every time the test suite is run.
func TestEncodeInt64(t *testing.T) {
	a := rand.Int63()
	log.Println("testing int64", a)
	i := big.NewInt(0)
	for v := a; v != 0; v = ^v >> 1 {
		i.SetInt64(v)
		d1 := Signed.EncodeToBytes(i)
		d2 := make([]byte, binary.MaxVarintLen64)
		l := binary.PutVarint(d2, v)
		if l != len(d1) {
			log.Println(v, d1, len(d1), d2, l)
			t.Fatal("bad int64 encoding length")
		}
		if bytes.Equal(d1, d2[:l]) {
			continue
		} else {
			log.Println(v, d1, len(d1), d2, l)
			t.Fatal("bad int64 encoding result")
		}
	}
}

// TestDecodeInt64 encodes a random int64 using binary.PutVarint and
// decodes it using Decode, and checks that the same value is recovered.
// Again, the randomness here is not real randomness.
func TestDecodeInt64(t *testing.T) {
	a := rand.Int63()
	for v := a; v != 0; v = ^v >> 1 {
		log.Println("testing int64", v)
		data := make([]byte, binary.MaxVarintLen64)
		length := binary.PutVarint(data, v)
		i, err := Signed.DecodeBytes(data)
		if err != nil {
			t.Fatal(err)
		} else if j := i.Int64(); j != v {
			log.Println(v, data, length, i)
			t.Fatal("bad int64 decoding result")
		}
	}
}

// TestEncodeDecodeInt128 encodes a random *big.Int in the 128-bit range
// and checks that decoding the result recovers the original value.
// Again, the randomness here is not real randomness.
func TestEncodeDecodeInt128(t *testing.T) {
	a, b := rand.Int63(), rand.Int63()
	i := big.NewInt(a)
	i = i.Lsh(i, 64)
	i = i.Add(i, big.NewInt(b))
	for i.BitLen() > 64 {
		log.Println("encode/decode int128", i)
		length := Signed.EncodedLen(i)
		data := Signed.EncodeToBytes(i)
		if len(data) != length {
			log.Println(i, length, data)
			t.Fatal("bad int128 encoding length")
		} else if j, err := Signed.DecodeBytes(data); err != nil {
			t.Fatal(err)
		} else if j.Cmp(i) != 0 {
			t.Fatal("bad int128 decoding result")
		}
		i = i.Rsh(i, 1)
		i = i.Not(i)
	}
}
