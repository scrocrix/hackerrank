package easy_test

import (
	"testing"

	"github.com/scrocrix/hackerrank/algorithms/easy"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type levelOrder struct {
	suite.Suite
}

func (this *levelOrder) TestLevelOrder() {
	this.Run("Return level order", func() {
		result := easy.LevelOrder(&easy.LevelOrderNode{
			Info: 1,
			Right: &easy.LevelOrderNode{
				Info: 2,
				Right: &easy.LevelOrderNode{
					Info: 5,
					Left: &easy.LevelOrderNode{
						Info: 3,
						Right: &easy.LevelOrderNode{
							Info: 4,
						},
					},
					Right: &easy.LevelOrderNode{
						Info: 6,
					},
				},
			},
		})

		require.Equal(this.T(), []int{1, 2, 5, 3, 6, 4}, result)
	})
}

func TestLevelOrder(t *testing.T) {
	suite.Run(t, new(levelOrder))
}
