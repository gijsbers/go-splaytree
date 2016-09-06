package splaytree

import "testing"

func TestMin(t *testing.T) {
	tree := NewSplayTree()
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	min := Int(99)
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		if item.Less(min) {
			min = item.(Int)
		}
		if tree.Min() != min {
			t.Errorf("tree min !%v", item)
		}
	}
}

func TestMax(t *testing.T) {
	tree := NewSplayTree()
	items := []Item{Int(3), Int(5), Int(7), Int(2), Int(4), Int(6), Int(8)}
	max := Int(0)
	for _, item := range items {
		if tree.Insert(item) != true {
			t.Errorf("tree insert %v", item)
		}
		if max.Less(item) {
			max = item.(Int)
		}
		if tree.Max() != max {
			t.Errorf("tree max !%v", item)
		}
	}
}
