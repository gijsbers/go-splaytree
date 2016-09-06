package splaytree

type node struct {
	left, right *node
	item        Item
}

func newNode(item Item) *node {
	return &node{item: item}
}

func (node *node) count() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.count() + node.right.count()
}

func (node *node) height() int {
	if node == nil {
		return 0
	}
	return 1 + max(node.left.height(), node.right.height())
}

func (node *node) duplicate() *node {
	if node == nil {
		return nil
	}
	copy := newNode(node.item)
	copy.left = node.left.duplicate()
	copy.right = node.right.duplicate()
	return copy
}
