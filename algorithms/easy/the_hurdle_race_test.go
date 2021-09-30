package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type hurdleRace struct {
	suite.Suite
}

func (this *hurdleRace) TestHurdleRace() {
	this.Run("Example 1", func() {
		result := easy.HurdleRace(4, []int32{1, 6, 3, 5, 2})
		require.Equal(this.T(), int32(2), result)
	})

	this.Run("Example 2", func() {
		result := easy.HurdleRace(7, []int32{2, 5, 4, 5, 2})
		require.Equal(this.T(), int32(0), result)
	})
}

func TestHurdleRace(t *testing.T) {
	suite.Run(t, new(hurdleRace))
}
