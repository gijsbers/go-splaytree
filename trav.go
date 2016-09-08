package splaytree

// Traverse a tree inorder with a function.
func (tree *SplayTree) Traverse(apply func(Item)) {
	tree.root.traverse(apply)
}

func (node *node) traverse(apply func(Item)) {
	if node != nil {
		node.left.traverse(apply)
		apply(node.item)
		node.right.traverse(apply)
	}
}
