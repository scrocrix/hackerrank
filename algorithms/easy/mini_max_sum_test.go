package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type miniMaxSum struct {
	suite.Suite
}

func (this *miniMaxSum) TestMiniMaxSum() {
	this.Run("Calculate example", func() {
		result := easy.MiniMaxSum([]int32{1, 3, 5, 7, 9})
		require.Equal(this.T(), []int{16, 24}, result)
	})
}

func TestMiniMaxSum(t *testing.T) {
	suite.Run(t, new(miniMaxSum))
}
