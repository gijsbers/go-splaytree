// Package splaytree implements splay trees.
// A splay tree is a self-balancing binary search tree.
// The tree contains only unique keys in sorted order.
// The self-balancing (called splaying) is done for every insert,
// lookup or remove operation. The key which is
// inserted/looked up/removed is splayed upwards to the root,
// by means of tree rotations. The effect is that future
// accesses to this key and to its neighbors become cheap,
// as they will be at or near the root of the tree.
// Accesses to other non-neighboring keys diminish this benefit over time.
// On average the cost of accesses is optimal: O(log N).
// Splay trees are especially beneficial in applications
// which exhibit locality of reference. I.e. when accesses
// to the tree are related in location or time.
// See https://en.wikipedia.org/wiki/Splay_tree
package splaytree
