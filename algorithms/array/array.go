/* Algorithms: Array Methods */

package array

import "math/rand/v2"

/* Apply a function to all entries of an Array. */
func Apply[T any](array []T, fn func(T)) {
	for i := range array {
		fn(array[i])
	}
}

/* Shuffle an Array using Fisher-Yate. */
func Shuffle[T any](array []T) {
	for i := 1; i < len(array); i++ {
		j := rand.IntN(i + 1)
		array[i], array[j] = array[j], array[i]
	}
}

/* Reverse an Array. */
func Reverse[T any](array []T) {
	L, R := 0, len(array)-1
	for L < R {
		array[L], array[R] = array[R], array[L]
		L++
		R--
	}
}
