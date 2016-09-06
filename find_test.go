package splaytree

import "testing"

func TestMin(t *testing.T) {
	tree := NewSplayTree()
	tree.CheckBinarySearchTree()
	if tree.Min() != nil {
		t.Errorf("tree min !nil")
	}
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	min := Int(99)
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.CheckBinarySearchTree()
		if item.Less(min) {
			min = item.(Int)
		}
		if tree.Min() != min {
			t.Errorf("tree min !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
}

func TestMax(t *testing.T) {
	tree := NewSplayTree()
	tree.CheckBinarySearchTree()
	if tree.Max() != nil {
		t.Errorf("tree max !nil")
	}
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	max := Int(0)
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.CheckBinarySearchTree()
		if max.Less(item) {
			max = item.(Int)
		}
		if tree.Max() != max {
			t.Errorf("tree max !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
}

func TestLookup(t *testing.T) {
	tree := NewSplayTree()
	tree.CheckBinarySearchTree()
	if i, b := tree.Lookup(nil); b || i != nil {
		t.Errorf("tree find !nil %v %v", b, i)
	}
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		tree.CheckBinarySearchTree()
		if i, b := tree.Lookup(item); !b || i == nil || i != item {
			t.Errorf("tree lookup !%v", item)
		}
		tree.CheckBinarySearchTree()
	}
	miss := []Item{Int(1), Int(0), Int(9)}
	for _, item := range miss {
		if i, b := tree.Lookup(item); b || i != nil {
			t.Errorf("tree lookup !!%v %v %v", item, b, i)
		}
		tree.CheckBinarySearchTree()
	}
}
