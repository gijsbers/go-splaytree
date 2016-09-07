package splaytree

// Check verifies that the tree
// is a proper binary search tree.
func (tree *SplayTree) Check() {
	if !tree.isBinarySearchTree() {
		panic("IsBinarySearchTree failed")
	}
}

// Return true if tree is a proper binary search tree.
func (tree *SplayTree) isBinarySearchTree() bool {
	if tree.root == nil {
		return true
	}
	min, max := tree.root.isBinarySearchTree()
	return min != nil && max != nil
}

// Return the minimum and maximum element at this node.
func (node *node) isBinarySearchTree() (Item, Item) {
	min, max := node.item, node.item
	if node.left != nil {
		lmin, lmax := node.left.isBinarySearchTree()
		if lmax == nil || !lmax.Less(min) {
			return nil, nil
		}
		min = lmin
	}
	if node.right != nil {
		rmin, rmax := node.right.isBinarySearchTree()
		if rmin == nil || !max.Less(rmin) {
			return nil, nil
		}
		max = rmax
	}
	return min, max
}
