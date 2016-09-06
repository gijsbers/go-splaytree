package splaytree

// Clone a tree. The internal tree structure is duplicated
// and the used items in the cloned tree are identical.
func (tree *SplayTree) Duplicate() Interface {
	copy := NewSplayTree()
	copy.root = tree.root.duplicate()
	return copy
}
