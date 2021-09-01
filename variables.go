package varint

import (
	"errors"
	"math/big"
)

var ErrLessThanZero = errors.New("unsigned varints must be non-negative")

var limit = big.NewInt(0x7F)
var two = big.NewInt(2)
