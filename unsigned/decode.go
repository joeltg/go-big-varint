package unsigned

import (
	"errors"
	"math/big"
)

var ErrOutOfRange = errors.New("out of range")

// Decode a big integer from a byte slice, returning the result and the number of bytes read.
// The result is guaranteed to be non-negative.
func Decode(data []byte) (*big.Int, int) {
	offset := 0
	i, shift := big.NewInt(0), uint(0)
	delta, c := big.NewInt(0), big.NewInt(0)
	for {
		if offset < len(data) {
			b := data[offset]
			delta = delta.SetUint64(uint64(b))
			delta = delta.And(delta, rest)
			if shift < 28 {
				delta = delta.Lsh(delta, shift)
			} else {
				c = c.SetUint64(1)
				c = c.Lsh(c, shift)
				delta = delta.Mul(delta, c)
			}
			i = i.Add(i, delta)
			shift += 7
			offset++
			if b < 0x80 {
				break
			} else {
				continue
			}
		} else {
			panic(ErrOutOfRange)
		}
	}

	return i, offset
}

/*
export function decode(data: Uint8Array, offset = 0): bigint {
	const l = data.length
	let b: bigint
	let res = 0n,
		shift = 0n,
		counter = offset

	do {
		if (counter >= l) {
			throw new RangeError("could not decode varint")
		}
		b = BigInt(data[counter++])
		res += shift < 28n ? (b & REST) << shift : (b & REST) * (1n << shift)
		shift += 7n
	} while (b >= MSB)

	return res
}




var MSB = 0x80
  , REST = 0x7F

function read(buf, offset) {
  var res    = 0
    , offset = offset || 0
    , shift  = 0
    , counter = offset
    , b
    , l = buf.length

  do {
    if (counter >= l || shift > 49) {
      read.bytes = 0
      throw new RangeError('Could not decode varint')
    }
    b = buf[counter++]
    res += shift < 28
      ? (b & REST) << shift
      : (b & REST) * Math.pow(2, shift)
    shift += 7
  } while (b >= MSB)

  read.bytes = counter - offset

  return res
}
*/
