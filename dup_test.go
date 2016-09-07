package splaytree

import "testing"

func TestDuplicate(t *testing.T) {
	tree := NewSplayTree()
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	tree.InsertAll(items)
	tree.Check()
	dup := tree.Duplicate().(*SplayTree)
	dup.Check()
	if tree.Count() != dup.Count() {
		t.Errorf("dup count")
	}
	if tree.Height() != dup.Height() {
		t.Errorf("dup height")
	}
	for _, item := range items {
		if tree.Lookup(item) == nil {
			t.Errorf("tree lookup %v", item)
		}
		if dup.Lookup(item) == nil {
			t.Errorf("dup lookup %v", item)
		}
		tree.Check()
		dup.Check()
	}
}
