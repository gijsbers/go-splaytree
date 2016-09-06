package splaytree

import "testing"

func TestNewSplayTree(t *testing.T) {
	tree := NewSplayTree()
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	if tree.root != nil {
		t.Errorf("new root != nil")
	}
}
