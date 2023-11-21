/* *  @Author: Sk Shahriar Ahmed Raka   * Email: skshahriarahmedraka@gmail.com  * Telegram: https://t.me/shahriarraka  * Github: https://github.com/skshahriarahmedraka  * StackOverflow: https://stackoverflow.com/users/12216779/  * Linkedin: https://linkedin.com/in/shahriarraka  * -----  * Last Modified:  * Modified By:  * -----  * Copyright (c) 2022 Your Company   * @Date: 2023-11-21 13:48:43  * @Last Modified by:   Sk Shahirar Ahmed Raka  * @Last Modified time: 2023-11-21 13:48:43  */
package main

import "fmt"

type LRUCache struct {
    capacity int
    cache    map[string]*Node
    head     *Node
    tail     *Node
}

type Node struct {
    key, value string
    prev, next *Node
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[string]*Node),
    }
}

func (lru *LRUCache) addToFront(node *Node) {
    if lru.head == nil {
        lru.head = node
        lru.tail = node
    } else {
        node.next = lru.head
        lru.head.prev = node
        lru.head = node
    }
}

func (lru *LRUCache) moveToHead(node *Node) {
    if node == lru.head {
        return
    }

    if node == lru.tail {
        lru.tail = node.prev
        lru.tail.next = nil
    } else {
        node.prev.next = node.next
        node.next.prev = node.prev
    }

    lru.addToFront(node)
}

func (lru *LRUCache) evict() {
    if lru.tail != nil {
        delete(lru.cache, lru.tail.key)
        lru.tail = lru.tail.prev

        if lru.tail != nil {
            lru.tail.next = nil
        } else {
            lru.head = nil
        }
    }
}

func (lru *LRUCache) Get(key string) string {
    if node, ok := lru.cache[key]; ok {
        lru.moveToHead(node)
        return node.value
    }
    return "Song not found"
}

func (lru *LRUCache) Put(key, value string) {
    if len(lru.cache) >= lru.capacity {
        lru.evict()
    }

    newNode := &Node{key: key, value: value}
    lru.addToFront(newNode)
    lru.cache[key] = newNode
}

func main() {
    lruCache := NewLRUCache(3)

    lruCache.Put("song1", "Artist1")
    lruCache.Put("song2", "Artist2")
    lruCache.Put("song3", "Artist3")

    fmt.Println("Recently played song:", lruCache.Get("song2"))

    lruCache.Put("song4", "Artist4")

    fmt.Println("Recently played song:", lruCache.Get("song1"))
    fmt.Println("Recently played song:", lruCache.Get("song4"))
}