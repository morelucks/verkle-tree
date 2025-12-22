package verkle

// Tree represents a Verkle tree data structure
type Tree struct {
	root *Node
}
 
type Node struct {
	children map[byte]*Node
	value    []byte
}
 
func NewTree() *Tree {
	return &Tree{root: &Node{children: make(map[byte]*Node)}}
}

// Insert adds a key-value pair to the tree
func (t *Tree) Insert(key, value []byte) {
	 
}
