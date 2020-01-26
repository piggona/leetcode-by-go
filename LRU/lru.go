package LRU

type Cache struct {
	Data int
	Key  int
	Prev *Cache
	Next *Cache
}

type LRUCache struct {
	Top      *Cache
	Rear     *Cache
	Pointer  map[int]*Cache
	Capacity int
	Length   int
}

func Constructor(capacity int) LRUCache {
	cache := &Cache{}
	return LRUCache{
		Top:      cache,
		Rear:     cache,
		Pointer:  make(map[int]*Cache),
		Capacity: capacity,
		Length:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	if cache, ok := this.Pointer[key]; ok {
		if this.Length > 1 && cache != this.Rear {
			rear := touch(cache, this.Rear)
			this.Rear = rear
		}
		return cache.Data
	}
	return -1
}

func touch(t *Cache, rear *Cache) *Cache {
	next := t.Next
	prev := t.Prev
	next.Prev = prev
	prev.Next = next
	t.Next = rear.Next
	t.Prev = rear
	rear.Next = t
	return t
}

func (this *LRUCache) Put(key int, value int) {
	cache := &Cache{
		Data: value,
		Key:  key,
	}
	if cache, ok := this.Pointer[key]; ok {
		cache.Data = value
		// 如果空间只有一个，或者调用的这个元素恰好就是rear，那么不用进行调换操作（touch）
		if this.Length > 1 && cache != this.Rear {
			rear := touch(cache, this.Rear)
			this.Rear = rear
		}
		return
	}
	if this.Length >= this.Capacity {
		// 先淘汰front
		top := this.Top
		front := top.Next
		frontkey := front.Key
		top.Next = front.Next
		if front.Next == nil {
			this.Rear = this.Top
		} else {
			front.Next.Prev = top
		}
		this.Length--
		delete(this.Pointer, frontkey)
	}
	// 加Rear上
	this.Rear.Next = cache
	cache.Prev = this.Rear
	this.Length++
	this.Pointer[key] = cache
	this.Rear = cache
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
