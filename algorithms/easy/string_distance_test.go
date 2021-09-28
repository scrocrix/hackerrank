package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type stringDistance struct {
	suite.Suite
}

func (this *stringDistance) TestStringDistance() {
	this.Run("Example", func() {
		result := easy.StringDistance("2\nFOOD\nMONEY\nCAT\nDOG")
		require.Equal(this.T(), "4\n3", result)

		result = easy.StringDistance("3\nDOG\nCAT\nELEPHANT\nSPIDER")
		require.Equal(this.T(), "3\n8", result)
	})
}

func TestStringDistance(t *testing.T) {
	suite.Run(t, new(stringDistance))
}
