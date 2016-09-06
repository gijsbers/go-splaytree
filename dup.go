package splaytree

// Duplicate clones a tree. The internal tree structure is
// duplicated, while the used items in the cloned tree
// are identical with the original tree.
func (tree *SplayTree) Duplicate() Interface {
	copy := NewSplayTree()
	copy.root = tree.root.duplicate()
	return copy
}
