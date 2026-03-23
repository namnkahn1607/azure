/* Data Structure: Hash Map */

package hashtable

import "iter"

const (
	minimalCapacity = 128
	fnvOffsetBasis  = 0xcbf29ce484222325
	fnvPrime        = 0x100000001b3
)

type HashMap[K comparable, V any] struct {
	size   int
	keys   []K
	values []V
	states []bool
	hashFn func(K) uint64
}

/* Create new Hashmap with the provided hash function. */
func NewHashMap[K comparable, V any](fn func(K) uint64) *HashMap[K, V] {
	if fn == nil {
		var zero K
		switch any(zero).(type) {
		case string:
			fn = func(arg K) uint64 {
				var hash uint64 = fnvOffsetBasis
				str := any(arg).(string)
				for i := range len(str) {
					hash ^= uint64(str[i])
					hash *= fnvPrime
				}

				return hash
			}

		case int:
			fn = func(arg K) uint64 {
				hash := uint64(any(arg).(int))

				// Thomas Wang's 64-bit integer hash / SplitMix64
				hash ^= hash >> 30
				hash *= 0xbf58476d1ce4e5b9 // Magic Prime 1
				hash ^= hash >> 27
				hash *= 0x94d049bb133111eb // Magic Prime 2
				hash ^= hash >> 31

				return hash
			}

		default:
			panic("unsupported key type")
		}
	}

	return &HashMap[K, V]{
		size:   0,
		keys:   make([]K, minimalCapacity),
		values: make([]V, minimalCapacity),
		states: make([]bool, minimalCapacity),
		hashFn: fn,
	}
}

/* Get value associating with given key in the Hashmap. */
func (st *HashMap[K, V]) Get(key K) (V, bool) {
	currCap := len(st.keys)
	i := st.index(key)
	for st.states[i] {
		if st.keys[i] == key {
			return st.values[i], true
		}

		i = (i + 1) & (currCap - 1)
	}

	var zero V
	return zero, false
}

/* Associate given key with value in the Hashmap. */
func (st *HashMap[K, V]) Put(key K, val V) {
	if st.size >= len(st.keys)/2 {
		st.grow()
	}

	currCap := len(st.keys)
	i := st.index(key)
	for st.states[i] {
		if st.keys[i] == key {
			st.values[i] = val
			return
		}

		i = (i + 1) & (currCap - 1)
	}

	st.keys[i] = key
	st.values[i] = val
	st.states[i] = true
	st.size++
}

/* Check if a specific key exists in Hashmap. */
func (st *HashMap[K, V]) Contains(key K) bool {
	currCap := len(st.keys)
	i := st.index(key)
	for st.states[i] {
		if st.keys[i] == key {
			return true
		}

		i = (i + 1) & (currCap - 1)
	}

	return false
}

/* Remove a key-value entry from Hashmap. */
func (st *HashMap[K, V]) Remove(key K) {
	currCap := len(st.keys)
	i := st.index(key)
	for st.states[i] {
		if st.keys[i] == key {
			st.size--
			st.states[i] = false
			next := (i + 1) & (currCap - 1)
			for st.states[next] {
				nextKey, nextVal := st.keys[next], st.values[next]
				st.states[next] = false

				j := st.index(nextKey)
				for st.states[j] {
					j = (j + 1) & (currCap - 1)
				}

				st.keys[j] = nextKey
				st.values[j] = nextVal
				st.states[j] = true
				next = (next + 1) & (currCap - 1)
			}

			if currCap > minimalCapacity && st.size <= currCap/8 {
				st.shrink()
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
	var (
		zero_key K
		zero_val V
	)

	for i, occupied := range st.states {
		if occupied {
			st.keys[i] = zero_key
			st.values[i] = zero_val
			st.states[i] = false
		}
	}

	st.size = 0
}

/* Iterate through all keys of the Hashmap. */
func (st *HashMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for i, occupied := range st.states {
			if occupied && !yield(st.keys[i]) {
				return
			}
		}
	}
}

/* Iterate through all values of the Hashmap. */
func (st *HashMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i, occupied := range st.states {
			if occupied && !yield(st.values[i]) {
				return
			}
		}
	}
}

/* Iterate through all entries of the Hashmap. */
func (st *HashMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i, occupied := range st.states {
			if occupied && !yield(st.keys[i], st.values[i]) {
				return
			}
		}
	}
}

func (st *HashMap[K, V]) index(key K) int {
	hash := st.hashFn(key)

	// Smearing / Bit mixing
	hash ^= hash >> 32
	hash ^= hash >> 16
	hash ^= hash >> 8

	return int(hash) & (len(st.keys) - 1)
}

func (st *HashMap[K, V]) grow() {
	currCap := len(st.keys)
	st.resize(2 * currCap)
}

func (st *HashMap[K, V]) shrink() {
	currCap := len(st.keys)
	st.resize(currCap / 2)
}

func (st *HashMap[K, V]) resize(newCap int) {
	newKeys := make([]K, newCap)
	newValues := make([]V, newCap)
	newStates := make([]bool, newCap)

	for i := range len(st.keys) {
		if st.states[i] {
			key, val := st.keys[i], st.values[i]

			hash := st.hashFn(key)
			hash ^= hash >> 32
			hash ^= hash >> 16
			hash ^= hash >> 8
			start := int(hash) & (newCap - 1)

			j := start
			for newStates[j] {
				j = (j + 1) & (newCap - 1)
			}

			newKeys[j] = key
			newValues[j] = val
			newStates[j] = true
		}
	}

	st.keys = newKeys
	st.values = newValues
	st.states = newStates
}
