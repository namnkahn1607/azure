/* Algorithm: MSD Radix Sort */

package radixsort

const (
	eos    = -1
	cutoff = 16
)

/*
Most Significant Digit - Stable Radix Sort.
- Time: O(W.N) & Space: O(N + D.R).
*/
func MSD(strs []string) {
	N := len(strs)
	if N <= 1 {
		return
	}

	aux := make([]string, N)
	sort(strs, 0, N-1, 0, aux)
}

func sort(strs []string, lo, hi, d int, aux []string) {
	if hi <= lo+cutoff {
		insertion(strs, lo, hi, d)
		return
	}

	/* count[0] must be 0 to preserve cumulation. */
	var count [R + 2]int

	for i := lo; i <= hi; i++ { // Count frequencies
		c := at(strs[i], d)
		count[c+2]++
	}

	for r := range R + 1 { // Compute cumulates
		count[r+1] += count[r]
	}

	for i := lo; i <= hi; i++ { // Move items
		c := at(strs[i], d)
		aux[lo+count[c+1]] = strs[i]
		count[c+1]++
	}

	copy(strs[lo:hi+1], aux[lo:hi+1]) // Copy back

	/* Sort R subarrays recursively. */
	for r := range R {
		newLo := lo + count[r]
		newHi := lo + count[r+1] - 1
		sort(strs, newLo, newHi, d+1, aux)
	}
}

func insertion(strs []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		key := strs[i]

		j := i
		for j > lo && less(key, strs[j-1], d) {
			strs[j] = strs[j-1]
			j--
		}

		strs[j] = key
	}
}

func at(s string, d int) int {
	/* Access range of KIC expanded to [-1, R-1]. */
	if d >= len(s) {
		return -1
	}

	return int(s[d])
}

func less(v, w string, d int) bool {
	for i := d; i < len(v) && i < len(w); i++ {
		if v[i] != w[i] {
			return v[i] < w[i]
		}
	}

	return len(v) < len(w)
}
