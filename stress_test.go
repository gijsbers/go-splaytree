package splaytree

import "testing"
import "time"
import "math/rand"
import "sort"

func TestStress(t *testing.T) {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rand := rand.New(source)
	atLeast := func(lower int) int {
		return lower + rand.Intn(lower)
	}
	findSorted := func(x int, a []int, n int) (int, bool) {
		i := sort.Search(n, func(i int) bool { return x <= a[i] })
		return i, i < n && x == a[i]
	}

	for outer := 0; outer < 10; outer++ {
		tree := NewSplayTree()
		size := atLeast(3000)
		nums := rand.Perm(size)
		for i := 0; i < size; i++ {
			nums[i] *= 2
		}

		for i := 0; i < size; i++ {
			if !tree.Insert(Int(nums[i])) {
				t.Errorf("insert failed %v %v", i, nums[i])
			}
		}
		sort.Ints(nums)

		doit := atLeast(2 * size)
		dels := make([]bool, size)
		for i := 0; i < doit; i++ {
			k := rand.Intn(2 * size)
			f, have := findSorted(k, nums, size)
			look := tree.Lookup(Int(k)) != nil
			if look != have && (!have || !dels[f]) {
				t.Errorf("lookup failed %v %v", look, have)
			} else if look {
				r := rand.Intn(1000)
				if r < 250 {
					del := tree.Delete(Int(k))
					if del == nil {
						t.Errorf("delete failed %v %v %v", look, have, dels[f])
					}
					if del != nil {
						dels[f] = true
					}
				}
			} else if have && dels[f] {
				r := rand.Intn(1000)
				if r < 100 {
					if !tree.Insert(Int(k)) {
						t.Errorf("insert failed %v %v %v", look, have, dels[f])
					} else {
						dels[f] = false
					}
				} else if r < 200 {
					if tree.Replace(Int(k)) {
						t.Errorf("replace failed %v %v %v", look, have, dels[f])
					} else {
						dels[f] = false
					}
				}
			} else if del := tree.Delete(Int(k)); del != nil {
				t.Errorf("delete should have failed %v %v %v", look, have, dels[f])
			}
		}
	}
}
