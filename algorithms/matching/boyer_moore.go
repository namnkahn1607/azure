/* Algorithm: Boyer-Moore */

package matching

const R = 256

/*
Boyer-Moore pattern matching.
- Time: O(N/M) average, O(N.M) worst & Space: O(R).
*/
func BoyerMoore(txt, pat string) int {
	N, M := len(txt), len(pat)

	// Calculate skipping heuristics.
	right := make([]int, R)
	for c := range R {
		right[c] = -1
	}

	for i, ch := range pat {
		right[ch] = i
	}

	skip := 0

	// Slide the pattern window from LEFT to RIGHT.
	for i := 0; i <= N-M; i += skip {
		skip = 0

		// Match scanning from RIGHT to LEFT
		// to find the rightmost mismatch, hence skipping more.
		for j := M - 1; j >= 0; j-- {
			if pat[j] != txt[i+j] {
				// Mismatch char not in pattern -> Jump over.
				// Otherwise unless 'backing up', align with rightmost occurence.
				skip = max(1, j-right[txt[i+j]])
				break
			}
		}

		// Not to jump -> Match found.
		if skip == 0 {
			return i
		}
	}

	return -1
}
