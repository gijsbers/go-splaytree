package splaytree

// This is the interface which is required for splay tree elements.
type Item interface {
	// Return true if this element is less than some other element.
	// Of course, if item.Less(other) is true, than other.Less(item) must be false,
	// and vice versa.
	Less(than Item) bool
}
