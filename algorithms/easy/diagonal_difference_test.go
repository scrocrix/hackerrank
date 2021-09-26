package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type diagonalDifference struct {
	suite.Suite
}

func (this *diagonalDifference) TestDiagonalDifference() {
	this.Run("Example", func() {
		result := easy.DiagonalDifference([][]int32{
			[]int32{1, 2, 3},
			[]int32{4, 5, 6},
			[]int32{9, 8, 9},
		})
		require.Equal(this.T(), int32(2), result)
	})

	this.Run("Example 1", func() {
		result := easy.DiagonalDifference([][]int32{
			[]int32{-1, 1, -7, -8},
			[]int32{-10, -8, -5, -2},
			[]int32{0, 9, 7, -1},
			[]int32{4, 4, -2, 1},
		})
		require.Equal(this.T(), int32(1), result)
	})
}

func TestDiagonalDifference(t *testing.T) {
	suite.Run(t, new(diagonalDifference))
}
