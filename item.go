package splaytree

// Item defines the minimal interface which is required for splay tree elements.
type Item interface {
	// Less returns true if this element is less than some other element.
	// Of course, when item.Less(other) is true,
	// then other.Less(item) must be false,
	// and vice versa.
	Less(than Item) bool
}
