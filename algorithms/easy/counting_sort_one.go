package easy

func CountringSort(arr []int32) []int32 {
	lookup := map[int32]int32{}

	for _, n := range arr {
		if _, exists := lookup[n]; !exists {
			lookup[n] = 1
			continue
		}

		lookup[n] += 1
	}

	items := make([]int32, 100)

	for i, _ := range items {
		if occurrences, exists := lookup[int32(i)]; exists {
			items[i] = occurrences
		}
	}

	return items
}
