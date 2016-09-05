package splaytree

func (tree *SplayTree) Insert(item Item) bool {
	if item == nil {
		panic("SplayTree Insert nil")
	}
	return tree.insertReplace(item, false)
}

func (tree *SplayTree) Replace(item Item) bool {
	if item == nil {
		panic("SplayTree Replace nil")
	}
	return tree.insertReplace(item, true)
}

func (tree *SplayTree) InsertAll(items []Item) int {
	if items == nil {
		panic("SplayTree InsertAll nil")
	}
	num := 0
	for _, item := range items {
		if tree.insertReplace(item, false) {
			num += 1
		}
	}
	return num
}

func (tree *SplayTree) ReplaceAll(items []Item) int {
	if items == nil {
		panic("SplayTree ReplaceAll nil")
	}
	num := 0
	for _, item := range items {
		if tree.insertReplace(item, true) {
			num += 1
		}
	}
	return num
}

func (tree *SplayTree) insertReplace(item Item, replace bool) bool {
	if tree.root == nil {
		tree.root = newNode(item)
		return !replace
	}
	tree.splay(item)
	root := tree.root
	if item.Less(root.item) {
		node := newNode(item)
		node.left = root.left
		node.right = root
		root.left = nil
		tree.root = node
		return !replace
	} else if root.item.Less(item) {
		node := newNode(item)
		node.right = root.right
		node.left = root
		root.right = nil
		tree.root = node
		return !replace
	} else if replace {
		root.item = item
		return true
	} else {
		return false
	}
}
