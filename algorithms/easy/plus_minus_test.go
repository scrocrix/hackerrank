package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type plusMinus struct {
	suite.Suite
}

func (this *plusMinus) TestPlusMinus() {
	this.Run("Calculate ratio", func() {
		result := easy.PlusMinus([]int32{-4, 3, -9, 0, 4, 1})
		require.Equal(this.T(), []float64{
			0.500000,
			0.3333333333333333,
			0.16666666666666666,
		}, result)
	})
}

func TestPlusMinus(t *testing.T) {
	suite.Run(t, new(plusMinus))
}
