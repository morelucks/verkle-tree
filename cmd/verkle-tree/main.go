package main

import (
	"fmt"
	"sort"

	"github.com/morelucks/verkle-tree/internal/verkle"
)

func main() {
	fmt.Println("=== Verkle Tree Implementation Tests ===")
	fmt.Println()

	// Create a new tree
	fmt.Println("Creating a new Verkle tree...")
	tree := verkle.NewTree()
	fmt.Println("a new tree", tree)
	fmt.Println("   Tree created successfully")
	fmt.Println()

	// KeyToStem with valid 32-byte key
	fmt.Println("Testing KeyToStem with 32-byte key...")
	key32 := make([]byte, 32)
	for i := range key32 {
		key32[i] = byte(i)
	}
	stem := verkle.KeyToStem(key32)
	fmt.Printf("Stem extracted: length=%d bytes\n", len(stem))
	fmt.Printf("First 8 bytes: %x\n\n", stem[:8])

	// offset2key
	fmt.Println("Testing offset2key...")
	offset := byte(5)
	keyByte := verkle.Offset2key(key32, offset)
	fmt.Printf("Key byte at offset %d: 0x%02x\n\n", offset, keyByte)

	//  equalPaths
	fmt.Println("Testing equalPaths...")
	key32_2 := make([]byte, 32)
	copy(key32_2, key32)
	key32_2[31] = 0xFF // Change byte 31 (outside stem, same stem)
	same := verkle.EqualPaths(key32, key32_2)
	fmt.Printf("Keys with same stem (different suffix): %v\n", same)

	key32_3 := make([]byte, 32)
	copy(key32_3, key32)
	key32_3[0] = 0xFF // Change first byte (different stem)
	same2 := verkle.EqualPaths(key32, key32_3)
	fmt.Printf("Keys with different stems: %v\n\n", same2)

	// Keylist sorting
	fmt.Println("Testing Keylist sorting...")
	keys := verkle.Keylist{
		{0x03, 0x04, 0x05},
		{0x01, 0x02, 0x03},
		{0x02, 0x03, 0x04},
	}
	fmt.Println("   Before sorting:")
	for i, k := range keys {
		fmt.Printf("     [%d]: %x\n", i, k)
	}
	sort.Sort(keys)
	fmt.Println("   After sorting:")
	for i, k := range keys {
		fmt.Printf("     [%d]: %x\n", i, k)
	}
	fmt.Println("Keylist sorted successfully")

}
