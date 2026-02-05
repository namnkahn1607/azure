/* Abstract Data Type: Deque */

package stackqueue

import "iter"

type Deque[T any] struct {
	head, size int
	array      []T
}

/* Create a Circular-buffer Deque with power of 2 capacity. */
func NewDeque[T any](initCap int) *Deque[T] {
	initCap = max(initCap, minimalCap)
	capacity := alignPowerOf2(uint(initCap))

	return &Deque[T]{
		head:  0,
		size:  0,
		array: make([]T, capacity),
	}
}

/* Create a Circular-buffer Deque with initial items. */
func NewDequeWith[T any](items ...T) *Deque[T] {
	capacity := alignPowerOf2(uint(len(items)))
	q := &Deque[T]{
		head:  0,
		size:  0,
		array: make([]T, capacity),
	}

	for i := range items {
		q.PushBack(items[i])
	}

	return q
}

/* Append an item to the front-side of Deque. */
func (dq *Deque[T]) PushFront(x T) {
	if dq.size == len(dq.array) {
		dq.grow()
	}

	dq.head = (dq.head - 1) & (len(dq.array) - 1)
	dq.array[dq.head] = x
	dq.size++
}

/* Append an item to the back-side of Deque. */
func (dq *Deque[T]) PushBack(x T) {
	if dq.size == len(dq.array) {
		dq.grow()
	}

	tail := (dq.head + dq.size) & (len(dq.array) - 1)
	dq.array[tail] = x
	dq.size++
}

/* Remove & return Deque's front-side item. */
func (dq *Deque[T]) PopFront() (T, bool) {
	var zero T

	if dq.size == 0 {
		return zero, false
	}

	capacity := len(dq.array)
	val := dq.array[dq.head]
	dq.array[dq.head] = zero
	dq.head = (dq.head - 1) & (capacity - 1)
	dq.size--

	if capacity > minimalCap && dq.size <= capacity/4 {
		dq.shrink()
	}

	return val, true
}

/* Remove & return Deque's back-side item. */
func (dq *Deque[T]) PopBack() (T, bool) {
	var zero T

	if dq.size == 0 {
		return zero, false
	}

	capacity := len(dq.array)
	last := (dq.head + dq.size - 1) & (capacity - 1)
	val := dq.array[last]
	dq.array[last] = zero
	dq.size--

	if capacity > minimalCap && dq.size <= capacity/4 {
		dq.shrink()
	}

	return val, true
}

/* Peek into Deque's front-side item. */
func (dq *Deque[T]) Front() (T, bool) {
	if dq.size == 0 {
		var zero T
		return zero, false
	}

	return dq.array[dq.head], true
}

/* Peek into Deque's back-side item. */
func (dq *Deque[T]) Back() (T, bool) {
	if dq.size == 0 {
		var zero T
		return zero, false
	}

	capacity := len(dq.array)
	last := (dq.head + dq.size - 1) & (capacity - 1)
	return dq.array[last], true
}

/* Clear all items of Deque. */
func (dq *Deque[T]) Clear() {
	capacity := len(dq.array)
	var zero T

	for i := range dq.size {
		idx := (dq.head + i) & (capacity - 1)
		dq.array[idx] = zero
	}

	dq.head = 0
	dq.size = 0
}

/* Calculate Deque's current size. */
func (dq *Deque[T]) Len() int { return dq.size }

/* Check is Deque is currently empty or not. */
func (dq *Deque[T]) IsEmpty() bool { return dq.size == 0 }

/* Iterator through all Deque's items. */
func (dq *Deque[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		N := dq.size
		capacity := len(dq.array)

		for i := range N {
			idx := (dq.head + i) & (capacity - 1)
			if !yield(dq.array[idx]) {
				return
			}
		}
	}
}

func (dq *Deque[T]) grow() {
	newCap := 2 * len(dq.array)
	dq.resize(newCap)
}

func (dq *Deque[T]) shrink() {
	newCap := len(dq.array) / 2
	dq.resize(newCap)
}

func (dq *Deque[T]) resize(newCap int) {
	newArray := make([]T, newCap)
	oldCap := len(dq.array)

	if dq.head+dq.size > oldCap {
		firstLen := oldCap - dq.head
		copy(newArray, dq.array[dq.head:])
		copy(newArray[firstLen:], dq.array[:dq.size-firstLen])
	} else {
		copy(newArray, dq.array[dq.head:dq.head+dq.size])
	}

	dq.array = newArray
	dq.head = 0
}
