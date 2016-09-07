package splaytree

// DeleteRoot removes the element which is currently at the root.
// Return (nil, false) if the tree was empty, else (root item, true).
func (tree *SplayTree) DeleteRoot() (Item, bool) {
	node := tree.root
	if node == nil {
		return nil, false
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
	return found, true
}

// Delete an item from the tree.
// Return a pair (item, true) if the item was found
// and deleted, else (nil, false).
func (tree *SplayTree) Delete(item Item) (Item, bool) {
	if item == nil || tree.root == nil {
		return nil, false
	}
	tree.splay(item)
	if item.Less(tree.root.item) || tree.root.item.Less(item) {
		return nil, false
	}
	return tree.DeleteRoot()
}

// DeleteAll deletes all given items.
// Return the number of deleted items.
func (tree *SplayTree) DeleteAll(items []Item) int {
	num := 0
	for _, item := range items {
		if _, bo := tree.Delete(item); bo {
			num++
		}
	}
	return num
}

// DeleteMin returns the smallest element.
// Return a pair (smallest, true) if the tree was non-empty.
func (tree *SplayTree) DeleteMin() (Item, bool) {
	node := tree.root
	if node == nil {
		return nil, false
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
	return node.item, true
}

// DeleteMax returns the largest element.
// Return a pair (largest, true) if the tree was non-empty.
func (tree *SplayTree) DeleteMax() (Item, bool) {
	node := tree.root
	if node == nil {
		return nil, false
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
	return node.item, true
}
