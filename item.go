package splaytree

// This is the interface which is required for tree elements.
type Item interface {
	// The tree element type must define this comparison.
	// Return true if this element is less than the given element.
	Less(than Item) bool
}
