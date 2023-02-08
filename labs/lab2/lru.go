package lru

import (
	"errors"
	"fmt"
)

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....
	if _, ok := (*lru).cache[key.(string)]; ok {
		return (*lru).cache[key.(string)], nil
	} else {
		return "", errors.New(fmt.Sprintf("%s does not exist in cache!\n", key.(string)))
	}
}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	key_str := key.(string)
	val_str := val.(string)
	if _, ok := (*lru).cache[key_str]; ok {
		(*lru).qDel(key_str)
		(*lru).queue = append((*lru).queue, key_str)
	} else {
		if (*lru).remaining != 0 {
			(*lru).remaining--
			(*lru).queue = append((*lru).queue, key_str)
			(*lru).cache[key_str] = val_str
		} else {
			delete((*lru).cache, (*lru).queue[0])
			(*lru).qDel((*lru).queue[0])
			(*lru).queue = append((*lru).queue, key_str)
			(*lru).cache[key_str] = val_str
		}
	}
	return nil
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
