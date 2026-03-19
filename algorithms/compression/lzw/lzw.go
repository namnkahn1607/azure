/* Algorithm: LZW Compression */

package lzw

import (
	binaryio "azure/data_structures/binary_io"
	"bufio"
	"io"
)

const (
	endMark = uint64(R)
)

/* LZW: Compress character stream into bitstream. */
func Encode(in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	writer := binaryio.NewBinaryWriter(out)
	dict := NewLZWDictionary()

	// Read the first character to determine the starting Trie.
	prevChar, readErr := reader.ReadByte()
	if readErr != nil {
		return readErr
	}

	currCode := uint16(prevChar)
	currNode := dict.roots[prevChar]

	for {
		// Read each character from the input stream.
		currChar, readErr := reader.ReadByte()
		if readErr != nil {
			if readErr == io.EOF {
				var writeErr error
				
				if writeErr = writer.WriteLSBOf(uint64(currCode), 16); writeErr != nil {
					return writeErr
				}

				if writeErr = writer.WriteLSBOf(endMark, 16); writeErr != nil {
					return writeErr
				}

				break
			}

			return readErr
		}

		var parent *LZWNode
		var match *LZWNode

		// Traverse the Trie to seek for matching Trie Node.
		node := currNode
		for node != nil {
			parent = node
			if currChar < node.Char {
				node = node.Left
			} else if currChar > node.Char {
				node = node.Right
			} else {
				match = node
				break
			}
		}

		if match != nil { // MATCHED
			currCode = match.Code
			currNode = match.Middle
		} else { // UNMATCHED
			// Write codeword of current longest prefix into bitstream.
			if writeErr := writer.WriteLSBOf(uint64(currCode), 16); writeErr != nil {
				return writeErr
			}

			// Set new codeword (if there's left any).
			if dict.nextCode < dict.maxCode {
				newNode := NewLZWNode(currChar, dict.nextCode)
				dict.nextCode++

				if parent == nil {
					dict.roots[prevChar] = newNode
				} else if currChar < parent.Char {
					parent.Left = newNode
				} else if currChar > parent.Char {
					parent.Right = newNode
				}
			}

			// State reset: new session with current character.
			prevChar = currChar
			currCode = uint16(currChar)
			currNode = dict.roots[currChar]
		}
	}

	return writer.Flush()
}
