package splaytree

import "testing"

// Iterate over items in a tree while observing lower and upper bounds.
func TestRangeIterator(t *testing.T) {
	tree := NewSplayTree()
	items := []Item{Int(2), Int(4), Int(6), Int(1), Int(5), Int(3), Int(0)}
	tree.InsertAll(items)
	for lkup := range items {
		tree.Lookup(Int(lkup))
		lower := Int(2)
		upper := Int(4)
		iter := tree.RangeIterator(lower, upper)
		for item := iter(); item != nil; item = iter() {
			if item.Less(lower) || upper.Less(item) {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = Int(-10)
		upper = Int(-1)
		iter = tree.RangeIterator(lower, upper)
		for item := iter(); item != nil; item = iter() {
			if item.Less(lower) || upper.Less(item) {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = Int(-1)
		upper = Int(3)
		iter = tree.RangeIterator(lower, upper)
		for item := iter(); item != nil; item = iter() {
			if item.Less(lower) || upper.Less(item) {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = Int(3)
		upper = Int(9)
		iter = tree.RangeIterator(lower, upper)
		for item := iter(); item != nil; item = iter() {
			if item.Less(lower) || upper.Less(item) {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
		lower = Int(9)
		upper = Int(29)
		iter = tree.RangeIterator(lower, upper)
		for item := iter(); item != nil; item = iter() {
			if item.Less(lower) || upper.Less(item) {
				t.Errorf("RangeIterator item %v ![%v, %v]", item, lower, upper)
			}
		}
	}
}
