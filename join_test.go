package splaytree

import "testing"

func TestJoin(t *testing.T) {
	fir := NewSplayTree()
	oak := NewSplayTree()
	it1 := []Item{Int(6), Int(4), Int(2), Int(5), Int(3), Int(7), Int(0)}
	it2 := []Item{Int(-1), Int(14), Int(-2), Int(15), Int(-3), Int(17), Int(10)}
	fir.InsertAll(it1)
	oak.InsertAll(it1)
	fir.Join(oak)
	if fir.Count() != len(it1) {
		t.Errorf("join count !it1")
	}
	if oak.NonEmpty() {
		t.Errorf("join oak !0")
	}
	fir.CheckBinarySearchTree()
	oak.InsertAll(it2)
	fir.Join(oak)
	if fir.Count() != len(it1)+len(it2) {
		t.Errorf("join count !it1+it2")
	}
	if oak.NonEmpty() {
		t.Errorf("join oak !0")
	}
	fir.CheckBinarySearchTree()
	for _, item := range append(it1, it2...) {
		if _, ok := fir.Lookup(item); !ok {
			t.Errorf("join lookup !%v", item)
		}
	}
	fir.CheckBinarySearchTree()
}
