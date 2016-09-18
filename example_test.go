package splaytree

import "fmt"

// The item type which we use in all our tests
type Int int

// Meet the contract obligations for our item type
func (i Int) Less(than Item) bool {
	return i < than.(Int)
}

func Example() {
	// The element type must define a function Less,
	// which compares two elements. For instance:
	// type Int int
	// func (i Int) Less(than Item) bool {
	// 	return i < than.(Int)
	// }
	tree := NewSplayTree()
	tree.Insert(Int(2))
	tree.InsertAll([]Item{Int(1), Int(4), Int(0), Int(3)})
	fmt.Printf("%v ", tree.Delete(tree.Root()))
	fmt.Printf("%v ", tree.DeleteMin())
	fmt.Printf("%v ", tree.DeleteMax())
	fmt.Printf("%v ", tree.Lookup(Int(3)))
	tree.Traverse(func(i Item) { fmt.Printf("%v ", i) })
	iter := tree.Iterator()
	for i := iter(); i != nil; i = iter() {
		fmt.Printf("%v ", i)
	}
	// Output: 3 0 4 <nil> 1 2 1 2
}

// Iterate over items in a tree. Abort the iteration halfway
// and do a lookup to preserve optimal theoretical properties.
func ExampleSplayTree_Iterator() {
	tree := NewSplayTree()
	tree.InsertAll([]Item{Int(2), Int(4), Int(1), Int(3), Int(6), Int(0)})
	iter := tree.Iterator()
	for item := iter(); item != nil; item = iter() {
		fmt.Printf("%v ", item)
		if item.(Int) > 2 {
			tree.Lookup(item)
			break
		}
	}
	// Output: 0 1 2 3
}

// Iterate over items in a tree in reverse order.
// Aborting the iteration halfway requires doing a lookup on the last item.
func ExampleSplayTree_ReverseIterator() {
	tree := NewSplayTree()
	tree.InsertAll([]Item{Int(10), Int(99), Int(42), Int(0), Int(66), Int(5)})
	iter := tree.ReverseIterator()
	for item := iter(); item != nil; item = iter() {
		fmt.Printf("%v ", item)
		if item.(Int) == 42 {
			tree.Lookup(item)
			break
		}
	}
	// Output: 99 66 42
}

// Iterate over items in a tree while observing lower and upper bounds.
func ExampleSplayTree_RangeIterator() {
	tree := NewSplayTree()
	tree.InsertAll([]Item{Int(2), Int(4), Int(6), Int(1), Int(5), Int(3), Int(0)})
	iter := tree.RangeIterator(Int(2), Int(4))
	var last Item
	for item := iter(); item != nil; item = iter() {
		fmt.Printf("%v ", item)
		last = item
	}
	if last != nil {
		tree.Lookup(last)
	}
	// Output: 2 3 4
}
