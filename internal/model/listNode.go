package model

type ListNode struct {
	Val  int
	Next *ListNode
}

// first char should be uppercase which it can be invoke from other package
func CreateLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}
