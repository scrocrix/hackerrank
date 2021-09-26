package easy

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	lr := 0
	rl := len(arr) - 1

	l := int32(0)
	r := int32(0)

	for _, item := range arr {
		l = l + item[lr]
		r = r + item[rl]
		lr = lr + 1
		rl = rl - 1
	}

	return int32(math.Abs(float64(l - r)))
}
