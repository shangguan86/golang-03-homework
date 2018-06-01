package go_lrucache

import "errors"

//定义Key，Value为空接口 ，空接口可以表示为任何类型
type Key interface{}
type Value interface{}


// The cache interface
type Cache interface {
	Set(key Key, value Value)
	Get(key Key) (Value, bool)
	Del(key Key)
	Size() int
}

//定义数据双链
type Node struct {
	next, prev *Node  //使用指针的方式，指定前后node
	key        Key
	value      Value
}

//定义缓存双链
type LRUCache struct {
	Capacity   int   //容量
	kvMap      map[Key]*Node   //使用map方式，便于查找，O1
	head, tail *Node
}


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
	if lru.kvMap == nil {//未初始化的map ，map==nil ===> true
		lru.kvMap = make(map[Key]*Node)
	}
	// key exists
	//当set的值存在时，移除node所在的原节点，添加node为head，表示最近一次使用
	if node, ok := lru.kvMap[key]; ok {
		lru.remove(node)
		node.value = value
		lru.setHead(node)
		return
	}
	//当set的值不存在时，添加node为head，表示最近一次使用

	var node = &Node{nil, nil, key, value}
	lru.setHead(node)
	lru.kvMap[key] = node

	// remove oldest node
	if len(lru.kvMap) > lru.Capacity {
		lru.Del(lru.tail.key)
	}
}




//实现Cache 接口的 Get方法
//get获取值，删除原位置的节点，将该节点设置为head （表示最近一次使用）
func (lru *LRUCache) Get(key Key) (value Value, ok bool) {
	if node, ok := lru.kvMap[key]; ok {//通过key，查询 缓存链表的map ，如果存在该值，返回true
		lru.remove(node)
		lru.setHead(node)
		//将该节点设置为头
		return node.value, ok
	}
	return -1, false
}

//
func (lru *LRUCache) setHead(node *Node) {
	node.next = lru.head
	//把该节点存放与当前链表头部的 前面 ，指定该节点的下一节点为链表的head
	node.prev = nil
	//设置该节点的上一节点为 nil
	if lru.head != nil {
		//双链，设置head 节点的上一节点为node
		lru.head.prev = node
	}
	if lru.tail == nil {
		lru.tail = node
	}
	//更新head为 node
	lru.head = node
}



//实现Cache 接口的 Del方法
//实现删除node，需要删除 数据链表的数据 及 缓存链表的数据
func (lru *LRUCache) Del(key Key) {
	if node, hit := lru.kvMap[key]; hit {
		lru.RemoveNode(node)
	}
}

func (lru *LRUCache) RemoveNode(node *Node) {
	// delete key from cache map
	delete(lru.kvMap, node.key)
	// remove node from cache list
	lru.remove(node)
}

//删除 缓存链表的数据
//删除 node 节点，该 node 节点的表示方式可以为 node ，node.prev.next , node.next.prev
//node 前一节点的表示方式为 node.prev  ,node 后一节点的表示方式为 node.next
func (lru *LRUCache) remove(node *Node) {
	if lru.isHead(node) {
		lru.head = node.next
		//如果删除的为头，则把头设置为当前节点的下一节点
	} else {
		node.prev.next = node.next
		//若果不为头，则把当前节点的前一节点的next ，直接指向当前节点的下一节点
	}

	if lru.isTail(node) {
		lru.tail = node.prev
		//如果删除的为尾，则把尾设置为当前节点的前一节点
	} else {
		node.next.prev = node.prev
		//若果不为尾，则把当前节点的下一节点的perv ，直接指向当前节点的前一节点
	}

}

//判断 当前节点的是否为头，如果为头，他的prev 为nil
func (lru *LRUCache) isHead(node *Node) bool {
	return node.prev == nil
}
//判断 当前节点的是否为尾，如果为尾，他的next 为nil
func (lru *LRUCache) isTail(node *Node) bool {
	return node.next == nil
}



//实现Cache 接口的 Size方法
func (lru *LRUCache) Size() int {
	return lru.Capacity
	//缓存链表的Capacity 字段即为 大小
}

