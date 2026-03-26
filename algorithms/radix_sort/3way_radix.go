/* Algorithm: 3-way Radix Quick Sort */

package radixsort

/*
3-way Radix Unstable Quick Sort.
- Time: O(N.logN) & Space: O(W + logN).
*/
func ThreeWay(strs []string) {
	quick(strs, 0, len(strs)-1, 0)
}

func quick(strs []string, lo, hi, d int) {
	if hi <= lo+cutoff {
		insertion(strs, lo, hi, d)
		return
	}

	left, right := lo, hi
	pivot := at(strs[lo], d)

	i := lo + 1
	for i <= right {
		c := at(strs[i], d)

		if c < pivot {
			strs[left], strs[i] = strs[i], strs[left]
			left++
			i++
		} else if c > pivot {
			strs[right], strs[i] = strs[i], strs[right]
			right--
		} else {
			i++
		}
	}

	quick(strs, lo, left-1, d)
	quick(strs, left, right, d+1)
	quick(strs, right+1, hi, d)
}
