/* Data Structure: Ternary Search Trie */

package trie

import "iter"

const (
	minPrealloc = 1024
	wildcard    = byte('.')
)

type TST[T any] struct {
	root *TSTNode[T]
	size int
}

/* Initialize a new Ternary Search Trie. */
func NewTST[T any]() *TST[T] {
	return &TST[T]{
		root: nil,
		size: 0,
	}
}

/* Associate a string key with a value */
func (tst *TST[T]) Insert(key string, val T) {
	tst.root = tst.insert(tst.root, key, val, 0)
}

func (tst *TST[T]) insert(node *TSTNode[T], key string, val T, d int) *TSTNode[T] {
	ch := key[d]

	if node == nil {
		node = NewTSTNode[T](ch)
	}

	if ch < node.Char {
		node.Left = tst.insert(node.Left, key, val, d)
	} else if ch > node.Char {
		node.Right = tst.insert(node.Right, key, val, d)
	} else if d < len(key)-1 {
		node.Middle = tst.insert(node.Middle, key, val, d+1)
	} else {
		if !node.IsEnd {
			node.IsEnd = true
			tst.size++
		}

		node.Value = val
	}

	return node
}

/* Remove a string key from the Trie. */
func (tst *TST[T]) Delete(key string) {
	tst.root = tst.delete(tst.root, key, 0)
}

func (tst *TST[T]) delete(node *TSTNode[T], key string, d int) *TSTNode[T] {
	if node == nil {
		return nil
	}

	ch := key[d]
	if ch < node.Char {
		node.Left = tst.delete(node.Left, key, d)
	} else if ch > node.Char {
		node.Right = tst.delete(node.Right, key, d)
	} else if d < len(key)-1 {
		node.Middle = tst.delete(node.Middle, key, d+1)
	} else if node.IsEnd {
		node.IsEnd = false
		tst.size--
	}

	if node.IsEnd || node.Middle != nil {
		return node
	}

	if node.Left == nil {
		return node.Right
	}

	if node.Right == nil {
		return node.Left
	}

	successor := node.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	node.Char = successor.Char
	node.Value = successor.Value
	node.IsEnd = successor.IsEnd
	node.Middle = successor.Middle
	node.Right = tst.deleteMin(node.Right)

	return node
}

func (tst *TST[T]) deleteMin(node *TSTNode[T]) *TSTNode[T] {
	if node.Left == nil {
		return node.Right
	}

	node.Left = tst.deleteMin(node.Left)
	return node
}

/* Check if a string key exists within the Trie. */
func (tst *TST[T]) Contains(key string) bool {
	curr := tst.root

	d := 0
	for d < len(key) {
		if curr == nil {
			return false
		}

		ch := key[d]
		if ch < curr.Char {
			curr = curr.Left
		} else if ch > curr.Char {
			curr = curr.Right
		} else {
			if d == len(key)-1 {
				break
			}

			curr = curr.Middle
			d++
		}
	}

	return curr.IsEnd
}

/* Ternary Search Trie's current size. */
func (tst *TST[T]) Len() int { return tst.size }

/* Check if Ternary Search Trie is empty or not. */
func (tst *TST[T]) IsEmpty() bool { return tst.size == 0 }

/* All string keys (and their values) of the TST. */
func (tst *TST[T]) Keys() iter.Seq2[string, T] {
	return tst.KeysWithPrefix("")
}

/* All string keys (and their values) that starts with given prefix */
func (tst *TST[T]) KeysWithPrefix(prefix string) iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		path := make([]byte, 0, minPrealloc)

		var dfs func(*TSTNode[T]) bool
		dfs = func(node *TSTNode[T]) bool {
			if node == nil {
				return true
			}

			// Traverse Left
			if !dfs(node.Left) {
				return false
			}

			// Traverse Middle
			path = append(path, node.Char)

			if node.IsEnd {
				if !yield(string(path), node.Value) {
					return false
				}
			}

			if !dfs(node.Middle) {
				return false
			}

			path = path[:len(path)-1]

			// Traverse Right
			if !dfs(node.Right) {
				return false
			}

			return true
		}

		// Handle all keys retrieval
		if len(prefix) == 0 {
			dfs(tst.root)
			return
		}

		// Prefix traversal
		curr := tst.root
		d := 0
		for curr != nil && d < len(prefix) {
			ch := prefix[d]
			if ch < curr.Char {
				curr = curr.Left
			} else if ch > curr.Char {
				curr = curr.Right
			} else {
				if d == len(prefix) - 1 {
					break
				}

				curr = curr.Middle
				d++
			}
		}

		if curr == nil { // Invalid prefix
			return
		}

		// Handle prefix cases
		path = append(path, prefix...)
		
		if curr.IsEnd && !yield(string(path), curr.Value) {
			return
		}

		dfs(curr.Middle)
	}
}

/* All string keys (and their values) that match given wilcard string */
func (tst *TST[T]) KeysThatMatch(pattern string) iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		if len(pattern) == 0 {
			return
		}

		path := make([]byte, 0, minPrealloc)

		var dfs func(*TSTNode[T], int) bool
		dfs = func(node *TSTNode[T], d int) bool {
			if node == nil {
				return true
			}

			ch := pattern[d]
			
			// Traverse Left
			if ch == wildcard || ch < node.Char {
				if !dfs(node.Left, d) {
					return false
				}
			}

			// Traverse Middle
			if ch == wildcard || ch == node.Char {
				path = append(path, node.Char)

				if d == len(pattern) - 1 {
					if node.IsEnd && !yield(string(path), node.Value) {
						return false
					}
				} else {
					if !dfs(node.Middle, d + 1) {
						return false
					}
				}

				path = path[:len(path) - 1]
			}

			// Traverse Right
			if ch == wildcard || ch > node.Char {
				if !dfs(node.Right, d) {
					return false
				}
			}

			return true
		}

		dfs(tst.root, 0)
	}
}

/* Longest prefix and its value of a given string */
func (tst *TST[T]) LongestPrefixOf(query string) (string, T, bool) {
	var longestLen int
	var value T
	var found bool

	curr := tst.root
	d := 0
	for d < len(query) {
		if curr == nil {
			break
		}

		ch := query[d]
		if ch < curr.Char {
			curr = curr.Left
		} else if ch > curr.Char {
			curr = curr.Right
		} else {
			if curr.IsEnd {
				longestLen = d + 1
				value = curr.Value
				found = true
			}

			curr = curr.Middle
			d++
		}
	}

	return query[:longestLen], value, found
}
