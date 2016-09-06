# go-splaytree

Splay tree in Go.

[Splay trees](https://en.wikipedia.org/wiki/Splay_tree)
are self-balancing binary search trees.
Accessing a node in the tree causes it and its neighbours
to be splayed upwards by means of rotations.
Future accesses to this node or to its neighbours will then be cheaper.
Splay trees perform best when there is temporal or spatial
[locality](https://en.wikipedia.org/wiki/Locality_of_reference)
in the access patterns (insertions, lookups or removals).
Say, you enter person data by date+time of birth and then
you investigate specific age groups.
Or you manage virtual memory pages and your applications tend
to focus on groups of related pages. Over time they will
abandon some groups of memory pages and move on to other
groups of memory pages.

On average the cost of accesses (insertions, lookups or removals)
is O(log n). The worst case cost for one access is O(n).
Inserting a sorted sequence of 'n' elements into an empty tree
has cost O(n), which is optimal.  Retrieving that same sequence
in either ascending or descending order has again cost O(n).

##References

["Performance Analysis of BSTs in System Software"](http://benpfaff.org/papers/libavl.pdf)
by Ben Pfaff, compares the performance of 20 variants of BSTs and
concludes that splay trees perform best when accesses
are sequential or clustered.

##Installation

        $ go get github.com/gijsbers/go-splaytree

##Copyright and License

Copyright ©2016 Bert Gijsbers except where otherwise noted. All rights reserved.
Use of this source code is governed by an Apache 2.0 license
that can be found in the LICENSE file.
