package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type timeConversion struct {
	suite.Suite
}

func (this *timeConversion) TestTimeConversion() {
	this.Run("Calculate example", func() {
		result := easy.TimeConversion("07:05:45PM")
		require.Equal(this.T(), "19:05:45", result)

		result = easy.TimeConversion("12:13:42AM")
		require.Equal(this.T(), "00:13:42", result)
	})
}

func TestTimeConversion(t *testing.T) {
	suite.Run(t, new(timeConversion))
}
