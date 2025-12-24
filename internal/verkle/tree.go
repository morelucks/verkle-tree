package verkle

import "bytes"

const (
	KeySize            = 32
	LeafValueSize      = 32
	NodeWidth          = 256
	NodeBitWidth  byte = 8
	StemSize           = 31
)

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

// KeyToStem extracts the stem (first 31 bytes) from a key
func KeyToStem(key []byte) []byte {
	if len(key) < StemSize {
		return key
	}
	return key[:StemSize]
}

// equalPaths checks if two keys have the same stem
func equalPaths(key1, key2 []byte) bool {
	return bytes.Equal(KeyToStem(key1), KeyToStem(key2))
}

// offset2key extracts the n bits of a key that correspond to the
// index of a child node.
func offset2key(key []byte, offset byte) byte {
	return key[offset]
}

// Insert adds a key-value pair to the tree
func (t *Tree) Insert(key, value []byte) {

}
