package medium

import "sort"

func ZigZagSequence(a []int, n int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	k := a[len(a)-1]

	a = a[:len(a)-1]

	a1 := a[:len(a)/2]
	sort.Slice(a1, func(i, j int) bool {
		return a1[i] < a1[j]
	})

	a2 := a[((len(a)/2)-1)+1:]
	sort.Slice(a2, func(i, j int) bool {
		return a2[i] > a2[j]
	})

	var r []int
	r = append(r, a1...)
	r = append(r, k)
	r = append(r, a2...)

	return r
}
