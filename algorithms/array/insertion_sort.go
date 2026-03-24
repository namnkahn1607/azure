/* Algorithm: Insertion Sort */

package array

const optimalThreshold = 12

/*
Stable Insertion Sort.
- Time: O(N) - O(N^2) & Space: O(1).
*/
func InsertionSort[T any](arr []T, compareFn func(a, b T) int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		j := i
		for j > 0 && compareFn(key, arr[j-1]) < 0 {
			arr[j] = arr[j-1]
			j--
		}

		arr[j] = key
	}
}
