package ds_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/ds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type printTheElementOfALinkedList struct {
	suite.Suite
}

func (this *printTheElementOfALinkedList) TestPrintTheElementOfALinkedList() {
	this.Run("Example", func() {
		result := ds.PrintTheElementsOfALinkedList(&ds.SinglyLikedListNode{
			Data: int32(2),
			Next: &ds.SinglyLikedListNode{
				Data: int32(16),
				Next: &ds.SinglyLikedListNode{
					Data: int32(13),
				},
			},
		})

		assert.Equal(this.T(), []int32{2, 16, 13}, result)
	})
}

func TestPrintTheElementOfALinkedList(t *testing.T) {
	suite.Run(t, new(printTheElementOfALinkedList))
}
