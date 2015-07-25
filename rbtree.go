package rbtree

import (
	"log"
)

type Key interface {
	Compare(Key) int
}

type Node struct {
	key    Key
	left   *Node
	right  *Node
	colour bool
	N      int
}

const (
	RED   bool = true
	BLACK bool = false
)

func NewNode(key Key, colour bool, N int) *Node {
	return &Node{
		key:    key,
		colour: colour,
		N:      N,
	}
}

func isRed(node *Node) bool {
	if node == nil {
		return false
	} else {
		return node.colour == RED
	}
}

type RBTree struct {
	root     *Node
	compares int
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) Size() int {
	return size(tree.root)
}

func (tree *RBTree) Put(key Key) {
	tree.root = put(tree.root, key)
}

func (tree *RBTree) Contains(key Key) bool {
	return tree.Get(key) != nil
}

func (tree *RBTree) IsEmpty() bool {
	return tree.root == nil
}

func compare(a string, b string) int {
	//compares += 1
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.N
}

func put(node *Node, key Key) *Node {
	// change key's value to value if key in subtree rooted at node
	// otherwise add a new node to subtree associating key with value.
	if node == nil {
		return NewNode(key, RED, 1)
	}
	cmp := key.Compare(node.key)
	if cmp < 0 {
		node.left = put(node.left, key)
	} else if cmp > 0 {
		node.right = put(node.right, key)
	} else {
		node.key = key
	}

	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColours(node)
	}
	node.N = 1 + size(node.left) + size(node.right)

	return node
}

func rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.colour = h.colour
	h.colour = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	return x
}

func rotateRight(h *Node) *Node {
	if h == nil {
		log.Println("1 nil h")
	}
	x := h.left
	if x == nil {
		log.Println("2 nil x")
	}
	//clog.Println(x)
	h.left = x.right
	x.right = h
	x.colour = h.colour
	h.colour = RED
	x.N = h.N
	h.N = 1 + size(h.left) + size(h.right)
	// rotateRights++
	return x
}

func flipColours(h *Node) {
	h.colour = !h.colour
	h.left.colour = !h.left.colour
	h.right.colour = !h.right.colour
	//flipColours++
}

func (tree *RBTree) Get(key Key) Key {
	node := get(tree.root, key)
	if node != nil {
		return node.key
	} else {
		return nil
	}
}

func get(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.Compare(node.key)
	if cmp < 0 {
		return get(node.left, key)
	} else if cmp > 0 {
		return get(node.right, key)
	} else {
		return node
	}
}

func (tree *RBTree) Delete(key Key) {
	if !tree.Contains(key) {
		return
	}
	// if both children red, set root black
	if !isRed(tree.root.left) && !isRed(tree.root.right) {
		tree.root.colour = RED
	}
	tree.root = deleteNode(tree.root, key)
	if !tree.IsEmpty() {
		tree.root.colour = BLACK
	}
}

func deleteNode(node *Node, key Key) *Node {

	if node == nil {
		return nil
	}

	if key.Compare(node.key) < 0 {
		if !isRed(node.left) && !isRed(node.left.left) {
			node = moveRedLeft(node)
		}
		node.left = deleteNode(node.left, key)
	} else {
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if key.Compare(node.key) == 0 && (node.right == nil) {
			return nil
		}
		if !isRed(node.right) && !isRed(node.right.left) {
			node = moveRedRight(node)
		}
		if key.Compare(node.key) == 0 {
			x := min(node.right)
			node.key = x.key
			node.right = deleteMin(node.right)
		} else {
			node.right = deleteNode(node.right, key)
		}
	}
	return balance(node)
}

func (tree *RBTree) Min() Key {
	if tree.IsEmpty() {
		return nil
	} else {
		return min(tree.root).key

	}
}

func min(node *Node) *Node {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	} else {
		return min(node.left)
	}
}

func (tree *RBTree) Max() Key {
	if tree.IsEmpty() {
		return nil
	} else {
		return max(tree.root).key
	}
}

func max(node *Node) *Node {
	if node == nil {
		return nil
	} else if node.right == nil {
		return node
	} else {
		return max(node.right)
	}
}

func (tree *RBTree) Floor(key Key) Key {
	node := floor(tree.root, key)
	if node == nil {
		return nil
	} else {
		return node.key
	}
}

func floor(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.Compare(node.key)
	if cmp == 0 {
		return node
	} else if cmp < 0 {
		return floor(node.left, key)
	} else {
		t := floor(node.right, key)
		if t != nil {
			return t
		} else {
			return node
		}
	}
}

func (tree *RBTree) Ceiling(key Key) Key {
	node := ceiling(tree.root, key)
	if node == nil {
		return nil
	} else {
		return node.key
	}
}

func ceiling(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}
	cmp := key.Compare(node.key)
	if cmp == 0 {
		return node
	} else if cmp > 0 {
		return ceiling(node.right, key)
	} else {
		t := ceiling(node.left, key)
		if t != nil {
			return t
		} else {
			return node
		}
	}
}

func (tree *RBTree) Rank(key Key) int {
	return rank(tree.root, key)
}

func rank(node *Node, key Key) int {
	if node == nil {
		return 0
	}
	cmp := key.Compare(node.key)
	if cmp == 0 {
		return size(node.left)
	} else if cmp < 0 {
		return rank(node.left, key)
	} else {
		return size(node.left) + 1 + rank(node.right, key)
	}
}

func (tree *RBTree) Select(k int) Key {
	if k < 0 || k >= tree.Size() {
		return nil
	}
	node := selectNode(tree.root, k)
	if node != nil {
		return node.key
	} else {
		return nil
	}
}

func selectNode(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	t := size(node.left)
	if t > k {
		return selectNode(node.left, k)
	} else if t < k {
		return selectNode(node.right, k-t-1)
	} else {
		return node
	}
}

func (tree *RBTree) DeleteMin() {
	if tree.IsEmpty() {
		return
	}
	// if both children of root are black, set root to red
	if !isRed(tree.root.left) && !isRed(tree.root.right) {
		tree.root.colour = RED
	}
	tree.root = deleteMin(tree.root)
	if !tree.IsEmpty() {
		tree.root.colour = BLACK
	}
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	node.left = deleteMin(node.left)
	return balance(node)
}

func balance(node *Node) *Node {
	if isRed(node.right) {
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) {
		flipColours(node)
	}
	node.N = size(node.left) + 1 + size(node.right)
	return node
}

func moveRedLeft(node *Node) *Node {
	// assuming that node is red and both node.left and node.left.left are black
	// make node.left or one of its children red
	flipColours(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		flipColours(node)
	}
	return node
}

func (tree *RBTree) DeleteMax() {
	if tree.IsEmpty() {
		return
	}
	// if both children black, set root red
	if !isRed(tree.root.left) && !isRed(tree.root.right) {
		tree.root.colour = RED
	}
	tree.root = deleteMax(tree.root)
	if !tree.IsEmpty() {
		tree.root.colour = BLACK
	}
}

func deleteMax(node *Node) *Node {
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	node.right = deleteMax(node.right)
	return balance(node)
}

func moveRedRight(node *Node) *Node {
	// assuming node is red and both node.right and node.right.left are black
	// make node.right or one of its children red
	flipColours(node)
	if isRed(node.left.left) {
		node = rotateRight(node)
		flipColours(node)
	}
	return node
}

func (tree *RBTree) Height() int {
	if tree.IsEmpty() {
		return 0
	}
	return height(tree.root)
}

func height(node *Node) int {
	if node == nil {
		return -1
	} else {
		leftHeight := height(node.left)
		rightHeight := height(node.right)
		if leftHeight >= rightHeight {
			return 1 + leftHeight
		} else {
			return 1 + rightHeight
		}
	}
}

func (tree *RBTree) Keys() []Key {
	return tree.KeysInRange(tree.Min(), tree.Max())
}

func (tree *RBTree) KeysInRange(lo Key, hi Key) []Key {
	queue := make([]Key, 0, 0)
	if tree.IsEmpty() || lo.Compare(hi) > 0 {
		return queue
	}

	keys(tree.root, &queue, lo, hi)
	return queue
}

func keys(node *Node, queue *[]Key, lo Key, hi Key) {
	if node == nil {
		return
	}
	cmplo := lo.Compare(node.key)
	cmphi := hi.Compare(node.key)
	if cmplo < 0 {
		keys(node.left, queue, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		*queue = append(*queue, node.key)
	}
	if cmphi > 0 {
		keys(node.right, queue, lo, hi)
	}
}

func (tree *RBTree) KeysCh(quit <-chan struct{}) <-chan Key {
	return tree.KeysInRangeCh(quit, tree.Min(), tree.Max())
}

func (tree *RBTree) KeysInRangeCh(quit <-chan struct{}, lo Key, hi Key) <-chan Key {

	out := make(chan Key)

	go func() {
		keysCh(quit, out, tree.root, lo, hi)
		close(out)
	}()

	return out
}

func keysCh(quit <-chan struct{}, ch chan Key, node *Node, lo Key, hi Key) {
	if node == nil {
		return
	}
	cmplo := lo.Compare(node.key)
	cmphi := hi.Compare(node.key)
	if cmplo < 0 {
		keysCh(quit, ch, node.left, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		select {
		case <-quit:
			return
		case ch <- node.key:
			break
		}
	}
	if cmphi > 0 {
		keysCh(quit, ch, node.right, lo, hi)
	}
}
