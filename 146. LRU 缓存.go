package main

import "fmt"

// 146. LRU 缓存
// 中等

// 请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
// 如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

// 哈希表 + 双向链表
// 时间复杂度：对于 put 和 get 都是 O(1)。
// 空间复杂度：O(capacity)，因为哈希表和双向链表最多存储 capacity+1 个元素
type Node struct {
	key       int
	value     int
	pre, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) deleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addNodeToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addNodeToHead(node)
}

func (this *LRUCache) deleteTail() *Node {
	node := this.tail.pre
	this.deleteNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		v.value = value
		this.moveToHead(v)
	}

	if len(this.cache) >= this.capacity {
		if this.capacity == 0 {
			return
		}
		node := this.deleteTail()
		delete(this.cache, node.key)
	}
	node := &Node{
		key:   key,
		value: value,
	}
	this.cache[key] = node
	this.addNodeToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func test01() {
	var lRUCache = Constructor(2)
	lRUCache.Put(1, 1)           // 缓存是 {1=1}
	lRUCache.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)           // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4

	fmt.Println("exit test01")
}

func main() {
	test01()
	fmt.Println("exit main")
}
