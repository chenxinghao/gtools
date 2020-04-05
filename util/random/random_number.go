package random

import (
	cryptoRand "crypto/rand"
	"math/big"
)

type Number struct {
}

func (r *Number) CryptoRandInt(min, max int) (int, error) {
	if min >= max {
		return max, nil
	}
	num, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return int(num.Int64()) + min, nil
}
