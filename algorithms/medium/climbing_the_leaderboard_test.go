package medium_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type climbingLeaderboard struct {
	suite.Suite
}

func (this *climbingLeaderboard) TestClimbingLeaderboard() {
	this.Run("Dummy example", func() {
		result := medium.ClimbingLeaderboard([]int32{100}, []int32{50})
		require.Equal(this.T(), []int32{2}, result)
	})

	this.Run("Example 1", func() {
		result := medium.ClimbingLeaderboard([]int32{
			100,
			100,
			50,
			40,
			40,
			20,
			10,
		}, []int32{
			5,
			25,
			50,
			120,
		})

		require.Equal(this.T(), []int32{6, 4, 2, 1}, result)
	})

	this.Run("Example 2", func() {
		result := medium.ClimbingLeaderboard([]int32{
			100,
			90,
			90,
			80,
			75,
			60,
		}, []int32{
			50,
			65,
			77,
			90,
			102,
		})

		require.Equal(this.T(), []int32{6, 5, 4, 2, 1}, result)
	})
}

func TestClimbingLeaderboard(t *testing.T) {
	suite.Run(t, new(climbingLeaderboard))
}
