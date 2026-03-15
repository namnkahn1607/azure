/* Data Structure: Non-deterministic Finite Automata */

package automata

import (
	"azure/data_structures/graph"
	stackqueue "azure/data_structures/stack_queue"
	"slices"
)

type NFA struct {
	Regex string
	G     *graph.Digraph // Regex graph machine
	M     int            // Number of states
}

// Create a Regex machine from given string expression.
func NewRegex(exp string) *NFA {
	N := len(exp)
	G := graph.NewDigraph(N + 1)
	ops := stackqueue.NewStack[int](N)

	for i := range N {
		lp := i

		// Remember left parentheses '(' or OR operator '|'
		switch exp[i] {
		case '(', '|':
			ops.Push(i)

		// Closing right parentheses ')'
		case ')':
			or, ok1 := ops.Pop()
			if !ok1 {
				panic("malformed regular expression")
			}

			// Process OR operator '|'
			switch exp[or] {
			case '|':
				var ok2 bool
				lp, ok2 = ops.Pop()
				if !ok2 {
					panic("malformed regular expression")
				}

				G.AddEdge(*graph.NewEdge(lp, or+1, 0))
				G.AddEdge(*graph.NewEdge(or, i, 0))
			case '(':
				lp = or
			}
		}

		// 1 character lookahead
		if i < N-1 {
			switch exp[i+1] {
			case '*': // Closure operator '*'
				G.AddEdge(*graph.NewEdge(lp, i+1, 0))
				G.AddEdge(*graph.NewEdge(i+1, lp, 0))

			case '+': // At least 1 operator '+'
				G.AddEdge(*graph.NewEdge(i+1, lp, 0))
			}
		}

		// Next epsilon transition for all meta-characters
		if exp[i] == '(' || exp[i] == '*' || exp[i] == '+' || exp[i] == ')' {
			G.AddEdge(*graph.NewEdge(i, i+1, 0))
		}

	}

	return &NFA{
		Regex: exp,
		G:     G,
		M:     N,
	}
}

// Check if the given string match the Regular Expression.
func (nfa *NFA) Recognizes(text string) bool {
	G := nfa.G
	pc := make([]int, 0, nfa.M)

	// States initially reachable from 0
	for v := range G.Reachable(0) {
		pc = append(pc, v)
	}

	for i := range text {
		// Set of matched states after scanning text[i].
		states := make([]int, 0, nfa.M)
		for v := range pc {
			if v == nfa.M {
				continue
			}

			// Encounter a match or a variable operator.
			if nfa.Regex[v] == text[i] || nfa.Regex[v] == '.' {
				states = append(states, v+1)
			}
		}

		// Follow epsilon transitions of all matched states.
		pc = pc[:0]
		for _, state := range states {
			for v := range G.Reachable(state) {
				pc = append(pc, v)
			}
		}
	}

	return slices.Contains(pc, nfa.M)
}
