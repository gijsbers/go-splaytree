package splaytree

import "testing"

func TestDeleteEmpty(t *testing.T) {
	tree := NewSplayTree()
	if tree.DeleteRoot() != nil {
		t.Errorf("DeleteRoot nil")
	}
	if tree.DeleteMax() != nil {
		t.Errorf("DeleteMax nil")
	}
	if tree.DeleteMin() != nil {
		t.Errorf("DeleteMin nil")
	}
	if tree.DeleteAll([]Item{}) != 0 {
		t.Errorf("DeleteAll nil")
	}
}

func TestDelete(t *testing.T) {
	tree := NewSplayTree()
	items := []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)}
	tree.InsertAll(items)
	tree.Check()
	if n := tree.Count(); n != len(items) {
		t.Errorf("delete all count != %v", n)
	}
	if i := tree.DeleteMax(); i == nil || i.(Int) != 7 || tree.Count() != len(items)-1 {
		t.Errorf("delete !max")
	}
	tree.Check()
	if i := tree.DeleteMin(); i == nil || i.(Int) != 0 || tree.Count() != len(items)-2 {
		t.Errorf("delete !min")
	}
	tree.Check()
	r := tree.root.item
	if i := tree.DeleteRoot(); i == nil || i != r || tree.Count() != len(items)-3 {
		t.Errorf("delete !root")
	}
	tree.Check()
	if n := tree.DeleteAll(items); n != len(items)-3 || tree.Count() != 0 {
		t.Errorf("delete !all")
	}
	tree.Check()
}

func TestDeleteMin(t *testing.T) {
	tree := NewSplayTree()
	if i := tree.DeleteMin(); i != nil {
		t.Errorf("DeleteMin nil")
	}
	tree.Insert(Int(3))
	if i := tree.DeleteMin(); i == nil || i.(Int) != 3 || tree.Count() != 0 {
		t.Errorf("DeleteMin 3")
	}
	tree.Insert(Int(3))
	tree.Insert(Int(2))
	if i := tree.DeleteMin(); i == nil || i.(Int) != 2 || tree.Count() != 1 {
		t.Errorf("DeleteMin 3 2")
	}
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	if i := tree.DeleteMin(); i == nil || i.(Int) != 2 || tree.Count() != 1 {
		t.Errorf("DeleteMin 2 3")
	}
	tree.Insert(Int(3))
	tree.Insert(Int(2))
	tree.Insert(Int(1))
	if i := tree.DeleteMin(); i == nil || i.(Int) != 1 || tree.Count() != 2 {
		t.Errorf("DeleteMin 3 2 1")
	}
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	tree.Insert(Int(1))
	if i := tree.DeleteMin(); i == nil || i.(Int) != 1 || tree.Count() != 2 {
		t.Errorf("DeleteMin 2 3 1")
	}
}

func TestDeleteMax(t *testing.T) {
	tree := NewSplayTree()
	if i := tree.DeleteMax(); i != nil {
		t.Errorf("DeleteMax nil")
	}
	tree.Insert(Int(3))
	if i := tree.DeleteMax(); i == nil || i.(Int) != 3 || tree.Count() != 0 {
		t.Errorf("DeleteMax 3")
	}
	tree.Insert(Int(3))
	tree.Insert(Int(2))
	if i := tree.DeleteMax(); i == nil || i.(Int) != 3 || tree.Count() != 1 {
		t.Errorf("DeleteMax 3 2")
	}
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	if i := tree.DeleteMax(); i == nil || i.(Int) != 3 || tree.Count() != 1 {
		t.Errorf("DeleteMax 2 3")
	}
	tree.Insert(Int(3))
	tree.Insert(Int(2))
	tree.Insert(Int(1))
	if i := tree.DeleteMax(); i == nil || i.(Int) != 3 || tree.Count() != 2 {
		t.Errorf("DeleteMax 3 2 1")
	}
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	tree.Insert(Int(1))
	if i := tree.DeleteMax(); i == nil || i.(Int) != 3 || tree.Count() != 2 {
		t.Errorf("DeleteMax 2 3 1")
	}
}
