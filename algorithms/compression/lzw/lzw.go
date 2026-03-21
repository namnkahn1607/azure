/* Algorithm: LZW Compression */

package lzw

import (
	binaryio "azure/data_structures/binary_io"
	"bufio"
	"errors"
	"io"
	"strings"
)

const (
	codeBitLen = 16
	endMark    = uint16(R)
)

var (
	ErrCodeToWordInit = errors.New("decoder: error initializing Code-to-Word dictionary")
	ErrUnexpectedCode = errors.New("decoder: encounter unexpected code")
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

				writeErr = writer.WriteLSBOf(uint64(currCode), codeBitLen)
				if writeErr != nil {
					return writeErr
				}

				writeErr = writer.WriteLSBOf(uint64(endMark), codeBitLen)
				if writeErr != nil {
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
			prevChar = currChar

		} else { // UNMATCHED
			// Write codeword of current longest prefix into bitstream.
			writeErr := writer.WriteLSBOf(uint64(currCode), codeBitLen)
			if writeErr != nil {
				return writeErr
			}

			// Set new codeword (if there's left any).
			if dict.nextCode < MaxCode {
				newNode := NewLZWNode(currChar, dict.nextCode)
				dict.nextCode++

				// parent == nil means TST of prevChar is empty.
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

/* LZW: Expand bitstream back into character stream. */
func Decode(in io.Reader, out io.Writer) error {
	reader := binaryio.NewBinaryReader(in)
	writer := bufio.NewWriter(out)

	// Code-to-Word dictionary initialization.
	dict := make([]string, MaxCode)
	for i := range R {
		dict[i] = string(rune(i))
	}

	// Read & Write the first word of the stream.
	prevCode, readErr := reader.ReadMultiBits(codeBitLen)
	if readErr != nil {
		return readErr
	}

	prevWord := dict[prevCode]
	if _, writeErr := writer.WriteString(prevWord); writeErr != nil {
		return writeErr
	}
	
	nextCode := R + 1

	for {
		// Read each 16-bit code from the input bitstream.
		code, readErr := reader.ReadMultiBits(codeBitLen)
		if readErr != nil {
			// Encounter unexpected natural EOF.
			if readErr == io.EOF {
				return io.ErrUnexpectedEOF
			}

			return readErr
		}

		currCode := uint16(code)

		// Encounter bitstream's manual EOF signature.
		if currCode == endMark {
			break
		}

		// Retrieve Word from dictionary.
		currWord := dict[currCode]
		if currWord == "" {
			// LOGICAL EDGE CASE: instant-used code in compression.
			if currCode == nextCode {
				var builder strings.Builder
				builder.Grow(len(prevWord) + 1)
				builder.WriteString(prevWord)
				builder.WriteByte(prevWord[0])
				currWord = builder.String()
			} else {
				return ErrUnexpectedCode
			}
		}

		// Write retrieved word from dictionary to output stream.
		_, writeErr := writer.WriteString(currWord)
		if writeErr != nil {
			return writeErr
		}

		// Insert new Code-to-Word into dictionary.
		var builder strings.Builder
		builder.Grow(len(prevWord) + 1)
		builder.WriteString(prevWord)
		builder.WriteByte(currWord[0])

		if nextCode < MaxCode {
			dict[nextCode] = builder.String()
			nextCode++
		}

		prevWord = currWord
	}

	return writer.Flush()
}
