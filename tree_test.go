package splaytree

import "testing"

func TestClear(t *testing.T) {
	tree := NewSplayTree()
	tree.InsertAll([]Item{Int(6), Int(4), Int(3), Int(1)})
	tree.Check()
	tree.Clear()
	if tree.root != nil {
		t.Errorf("root != nil")
	}
}

func TestCount(t *testing.T) {
	tree := NewSplayTree()
	if tree.Count() != 0 {
		t.Errorf("count !0")
	}
	tree.InsertAll([]Item{Int(6), Int(4), Int(3), Int(1)})
	tree.Check()
	if tree.Count() != 4 {
		t.Errorf("count !4")
	}
	tree.DeleteAll([]Item{Int(3), Int(2)})
	if tree.Count() != 3 {
		t.Errorf("count !3")
	}
	tree.Clear()
	if tree.Count() != 0 {
		t.Errorf("count !0")
	}
}

func TestHeight(t *testing.T) {
	tree := NewSplayTree()
	if tree.Height() != -1 {
		t.Errorf("height !-1")
	}
	tree.Insert(Int(3))
	if tree.Height() != 0 {
		t.Errorf("height !0")
	}
	tree.Insert(Int(3))
	if tree.Height() != 0 {
		t.Errorf("height !0")
	}
	tree.Insert(Int(1))
	if h := tree.Height(); h != 1 {
		t.Errorf("height !1 %v %v", h, tree.Count())
	}
	tree.Insert(Int(2))
	if h := tree.Height(); h != 1 {
		t.Errorf("height !1 %v %v", h, tree.Count())
	}
	tree.Insert(Int(9))
	if h := tree.Height(); h != 2 && h != 3 {
		t.Errorf("height !<2 %v %v", h, tree.Count())
	}
	tree.Check()
}

func TestNonEmpty(t *testing.T) {
	tree := NewSplayTree()
	if tree.NonEmpty() != false {
		t.Errorf("NonEmpty !false")
	}
	tree.Insert(Int(2))
	if tree.NonEmpty() != true {
		t.Errorf("NonEmpty !true")
	}
	tree.Clear()
	if tree.NonEmpty() {
		t.Errorf("NonEmpty clear")
	}
}

func TestRoot(t *testing.T) {
	tree := NewSplayTree()
	if r, b := tree.Root(); b != false || r != nil {
		t.Errorf("Root !false")
	}
	tree.Insert(Int(2))
	if r, b := tree.Root(); b == false || r == nil {
		t.Errorf("Root !true")
	}
	tree.Clear()
	if r, b := tree.Root(); b != false || r != nil {
		t.Errorf("Root !false")
	}
}
