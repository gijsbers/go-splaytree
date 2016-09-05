package splaytree

type Interface interface {
	Clear()
	Count() int
	Delete(item Item) (Item, bool)
	DeleteAll(items []Item) int
	DeleteMax() (Item, bool)
	DeleteMin() (Item, bool)
	DeleteRoot() (Item, bool)
	Duplicate() Interface
	Height() int
	Insert(item Item) bool
	InsertAll(items []Item) int
	Lookup(item Item) (Item, bool)
	Max() Item
	Min() Item
	NonEmpty() bool
	Replace(item Item) bool
	ReplaceAll(items []Item) int
	Root() (Item, bool)
}
