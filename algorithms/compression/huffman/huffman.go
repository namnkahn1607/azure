/* Algorithm: Huffman Encoding */

package huffman

import (
	binaryio "azure/data_structures/binary_io"
	pq "azure/data_structures/priority_queue"
	"bufio"
	"io"
)

/* Huffman: Compress character stream into bitstream. */
func Encode(in io.ReadSeeker, out io.Writer) error {
	// Tabulate character frequencies.
	freqs, length, freqErr := buildFreqTable(in)
	if freqErr != nil {
		return freqErr
	}

	// Construct Binary Trie from frequency table.
	root := constructTrie(freqs)

	// Build Code Table from Binary Trie traversal.
	var codeTable CodeTable
	buildCodeTable(root, &codeTable, 0, 0)

	// Write compressed Binary Trie for Decode().
	writer := binaryio.NewBinaryWriter(out)
	if trieIOErr := writeTrie(root, writer); trieIOErr != nil {
		return trieIOErr
	}

	// Write total number of encoding characters.
	if totalErr := writer.WriteLSBOf(length, 64); totalErr != nil {
		return totalErr
	}

	// Read each from character stream & write bitcodes.
	reader := bufio.NewReader(in)
	for {
		ch, readErr := reader.ReadByte()
		if readErr != nil {
			if readErr == io.EOF {
				break
			}

			return readErr
		}

		code := codeTable[int(ch)]
		if codeWriErr := writer.WriteLSBOf(code.Bits, code.Len); codeWriErr != nil {
			return codeWriErr
		}
	}

	return writer.Flush()
}

/* Huffman: Expand bitstream back into character stream. */
func Decode(in io.Reader, out io.Writer) error {
	reader := binaryio.NewBinaryReader(in)

	// Reconstruct Binary Trie from it compressed form.
	root, trieIOErr := readTrie(reader)
	if trieIOErr != nil {
		return trieIOErr
	}

	// Retrieve total number of compressed chars.
	length, lenReadErr := reader.ReadMultiBits(64)
	if lenReadErr != nil {
		return lenReadErr
	}

	// Read each compressed bit and write chars.
	writer := bufio.NewWriter(out)
	for range length {
		curr := root
		for !curr.IsLeaf() {
			bit, bitIOErr := reader.ReadBit()
			if bitIOErr != nil {
				if bitIOErr == io.EOF {
					return io.ErrUnexpectedEOF
				}

				return bitIOErr
			}

			if bit {
				curr = curr.Right
			} else {
				curr = curr.Left
			}
		}

		if writeErr := writer.WriteByte(curr.Char); writeErr != nil {
			return writeErr
		}
	}

	return writer.Flush()
}

func buildFreqTable(in io.ReadSeeker) ([R]int, uint64, error) {
	reader := bufio.NewReader(in)
	var freqs [R]int
	var charCount uint64

	for {
		ch, readErr := reader.ReadByte()
		if readErr != nil {
			if readErr == io.EOF {
				break
			}

			return freqs, 0, readErr
		}

		freqs[ch]++
		charCount++
	}

	_, seekErr := in.Seek(0, io.SeekStart)
	if seekErr != nil {
		return freqs, 0, seekErr
	}

	return freqs, charCount, nil
}

func constructTrie(freqs [R]int) *Node {
	minpq := pq.NewPQ(func(a, b *Node) bool {
		return a.Freq < b.Freq
	})

	for c := range R {
		if freqs[c] > 0 {
			minpq.Enqueue(NewNode(byte(c), freqs[c], nil, nil))
		}
	}

	if minpq.Len() == 1 {
		minpq.Enqueue(NewNode('\u0000', 0, nil, nil))
	}

	for minpq.Len() > 1 {
		left := minpq.Dequeue()
		right := minpq.Dequeue()

		parent := NewNode('\u0000', left.Freq+right.Freq, left, right)
		minpq.Enqueue(parent)
	}

	return minpq.Dequeue()
}

func buildCodeTable(node *Node, table *CodeTable, currBits uint64, currLen uint8) {
	if node.IsLeaf() {
		table[int(node.Char)] = BitCode{Bits: currBits, Len: currLen}
		return
	}

	buildCodeTable(node.Left, table, currBits<<1, currLen+1)
	buildCodeTable(node.Right, table, (currBits<<1)|1, currLen+1)
}

func writeTrie(node *Node, bw *binaryio.BinaryWriter) error {
	if node.IsLeaf() {
		bitErr := bw.WriteBit(true)
		if bitErr != nil {
			return bitErr
		}

		byteErr := bw.WriteByte(node.Char)
		if byteErr != nil {
			return byteErr
		}

		return nil
	}

	bitErr := bw.WriteBit(false)
	if bitErr != nil {
		return bitErr
	}

	leftErr := writeTrie(node.Left, bw)
	if leftErr != nil {
		return leftErr
	}

	rightErr := writeTrie(node.Right, bw)
	if rightErr != nil {
		return rightErr
	}

	return nil
}

func readTrie(br *binaryio.BinaryReader) (*Node, error) {
	bit, bitErr := br.ReadBit()
	if bitErr != nil {
		return nil, bitErr
	}

	if bit {
		ch, byteErr := br.ReadByte()
		if byteErr != nil {
			return nil, byteErr
		}

		return NewNode(ch, -1, nil, nil), nil
	}

	left, leftErr := readTrie(br)
	if leftErr != nil {
		return nil, leftErr
	}

	right, rightErr := readTrie(br)
	if rightErr != nil {
		return nil, rightErr
	}

	return NewNode('\u0000', -1, left, right), nil
}
