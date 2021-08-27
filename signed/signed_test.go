package signed

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
	"math/rand"
	"testing"
)

func TestEncodeZero(t *testing.T) {
	i := big.NewInt(0)
	data := make([]byte, 10)
	l := Encode(data, i)
	if l != 1 {
		log.Println(data, i, l)
		t.Fatal("bad int64 encoding length")
	}

	if bytes.Equal(data, make([]byte, 10)) {
		return
	} else {
		log.Println(data, i, l)
		t.Fatal("bad int64 encoding result")
	}
}

func TestDecodeZero(t *testing.T) {
	data := make([]byte, 10)
	i, l := Decode(data)
	if l != 1 {
		log.Println(data, i, l)
		t.Fatal("bad int64 decoding length")
	}
	if i.Sign() != 0 {
		log.Println(data, i, l)
		t.Fatal("bad int64 decoding result")
	}
}

// TestEncodeInt64 encodes a random int64 using both binary.PutVarint and Encode
// and checks that the encoded results are equal using bytes.Equal.
// The randomness here is not real randomness; rand.Int63() yields
// the same sequence of values every time the test suite is run.
func TestEncodeInt64(t *testing.T) {
	a := rand.Int63()
	log.Println("using int64", a)
	i := big.NewInt(0)
	for v := a; v > 0; v = v >> 1 {
		i.SetInt64(v)
		d1 := make([]byte, binary.MaxVarintLen64)
		d2 := make([]byte, binary.MaxVarintLen64)
		l1 := Encode(d1, i)
		l2 := binary.PutVarint(d2, i.Int64())
		if l1 != l2 {
			log.Println(v, d1, l1, d2, l2)
			t.Fatal("bad int64 encoding length")
		}
		if bytes.Equal(d1, d2) {
			continue
		} else {
			log.Println(v, d1, l1, d2, l2)
			t.Fatal("bad int64 encoding result")
		}
	}
}

// TestDecodeInt64 encodes a random int64 using binary.PutVarint and
// decodes it using Decode, and checks that the same value is recovered.
// Again the randomness here is not real randomness.
func TestDecodeInt64(t *testing.T) {
	a := rand.Int63()
	log.Println("using int64", a)
	for v := a; v > 0; v = v >> 1 {
		data := make([]byte, binary.MaxVarintLen64)
		length := binary.PutVarint(data, v)
		i, l := Decode(data)
		if l != length {
			log.Println(v, data, length, i, l)
			t.Fatal("bad int64 decoding length")
		} else if j := i.Int64(); j != v {
			log.Println(v, data, length, i, l)
			t.Fatal("bad int64 decoding result")
		}
	}
}

// TestEncodeDecodeInt128 encodes a random *big.Int in the 128-bit range
// and checks that decoding the result recovers the original value.
// Again the randomness here is not real randomness.
func TestEncodeDecodeInt128(t *testing.T) {
	a, b := rand.Int63(), rand.Int63()
	log.Println("using int64s", a, b)
	i := big.NewInt(a)
	i = i.Lsh(i, 64)
	i = i.Add(i, big.NewInt(b))
	for i.BitLen() > 0 {
		length := EncodingLength(i)
		data := make([]byte, length)
		if Encode(data, i) != length {
			log.Println(i, length, data)
			t.Fatal("bad int128 encoding length")
		} else if j, l := Decode(data); l != length {
			t.Fatal("bad int128 decoding length")
		} else if j.Cmp(i) != 0 {
			t.Fatal("bad int128 decoding result")
		}
		i = i.Rsh(i, 1)
	}
}
