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

func (this *LRUCache) Get(key Key) Value {
	if node, ok := this.kvMap[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key Key, value Value) {
	if node, ok := this.kvMap[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node = new(Node)
		node.key = key
		node.value = value
		this.addNode(node)

		if this.Capacity < len(this.kvMap) {
			this.popTail()
		}

	}

}

func (this *LRUCache) popTail() {
	this.removeNode(this.tail.prev)
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) removeNode(node *Node) {
	delete(this.kvMap, node.key)

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addNode(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node

	this.kvMap[node.key] = node
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
