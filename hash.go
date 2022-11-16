package datastruct

// IMPORTANT: it's a rough implementation for demo purposes. Don't use it in production.

import (
	"fmt"
	"sync"
)

// NewHashTable creates a new instance of the hast table.
func NewHashTable[KT comparable, VT any](cap uint) *HashTable[KT, VT] {
	if cap < 1 {
		cap = 1
	}
	return &HashTable[KT, VT]{cap: cap, bucket: make([]htSubBucket[KT, VT], cap)}
}

type HashTable[KT comparable, VT any] struct {
	cap uint

	mux    sync.RWMutex
	bucket []htSubBucket[KT, VT]
}

// Set a key/value pair.
func (ht *HashTable[KT, VT]) Set(key KT, val VT) {
	i := ht.bucketIndex(key)
	ht.mux.Lock()
	defer ht.mux.Unlock()
	ht.bucket[i].set(key, val)
}

// Del key.
func (ht *HashTable[KT, VT]) Del(key KT) {
	i := ht.bucketIndex(key)
	ht.mux.Lock()
	defer ht.mux.Unlock()
	ht.bucket[i].del(key)
}

// Has returns a boolean flag that indicates if the key exists.
func (ht *HashTable[KT, VT]) Has(key KT) bool {
	i := ht.bucketIndex(key)
	ht.mux.RLock()
	defer ht.mux.RUnlock()
	for j := 0; j < ht.bucket[i].len; j++ {
		if ht.bucket[i].bucket[j].key == key {
			return true
		}
	}
	return false
}

// Get value from the table. It additionally returns a boolean flag that indicates if the key exists.
func (ht *HashTable[KT, VT]) Get(key KT) (val VT, exists bool) {
	i := ht.bucketIndex(key)
	ht.mux.RLock()
	defer ht.mux.RUnlock()
	for j := 0; j < ht.bucket[i].len; j++ {
		if ht.bucket[i].bucket[j].key == key {
			val = ht.bucket[i].bucket[j].val
			exists = true
			break
		}
	}
	return
}

// bucketIndex returns the index of the sub-bucket for the particular key.
func (ht *HashTable[KT, VT]) bucketIndex(key KT) uint {
	return hashAny(key) % ht.cap
}

// hashAny hashes an arbitrary value to integer.
// it's perhaps the worst hash function in the world, but it's ok for demo purposes.
func hashAny(v any) (hash uint) {
	for _, b := range anyToBytes(v) {
		hash += uint(b)
	}
	return
}

// anyToBytes converts an arbitrary value to the sequence of bytes.
// it's not the best implementation ever, but it's ok for demo purposes.
func anyToBytes(v any) []byte {
	return []byte(fmt.Sprintf("%+v", v))
}

// htItem is a key/value pair stored in the hash table.
type htItem[KT comparable, VT any] struct {
	key KT
	val VT
}

// htSubBucket is a collection of the values in which keys collide.
type htSubBucket[KT comparable, VT any] struct {
	len    int
	bucket []htItem[KT, VT]
}

// set a key/value pair.
func (sb *htSubBucket[KT, VT]) set(key KT, val VT) {
	// if first item in the sub-bucket
	if sb.bucket == nil {
		sb.bucket = []htItem[KT, VT]{
			{
				key: key,
				val: val,
			},
		}
		sb.len = 1
		return
	}
	// if already exists in the sub-bucket
	for j := 0; j < sb.len; j++ {
		if sb.bucket[j].key == key {
			sb.bucket[j].val = val
			return
		}
	}
	// add to the sub-bucket if it's length is less than the length of the underlying slice
	if sb.len < len(sb.bucket) {
		sb.bucket[sb.len].key = key
		sb.bucket[sb.len].val = val
		sb.len++
		return
	}
	// add new item to the sub-bucket
	sb.bucket = append(sb.bucket, htItem[KT, VT]{
		key: key,
		val: val,
	})
	sb.len++
}

// del key.
func (sb *htSubBucket[KT, VT]) del(key KT) {
	for i := 0; i < sb.len; i++ {
		if sb.bucket[i].key == key {
			// if last item then decrement length
			if i == sb.len-1 {
				sb.len--
				return
			}
			// otherwise move last item to the current position and decrement length
			sb.bucket[i] = sb.bucket[sb.len-1]
			sb.len--
			return
		}
	}
}
