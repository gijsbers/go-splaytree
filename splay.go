package splaytree

/* Translated from a public domain Java version by Danny Sleator. */

/**
 * Internal method to perform a top-down splay.
 *
 *   splay(key) does the splay operation on the given key.
 *   If key is in the tree, then the node containing that
 *   key becomes the new root.  If key is not in the tree,
 *   then after the splay, root contains either the greatest
 *   key in the tree which is less than key, or the smallest
 *   key in the tree which is bigger than key.
 *
 *   This means, among other things, that if you splay with
 *   a key that's larger than any in the tree, the rightmost
 *   node of the tree becomes the root.  This property is used
 *   in the delete() method.
 */
func (tree *SplayTree) splay(item Item) {
	var temp node = node{}
	var lef *node = &temp
	var rig *node = &temp
	var top *node = tree.root
	var yes *node = nil
	for {
		if item.Less(top.item) {
			if top.left == nil {
				break
			}
			if item.Less(top.left.item) {
				// rotate right
				yes = top.left
				top.left = yes.right
				yes.right = top
				top = yes
				if top.left == nil {
					break
				}
			}
			// link right
			rig.left = top
			rig = top
			top = top.left
		} else if top.item.Less(item) {
			if top.right == nil {
				break
			}
			if top.right.item.Less(item) {
				// rotate left
				yes = top.right
				top.right = yes.left
				yes.left = top
				top = yes
				if top.right == nil {
					break
				}
			}
			// link left
			lef.right = top
			lef = top
			top = top.right
		} else {
			break
		}
	}
	// assemble
	lef.right = top.left
	rig.left = top.right
	top.left = temp.right
	top.right = temp.left
	tree.root = top
}
