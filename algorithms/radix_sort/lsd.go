/* Algorithm: LSD Radix Sort */

package radixsort

import "errors"

const R = 256

var ErrVariableLenString = errors.New("strings must be of same length")

/*
Least Significant Digit Radix Sort.
- Time: O(W.N) & Space: O(N + R).
*/
func LSD(strs []string, W int) error {
	for i := range strs {
		if len(strs[i]) != W {
			return ErrVariableLenString
		}
	}

	N := len(strs)
	aux := make([]string, N)
	count := make([]int, R+1)

	for d := W - 1; d >= 0; d-- {
		clear(count)

		for i := range N { // Count frequency
			count[strs[i][d]+1]++
		}

		for r := range R { // Compute cumulates
			count[r+1] += count[r]
		}

		for i := range N { // Move items
			aux[count[strs[i][d]]] = strs[i]
			count[strs[i][d]]++
		}

		copy(strs, aux) // Copy back
	}

	return nil
}
