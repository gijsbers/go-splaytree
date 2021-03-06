package splaytree

// SplayTree defines the splay tree type.
type SplayTree struct {
	root *node
}

// Compile time test if SplayTree fully implements Interface.
var _ Interface = (*SplayTree)(nil)

// Clear all elements from the tree.
func (tree *SplayTree) Clear() {
	tree.root = nil
}

// Count the number of elements in the tree.
func (tree *SplayTree) Count() int {
	return tree.root.count()
}

// Height computes the height of the tree.
// This is the number of steps from
// the root to the farthest element.
// An empty tree has height -1.
func (tree *SplayTree) Height() int {
	return -1 + tree.root.height()
}

// NonEmpty tests if the tree contains at least one element.
func (tree *SplayTree) NonEmpty() bool {
	return tree.root != nil
}

// Root gives the element which is currently at the root
// of the tree. If the tree is empty then nil is returned.
func (tree *SplayTree) Root() Item {
	if tree.root == nil {
		return nil
	}
	return tree.root.item
}
