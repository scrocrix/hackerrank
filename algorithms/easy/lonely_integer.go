package easy

func LonelyInteger(a []int32) int32 {
	lookup := map[int32]int32{}

	for _, n := range a {
		if _, exists := lookup[n]; !exists {
			lookup[n] = 1
			continue
		}

		lookup[n] += 1
	}

	for i, n := range lookup {
		if n == 1 {
			return i
		}
	}

	return -1
}
