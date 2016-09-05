package splaytree

import "testing"

func TestMaxMax(t *testing.T) {
	if max(1, 3) != 3 {
		t.Errorf("max 3")
	}
	if max(4, 2) != 4 {
		t.Errorf("max 4")
	}
}
