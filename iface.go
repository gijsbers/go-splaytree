package splaytree

// Interface defines the public interface of splay trees
type Interface interface {
	// Remove all elements of the tree
	Clear()
	// Count the number of elements in the tree
	Count() int
	// Remove an item from the tree
	Delete(item Item) (Item, bool)
	// Remove all given items from the tree
	DeleteAll(items []Item) int
	// Remove the largest element of the tree
	DeleteMax() (Item, bool)
	// Remove the smallest element of the tree
	DeleteMin() (Item, bool)
	// Remove and return the element at the root
	DeleteRoot() (Item, bool)
	// Clone the tree
	Duplicate() Interface
	// Compute the height of the tree
	Height() int
	// Add a new item to the tree (if unique)
	Insert(item Item) bool
	// Add a number of items to the tree
	InsertAll(items []Item) int
	// Lookup an item
	Lookup(item Item) (Item, bool)
	// Give the largest element
	Max() Item
	// Give the smallest element
	Min() Item
	// Test if the tree contains at least one element
	NonEmpty() bool
	// Insert or replace an element
	Replace(item Item) bool
	// Insert or replace a number of elements
	ReplaceAll(items []Item) int
	// Give the current root element
	Root() (Item, bool)
}
