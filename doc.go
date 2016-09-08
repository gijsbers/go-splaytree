// Package splaytree implements the splay tree data structure,
// which is a self-balancing binary search tree, that is,
// it adapts its internal tree structure to how it is
// being used in order to optimize future operations.
// The tree contains only unique keys in sorted order.
// The self-balancing (called splaying) is done for every
// insert, lookup or remove operation. Splaying is heavily
// optimized in a single loop. The key which is
// inserted/looked up/removed is splayed upwards to the
// root, by means of rotations over two or three nodes.
// The effect is that future accesses to this key and to
// its neighbors become cheap, as they will be at or near
// the root of the tree. Accesses to other non-neighboring
// keys diminish this benefit over time. On average the
// cost of accesses is optimal: O(log N). Splay trees
// are especially beneficial in applications which exhibit
// locality of reference. I.e. when accesses to the tree are
// related in location or time. This happens for instance
// for sequential (sorted) or clustered access patterns.
// See https://en.wikipedia.org/wiki/Splay_tree for details.
package splaytree
