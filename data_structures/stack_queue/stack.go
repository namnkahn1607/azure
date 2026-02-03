/* Abstract Data Type: Stack */

package stackqueue

import (
	"iter"
	"math/bits"
)

type Stack[T any] struct {
	array []T
}

/* Create a Stack with power of 2 capacity. */
func NewStack[T any](initCap int) *Stack[T] {
	initCap = max(initCap, minimalCap)
	capacity := alignPowerOf2(uint(initCap))

	return &Stack[T]{
		array: make([]T, 0, capacity),
	}
}

/* Create a Stack with initial items. */
func NewStackWith[T any](items ...T) *Stack[T] {
	capacity := alignPowerOf2(uint(len(items)))
	array := make([]T, 0, capacity)
	array = append(array, items...)

	return &Stack[T]{
		array: array,
	}
}

/* Append an item onto Stack. */
func (s *Stack[T]) Push(x T) {
	s.array = append(s.array, x)
}

/* Remove & return the top item of Stack. */
func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if len(s.array) == 0 {
		return zero, false
	}

	N := len(s.array)
	val := s.array[N-1]
	s.array[N-1] = zero
	s.array = s.array[:N-1]
	return val, true
}

/* Peek into the top item of Stack. */
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.array) == 0 {
		var zero T
		return zero, false
	}

	return s.array[len(s.array)-1], true
}

/* Calculate Stack's current size. */
func (s *Stack[T]) Len() int { return len(s.array) }

/* Check is Stack is currently empty or not. */
func (s *Stack[T]) IsEmpty() bool { return len(s.array) == 0 }

/* Iterator through all Stack's items. */
func (s *Stack[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		N := len(s.array)
		for i := range N {
			if !yield(s.array[i]) {
				return
			}
		}
	}
}

func alignPowerOf2(num uint) uint {
	const minPowerOf2 = 62

	if bits.OnesCount(num) != 1 {
		counts := bits.Len(num)
		counts = min(counts, minPowerOf2)
		num = 1 << counts
	}

	return num
}
