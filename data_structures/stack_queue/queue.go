/* Abstract Data Type: Queue */

package stackqueue

import "iter"

const minimalCap = 64

type Queue[T any] struct {
	head, size int
	array      []T
}

/* Create a Circular-buffer Queue with power of 2 capacity. */
func NewQueue[T any](initCap int) *Queue[T] {
	initCap = max(initCap, minimalCap)
	capacity := alignPowerOf2(uint(initCap))

	return &Queue[T]{
		head:  0,
		size:  0,
		array: make([]T, capacity),
	}
}

/* Create a Circular-buffer Queue with initial items. */
func NewQueueWith[T any](items ...T) *Queue[T] {
	capacity := alignPowerOf2(uint(len(items)))
	q := &Queue[T]{
		head: 0,
		size: 0,
		array: make([]T, capacity),
	}

	for i := range items {
		q.Enqueue(items[i])
	}

	return q
}

/* Append an item onto Queue. */ 
func (q *Queue[T]) Enqueue(x T) {
	if q.size == len(q.array) {
		q.grow()
	}

	capacity := len(q.array)
	tail := (q.head + q.size) & (capacity - 1)
	q.array[tail] = x
	q.size++
}

/* Remove & return Queue's frontier item. */
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T

	if q.size == 0 {
		return zero, false
	}

	capacity := len(q.array)
	val := q.array[q.head]
	q.array[q.head] = zero
	q.head = (q.head + 1) & (capacity - 1)
	q.size--

	if capacity > minimalCap && q.size <= capacity/4 {
		q.shrink()
	}

	return val, true
}

/* Access Queue's frontier item. */
func (q *Queue[T]) Front() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	return q.array[q.head], true
}

/* Clear all items of Queue. */
func (q *Queue[T]) Clear() {
	capacity := len(q.array)
	var zero T

	for i := range q.size {
		idx := (q.head + i) & (capacity - 1)
		q.array[idx] = zero
	}

	q.head = 0
	q.size = 0
}

/* Calculate Queue's current size. */
func (q *Queue[T]) Len() int { return q.size }

/* Check is Queue is currently empty or not. */
func (q *Queue[T]) IsEmpty() bool { return q.size == 0 }

/* Iterator through all Queue's items. */
func (q *Queue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		N := q.size
		capacity := len(q.array)
		for i := range N {
			idx := (q.head + i) & (capacity - 1)
			if !yield(q.array[idx]) {
				return
			}
		}
	}
}

func (q *Queue[T]) grow() {
	newCap := 2 * len(q.array)
	q.resize(newCap)
}

func (q *Queue[T]) shrink() {
	newCap := len(q.array) / 2
	q.resize(newCap)
}

func (q *Queue[T]) resize(newCap int) {
	newArray := make([]T, newCap)
	oldCap := len(q.array)

	if q.head+q.size > oldCap {
		firstLen := oldCap - q.head
		copy(newArray, q.array[q.head:])
		copy(newArray[firstLen:], q.array[:q.size-firstLen])
	} else {
		copy(newArray, q.array[q.head:q.head+q.size])
	}

	q.array = newArray
	q.head = 0
}
