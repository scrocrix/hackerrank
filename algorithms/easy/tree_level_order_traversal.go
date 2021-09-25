package easy

import (
	"github.com/Workiva/go-datastructures/queue"
)

type LevelOrderNode struct {
	Left  *LevelOrderNode
	Right *LevelOrderNode
	Info  int
}

func LevelOrder(root *LevelOrderNode) []int {
	var s []int

	q := queue.New(500)

	q.Put(root)

	for q.Len() != 0 {
		item, _ := q.Get(1)

		n := item[0].(*LevelOrderNode)

		s = append(s, n.Info)

		if n.Left != nil {
			_ = q.Put(n.Left)
		}

		if n.Right != nil {
			_ = q.Put(n.Right)
		}
	}

	return s
}
