/* API: Binary Trie Node */

package huffman

type Node struct {
	Char  byte
	Freq  int
	Left  *Node
	Right *Node
}

/* Create a new Binary Trie Node. */
func NewNode(char byte, freq int, left, right *Node) *Node {
	return &Node{
		Char:  char,
		Freq:  freq,
		Left:  left,
		Right: right,
	}
}

/* Check if a Trie Node is a leaf one. */
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}
