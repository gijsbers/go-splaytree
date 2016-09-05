package splaytree

func (tree *SplayTree) Delete(item Item) (Item, bool) {
	if item == nil || tree.root == nil {
		return nil, false
	}
	tree.splay(item)
	if item.Less(tree.root.item) || tree.root.item.Less(item) {
		return nil, false
	}
	found := tree.root.item
	if tree.root.left == nil {
		tree.root = tree.root.right
	} else {
		temp := tree.root.right
		tree.root = tree.root.left
		tree.splay(item)
		tree.root.right = temp
	}
	return found, true
}

func (tree *SplayTree) DeleteAll(items []Item) int {
	num := 0
	for _, item := range items {
		if _, bo := tree.Delete(item); bo {
			num += 1
		}
	}
	return num
}

func (tree *SplayTree) DeleteMin() (Item, bool) {
	node := tree.root
	if node == nil {
		return nil, false
	}
	for node.left != nil {
		node = node.left
	}
	return tree.Delete(node.item)
}

func (tree *SplayTree) DeleteMax() (Item, bool) {
	node := tree.root
	if node == nil {
		return nil, false
	}
	for node.right != nil {
		node = node.right
	}
	return tree.Delete(node.item)
}

func (tree *SplayTree) DeleteRoot() (Item, bool) {
	if tree.root == nil {
		return nil, false
	}
	return tree.Delete(tree.root.item)
}
