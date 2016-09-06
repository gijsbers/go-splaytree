package splaytree

// Clone a tree.
func (tree *SplayTree) Duplicate() Interface {
	copy := NewTree()
	copy.root = tree.root.duplicate()
	return copy
}
