# rbtree
Red Black Tree in Go

Based on Algorithms 4th Ed. by Sedgewick

Check out rbtree_test.go for usage.

I'm adding concurrency features, aim is to have most operations via channels for thread safety. I'll probably end up having another package for concurrent RBTrees later.


Right now not thread safe, but well tested for single thread usage.

rbtree.Keys() and rbtree.KeysInRange() is good for returning an ordered slice of whatever you've saved in the tree.

rbtree.KeysCh() and rbtree.KeysInRangeCh() is good for iterating through the tree in order, without the cost of creating a slice.

