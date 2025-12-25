package verkle

import (
	"bytes"
	"fmt"
)

const (
	KeySize            = 32
	LeafValueSize      = 32
	NodeWidth          = 256
	NodeBitWidth  byte = 8
	StemSize           = 31
)

type (
	NodeFlushFn    func([]byte, VerkleNode)
	NodeResolverFn func([]byte) ([]byte, error)
)

type Keylist [][]byte

func (kl Keylist) Len() int {
	return len(kl)
}

func (kl Keylist) Less(i, j int) bool {
	return bytes.Compare(kl[i], kl[j]) == -1
}

func (kl Keylist) Swap(i, j int) {
	kl[i], kl[j] = kl[j], kl[i]
}

type Stem []byte

// Point represents a curve point (placeholder for cryptographic implementation)
type Point struct{}

// Fr represents a field element (placeholder for cryptographic implementation)
type Fr struct{}

// ProofElements represents proof elements (placeholder)
type ProofElements struct{}

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
func KeyToStem(key []byte) Stem {
	if len(key) < StemSize {
		panic(fmt.Errorf("key length (%d) is shorter than the expected stem size (%d)", len(key), StemSize))
	}
	return Stem(key[:StemSize])
}

// EqualPaths checks if two keys have the same stem
func EqualPaths(key1, key2 []byte) bool {
	return bytes.Equal(KeyToStem(key1), KeyToStem(key2))
}

// VerkleNode interface defines the operations for Verkle tree nodes
type VerkleNode interface {
	// Insert or Update value into the tree
	Insert([]byte, []byte, NodeResolverFn) error

	// Delete a leaf with the given key
	Delete([]byte, NodeResolverFn) (bool, error)

	// Get value at a given key
	Get([]byte, NodeResolverFn) ([]byte, error)

	// Commit computes the commitment of the node. The
	// result (the curve point) is cached.
	Commit() *Point

	// Commitment is a getter for the cached commitment
	// to this node.
	Commitment() *Point

	// Hash returns the field representation of the commitment.
	Hash() *Fr

	// GetProofItems collects the various proof elements, and
	// returns them breadth-first. On top of that, it returns
	// one "extension status" per stem, and an alternate stem
	// if the key is missing but another stem has been found.
	GetProofItems(Keylist, NodeResolverFn) (*ProofElements, []byte, []Stem, error)

	// Serialize encodes the node to RLP.
	Serialize() ([]byte, error)

	// Copy a node and its children
	Copy() VerkleNode

	// toDot returns a string representing this subtree in DOT language
	toDot(string, string) string

	setDepth(depth byte)
}

// Offset2key extracts the n bits of a key that correspond to the
// index of a child node.
func Offset2key(key []byte, offset byte) byte {
	return key[offset]
}

// Insert adds a key-value pair to the tree
func (t *Tree) Insert(key, value []byte) {

}
