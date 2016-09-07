package splaytree

// Split a tree in two at a given boundary.
// Return two trees as a pair (left, right),
// where left contains all items less than the boundary
// and right contains only items at or beyond the boundary.
func (tree *SplayTree) Split(item Item) (*SplayTree, *SplayTree) {
	more := NewSplayTree()
	if item == nil || tree.root == nil {
		return tree, more
	}
	tree.splay(item)
	if tree.root.item.Less(item) {
		more.root = tree.root.right
		tree.root.right = nil
	} else {
		more.root = tree.root
		tree.root = more.root.left
		more.root.left = nil
	}
	return tree, more
}
