package medium_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type timeInWords struct {
	suite.Suite
}

func (this *timeInWords) TestTimeInWords() {
	this.Run("Example", func() {
		result := medium.TimeInWords(5, 47)
		require.Equal(this.T(), "thirteen minutes to six", result)

		result = medium.TimeInWords(6, 00)
		require.Equal(this.T(), "six o' clock", result)

		result = medium.TimeInWords(3, 00)
		require.Equal(this.T(), "three o' clock", result)

		result = medium.TimeInWords(7, 15)
		require.Equal(this.T(), "quarter past seven", result)

		result = medium.TimeInWords(7, 53)
		require.Equal(this.T(), "seven minutes to eight", result)

		result = medium.TimeInWords(5, 30)
		require.Equal(this.T(), "half past five", result)

		result = medium.TimeInWords(4, 15)
		require.Equal(this.T(), "quarter past four", result)
	})
}

func TestTimeInWords(t *testing.T) {
	suite.Run(t, new(timeInWords))
}
