/* Data Structure: Deterministic Matcher */

package automata

import (
	"bufio"
	"errors"
	"io"
)

var ErrNonASCIIChar = errors.New("non-ASCII character encountered")

type matcher struct {
	dfa    *DFA
	reader *bufio.Reader
	offset int64
	state  int
	err    error
}

/* Next pattern match position in the byte stream. */
func (mat *matcher) Next() int64 {
	if mat.err != nil {
		return -1
	}

	accept := mat.dfa.length

	for {
		char, readErr := mat.reader.ReadByte()
		if readErr != nil {
			mat.err = readErr
			return -1
		}

		if int(char) >= R {
			mat.err = ErrNonASCIIChar
			return -1
		}

		// Proceed to next state of machine.
		mat.state = mat.dfa.table[char][mat.state]
		mat.offset++ // Increment 'j'

		// End of pattern reached.
		if mat.state == accept {
			// Avoid overlapping matches by reseting state.
			mat.state = 0

			matchPos := mat.offset - int64(accept)
			return matchPos
		}
	}
}

/* Check the sticky error on the matcher. */
func (mat *matcher) Error() error {
	if mat.err == io.EOF { // Ignore EOF
		return nil
	}

	return mat.err
}
