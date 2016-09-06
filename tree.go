package splaytree

// A splay tree.
type SplayTree struct {
	root *node
}

// Compile time test if SplayTree fully implements Interface.
var _ Interface = (*SplayTree)(nil)

// Remove all elements of the tree.
func (tree *SplayTree) Clear() {
	tree.root = nil
}

// Count the number of elements in the tree.
func (tree *SplayTree) Count() int {
	return tree.root.count()
}

// Compute the height of the tree.
// This is the number of steps from
// the root to the farthest element.
// An empty tree has height -1.
func (tree *SplayTree) Height() int {
	return -1 + tree.root.height()
}

// Test if the tree contains at least one element.
func (tree *SplayTree) NonEmpty() bool {
	return tree.root != nil
}

// Give the element which is currently at the root.
// Return a pair (nil, false) if the tree is empty.
func (tree *SplayTree) Root() (Item, bool) {
	if tree.root == nil {
		return nil, false
	}
	return tree.root.item, true
}
