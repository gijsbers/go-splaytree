package splaytree

// Join two trees with cost O(N + M), which is optimal.
// After the join all unique elements of both trees are
// in the first tree. The second tree becomes empty.
func (tree *SplayTree) Join(oak *SplayTree) {
	if oak == nil || tree == oak {
		return
	}
	tree.root = join(tree.root, oak.root)
	oak.root = nil
}

func join(fir, oak *node) *node {
	if fir == nil {
		return oak
	} else if oak == nil {
		return fir
	}
	if fir.item.Less(oak.item) {
		oak = join(fir.right, oak)
		fir.right = nil
		oak.left = join(fir, oak.left)
		return oak
	} else if oak.item.Less(fir.item) {
		fir = join(fir, oak.right)
		oak.right = nil
		fir.left = join(fir.left, oak)
		return fir
	} else {
		fir.left = join(fir.left, oak.left)
		fir.right = join(fir.right, oak.right)
		return fir
	}
}
