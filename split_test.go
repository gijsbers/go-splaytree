package splaytree

import "testing"
import "time"
import "math/rand"

func TestSplit(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []Item {
		nums := rand.Perm(upper)
		items := make([]Item, upper)
		for i := 0; i < upper; i++ {
			items[i] = Int(nums[i])
		}
		return items
	}
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

	tree := NewSplayTree()
	lef, rig := tree.Split(Int(4))
	if lef.NonEmpty() || rig.NonEmpty() {
		t.Errorf("split nonempty")
	}

	for k := 1; k < 20; k++ {
		items := perm(k)
		for i := 0; i <= k; i++ {
			tree := NewSplayTree()
			tree.InsertAll(items)
			lef, rig := tree.Split(Int(i))
			lef.Traverse(lessThan(i))
			rig.Traverse(atLeast(i))
		}
	}
}
