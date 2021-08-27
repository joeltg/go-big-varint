package signed

import (
	"math/big"

	"github.com/joeltg/go-big-varint/unsigned"
)

// Decode a big integer from a byte slice, returning the result and the number of bytes read
func Decode(data []byte) (*big.Int, int) {
	i, l := unsigned.Decode(data)
	if i.Bit(0) == 1 {
		i = i.Not(i)
	}
	i = i.Div(i, two)
	return i, l
}
