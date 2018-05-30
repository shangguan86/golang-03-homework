package main

import (
	"errors"
	"fmt"
	"log"
)

//定义Key Value两个接口类型，可用于存放所有基本类型的数据，如Set\Get可传入随意类型
type Key interface{}
type Value interface{}

//定义一个结构体，存储一个节点的数据，每个节点数据包括：两个Node类型的指针(一个指向前面的Node,一个指向后面的Node),Key 键值，value 键值对应的值
type Node struct {
	next, prev *Node
	key        Key
	value      Value
}

//定义一个名为Cache的接口，内有四种方法:  Set Get Del Size
// The cache interface
type Cache interface {
	Set(key Key, value Value)
	Get(key Key) (Value, bool)
	Del(key Key)
	Size() (int, int)
}

//定义结构体，用来管理链表的容量、开始节点信息、结束节点信息、使用map对每个节点的key进行管理，方便Node内容提取，因map的时间复杂度为O(1)
type LRUCache struct {
	Capacity   int
	kvMap      map[Key]*Node
	head, tail *Node
}

//初始化链表函数、返回链表初始化指针

func New(capacity int) (*LRUCache, error) {
	if capacity == 0 {
		return nil, errors.New("Capacity can not be 0")
	}
	return &LRUCache{
		Capacity: capacity,
		kvMap:    make(map[Key]*Node),
	}, nil
}

/**
 * If map contains key, then update map and replace key to list head
 * Else set key to list head, check is map.size > capacity, if true, remove list tail
 */
func (lru *LRUCache) Set(key Key, value Value) {
	if lru.kvMap == nil {
		lru.kvMap = make(map[Key]*Node)
	}
	// key exists
	if node, ok := lru.kvMap[key]; ok {
		lru.remove(node)
		node.value = value
		lru.setHead(node)
		return
	}

	var node = &Node{nil, nil, key, value}
	lru.setHead(node)
	lru.kvMap[key] = node

	// remove oldest node
	if len(lru.kvMap) > lru.Capacity {
		lru.Del(lru.tail.key)
	}
}

func (lru *LRUCache) Get(key Key) (value Value, ok bool) {
	if node, ok := lru.kvMap[key]; ok {
		lru.remove(node)
		lru.setHead(node)
		return node.value, ok
	}
	return -1, false
}

func (lru *LRUCache) Del(key Key) {
	if node, hit := lru.kvMap[key]; hit {
		lru.RemoveNode(node)
	}
}

func (lru *LRUCache) Size() (int, int) {
	return lru.Capacity, len(lru.kvMap)
}

func (lru *LRUCache) RemoveNode(node *Node) {
	// delete key from cache map
	delete(lru.kvMap, node.key)
	// remove node from cache list
	lru.remove(node)
}

func (lru *LRUCache) setHead(node *Node) {
	node.next = lru.head
	node.prev = nil
	if lru.head != nil {
		lru.head.prev = node
	}
	if lru.tail == nil {
		lru.tail = node
	}
	lru.head = node
}

func (lru *LRUCache) isHead(node *Node) bool {
	return node.prev == nil
}

func (lru *LRUCache) isTail(node *Node) bool {
	return node.next == nil
}

func (lru *LRUCache) remove(node *Node) {
	if lru.isHead(node) {
		lru.head = node.next
	} else {
		node.prev.next = node.next
	}

	if lru.isTail(node) {
		lru.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

}

func main() {

	var golru Cache = new(LRUCache)
	golru, err := New(10)
	if err != nil {
		log.Fatal(err)
	}
	golru.Set(10, "shen")
	fmt.Println(golru.Size())
	fmt.Println(golru.Get(10))

	golru.Set(20, "pp")
	golru.Set(30, "ritu")
	golru.Set(40, "amen")
	cap, len := golru.Size()
	fmt.Printf("Lru Length:%d,%d\n", cap, len)
	fmt.Println(golru.Get(40))
	golru.Del(30)
	cap, len = golru.Size()
	fmt.Printf("Lru Length:%d,%d\n", cap, len)
	fmt.Println(golru)

}
