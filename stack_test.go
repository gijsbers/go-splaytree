package splaytree

import "testing"

func TestStack(t *testing.T) {
	stack := newStack()
	if stack == nil {
		t.Errorf("newStack nil")
	}
	if stack.length() != 0 {
		t.Errorf("newStack length !0")
	}
	if stack.empty() != true {
		t.Errorf("newStack empty !1")
	}
	head := &node{}
	stack.push(head)
	if stack.length() != 1 {
		t.Errorf("newStack length !1")
	}
	if stack.empty() != false {
		t.Errorf("newStack empty !0")
	}
	if stack.tos() != head {
		t.Errorf("newStack tos !head")
	}
	tail := &node{}
	stack.push(tail)
	if stack.length() != 2 {
		t.Errorf("newStack length !2")
	}
	if stack.empty() != false {
		t.Errorf("newStack empty !0")
	}
	if stack.tos() != tail {
		t.Errorf("newStack tos !tail")
	}
	if stack.pop() != tail {
		t.Errorf("newStack pop !tail")
	}
	if stack.length() != 1 {
		t.Errorf("newStack length !1")
	}
	if stack.empty() != false {
		t.Errorf("newStack empty !0")
	}
	if stack.tos() != head {
		t.Errorf("newStack tos !head")
	}
	if stack.pop() != head {
		t.Errorf("newStack pop !head")
	}
	if stack.length() != 0 {
		t.Errorf("newStack length !0")
	}
	if stack.empty() != true {
		t.Errorf("newStack empty !2")
	}
}
