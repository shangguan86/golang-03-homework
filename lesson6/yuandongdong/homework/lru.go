package main

import "fmt"

type Key interface{}
type Value interface{}

type Node struct {
	key   Key
	value Value
	prev  *Node
	next  *Node
}

type LRUCache struct {
	Capacity int
	kvMap    map[Key]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) *LRUCache {
	if capacity == 0 {
		fmt.Println("capacity can not be 0")
		return nil
	}

	lru := &LRUCache{
		Capacity: capacity,
		kvMap:    make(map[Key]*Node),
		head:     new(Node),
		tail:     new(Node),
	}

	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key Key) Value {
	if node, ok := lru.kvMap[key]; ok {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCache) Put(key Key, value Value) {
	if node, ok := lru.kvMap[key]; ok {
		node.value = value
		lru.moveToHead(node)
	} else {
		node = new(Node)
		node.key = key
		node.value = value
		lru.addNode(node)

		if lru.Capacity < len(lru.kvMap) {
			lru.popTail()
		}

	}

}

func (lru *LRUCache) popTail() {
	lru.removeNode(lru.tail.prev)
}

func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addNode(node)
}

//remove node
func (lru *LRUCache) removeNode(node *Node) {
	delete(lru.kvMap, node.key)

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

//add node
//always add node behind head
func (lru *LRUCache) addNode(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next

	lru.head.next.prev = node
	lru.head.next = node

	lru.kvMap[node.key] = node
}

func main() {
	fmt.Println("Initlize lrucahe capicity 2")
	lru := Constructor(2)
	fmt.Println("Get(2) :", lru.Get(2))

	lru.Put(1, 1)
	fmt.Println("Put(1, 1)")

	lru.Put(2, 2)
	fmt.Println("Put(2, 2)")

	fmt.Println("Get(1)", lru.Get(1))

	lru.Put(3, 3)
	fmt.Println("Put(3, 3)")

	fmt.Println("Get(2) :", lru.Get(2))

	lru.Put(4, 4)
	fmt.Println("Put(4, 4)")

	fmt.Println("Get(1)", lru.Get(1))

	fmt.Println("Get(3) :", lru.Get(3))
	fmt.Println("Get(4) :", lru.Get(4))

}
