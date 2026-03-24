/* Algorithm: Merge Sort */

package array

/*
Stable Top Down Mergesort.
- Time: O(N.logN) & Space: O(N).
*/
func MergeSort[T any](arr []T, compareFn func(a, b T) int) {
	N := len(arr)
	if N == 1 {
		return
	}

	aux := make([]T, N)

	merge := func(lo, mid, hi int) {
		copy(aux[lo:hi+1], arr[lo:hi+1])
		
		i, j := lo, mid+1
		for k := lo; k <= hi; k++ {
			if i > mid {
				arr[k] = aux[j]
				j++
			} else if j > hi {
				arr[k] = aux[i]
				i++
			} else if compareFn(aux[i], aux[j]) <= 0 {
				arr[k] = aux[i]
				i++
			} else {
				arr[k] = aux[j]
				j++
			}
		}
	}

	var sort func(int, int)
	sort = func(lo, hi int) {
		if hi-lo <= optimalThreshold {
			InsertionSort(arr[lo:hi+1], compareFn)
			return
		}

		mid := lo + (hi-lo)/2
		sort(lo, mid)
		sort(mid+1, hi)

		if compareFn(arr[mid], arr[mid+1]) <= 0 {
			return
		}

		merge(lo, mid, hi)
	}

	sort(0, N-1)
}
