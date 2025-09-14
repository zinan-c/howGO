package model

type DoublyListNode struct {
	Val        int
	Prev, Next *DoublyListNode
}

func NewDoublyListNode(x int) *DoublyListNode {
	return &DoublyListNode{Val: x}
}

func CreateDoublyLinkedList(arr []int) *DoublyListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := NewDoublyListNode(arr[0])
	cur := head
	for i := 1; i < len(arr); i++ {
		newNode := NewDoublyListNode(arr[i])
		cur.Next = newNode
		newNode.Prev = cur
		cur = cur.Next
	}
	return head
}
