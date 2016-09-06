package splaytree

// Add a new item to the tree, if it is not yet in the tree.
// Return true if the element was added, else false.
func (tree *SplayTree) Insert(item Item) bool {
	if item == nil {
		panic("SplayTree Insert nil")
	}
	return tree.insertReplace(item, false)
}

// Insert or replace an element.
// Return true if it was replaced, false if inserted.
func (tree *SplayTree) Replace(item Item) bool {
	if item == nil {
		panic("SplayTree Replace nil")
	}
	return tree.insertReplace(item, true)
}

// Insert a number of elements.
// Return the number of added items,
// i.e. the increase in size of the tree.
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

// Insert or replace a number of elements.
// Return the number of elements which replaced existing items.
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
