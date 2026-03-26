/* Data Structure: Red-black Tree */

package binarytree

import "iter"

type RedBlackTree[T any] struct {
	root      *TreeNode[T]
	compareFn func(a, b T) int
}

/* Create a new Red-black Tree with specified comparator. */
func NewRedBlackTree[T any](fn func(a, b T) int) *RedBlackTree[T] {
	if fn == nil {
		panic("comparator must be specified")
	}

	return &RedBlackTree[T]{
		root:      nil,
		compareFn: fn,
	}
}

/* Append a key to Red-black Tree. */
func (rbt *RedBlackTree[T]) Insert(x T) {
	if !rbt.Contains(x) {
		rbt.root = rbt.insert(rbt.root, x)
		if rbt.root != nil {
			rbt.root.color = BLACK
		}
	}
}

func (rbt *RedBlackTree[T]) insert(node *TreeNode[T], x T) *TreeNode[T] {
	if node == nil {
		return NewTreeNode(x)
	}

	cmp := rbt.compareFn(node.val, x)
	if cmp > 0 {
		node.left = rbt.insert(node.left, x)
	} else if cmp < 0 {
		node.right = rbt.insert(node.right, x)
	}

	node = rbt.balance(node)
	node.size = node.left.Size() + 1 + node.right.Size()
	return node
}

/* Remove a key from Red-black Tree (if there exists any). */
func (rbt *RedBlackTree[T]) Delete(x T) {
	if rbt.Contains(x) {
		if !rbt.root.left.IsRed() && !rbt.root.right.IsRed() {
			rbt.root.color = RED
		}

		rbt.root = rbt.delete(rbt.root, x)
		if rbt.root != nil {
			rbt.root.color = BLACK
		}
	}
}

func (rbt *RedBlackTree[T]) delete(node *TreeNode[T], x T) *TreeNode[T] {
	if node == nil {
		return nil
	}

	cmp := rbt.compareFn(node.val, x)
	if cmp > 0 {
		if !node.left.IsRed() && !node.left.left.IsRed() {
			node = rbt.moveRedLeft(node)
		}

		node.left = rbt.delete(node.left, x)

	} else {
		if node.left.IsRed() {
			node = rbt.rotateRight(node)
		}

		if rbt.compareFn(node.val, x) == 0 && node.right == nil {
			return nil
		}

		if !node.right.IsRed() && !node.right.left.IsRed() {
			node = rbt.moveRedRight(node)
		}

		if rbt.compareFn(node.val, x) == 0 {
			successor := node.right
			for successor.left != nil {
				successor = successor.left
			}

			node.val = successor.val
			node.right = rbt.deleteMin(node.right)
		} else {
			node.right = rbt.delete(node.right, x)
		}
	}

	node = rbt.balance(node)
	node.size = node.left.Size() + 1 + node.right.Size()
	return node
}

func (rbt *RedBlackTree[T]) deleteMin(h *TreeNode[T]) *TreeNode[T] {
	if h.left == nil {
		return nil
	}

	if !h.left.IsRed() && !h.left.left.IsRed() {
		h = rbt.moveRedLeft(h)
	}

	h.left = rbt.deleteMin(h.left)
	return rbt.balance(h)
}

func (rbt *RedBlackTree[T]) moveRedLeft(h *TreeNode[T]) *TreeNode[T] {
	rbt.flipColors(h)

	if h.right.left.IsRed() {
		h.right = rbt.rotateRight(h.right)
		h = rbt.rotateLeft(h)
		rbt.flipColors(h)
	}

	return h
}

func (rbt *RedBlackTree[T]) moveRedRight(h *TreeNode[T]) *TreeNode[T] {
	rbt.flipColors(h)

	if h.left.left.IsRed() {
		h = rbt.rotateRight(h)
		rbt.flipColors(h)
	}

	return h
}

func (rbt *RedBlackTree[T]) balance(h *TreeNode[T]) *TreeNode[T] {
	if h.right.IsRed() && !h.left.IsRed() {
		h = rbt.rotateLeft(h)
	}

	if h.left.IsRed() && h.left.left.IsRed() {
		h = rbt.rotateRight(h)
	}

	if h.left.IsRed() && h.right.IsRed() {
		rbt.flipColors(h)
	}

	return h
}

func (rbt *RedBlackTree[T]) rotateLeft(h *TreeNode[T]) *TreeNode[T] {
	x := h.right

	h.right = x.left
	x.left = h

	x.color = h.color
	h.color = RED

	h.size = h.left.Size() + 1 + h.right.Size()
	x.size = x.left.Size() + 1 + x.right.Size()

	return x
}

func (rbt *RedBlackTree[T]) rotateRight(h *TreeNode[T]) *TreeNode[T] {
	x := h.left

	h.left = x.right
	x.right = h

	x.color = h.color
	h.color = RED

	h.size = h.left.Size() + 1 + h.right.Size()
	x.size = x.left.Size() + 1 + x.right.Size()

	return x
}

func (rbt *RedBlackTree[T]) flipColors(h *TreeNode[T]) {
	h.Flip()
	h.left.Flip()
	h.right.Flip()
}

/* Check if BST contains a specified key */
func (rbt *RedBlackTree[T]) Contains(x T) bool {
	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)

		if cmp > 0 {
			curr = curr.left
		} else if cmp < 0 {
			curr = curr.right
		} else {
			return true
		}
	}

	return false
}

/* Clear all current keys in Red-black Tree. */
func (rbt *RedBlackTree[T]) Clear() {
	rbt.root = nil
}

/* Calculate current height of Red-black Tree. */
func (rbt *RedBlackTree[T]) Height() int {
	return rbt.height(rbt.root)
}

func (rbt *RedBlackTree[T]) height(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}

	return 1 + max(rbt.height(node.left), rbt.height(node.right))
}

/* Total number of nodes in Red-black Tree. */
func (rbt *RedBlackTree[T]) Size() int { return rbt.root.Size() }

/* Check if Red-black Tree is empty? */
func (rbt *RedBlackTree[T]) IsEmpty() bool { return rbt.root.Size() == 0 }

/* Maximum key of Red-black Tree. */
func (rbt *RedBlackTree[T]) Max() (T, bool) {
	if rbt.root == nil {
		var zero T
		return zero, false
	}

	curr := rbt.root
	for curr.right != nil {
		curr = curr.right
	}

	return curr.val, true
}

/* Minimum key of Red-black Tree. */
func (rbt *RedBlackTree[T]) Min() (T, bool) {
	if rbt.root == nil {
		var zero T
		return zero, false
	}

	curr := rbt.root
	for curr.left != nil {
		curr = curr.left
	}

	return curr.val, true
}

/* Largest key smaller than specified key. */
func (rbt *RedBlackTree[T]) Lower(x T) (T, bool) {
	var lower T
	found := false

	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)
		if cmp < 0 {
			lower = curr.val
			curr = curr.right
		} else {
			curr = curr.left
		}
	}

	return lower, found
}

/* Smallest key larger than specified key. */
func (rbt *RedBlackTree[T]) Higher(x T) (T, bool) {
	var higher T
	found := false

	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)
		if cmp > 0 {
			higher = curr.val
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return higher, found
}

/* Largest key smaller than/equal to specified key. */
func (rbt *RedBlackTree[T]) Floor(x T) (T, bool) {
	var floor T
	found := false

	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)
		if cmp <= 0 {
			floor = curr.val
			found = true
			curr = curr.right
		} else {
			curr = curr.left
		}
	}

	return floor, found
}

/* Smallest key larger than/equal to specified key. */
func (rbt *RedBlackTree[T]) Ceiling(x T) (T, bool) {
	var ceil T
	found := false

	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)
		if cmp >= 0 {
			ceil = curr.val
			found = true
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return ceil, found
}

/* Count number of keys smaller than specified key. */
func (rbt *RedBlackTree[T]) Rank(x T) int {
	rank := 0
	curr := rbt.root
	for curr != nil {
		cmp := rbt.compareFn(curr.val, x)
		if cmp > 0 {
			curr = curr.left
		} else if cmp < 0 {
			rank += (1 + curr.left.Size())
			curr = curr.right
		} else {
			return rank + curr.left.Size()
		}
	}

	return rank
}

/* Select k-th smallest key in BST (1-indexed). */
func (rbt *RedBlackTree[T]) Select(k int) T {
	if k <= 0 || k > rbt.Size() {
		panic("rank selected out of bounds")
	}

	curr := rbt.root
	for curr != nil {
		rank := 1 + curr.left.Size()
		if rank > k {
			curr = curr.left
		} else if rank < k {
			k -= rank
			curr = curr.right
		} else {
			return curr.val
		}
	}

	var zero T
	return zero
}

/* Count number of keys between 'low' and 'high' keys. */
func (rbt *RedBlackTree[T]) Count(lo, hi T) int {
	if rbt.compareFn(lo, hi) > 0 {
		panic("lowerbound is larger than upperbound")
	}

	count := rbt.Rank(hi) - rbt.Rank(lo)
	if rbt.Contains(hi) {
		count++
	}

	return count
}

/* Iterate through items between 'low' and 'high' keys. */
func (rbt *RedBlackTree[T]) Between(lo, hi T) iter.Seq[T] {
	if rbt.compareFn(lo, hi) > 0 {
		panic("lowerbound is larger than upperbound")
	}

	return func(yield func(T) bool) {
		var dfs func(*TreeNode[T]) bool
		dfs = func(node *TreeNode[T]) bool {
			if node == nil {
				return true
			}

			cmpLow := rbt.compareFn(lo, node.val)
			if cmpLow < 0 && !dfs(node.left) {
				return false
			}

			cmpHigh := rbt.compareFn(node.val, hi)
			if cmpLow <= 0 && cmpHigh <= 0 {
				if !yield(node.val) {
					return false
				}
			}

			if cmpHigh > 0 && !dfs(node.right) {
				return false
			}

			return true
		}

		dfs(rbt.root)
	}
}

/* Inorder Traversal on Red-black Tree. */
func (rbt *RedBlackTree[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		var dfs func(*TreeNode[T]) bool
		dfs = func(node *TreeNode[T]) bool {
			if node == nil {
				return true
			}

			if !dfs(node.left) {
				return false
			}

			if !yield(node.val) {
				return false
			}

			if !dfs(node.right) {
				return false
			}

			return true
		}

		dfs(rbt.root)
	}
}
