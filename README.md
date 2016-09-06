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

##References

"Performance Analysis of BSTs in System Software"
[link](http://benpfaff.org/papers/libavl.pdf) by Ben Pfaff
compares 20 variants of BSTs and concludes that splay trees
performs best when accesses are sequential or clustered.

##Installation

        $ go get github.com/gijsbers/go-splaytree

##Copyright and License

Copyright ©2016 Bert Gijsbers except where otherwise noted. All rights reserved.
Use of this source code is governed by an Apache 2.0 license
that can be found in the LICENSE file.
