package splaytree

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

func (tree *SplayTree) Lookup(item Item) (Item, bool) {
	if item == nil || tree.root == nil {
		return nil, false
	}
	tree.splay(item)
	if item.Less(tree.root.item) || tree.root.item.Less(item) {
		return nil, false
	}
	return tree.root.item, true
}
