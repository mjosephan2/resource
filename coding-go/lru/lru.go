package lru

type Node struct {
	prev  *Node
	next  *Node
	key   int
	value int
}

type LRUCache struct {
	cap   int
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (l *LRUCache) Put(key int, value int) {
	node, exist := l.cache[key]
	if exist {
		l.moveToHead(node)
		return
	}
	if len(l.cache) == l.cap {
		lruNode := l.getLRU()
		l.removeNode(lruNode)
		delete(l.cache, lruNode.key)
	}
	l.putNew(key, value)
}

func (l *LRUCache) Get(key int) int {
	node, exist := l.cache[key]
	if !exist {
		return -1
	}
	l.moveToHead(node)
	return node.value
}

//////////
// Helper
//////////

func (l *LRUCache) removeNode(node *Node) {
	prevNode := node.prev
	nextNode := node.next
	if prevNode != nil {
		prevNode.next = node.next
	}
	if nextNode != nil {
		nextNode.prev = node.prev
	}
	node.next = nil
	node.prev = nil
}

func (l *LRUCache) getLRU() *Node {
	return l.tail.prev
}

func (l *LRUCache) addToHead(node *Node) {
	node.next = l.head.next
	node.prev = l.head

	// next node perspective
	oldHeadNext := l.head.next
	oldHeadNext.prev = node

	// head perspective
	l.head.next = node
}
func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) putNew(key, value int) {
	newNode := &Node{
		key:   key,
		value: value,
	}
	l.cache[key] = newNode
	l.addToHead(newNode)
}
