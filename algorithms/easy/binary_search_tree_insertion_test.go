package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type binarySearchTree struct {
	suite.Suite
}

func (this *binarySearchTree) TestInsert() {
	this.Run("Insert example", func() {
		sut := &easy.BinarySearchTree{}

		sut.Insert(4)
		sut.Insert(2)
		sut.Insert(3)
		sut.Insert(1)
		sut.Insert(7)
		sut.Insert(6)

		require.Equal(this.T(), 4, sut.Root.Info)
		require.Equal(this.T(), 2, sut.Root.Left.Info)
		require.Equal(this.T(), 7, sut.Root.Right.Info)
		require.Equal(this.T(), 1, sut.Root.Left.Left.Info)
		require.Equal(this.T(), 3, sut.Root.Left.Right.Info)
		require.Equal(this.T(), 6, sut.Root.Right.Left.Info)
	})
}

func TestBinarySearchTree(t *testing.T) {
	suite.Run(t, new(binarySearchTree))
}
