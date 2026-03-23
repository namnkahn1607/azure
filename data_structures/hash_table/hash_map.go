/* Data Structure: Hash Map */

package hashtable

import (
	"hash/maphash"
	"iter"
)

const minimalCapacity = 128

type Entry[K comparable, V any] struct {
	key      K
	value    V
	occupied bool
}

type HashMap[K comparable, V any] struct {
	size    int
	entries []Entry[K, V]
	hashFn  func(K) uint64
}

/* Create new Hashmap with the provided hash function. */
func NewHashMap[K comparable, V any](fn func(K) uint64) *HashMap[K, V] {
	if fn == nil {
		seed := maphash.MakeSeed()
		fn = func(key K) uint64 {
			return maphash.Comparable(seed, key)
		}
	}

	return &HashMap[K, V]{
		size:    0,
		entries: make([]Entry[K, V], minimalCapacity),
		hashFn:  fn,
	}
}

/* Get value associating with given key in the Hashmap. */
func (st *HashMap[K, V]) Get(key K) (V, bool) {
	ent := st.entries
	currCap := len(ent)
	i := st.index(key, currCap)

	for ent[i].occupied {
		if ent[i].key == key {
			return ent[i].value, true
		}

		i = (i + 1) & (currCap - 1)
	}

	var zero_val V
	return zero_val, false
}

/* Associate given key with value in the Hashmap. */
func (st *HashMap[K, V]) Put(key K, val V) {
	if st.size >= len(st.entries)/2 {
		st.resize(2 * len(st.entries))
	}

	ent := st.entries
	currCap := len(ent)
	i := st.index(key, currCap)

	for ent[i].occupied {
		if ent[i].key == key {
			ent[i].value = val
			return
		}

		i = (i + 1) & (currCap - 1)
	}

	ent[i].key = key
	ent[i].value = val
	ent[i].occupied = true
	st.size++
}

/* Check if a specific key exists in Hashmap. */
func (st *HashMap[K, V]) Contains(key K) bool {
	ent := st.entries
	currCap := len(ent)
	i := st.index(key, currCap)

	for ent[i].occupied {
		if ent[i].key == key {
			return true
		}

		i = (i + 1) & (currCap - 1)
	}

	return false
}

/* Remove a specified key-value entry from Hashmap. */
func (st *HashMap[K, V]) Remove(key K) {
	var zero Entry[K, V]

	ent := st.entries
	currCap := len(ent)
	i := st.index(key, currCap)
	for ent[i].occupied {
		if ent[i].key == key {
			st.size--
			ent[i] = zero

			next := (i + 1) & (currCap - 1)
			for ent[next].occupied {
				nextEn := ent[next]
				ent[next] = zero

				j := st.index(nextEn.key, currCap)
				for ent[j].occupied {
					j = (j + 1) & (currCap - 1)
				}

				ent[j] = nextEn
				next = (next + 1) & (currCap - 1)
			}

			if currCap > minimalCapacity && st.size <= currCap/8 {
				st.resize(currCap / 2)
			}

			return
		}

		i = (i + 1) & (currCap - 1)
	}
}

/* Current size of the Hashmap. */
func (st *HashMap[K, V]) Len() int { return st.size }

/* Check whether the Hashmap is empty or not. */
func (st *HashMap[K, V]) IsEmpty() bool { return st.size == 0 }

/* Reset the Hashmap into its empty state. */
func (st *HashMap[K, V]) Clear() {
	var zero Entry[K, V]

	ent := st.entries
	for i := range ent {
		if ent[i].occupied {
			ent[i] = zero
		}
	}

	st.size = 0
}

/* Iterate through all keys of the Hashmap. */
func (st *HashMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		ent := st.entries
		for i := range ent {
			if ent[i].occupied && !yield(ent[i].key) {
				return
			}
		}
	}
}

/* Iterate through all values of the Hashmap. */
func (st *HashMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		ent := st.entries
		for i := range ent {
			if ent[i].occupied && !yield(ent[i].value) {
				return
			}
		}
	}
}

/* Iterate through all entries of the Hashmap. */
func (st *HashMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		ent := st.entries
		for i := range ent {
			if ent[i].occupied && !yield(ent[i].key, ent[i].value) {
				return
			}
		}
	}
}

func (st *HashMap[K, V]) index(key K, capacity int) int {
	hash := st.hashFn(key)

	// Smearing / Bit mixing
	hash ^= hash >> 32
	hash ^= hash >> 16
	hash ^= hash >> 8

	return int(hash) & (capacity - 1)
}

func (st *HashMap[K, V]) resize(newCap int) {
	oldEnt := st.entries
	newEnt := make([]Entry[K, V], newCap)

	for i := range oldEnt {
		if oldEnt[i].occupied {
			currEn := oldEnt[i]
			j := st.index(currEn.key, newCap)

			for newEnt[j].occupied {
				j = (j + 1) & (newCap - 1)
			}

			newEnt[j] = currEn
		}
	}

	st.entries = newEnt
}
