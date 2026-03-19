/* API: LZW Dictionary */

package lzw

const R = 256

type LZWNode struct {
	Char                byte
	Code                uint16
	Left, Middle, Right *LZWNode
}

/* Initialize a new LZW Trie Node. */
func NewLZWNode(char byte, code uint16) *LZWNode {
	return &LZWNode{
		Char:   char,
		Code:   code,
		Left:   nil,
		Middle: nil,
		Right:  nil,
	}
}

type LZWDict struct {
	roots    [R]*LZWNode
	nextCode uint16
	maxCode  uint16
}

/* Initialize a new LZW Dictionary. */
func NewLZWDictionary() *LZWDict {
	return &LZWDict{
		roots:    [R]*LZWNode{},
		nextCode: R + 1,
		maxCode:  (1 << 16) - 1,
	}
}
