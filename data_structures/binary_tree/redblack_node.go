/* Data Structure: Red-black Tree Node */

package binarytree

type color bool

const (
	BLACK color = false
	RED   color = true
)

type TreeNode[T any] struct {
	val   T
	left  *TreeNode[T]
	right *TreeNode[T]
	color color
	size  int
}

/* Create a new Red-black Tree Node from specified value. */
func NewTreeNode[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{
		val:   val,
		left:  nil,
		right: nil,
		color: BLACK,
		size:  1,
	}
}

/* Check if a node has a Red link to it. */
func (node *TreeNode[T]) IsRed() bool {
	if node == nil {
		return false
	}

	return node.color == RED
}

/* Flip the link's color to the specified node. */
func (node *TreeNode[T]) Flip() {
	node.color = !node.color
}

/* Size of the Subtree with specified node as its root. */
func (node *TreeNode[T]) Size() int {
	if node == nil {
		return 0
	}
	
	return node.size
}
