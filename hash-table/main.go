// HashTable (ハッシュテーブル)
package main

import (
	"fmt"
	"sync"
)

type HashTable struct {
	table map[int]interface{}
	lock  sync.RWMutex
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (h *HashTable) hashKey(key string) (hash int) {
	for i := 0; i < len(key); i++ {
		// Horner's 法
		hash = 31*hash + int(key[i])
	}
	return
}

func (h *HashTable) Put(key string, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.table == nil {
		h.table = make(map[int]interface{})
	}
	i := h.hashKey(key)
	h.table[i] = value
}

func (h *HashTable) Remove(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	i := h.hashKey(key)
	delete(h.table, i)
}

func (h *HashTable) Get(key string) (interface{}, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	i := h.hashKey(key)
	ret, ok := h.table[i]
	return ret, ok
}

func (h *HashTable) Size() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return len(h.table)
}

func main() {
	ht := NewHashTable()
	ht.Put("car", "TOYOTA")
	ht.Put("pc", "Mac")
	ht.Put("sns", "Twitter")
	fmt.Println(ht.Size())
	ht.Remove("sns")
	ht.Remove("car")
	fmt.Println(ht.Size())
}
