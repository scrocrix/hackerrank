package medium

import (
	"sort"
)

func ClimbingLeaderboard(ranked, player []int32) []int32 {
	cr := []int{}
	for _, rank := range ranked {
		if len(cr) == 0 {
			cr = append(cr, int(rank))
			continue
		}

		if rank == int32(cr[len(cr)-1]) {
			continue
		}

		cr = append(cr, int(rank))
	}

	positions := []int32{}
	for _, score := range player {
		position := sort.Search(len(cr), func(i int) bool {
			return int(score) > int(cr[i]) || int(score) == int(cr[i])
		}) + 1

		if position == -1 {
			positions = append(positions, int32(len(cr)+1))
		} else {
			positions = append(positions, int32(position))
		}
	}

	return positions
}
