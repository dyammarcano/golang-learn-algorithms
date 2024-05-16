# Linked list

## Description

A linked list is a data structure that consists of a sequence of elements, each of which points to the next element in the sequence. Each element in a linked list is called a node. A node consists of two parts: data and a reference to the next node in the sequence.

## Comparison of Singly Linked Lists with Arrays & Dynamic Arrays

| Operation | Linked List | Array | Dynamic Array |
| --- | --- | --- | --- |
| Access | O(n) | O(1) | O(1) |
| Insertion/Deletion at beginning | O(1) | O(n) | O(n) |
| Insertion/Deletion at end | O(n) | O(1) | O(1) |
| Insertion/Deletion in middle | O(n) | O(n) | O(n) |
| Size | O(n) | O(n) | O(n) |

Based on the comparison, linked lists are efficient for insertion and deletion operations but inefficient for access operations. Arrays are efficient for access operations but inefficient for insertion and deletion operations. Dynamic arrays provide a balance between access and insertion/deletion operations.

## Types of Linked Lists

- [x] **Singly Linked List**: Each node has a data field and a reference to the next node.
- [x] **Doubly Linked List**: Each node has a data field, a reference to the next node, and a reference to the previous node.
- [x] **Circular Linked List**: The last node points back to the first node, forming a circle.
- [ ] **Doubly Circular Linked List**: A combination of doubly linked list and circular linked list.
- [ ] **Skip List**: A linked list with multiple levels to allow for faster search operations.
- [ ] **Self-organizing List**: A linked list that reorders its elements based on access frequency.
- [ ] **Unrolled Linked List**: A linked list that stores multiple elements in each node.
- [ ] **XOR Linked List**: A linked list that uses bitwise XOR to store the address of the next node.
- [ ] **Multi-level Linked List**: A linked list with multiple levels of nodes.
- [ ] **Sparse Linked List**: A linked list that stores only non-null values.
- [ ] **Hybrid Linked List**: A linked list that combines multiple types of linked lists.
- [ ] **Hashed Linked List**: A linked list that uses a hash table for faster access.
- [ ] **Weighted Linked List**: A linked list that assigns weights to nodes for priority-based operations.
- [ ] **Sorted Linked List**: A linked list that maintains elements in sorted order.
- [ ] **Unrolled Linked List**: A linked list that stores multiple elements in each node.
