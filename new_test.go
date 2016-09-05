package splaytree

import "testing"

func TestNewTree(t *testing.T) {
	tree := NewTree()
	if tree == nil {
		t.Errorf("new tree == nil")
	}
	if tree.root != nil {
		t.Errorf("new root != nil")
	}
}
