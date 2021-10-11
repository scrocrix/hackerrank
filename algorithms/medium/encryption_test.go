package medium_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/medium"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type encryption struct {
	suite.Suite
}

func (this *encryption) TestEncryption() {
	this.Run("Example", func() {
		assert.Equal(this.T(), "hae and via ecy", medium.Encryption("haveaniceday"))
		assert.Equal(this.T(), "fto ehg ee dd", medium.Encryption("feedthedog"))
		assert.Equal(this.T(), "clu hlt io", medium.Encryption("chillout"))
	})
}

func TestEncryption(t *testing.T) {
	suite.Run(t, new(encryption))
}
