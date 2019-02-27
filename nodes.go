package doublearray

// Node is a node of tree
type Node struct {
	Base  int
	Check int
}

// Nodes is that all trie tree nodes
type Nodes []Node

// At returns a node from cursor
func (n *Nodes) At(cursor int) *Node {
	return &(*n)[cursor-1]
}
