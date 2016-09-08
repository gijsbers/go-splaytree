package splaytree

// Iterator creates a new iterator for this tree.
// On each call, the iterator gives the current item
// and advances to the next node.
// The items are returned sorted
// from the smallest to the largest item.
// When all items have been visited nil is returned.
//
func (tree *SplayTree) Iterator() func() Item {
	stack := newStack()
	for node := tree.root; node != nil; node = node.left {
		stack.push(node)
	}
	return func() Item {
		if stack.empty() {
			return nil
		}
		node := stack.pop()
		item := node.item
		node = node.right
		if node != nil {
			for node.left != nil {
				node = node.left
			}
			stack.push(node)
		}
		return item
	}
}

// ReverseIterator creates a new reverse iterator for this tree.
// On each call, the iterator gives the current item
// and advances to the next node.
// The items are returned in reverse sorted order
// from the largest to the smallest item.
// When all items have been visited nil is returned.
//
func (tree *SplayTree) ReverseIterator() func() Item {
	stack := newStack()
	for node := tree.root; node != nil; node = node.right {
		stack.push(node)
	}
	return func() Item {
		if stack.empty() {
			return nil
		}
		node := stack.pop()
		item := node.item
		node = node.left
		if node != nil {
			for node.right != nil {
				node = node.right
			}
			stack.push(node)
		}
		return item
	}
}

// RangeIterator creates a new iterator for this tree which
// observes lower and upper bounds on all returned items.
// A RangeIterator traverses only those parts of the tree
// which are in the given range, or leading to it.
// On each call, the iterator gives the current item,
// which will be in the range [lower, upper],
// and advances the internal state to the next node.
// The items are returned in ascending sorted sequence.
// When all items have been visited nil is returned.
// After having used a range iterator one should do a tree
// lookup on the last returned non-nil item in order
// to preserve optimal theoretical properties.
// The iterator examples illustrate this.
//
func (tree *SplayTree) RangeIterator(lower Item, upper Item) func() Item {
	inRange := func(item Item) bool {
		return !item.Less(lower) && !upper.Less(item)
	}
	stack := newStack()
	if !upper.Less(lower) && tree.root != nil {
		for node := tree.root; ; {
			if node.item.Less(lower) {
				if node.right == nil {
					tree.splay(node.item)
					break
				}
				node = node.right
			} else if upper.Less(node.item) {
				if node.left == nil {
					tree.splay(node.item)
					break
				}
				node = node.left
			} else {
				for node != nil && inRange(node.item) {
					stack.push(node)
					node = node.left
				}
				break
			}
		}
	}
	return func() Item {
		if stack.empty() {
			return nil
		}
		node := stack.pop()
		item := node.item
		node = node.right
		if node != nil && inRange(node.item) {
			for node.left != nil && inRange(node.left.item) {
				node = node.left
			}
			stack.push(node)
		}
		return item
	}
}
