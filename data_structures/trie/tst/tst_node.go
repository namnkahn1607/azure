/* API: TST Node */

package trie

type TSTNode[T any] struct {
	Char                byte
	Left, Middle, Right *TSTNode[T]
	Value               T
	IsEnd               bool
}

/* Create a new Ternary Search Trie Node. */
func NewTSTNode[T any](char byte) *TSTNode[T] {
	var zero T
	return &TSTNode[T]{
		Char:   char,
		Left:   nil,
		Middle: nil,
		Right:  nil,
		Value:  zero,
		IsEnd:  false,
	}
}
