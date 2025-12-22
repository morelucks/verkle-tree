# Go Verkle Tree

A **Go implementation of the Verkle Tree data structure**, built to closely follow the Verkle Tree specification and support efficient state representation and proof generation for blockchain systems.

This repository is primarily a **learning, research, and engineering-focused implementation**, exploring how Verkle Trees work internally and how they can be implemented in Go.


**Verkle Tree Spec (reference):** https://github.com/crate-crypto/verkle-trie-ref/tree/master/verkle

---

## What is a Verkle Tree?

A Verkle Tree is a cryptographic data structure that improves on traditional Merkle trees by using **vector commitments** instead of simple hashes. This allows many child nodes to be committed to at once, resulting in **significantly smaller proofs**.

Verkle Trees are a key component in Ethereum’s roadmap toward **stateless clients**, enabling nodes to verify state without storing the entire state tree.

---

## Why this project

- Learn and understand Verkle Trees at a protocol level
- Explore Go-based implementations of advanced cryptographic data structures
- Study how Verkle Trees improve proof size and scalability
- Gain hands-on experience relevant to Ethereum core development

---

## Project Structure

```
verkel-tree/
├── cmd/
│   └── verkle-tree/
│       └── main.go          # Application entry point
├── internal/
│   └── verkle/
│       └── tree.go          # Verkle tree implementation
├── go.mod                    # Go module definition
├── .gitignore
├── LICENSE
└── README.md
```
 
 

## Getting Started

### Prerequisites

- Go 1.25.5 or later

### Building

Build the application:

```bash
go build ./cmd/verkle-tree
```

This will create an executable `verkle-tree` in the current directory.

### Running

Run the application:

```bash
go run ./cmd/verkle-tree
```

Or run the built binary:

```bash
./verkle-tree
```

### Development

To work on the project:

```bash
# Clone the repository
git clone https://github.com/morelucks/verkle-tree.git
cd verkle-tree

# Run the application
go run ./cmd/verkle-tree
```

---

## Features

- Written entirely in **Go**
- Verkle Tree construction and traversal
- Key–value storage model
- Proof generation and verification (work-in-progress / experimental)
- Focus on clarity and correctness

---
 
