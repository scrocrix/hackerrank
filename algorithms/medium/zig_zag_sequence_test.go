package medium_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type zigZagSequence struct {
	suite.Suite
}

func (this *zigZagSequence) TestZigZagSequence() {
	this.Run("Example", func() {
		result := medium.ZigZagSequence([]int{1, 2, 3, 4, 5, 6, 7}, 7)
		require.Equal(this.T(), []int{1, 2, 3, 7, 6, 5, 4}, result)
	})

	this.Run("Test case 4", func() {
		result := medium.ZigZagSequence([]int{25, 8, 20, 45, 12, 35, 36, 39, 31, 49, 3, 47, 30, 33, 37, 23, 14, 9, 27, 6, 10, 22, 24, 15, 43, 28, 38, 40, 17, 34, 21, 16, 4, 46, 7, 5, 44, 50, 26, 18, 32, 29, 42, 1, 48, 41, 19, 11, 2}, 49)
		require.Equal(this.T(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26}, result)
	})

	this.Run("Test case 5", func() {
		result := medium.ZigZagSequence([]int{9, 47, 14, 57, 26, 80, 61, 76, 64, 99, 39, 42, 44, 58, 1, 81, 100, 18, 62, 63, 72, 87, 85, 65, 97, 46, 96, 27, 4, 45, 28, 38, 89, 3, 101, 17, 92, 84, 52, 16, 20, 86, 77, 32, 53, 79, 23, 8, 36, 82, 13, 34, 90, 88, 91, 71, 31, 6, 29, 70, 73, 21, 98, 37, 24, 33, 67, 25, 75, 19, 40, 94, 12, 49, 54, 68, 15, 59, 83, 66, 78, 95, 69, 2, 51, 11, 43, 22, 50, 30, 41, 104, 10, 60, 5, 102, 48, 74, 103, 105, 93, 35, 7, 56, 55}, 105)

		require.Equal(this.T(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 105, 104, 103, 102, 101, 100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53}, result)
	})
}

func TestZigZagSequence(t *testing.T) {
	suite.Run(t, new(zigZagSequence))
}
