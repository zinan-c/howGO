package model

type Node struct {
	val  interface{}
	next *Node
	prev *Node
}
