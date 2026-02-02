/* Abstract Data Type: Index Priority Queue */

package pq

const minimalCap = 64

type IndexPQ[Key any] struct {
	cap       int
	size      int
	pq        []int // heappos -> ID
	qp        []int // ID -> heappos
	keys      []Key // ID -> Key
	compareFn func(a, b Key) bool
}

/* Create a fixed-capacity Index Priority Queue. */
func NewIndexPQ[Key any](initCap int, cmp func(Key, Key) bool) *IndexPQ[Key] {
	if cmp == nil {
		panic("comparator unspecified")
	}

	finalCap := max(initCap, minimalCap)

	qp := make([]int, finalCap)
	for i := range finalCap {
		qp[i] = -1
	}

	return &IndexPQ[Key]{
		cap:       finalCap,
		size:      0,
		pq:        make([]int, finalCap+1),
		qp:        qp,
		keys:      make([]Key, finalCap),
		compareFn: cmp,
	}
}

/* Associate a key with given ID in IndexPQ. */
func (PQ *IndexPQ[Key]) Enqueue(id int, key Key) bool {
	if PQ.Contains(id) || PQ.size >= PQ.cap {
		return false
	}

	PQ.size++
	PQ.pq[PQ.size] = id
	PQ.qp[id] = PQ.size
	PQ.keys[id] = key
	PQ.swim(PQ.size)
	return true
}

/* Pop the current best entry from the IndexPQ. */
func (PQ *IndexPQ[Key]) Dequeue() (int, Key, bool) {
	var zero Key

	if PQ.size == 0 {
		return -1, zero, false
	}

	bestID := PQ.pq[1]
	bestKey := PQ.keys[bestID]
	PQ.swap(1, PQ.size)
	PQ.size--
	PQ.sink(1)

	PQ.qp[bestID] = -1
	PQ.keys[bestID] = zero

	return bestID, bestKey, true
}

/* Peek the current best entry of the IndexPQ. */
func (PQ *IndexPQ[Key]) Peek() (int, Key, bool) {
	if PQ.size == 0 {
		var zero Key
		return -1, zero, false
	}

	id := PQ.pq[1]

	return id, PQ.keys[id], true
}

/* Delete the key associated with given ID in IndexPQ. */
func (PQ *IndexPQ[Key]) Remove(id int) (Key, bool) {
	var zero Key

	if !PQ.Contains(id) {
		return zero, false
	}

	key := PQ.keys[id]
	heappos := PQ.qp[id]
	PQ.swap(heappos, PQ.size)
	PQ.size--

	if heappos <= PQ.size {
		PQ.swim(heappos)
		PQ.sink(heappos)
	}

	PQ.qp[id] = -1
	PQ.keys[id] = zero

	return key, true
}

/* Changing key associated with given ID in IndexPQ. */
func (PQ *IndexPQ[Key]) ChangeKey(id int, key Key) bool {
	if !PQ.Contains(id) {
		return false
	}

	PQ.keys[id] = key
	heappos := PQ.qp[id]
	PQ.swim(heappos)
	PQ.sink(heappos)
	return true
}

/* Check if an ID is occupied in the IndexPQ. */
func (PQ *IndexPQ[Key]) Contains(id int) bool {
	PQ.validate(id)
	return PQ.qp[id] != -1
}

/* Current size of the IndexPQ. */
func (PQ *IndexPQ[Key]) Len() int { return PQ.size }

/* Check if the IndexPQ is empty or not. */
func (PQ *IndexPQ[Key]) IsEmpty() bool { return PQ.size == 0 }

/* Return the associated key with given ID in IndexPQ. */
func (PQ *IndexPQ[Key]) KeyOf(id int) (Key, bool) {
	PQ.validate(id)

	if PQ.qp[id] == -1 {
		var zero Key
		return zero, false
	}

	return PQ.keys[id], true
}

func (PQ *IndexPQ[Key]) swim(k int) {
	for k > 1 {
		parent := k / 2
		if PQ.better(k, parent) {
			PQ.swap(k, parent)
			k = parent
		} else {
			break
		}
	}
}

func (PQ *IndexPQ[Key]) sink(k int) {
	for 2*k <= PQ.size {
		left, right := 2*k, 2*k+1
		child := left

		if right <= PQ.size && PQ.better(right, child) {
			child = right
		}

		if PQ.better(child, k) {
			PQ.swap(child, k)
			k = child
		} else {
			break
		}
	}
}

func (PQ *IndexPQ[Key]) better(a, b int) bool {
	idA := PQ.pq[a]
	idB := PQ.pq[b]
	return PQ.compareFn(PQ.keys[idA], PQ.keys[idB])
}

func (PQ *IndexPQ[Key]) swap(a, b int) {
	PQ.pq[a], PQ.pq[b] = PQ.pq[b], PQ.pq[a]
	PQ.qp[PQ.pq[a]] = a
	PQ.qp[PQ.pq[b]] = b
}

func (PQ *IndexPQ[Key]) validate(id int) {
	if id < 0 || id >= PQ.cap {
		panic("ID is out of bounds")
	}
}
