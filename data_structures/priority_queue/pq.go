/* Data Structure: Priority Queue */

package pq

import (
	"container/heap"
	"iter"
)

type PQ[T any] struct {
	array     []T
	compareFn func(a, b T) bool
}

/* Create a Priority Queue using specified comparator. */
func NewPQ[T any](cmp func(T, T) bool) *PQ[T] {
	if cmp == nil {
		panic("comparator unspecified")
	}

	pq := &PQ[T]{
		array:     make([]T, 0, minimalCap),
		compareFn: cmp,
	}

	heap.Init(pq)
	return pq
}

/* Append an Item into the Priority Queue. */
func (pq *PQ[T]) Enqueue(x T) {
	heap.Push(pq, x)
}

/* Delete & return the best Item in the Priority Queue. */
func (pq *PQ[T]) Dequeue() T {
	return heap.Pop(pq).(T)
}

/* Heap operation. Get called before the heapify. */
func (pq *PQ[T]) Push(x any) {
	pq.array = append(pq.array, x.(T))
}

/* Heap operation. Get called after the heapify. */
func (pq *PQ[T]) Pop() any {
	N := len(pq.array)
	key := pq.array[N-1]

	var zero T
	pq.array[N-1] = zero

	pq.array = pq.array[:N-1]
	return key
}

/* Get current size of the Priority Queue. */
func (pq *PQ[T]) Len() int { return len(pq.array) }

/* Check if the Priority Queue is empty or not. */
func (pq *PQ[T]) IsEmpty() bool { return len(pq.array) == 0 }

/* Heap operation. Check for better Item of 2. */
func (pq *PQ[T]) Less(a, b int) bool {
	return pq.compareFn(pq.array[a], pq.array[b])
}

/* Heap operation. Swap 2 Items' positions in internal array. */
func (pq *PQ[T]) Swap(a, b int) {
	pq.array[a], pq.array[b] = pq.array[b], pq.array[a]
}

/* All Items in the Priority Queue in heap order. */
func (pq *PQ[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		N := pq.Len()
		for i := range N {
			if !yield(pq.array[i]) {
				return
			}
		}
	}
}
