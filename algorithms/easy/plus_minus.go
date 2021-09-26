package easy

func PlusMinus(arr []int32) []float64 {
	ratios := map[string]float64{
		"negative": 0,
		"positive": 0,
		"zero":     0,
	}

	for _, item := range arr {
		if item < 0 {
			ratios["negative"] += 1
		}

		if item > 0 {
			ratios["positive"] += 1
		}

		if item == 0 {
			ratios["zero"] += 1
		}
	}

	s := len(arr)

	return []float64{
		ratios["positive"] / float64(s),
		ratios["negative"] / float64(s),
		ratios["zero"] / float64(s),
	}
}
