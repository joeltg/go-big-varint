package varint

import (
	"errors"
	"math/big"
)

var ErrLessThanZero = errors.New("unsigned varints must be non-negative")

var msb = big.NewInt(0x80)
var rest = big.NewInt(0x7F)
var notRest = big.NewInt(0).Not(rest)
var limit = big.NewInt(1 << 31)
var pad = big.NewInt(0xFF)

var two = big.NewInt(2)
