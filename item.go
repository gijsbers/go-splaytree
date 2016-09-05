package splaytree

type Item interface {
	Less(than Item) bool
}
