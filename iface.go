package splaytree

// Interface defines the public interface of splay trees
type Interface interface {
	// Remove all elements of the tree
	Clear()
	// Count the number of elements in the tree
	Count() int
	// Remove an item from the tree
	Delete(item Item) Item
	// Remove all given items from the tree
	DeleteAll(items []Item) int
	// Remove the largest element of the tree
	DeleteMax() Item
	// Remove the smallest element of the tree
	DeleteMin() Item
	// Remove and return the element at the root
	DeleteRoot() Item
	// Clone the tree
	Duplicate() Interface
	// Compute the height of the tree
	Height() int
	// Infimum gives the closest smaller item
	Infimum(item Item) Item
	// Add a new item to the tree (if unique)
	Insert(item Item) bool
	// Add a number of items to the tree
	InsertAll(items []Item) int
	// Create a new iterator over this tree.
	Iterator() func() Item
	// Join two trees with cost O(N + M), which is optimal.
	Join(other Interface)
	// Lookup an item
	Lookup(item Item) Item
	// Give the largest element
	Max() Item
	// Give the smallest element
	Min() Item
	// Test if the tree contains at least one element
	NonEmpty() bool
	// Create an iterator for a limited range of items
	RangeIterator(lower Item, upper Item) func() Item
	// Insert or replace an element
	Replace(item Item) bool
	// Insert or replace a number of elements
	ReplaceAll(items []Item) int
	// Create an iterator which iterates in descending order
	ReverseIterator() func() Item
	// Give the current root element
	Root() Item
	// Supremum gives the closest larger item
	Supremum(item Item) Item
	// Traverse a tree inorder with a function
	Traverse(func(Item))
}
