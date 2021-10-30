package ds

func PrintTheElementsOfALinkedList(head *SinglyLikedListNode) []int32 {
	c := head
	data := []int32{}
	for c != nil {
		data = append(data, c.Data.(int32))
		c = c.Next
	}
	return data
}
