package easy

func HurdleRace(k int32, height []int32) int32 {
	max := int32(0)
	for _, h := range height {
		if h > max {
			max = h
		}
	}

	if max-k > 0 {
		return max - k
	}

	return int32(0)
}
