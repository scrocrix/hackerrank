package easy

type BinarySearchTree struct {
	Root *LevelOrderNode
}

func (this *BinarySearchTree) Insert(val int) {
	if this.Root == nil {
		this.Root = &LevelOrderNode{
			Info: val,
		}
		return
	}

	node := this.Root

	for node != nil {
		if val > node.Info {
			if node.Right != nil {
				node = node.Right
				continue
			} else {
				node.Right = &LevelOrderNode{
					Info: val,
				}

				break
			}
		}

		if val < node.Info {
			if node.Left != nil {
				node = node.Left
				continue
			} else {
				node.Left = &LevelOrderNode{
					Info: val,
				}

				break
			}
		}
	}
}
