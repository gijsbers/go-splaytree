package splaytree

type SplayTree struct {
	root *node
}

var _ Interface = (*SplayTree)(nil)

func (tree *SplayTree) Clear() {
	tree.root = nil
}

func (tree *SplayTree) Count() int {
	return tree.root.count()
}

func (tree *SplayTree) Height() int {
	return -1 + tree.root.height()
}

func (tree *SplayTree) NonEmpty() bool {
	return tree.root != nil
}

func (tree *SplayTree) Root() (Item, bool) {
	if tree.root == nil {
		return nil, false
	}
	return tree.root.item, true
}
