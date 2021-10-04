package medium

import "math/big"

func ExtraLongFactorials(n int32) *big.Int {
	r := big.NewInt(1)
	nn := big.NewInt(int64(n))
	for nn.Int64() > int64(0) {
		r = r.Mul(r, nn)
		nn = nn.Sub(nn, big.NewInt(1))
	}
	return r
}
