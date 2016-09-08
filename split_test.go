package splaytree

import "testing"

func TestSplit(t *testing.T) {
	lessThan := func(bound int) func(Item) {
		return func(item Item) {
			if Int(bound).Less(item) {
				t.Errorf("split %v is more than %v", item, bound)
			}
		}
	}
	atLeast := func(bound int) func(Item) {
		return func(item Item) {
			if item.Less(Int(bound)) {
				t.Errorf("split %v is less than %v", item, bound)
			}
		}
	}

	items := []Item{Int(1), Int(5), Int(3), Int(8), Int(4), Int(9), Int(0)}
	for i := 0; i < 10; i++ {
		tree := NewSplayTree()
		tree.InsertAll(items)
		lef, rig := tree.Split(Int(i))
		lef.Traverse(lessThan(i))
		rig.Traverse(atLeast(i))
	}

	tree := NewSplayTree()
	lef, rig := tree.Split(Int(4))
	if lef.NonEmpty() || rig.NonEmpty() {
		t.Errorf("split nonempty")
	}
}
