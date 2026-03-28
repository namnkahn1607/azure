/* Data Structure: Deterministic Finite Automata */

package automata

import (
	"bufio"
	"io"
)

const R = 256

type DFA struct {
	pattern string
	length  int
	table   [R][]int
}

/* Construct a new DFA based on given pattern string. */
func NewDFA(pat string) *DFA {
	dfa := &DFA{
		pattern: pat,
		length:  len(pat),
	}

	for i := range R {
		dfa.table[i] = make([]int, len(pat))
	}

	dfa.table[pat[0]][0] = 1
	X := 0 // Shadow state X

	for j := 1; j < len(pat); j++ {
		// Copy mismatch case.
		// Pretend if current state 'j' acts like X.
		for c := range R {
			dfa.table[c][j] = dfa.table[c][X]
		}

		char := pat[j]
		// Overwrite match case (by next state).
		dfa.table[char][j] = j + 1

		// Update X - like it chose current char.
		X = dfa.table[char][X]
	}

	return dfa
}

/* Attach the DFA to a byte stream to find matching pattern. */ 
func (dfa *DFA) Search(r io.Reader) *matcher {
	return &matcher{
		dfa:    dfa,
		reader: bufio.NewReader(r),
	}
}
