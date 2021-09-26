package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type lonelyInteger struct {
	suite.Suite
}

func (this *lonelyInteger) TestLonelyInteger() {
	this.Run("Example", func() {
		result := easy.LonelyInteger([]int32{1, 2, 3, 4, 3, 2, 1})
		require.Equal(this.T(), int32(4), result)
	})
}

func TestLonelyInteger(t *testing.T) {
	suite.Run(t, new(lonelyInteger))
}
