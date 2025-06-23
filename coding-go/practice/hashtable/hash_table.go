package hashtable

import (
	"encoding/binary"
	"hash"

	"github.com/cespare/xxhash"
)

// HashTable is an implementation of map functionality which allows O(1) set and get operation with unique key constraint
type HashTable struct {
	capacity int
	buckets  []*node
	hashFunc hash.Hash
	// alpha is the load factor before rehashing
	alpha float64
	size  int
}

type node struct {
	next *node
	key  string
	val  string
}

func NewHashTable(capacity int) *HashTable {
	hashFunc := xxhash.New()
	return &HashTable{
		capacity: capacity,
		buckets:  make([]*node, capacity),
		alpha:    0.75,
		hashFunc: hashFunc,
	}
}

func (h *HashTable) hash(key string) int {
	h.hashFunc.Reset() // Reset the hash state before each use
	h.hashFunc.Write([]byte(key))
	hashBytes := h.hashFunc.Sum(nil)
	// Convert first 8 bytes to uint64 and then to int
	hashValue := binary.BigEndian.Uint64(hashBytes[:8])
	return int(hashValue % uint64(h.capacity))
}

func (h *HashTable) Set(key string, value string) {
	h.upsize()
	hashIndex := h.hash(key)
	current := h.buckets[hashIndex]
	for current != nil {
		if current.key == key {
			current.val = value
			return
		}
		current = current.next
	}
	newNode := &node{
		key: key,
		val: value,
		// add to the start of the linked list
		next: h.buckets[hashIndex],
	}
	h.buckets[hashIndex] = newNode
	h.size++
}

func (h *HashTable) upsize() {
	factor := float64(h.size) / float64(h.capacity)
	if factor > h.alpha {
		// reinitialized the hash table with double capacity
		newHashTable := NewHashTable(h.capacity * 2)
		for _, bucket := range h.buckets {
			for bucket != nil {
				newHashTable.Set(bucket.key, bucket.val)
				bucket = bucket.next
			}
		}
		h = newHashTable
	}
}

func (h *HashTable) Get(key string) string {
	hashIndex := h.hash(key)
	bucket := h.buckets[hashIndex]
	if bucket == nil {
		return ""
	}
	for bucket != nil {
		if bucket.key == key {
			return bucket.val
		}
		bucket = bucket.next
	}
	return ""
}
