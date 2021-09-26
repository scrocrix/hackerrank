package easy

import "sort"

func MiniMaxSum(arr []int32) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	t1 := int(arr[0])
	t2 := int(arr[1])
	t3 := int(arr[2])
	t4 := int(arr[3])
	t5 := int(arr[4])

	return []int{
		t1 + t2 + t3 + t4,
		t2 + t3 + t4 + t5,
	}
}
