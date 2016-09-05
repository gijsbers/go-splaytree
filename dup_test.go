package splaytree

import "testing"

func TestDuplicate(t *testing.T) {
	tree := NewTree()
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	tree.InsertAll(items)
	dup := tree.Duplicate()
	if tree.Count() != dup.Count() {
		t.Errorf("dup count")
	}
	if tree.Height() != dup.Height() {
		t.Errorf("dup height")
	}
	for _, item := range items {
		if _, bl := tree.Lookup(item); !bl {
			t.Errorf("tree lookup %v", item)
		}
		if _, bl := dup.Lookup(item); !bl {
			t.Errorf("dup lookup %v", item)
		}
	}
}
