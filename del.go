package splaytree

// DeleteRoot removes the element which is currently at the root.
// Return nil if the tree was empty, else the root item.
func (tree *SplayTree) DeleteRoot() Item {
	node := tree.root
	if node == nil {
		return nil
	}
	found := node.item
	if node.left == nil {
		tree.root = node.right
	} else if node.right == nil {
		tree.root = node.left
	} else {
		temp := node.right
		tree.root = node.left
		tree.splay(found)
		tree.root.right = temp
	}
	return found
}

// Delete an item from the tree.
// Return the deleted item if it was found, else nil.
func (tree *SplayTree) Delete(item Item) Item {
	if item == nil || tree.root == nil {
		return nil
	}
	tree.splay(item)
	if item.Less(tree.root.item) || tree.root.item.Less(item) {
		return nil
	}
	return tree.DeleteRoot()
}

// DeleteAll deletes all given items.
// Return the number of deleted items.
func (tree *SplayTree) DeleteAll(items []Item) int {
	num := 0
	for _, item := range items {
		if tree.Delete(item) != nil {
			num++
		}
	}
	return num
}

// DeleteMin deletes the smallest element from the tree.
// Return the smallest item if the tree was non-empty, else nil.
func (tree *SplayTree) DeleteMin() Item {
	node := tree.root
	if node == nil {
		return nil
	}
	if node.left == nil {
		tree.root = node.right
	} else {
		var parent = node
		node = node.left
		for node.left != nil {
			parent = node
			node = node.left
		}
		parent.left = node.right
		tree.splay(parent.item)
	}
	return node.item
}

// DeleteMax deletes the largest element from the tree.
// Return the largest item if the tree was non-empty, else nil.
func (tree *SplayTree) DeleteMax() Item {
	node := tree.root
	if node == nil {
		return nil
	}
	if node.right == nil {
		tree.root = node.left
	} else {
		var parent = node
		node = node.right
		for node.right != nil {
			parent = node
			node = node.right
		}
		parent.right = node.left
		tree.splay(parent.item)
	}
	return node.item
}
