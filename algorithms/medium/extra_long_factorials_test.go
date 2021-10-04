package medium_test

import (
	"math/big"
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type extraLongFactorials struct {
	suite.Suite
}

func (this *extraLongFactorials) TestExtraLongFactorials() {
	this.Run("Example 1", func() {
		result := medium.ExtraLongFactorials(25)
		e, _ := big.NewInt(0).SetString("15511210043330985984000000", 10)
		assert.Equal(this.T(), e, result)
	})
}

func TestExtraLongFactorials(t *testing.T) {
	suite.Run(t, new(extraLongFactorials))
}
