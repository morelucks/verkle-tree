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

## Features

- Written entirely in **Go**
- Verkle Tree construction and traversal
- Key–value storage model
- Proof generation and verification (work-in-progress / experimental)
- Focus on clarity and correctness

---
 
