package splaytree

// Min gives the smallest element in the tree.
// If the tree is empty then nil is returned.
func (tree *SplayTree) Min() Item {
	node := tree.root
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	item := node.item
	tree.splay(item)
	return item
}

// Max gives the largest element in the tree.
// If the tree is empty then nil is returned.
func (tree *SplayTree) Max() Item {
	node := tree.root
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	item := node.item
	tree.splay(item)
	return item
}

// Lookup an item and return the found item.
// If the tree is empty then nil is returned.
func (tree *SplayTree) Lookup(item Item) Item {
	if item == nil || tree.root == nil {
		return nil
	}
	tree.splay(item)
	if item.Less(tree.root.item) || tree.root.item.Less(item) {
		return nil
	}
	return tree.root.item
}
