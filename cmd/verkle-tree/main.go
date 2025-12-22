package main

import "github.com/morelucks/verkle-tree/internal/verkle"

func main() {
	tree := verkle.NewTree()
	tree.Insert([]byte("key1"), []byte("value1"))
}
