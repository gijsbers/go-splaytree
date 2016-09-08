package splaytree

import "testing"
import "math/rand"
import "time"

func TestSupremum(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []Item {
		nums := rand.Perm(upper)
		items := make([]Item, upper)
		for i := 0; i < upper; i++ {
			items[i] = Int(nums[i])
		}
		return items
	}

	tree := NewSplayTree()
	if tree.Supremum(Int(-1)) != nil {
		t.Errorf("supremum -1")
	}

	for i := 1; i < 12; i++ {
		tree = NewSplayTree()
		tree.InsertAll(perm(i))
		for k := -1; k <= i; k++ {
			sup := tree.Supremum(Int(k))
			if sup == nil && i > 0 && k+1 < i {
				t.Errorf("supremum %v %v nil", i, k)
			}
			if sup != nil && sup.(Int) != Int(k+1) {
				t.Errorf("supremum %v %v = %v", i, k, sup)
			}
		}
	}
}

func TestInfimum(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := func(upper int) []Item {
		nums := rand.Perm(upper)
		items := make([]Item, upper)
		for i := 0; i < upper; i++ {
			items[i] = Int(nums[i])
		}
		return items
	}

	tree := NewSplayTree()
	if tree.Infimum(Int(-1)) != nil {
		t.Errorf("infimum -1")
	}

	for i := 1; i < 12; i++ {
		tree = NewSplayTree()
		tree.InsertAll(perm(i))
		for k := -1; k <= i; k++ {
			inf := tree.Infimum(Int(k))
			if inf == nil && k > 0 {
				t.Errorf("infimum %v %v nil", i, k)
			}
			if inf != nil && inf.(Int) != Int(k-1) {
				t.Errorf("infimum %v %v = %v", i, k, inf)
			}
		}
	}
}
