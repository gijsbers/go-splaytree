package splaytree

// Supremum finds the smallest element greater than a given.
// Return the supremum if present, else nil.
func (tree *SplayTree) Supremum(item Item) Item {
	if item == nil || tree.root == nil {
		return nil
	}
	tree.splay(item)
	if item.Less(tree.root.item) {
		return tree.root.item
	}
	if tree.root.right != nil {
		node := tree.root.right
		for node.left != nil {
			node = node.left
		}
		return node.item
	}
	return nil
}

// Infimum finds the largest element smaller than a given.
// Return the infimum if present, else nil.
func (tree *SplayTree) Infimum(item Item) Item {
	if item == nil || tree.root == nil {
		return nil
	}
	tree.splay(item)
	if tree.root.item.Less(item) {
		return tree.root.item
	}
	if tree.root.left != nil {
		node := tree.root.left
		for node.right != nil {
			node = node.right
		}
		return node.item
	}
	return nil
}
