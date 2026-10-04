package types

type Dll struct {
	Head *Node
}

type Node struct {
	Previous *Node
	Data int
	Next *Node
}
